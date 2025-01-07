from fastapi import FastAPI, WebSocket, WebSocketDisconnect
import jwt
import os
import datetime

class SubscriptionManager:
    def __init__(self):
        self.topic_subscriptions:dict[str, set[str]] = dict() # topic -> set(user id)
    def subscribe(self, topic:str, user_id:str):
        if topic not in self.topic_subscriptions:
            self.topic_subscriptions[topic] = set()
        self.topic_subscriptions[topic].add(user_id)
    def unsubscribe(self, topic:str, user_id:str):
        if topic not in self.topic_subscriptions: return
        self.topic_subscriptions[topic].remove(user_id)
        if len(self.topic_subscriptions[topic]) == 0:
            self.topic_subscriptions.pop(topic)
    def getSubscribedUsers(self, topic:str) -> set[str]:
        if topic not in self.topic_subscriptions: return set()
        return self.topic_subscriptions[topic]

class NotificationService:
    def __init__(self):
        self.fast_api_app = FastAPI()
        self.client_sockets = dict() # id -> socket
        self.subscriptionManager = SubscriptionManager()

        @self.fast_api_app.websocket('/ws')
        async def endpoint(websocket: WebSocket): await self._handleEndpoint(websocket)

    def subscribe(self, user_id:str, topics:set[str]):
        for topic in topics:
            self.subscriptionManager.subscribe(topic, user_id)

    def unsubscribe(self, user_id:str, topics:set[str]):
        for topic in topics:
            self.subscriptionManager.unsubscribe(topic, user_id)

    async def publishUpdate(self, topic:str, data):
        await self._broadcastJSON("update", data, self.subscriptionManager.getSubscribedUsers(topic))

    async def publishNotification(self, topic:str, message:str):
        timestamp = datetime.datetime.now().strftime("%H:%M")
        data = {"message": message, "time":timestamp}
        await self._broadcastJSON("notification", data, self.subscriptionManager.getSubscribedUsers(topic))

    async def _handleEndpoint(self, websocket: WebSocket):
        await websocket.accept()
        await websocket.send_json({"type":"connection_status", "data":{"message":"Connected. Send token to authorise connection.", "ok":True}})
        try:
            authenticated = False
            while True:
                message = await websocket.receive_text()
                if message and not authenticated:
                    authenticated, user_id = self._handle_registration(message, websocket)
                    await websocket.send_json({"type":"connection_status", "data":{"message":f"Connection fully registered. UserID:{user_id}", "ok":True}})
                else:
                    await websocket.send_json({"type":"connection_status", "data":{"message":f"This socket only accepts one message containing your authentication token.", "ok":False}})

        except WebSocketDisconnect:
            if user_id in self.client_sockets:
                self.client_sockets.pop(user_id)
            print(f"Client {user_id} disconnected")

    async def _broadcastJSON(self, type:str, data, users:set[str]):
        for user_id in users:
            webSocket:WebSocket = self.client_sockets[user_id]
            try:
                await webSocket.send_json({"type": type, "data": data})
            except Exception as e:
                print(f"Error sending message to {user_id}: {e}")

    def _handle_registration(self, token:str, websocket:WebSocket) -> tuple[bool, str]:
        try:
            jwt_data = jwt.decode(token, os.getenv('JWT_SECRET'), algorithms=['HS256'])
            user_id = jwt_data['user_id']
            self.client_sockets[user_id] = websocket
            print(f"Client {user_id} connected and authenticated")
            return True, user_id
        except: return False, ''

