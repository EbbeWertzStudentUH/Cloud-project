from fastapi import FastAPI, WebSocket, WebSocketDisconnect
import jwt
import os
import datetime

class SubscriptionManager:
    def __init__(self):
        self.topic_subscriptions: dict[tuple[str, str], set[str]] = dict() # (topic name, topic id) -> set(user id)
        self.user_id_subscriptions: dict[str, set[tuple[str, str]]] = dict() # topic id -> set(topic name, user id)

    def subscribe(self, topic:tuple[str, str], user_id:str):
        topic_name, topic_id = topic
        # voeg subscription toe aan de map
        if topic not in self.topic_subscriptions:
            self.topic_subscriptions[(topic_name, topic_id)] = set()
        self.topic_subscriptions[(topic_name, topic_id)].add(user_id)
        # voeg subscription toe aan de reverse map
        if user_id not in self.user_id_subscriptions:
            self.user_id_subscriptions[topic_id] = set()
        self.user_id_subscriptions[topic_id].add((topic_name, user_id))

    def unsubscribe(self, topic:tuple[str, str], user_id:str):
        topic_name, topic_id = topic
        # verwijder subscription uit de map
        if topic not in self.topic_subscriptions: return
        self.topic_subscriptions[(topic_name, topic_id)].remove(user_id)
        if len(self.topic_subscriptions[(topic_name, topic_id)]) == 0:
            self.topic_subscriptions.pop(topic)
        # verwijder subscription uit de reverse map
        if user_id in self.user_id_subscriptions:
            self.user_id_subscriptions[topic_id].remove((topic_name, user_id))
        if len(self.user_id_subscriptions[topic_id]) == 0:
            self.user_id_subscriptions.pop(topic_id)

    def getSubscribedUsers(self, topic:tuple[str, str]) -> set[str]:
        if topic not in self.topic_subscriptions: return set()
        return self.topic_subscriptions[topic]
    

class NotificationService:
    def __init__(self):
        self.fast_api_app = FastAPI()
        self.client_sockets = dict() # id -> socket
        self.subscriptionManager = SubscriptionManager()

        @self.fast_api_app.websocket('/ws')
        async def endpoint(websocket: WebSocket): await self._handleEndpoint(websocket)

    async def subscribe(self, user_id:str, topics:set[tuple[str, str]]):
        online_users = set() # users die dezelfde uuid hebben als de topic
        topic_name = ''
        for topic in topics:
            self.subscriptionManager.subscribe(topic, user_id)
            topic_name, topic_id = topic
            if topic_id in self.client_sockets:
                online_users.add(topic_id)
        if len(online_users) > 0:
            await self._send_topic_equals_user_message(user_id, online_users, topic_name)

    def unsubscribe(self, user_id:str, topics:set[tuple[str, str]]):
        for topic in topics:
            self.subscriptionManager.unsubscribe(topic, user_id)

    async def publishUpdate(self, topics:set[tuple[str, str]], update):
        for topic in topics:
            await self._broadcastJSON("update", update, self.subscriptionManager.getSubscribedUsers(topic))

    async def publishNotification(self, topics:set[tuple[str, str]], notification):
        timestamp = datetime.datetime.now().strftime("%H:%M")
        data = {"message": notification.message, "time":timestamp}
        for topic in topics:
            await self._broadcastJSON("notification", data, self.subscriptionManager.getSubscribedUsers(topic))
    
    async def sendUpdate(self, user_id:str, update):
        await self._broadcastJSON("update", update, [user_id])

    async def sendNotification(self, user_id:str, notification):
        timestamp = datetime.datetime.now().strftime("%H:%M")
        data = {"message": notification.message, "time":timestamp}
        await self._broadcastJSON("notification", data, [user_id])

    async def _handleEndpoint(self, websocket: WebSocket):
        user_id = None
        await websocket.accept()
        await websocket.send_json({"type":"connection_status", "data":{"message":"Connected. Send token to authorise connection.", "ok":True}})
        try:
            authenticated = False
            while True:
                message = await websocket.receive_text()
                if message and not authenticated:
                    authenticated, user_id = self._handle_registration(message, websocket)
                    await websocket.send_json({"type":"connection_status", "data":{"message":f"Connection fully registered. UserID:{user_id}", "ok":True}})
                    # update degenen die naar jou subscribed zijn (hun topic id = jouw user_id)
                    await self._notify_other_users_about_status_change(user_id, "online")
                else:
                    await websocket.send_json({"type":"connection_status", "data":{"message":f"This socket only accepts one message containing your authentication token.", "ok":False}})

        except WebSocketDisconnect:
            if user_id and user_id in self.client_sockets:
                await self._notify_other_users_about_status_change(user_id, "offline")
                self.client_sockets.pop(user_id)
            print(f"Client {user_id} disconnected")

    async def _broadcastJSON(self, type:str, data, users:set[str]):
        for user_id in users:
            webSocket:WebSocket = self.client_sockets[user_id]
            try:
                await webSocket.send_json({"type": type, "data": data})
            except Exception as e:
                print(f"Error sending message to {user_id}: {e}")

    async def _send_topic_equals_user_message(self, user_id:str, topic_ids:set[str], topic_name:str):
        message = {
            "message":"You are subscribing to a topic which contains active user id's",
            "topic":topic_name,
            "users":[{"id":id, "status":"online"} for id in topic_ids]
        }
        await self._broadcastJSON("subscribed_users_status", message, [user_id])

    async def _notify_other_users_about_status_change(self, user_id:str, status:str):
        if user_id in self.subscriptionManager.user_id_subscriptions:
            topic_and_user_set = self.subscriptionManager.user_id_subscriptions[user_id]
            for topic_and_user in topic_and_user_set:
                topic_name, other_user = topic_and_user
                message = {
                    "message":"You are subscribed to a topic which contains a user id that changed status",
                    "topic":topic_name,
                    "users":{"id":user_id, "status":status},
                }
                await self._broadcastJSON("subscribed_to_users", message, [other_user])
        

    def _handle_registration(self, token:str, websocket:WebSocket) -> tuple[bool, str]:
        try:
            jwt_data = jwt.decode(token, os.getenv('JWT_SECRET'), algorithms=['HS256'])
            user_id = jwt_data['user_id']
            self.client_sockets[user_id] = websocket
            print(f"Client {user_id} connected and authenticated")
            return True, user_id
        except: return False, ''
