from fastapi import FastAPI, HTTPException
from pymongo import MongoClient
from pydantic import BaseModel, Field
from typing import List, Optional
from bson import ObjectId
from datetime import datetime
import uvicorn
import logging
from dotenv import load_dotenv
import os

# logging.basicConfig(level=logging.DEBUG, format="%(levelname)s: %(message)s")
load_dotenv()
print("trying to connect to db...")
mongo_url = os.getenv('MONGO_URL')
client = MongoClient(mongo_url)
try:
    client.admin.command('ping')
    print("Connected to MongoDB!")
except Exception as e:
    print("Failed to connect to MongoDB:", e)
    
db = client["project_management"]


app = FastAPI()

def object_id_to_str(obj):
    if isinstance(obj, ObjectId):
        return str(obj)
    return obj

class Problem(BaseModel):
    name: str
    resolved: Optional[bool] = False
    resolved_at: Optional[datetime] = None


class Task(BaseModel):
    title: str
    description: Optional[str] = None
    user_id: Optional[str] = None
    problems: List[Problem] = []
    status: str = "Open"
    active_start: Optional[datetime] = None
    active_end: Optional[datetime] = None


class Milestone(BaseModel):
    title: str
    deadline: datetime
    tasks: List[Task] = []


class Project(BaseModel):
    name: str
    description: Optional[str] = None
    deadline: datetime
    user_ids: List[str]
    github_repo: str
    milestones: List[Milestone] = []


@app.post("/projects/")
async def create_project(project: Project):
    project_dict = project.dict()
    result = db.projects.insert_one(project_dict)
    project_dict["_id"] = str(result.inserted_id)
    return project_dict


@app.get("/projects/{project_id}")
async def get_project(project_id: str):
    project = db.projects.find_one({"_id": ObjectId(project_id)})
    if not project:
        raise HTTPException(status_code=404, detail="Project not found")
    project["_id"] = object_id_to_str(project["_id"])
    return project


@app.put("/projects/{project_id}")
async def update_project(project_id: str, project: Project):
    project_dict = project.dict()
    result = db.projects.update_one({"_id": ObjectId(project_id)}, {"$set": project_dict})
    if result.matched_count == 0:
        raise HTTPException(status_code=404, detail="Project not found")
    return {"message": "Project updated successfully"}


@app.delete("/projects/{project_id}")
async def delete_project(project_id: str):
    result = db.projects.delete_one({"_id": ObjectId(project_id)})
    if result.deleted_count == 0:
        raise HTTPException(status_code=404, detail="Project not found")
    return {"message": "Project deleted successfully"}


@app.post("/projects/{project_id}/milestones/")
async def add_milestone(project_id: str, milestone: Milestone):
    milestone_dict = milestone.dict()
    result = db.projects.update_one(
        {"_id": ObjectId(project_id)},
        {"$push": {"milestones": milestone_dict}}
    )
    if result.matched_count == 0:
        raise HTTPException(status_code=404, detail="Project not found")
    return {"message": "Milestone added successfully"}


@app.get("/milestones/{milestone_id}")
async def get_milestone(project_id: str, milestone_id: str):
    project = db.projects.find_one({"_id": ObjectId(project_id)})
    if not project:
        raise HTTPException(status_code=404, detail="Project not found")
    milestones = project.get("milestones", [])
    milestone = next((m for m in milestones if m["_id"] == milestone_id), None)
    if not milestone:
        raise HTTPException(status_code=404, detail="Milestone not found")
    # Add stats
    milestone["total_tasks"] = len(milestone["tasks"])
    milestone["completed_tasks"] = sum(
        1 for t in milestone["tasks"] if t["status"] == "closed"
    )
    milestone["total_problems"] = sum(len(t["problems"]) for t in milestone["tasks"])
    return milestone


@app.post("/tasks/")
async def add_task_to_milestone(project_id: str, milestone_id: str, task: Task):
    task_dict = task.dict()
    result = db.projects.update_one(
        {"_id": ObjectId(project_id), "milestones._id": milestone_id},
        {"$push": {"milestones.$.tasks": task_dict}}
    )
    if result.matched_count == 0:
        raise HTTPException(status_code=404, detail="Project or milestone not found")
    return {"message": "Task added successfully"}


@app.patch("/tasks/{task_id}/status")
async def update_task_status(project_id: str, milestone_id: str, task_id: str, status: str):
    valid_status = {"open", "active", "closed"}
    if status not in valid_status:
        raise HTTPException(status_code=400, detail="Invalid status")
    now = datetime.utcnow()
    update_fields = {"status": status}
    if status == "active":
        update_fields["active_start"] = now
    elif status == "closed":
        update_fields["active_end"] = now
    result = db.projects.update_one(
        {"_id": ObjectId(project_id), "milestones._id": milestone_id, "milestones.tasks._id": task_id},
        {"$set": {f"milestones.$.tasks.$.status": update_fields}}
    )
    if result.matched_count == 0:
        raise HTTPException(status_code=404, detail="Task not found")
    return {"message": "Task status updated successfully"}



if __name__ == "__main__":
    port = os.getenv('LISTEN_PORT')
    logging.info(f"running on http://0.0.0.0:{port}")
    uvicorn.run("main:app", host="0.0.0.0", port=int(port), reload=True)