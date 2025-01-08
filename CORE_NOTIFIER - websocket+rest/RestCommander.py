from fastapi import FastAPI
from pydantic import BaseModel
from NotificationService import NotificationService


class MultiTopicJSON(BaseModel):
    name: str
    ids: list[str]

class TopicJSON(BaseModel):
    name: str
    id: str

class SubscriptionJSON(BaseModel):
    user_id: str
    topic: MultiTopicJSON

class UpdateJSON(BaseModel):
    type: str # edit/ status change / delete / new
    subject: str # edited/new/changed object id
    data: str # wat new? wat changed?

class UpdateMessageJSON(BaseModel):
    topic: MultiTopicJSON
    update: UpdateJSON


class NotificationJSON(BaseModel):
    topic: TopicJSON
    message: str


class RestCommander:
    def __init__(self, notificationService: NotificationService):
        self.notificationService = notificationService
        self.fast_api_app = FastAPI()

        @self.fast_api_app.post("/subscribe")
        def subscribe(data: SubscriptionJSON): return self.subscribe(data)
        @self.fast_api_app.delete("/subscribe")
        def unsubscribe(data: SubscriptionJSON): return self.unsubscribe(data)
        @self.fast_api_app.post("/publish/notification")
        async def publish_notification(data: NotificationJSON): return await self.publishNotification(data)
        @self.fast_api_app.post("/publish/update")
        async def publishUpdate(data:UpdateMessageJSON): return await self.publishUpdate(data)

    def subscribe(self, data:SubscriptionJSON):
        flat_topics = self._flattenMultiTopic(data.topic)
        self.notificationService.subscribe(data.user_id, flat_topics)
        return {"message":f"subscribed user {data.user_id} to topic", "topic":data.topic}
    
    def unsubscribe(self, data:SubscriptionJSON):
        flat_topics = self._flattenMultiTopic(data.topic)
        self.notificationService.unsubscribe(data.user_id, flat_topics)
        return {"message":f"unsubscribed user {data.user_id} from topic", "topic":data.topic}

    async def publishNotification(self, data:NotificationJSON):
        topic = self._flattenTopic(data.topic)
        await self.notificationService.publishNotification(topic, data.message)
        return {"message":f"sent notification to users subscribed to topic", "topic":data.topic}
    
    async def publishUpdate(self, data:UpdateMessageJSON):
        topic = self._flattenTopic(data.topic)
        await self.notificationService.publishUpdate(topic, data.update)
        return {"message":f"sent updates to users subscribed to topic", "topic":data.topic}

    def _flattenTopic(self, topic:TopicJSON) -> str:
        return f"{topic.name}_{topic.id}"

    def _flattenMultiTopic(self, topic:MultiTopicJSON) -> list[str]:
        return [f"{topic.name}_{id}" for id in topic.ids]
