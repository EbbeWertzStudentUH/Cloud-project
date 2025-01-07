from fastapi import FastAPI
from pydantic import BaseModel
from NotificationService import NotificationService


class MessagePayload(BaseModel):
    message: str


class RestCommander:
    def __init__(self, notificationService: NotificationService):
        self.notificationService = notificationService
        self.fast_api_app = FastAPI()

        @self.fast_api_app.post("/broadcast")
        async def broadcast(payload: MessagePayload): return await self.broadcast(payload.message)

    async def broadcast(self, message: str):
        """Broadcast a message using NotificationService."""
        await self.notificationService._broadcastUpdateJSON(message)
        return {"status": "success", "message": f"Broadcasted message: {message}"}
