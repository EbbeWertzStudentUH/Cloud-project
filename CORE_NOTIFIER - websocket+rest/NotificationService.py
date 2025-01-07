from fastapi import FastAPI, WebSocket, WebSocketDisconnect
import jwt
import os

class NotificationService:
    def __init__(self):
        self.fast_api_app = FastAPI()
        self.clients = {}

        @self.fast_api_app.websocket("/ws")
        async def websocket_endpoint(websocket: WebSocket): await self.handle_endpoint(websocket)

    async def handle_endpoint(self, websocket: WebSocket):
        await websocket.accept()
        await websocket.send_json({"type":"connection_status", "data":{"message":"Connected. Send token to authorise connection.", "ok":True}})
        try:
            authenticated = False
            while True:
                message = await websocket.receive_text()
                if message and not authenticated:
                    token = message
                    jwt_data = jwt.decode(token, os.getenv('JWT_SECRET'), algorithms=["HS256"])
                    user_id = jwt_data["user_id"]
                    self.clients[user_id] = websocket
                    authenticated = True
                    await websocket.send_json({"type":"connection_status", "data":{"message":f"Connection fully registered. UserID:{user_id}", "ok":True}})
                else:
                    await websocket.send_json({"type":"connection_status", "data":{"message":f"This socket only accepts one message containing your authentication token.", "ok":False}})

        except WebSocketDisconnect:
            if user_id in self.clients:
                del self.clients[user_id]
            print(f"Client {user_id} disconnected")

    async def broadcast(self, message):
        for user_id, websocket in self.clients.items():
            try:
                await websocket.send_text(message)
            except Exception as e:
                print(f"Error sending message to {user_id}: {e}")
