from fastapi import FastAPI, HTTPException, Body
from pymongo import MongoClient
from pydantic import BaseModel, Field
from typing import List, Optional
from bson import ObjectId
import uuid
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


# FastAPI App
app = FastAPI()

# Models
class ActivePeriod(BaseModel):
    start: Optional[str]
    end: Optional[str] = None

class Problem(BaseModel):
    name: str
    posted_at: str

class Task(BaseModel):
    _id: Optional[ObjectId] = None
    name: str
    status: str
    user: Optional[str] = None
    active_period: Optional[ActivePeriod] = None
    problems: Optional[List[Problem]] = []

class Milestone(BaseModel):
    _id: Optional[ObjectId] = None
    name: str
    deadline: str

class Project(BaseModel):
    id: str = str(uuid.uuid4())
    name: str
    deadline: str
    github_repo: str
    users: List[str] #ID's
    milestones: Optional[List[str]] = [] #ID's
    
class AddUserRequest(BaseModel):
    user_id: str

# ======================================================================
# 
#       PROJECTS
# 
# ======================================================================

@app.post("/projects")
def create_new_project(project: Project):
    db.projects.insert_one(project.model_dump())
    return project

@app.get("/projects/{project_id}")
def get_project_by_id(project_id: str):
    project = db.projects.find_one({"id": project_id}, {'_id':0})
    if not project:
        raise HTTPException(status_code=404, detail="Project not found")
    return project

@app.get("/projects/user/{user_id}")
def get_projects_from_user(user_id: str):
    projects = db.projects.find({"users": user_id}, {'_id':0})
    return list(projects)

@app.post("/projects/{project_id}/users")
def add_user_to_project(project_id: str, add_req: AddUserRequest):
    result = db.projects.update_one({"id": project_id}, {"$addToSet": {"users": add_req.user_id}})
    if result.matched_count == 0:
        raise HTTPException(status_code=404, detail="Project not found")
    return {"message": "User added to project"}

# ======================================================================
# 
#       MILESTONES
# 
# ======================================================================

# ======================================================================
# 
#       TASKS
# 
# ======================================================================

# @app.get("/milestones/{milestone_id}", response_model=Milestone)
# def get_milestone(milestone_id: str):
#     milestone = db.milestones.find_one({"_id": milestone_id})
#     if not milestone:
#         raise HTTPException(status_code=404, detail="Milestone not found")
#     return milestone

# @app.get("/milestones/project/{project_id}", response_model=List[Milestone])
# def list_milestones_in_project(project_id: str):
#     """List all milestones in a project."""
#     milestones = db.milestones.find({"project_id": project_id}, {"_id": 1, "name": 1})
#     return [{"id": str(m["_id"]), "name": m["name"]} for m in milestones]

# @app.get("/tasks/{task_id}", response_model=Task)
# def get_task(task_id: str):
#     """Get full details of a specific task."""
#     task = db.tasks.find_one({"_id": task_id})
#     if not task:
#         raise HTTPException(status_code=404, detail="Task not found")
#     return task

# @app.get("/tasks/milestone/{milestone_id}", response_model=List[Task])
# def list_tasks_in_milestone(milestone_id: str):
#     """List all tasks in a milestone."""
#     tasks = db.tasks.find({"milestone_id": milestone_id}, {"_id": 1, "name": 1})
#     return [{"id": str(t["_id"]), "name": t["name"]} for t in tasks]

# @app.patch("/tasks/{task_id}/status")
# def update_task_status(task_id: str, status: str = Body(...)):
#     """Update the status of a task."""
#     result = db.tasks.update_one({"_id": task_id}, {"$set": {"status": status}})
#     if result.matched_count == 0:
#         raise HTTPException(status_code=404, detail="Task not found")
#     return {"message": "Task status updated"}

# @app.patch("/tasks/{task_id}/active_period")
# def update_task_active_period(task_id: str, start: Optional[datetime] = None, end: Optional[datetime] = None):
#     """Update the active period of a task."""
#     update_fields = {}
#     if start:
#         update_fields["active_period.start"] = start
#     if end:
#         update_fields["active_period.end"] = end
#     if not update_fields:
#         raise HTTPException(status_code=400, detail="No fields to update")

#     result = db.tasks.update_one({"_id": task_id}, {"$set": update_fields})
#     if result.matched_count == 0:
#         raise HTTPException(status_code=404, detail="Task not found")
#     return {"message": "Task active period updated"}

# @app.delete("/tasks/{task_id}/problems")
# def delete_all_problems_from_task(task_id: str):
#     """Delete all problems from a task."""
#     result = db.tasks.update_one({"_id": task_id}, {"$set": {"problems": []}})
#     if result.matched_count == 0:
#         raise HTTPException(status_code=404, detail="Task not found")
#     return {"message": "All problems deleted from task"}


if __name__ == "__main__":
    port = os.getenv('LISTEN_PORT')
    logging.info(f"running on http://0.0.0.0:{port}")
    uvicorn.run("main:app", host="0.0.0.0", port=int(port), reload=True)