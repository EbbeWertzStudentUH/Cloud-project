from typing import Any
from fastapi import FastAPI
from pydantic import BaseModel
from NotificationService import NotificationService

class TopicJSON(BaseModel):
    name: str   # friends / projects
    ids: list[str] # user id's project id's


class UpdateJSON(BaseModel):
    type: str # edit/ status change / delete / new
    subject: str # edited/new/changed object id
    data: dict[str, Any] # wat new? wat changed?

class NotificationJSON(BaseModel):
    message: str

class NotificationSendRequestJSON(BaseModel):
    user_id: str
    notification: NotificationJSON

class UpdateSendRequestJSON(BaseModel):
    user_id: str
    update: UpdateJSON

class NotificationPublishRequestJSON(BaseModel):
    topic: TopicJSON
    notification: NotificationJSON

class UpdatePublishRequestJSON(BaseModel):
    topic: TopicJSON
    update: UpdateJSON

class SubscriptionRequestJSON(BaseModel):
    user_id: str
    topic: TopicJSON


class RestCommander:
    def __init__(self, notificationService: NotificationService):
        self.notificationService = notificationService
        self.fast_api_app = FastAPI()

        #  subscriben voor topics
        @self.fast_api_app.post("/subscribe")
        async def subscribe(data: SubscriptionRequestJSON): return await self.subscribe(data)
        @self.fast_api_app.delete("/subscribe")
        def unsubscribe(data: SubscriptionRequestJSON): return self.unsubscribe(data)
        #  messages naar subscribed users voor topic
        @self.fast_api_app.post("/publish/notification")
        async def publish_notification(data: NotificationPublishRequestJSON): return await self.publishNotification(data)
        @self.fast_api_app.post("/publish/update")
        async def publish_pdate(data:UpdatePublishRequestJSON): return await self.publishUpdate(data)
        # message naar specifieke user
        @self.fast_api_app.post("/send/notification")
        async def send_notification(data: NotificationSendRequestJSON): return await self.sendNotification(data)
        @self.fast_api_app.post("/send/update")
        async def send_update(data: UpdateSendRequestJSON): return await self.sendUpdate(data)

    async def subscribe(self, data:SubscriptionRequestJSON):
        flat_topics = self._tuplifyTopic(data.topic)
        await self.notificationService.subscribe(data.user_id, flat_topics)
        return {"message":f"subscribed user {data.user_id} to topic", "topic":data.topic}
    
    def unsubscribe(self, data:SubscriptionRequestJSON):
        flat_topics = self._tuplifyTopic(data.topic)
        self.notificationService.unsubscribe(data.user_id, flat_topics)
        return {"message":f"unsubscribed user {data.user_id} from topic", "topic":data.topic}

    async def publishNotification(self, data:NotificationPublishRequestJSON):
        topic = self._tuplifyTopic(data.topic)
        await self.notificationService.publishNotification(topic, data.notification.model_dump())
        return {"message":f"sent notification to users subscribed to topic", "topic":data.topic}
    
    async def publishUpdate(self, data:UpdatePublishRequestJSON):
        topic = self._tuplifyTopic(data.topic)
        await self.notificationService.publishUpdate(topic, data.update.model_dump())
        return {"message":f"sent updates to users subscribed to topic", "topic":data.topic}
    
    async def sendNotification(self, data:NotificationSendRequestJSON):
        await self.notificationService.sendNotification(data.user_id, data.notification.model_dump())
        return {"message":f"sent notification to user {data.user_id}"}
    
    async def sendUpdate(self, data:UpdateSendRequestJSON):
        await self.notificationService.sendUpdate(data.user_id, data.update.model_dump())
        return {"message":f"sent update to user {data.user_id}"}

    def _tuplifyTopic(self, topic:TopicJSON) -> list[tuple[str, str]]:
        return [(topic.name,id) for id in topic.ids]
