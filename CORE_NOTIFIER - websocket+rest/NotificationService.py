from fastapi import FastAPI, WebSocket, WebSocketDisconnect


class NotificationService:
    def __init__(self):
        self.fast_api_app = FastAPI()
        self.clients = {}

        @self.fast_api_app.websocket("/ws")
        async def websocket_endpoint(websocket: WebSocket): await self.handle_endpoint(websocket)

    async def handle_endpoint(self, websocket: WebSocket):
        await websocket.accept()
        await websocket.send_text("Welcome! Please send your user ID.")
        user_id = None
        try:
            while True:
                message = await websocket.receive_text()
                if message and not user_id:
                    user_id = message
                    self.clients[user_id] = websocket
                    await websocket.send_text(f"User ID {user_id} received. You're now connected.")
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
