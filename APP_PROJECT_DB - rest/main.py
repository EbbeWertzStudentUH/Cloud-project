from fastapi import FastAPI, HTTPException, Body
from pymongo import MongoClient
from pydantic import BaseModel, Field
from typing import List, Optional
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
    id: str = str(uuid.uuid4())
    name: str
    posted_at: str

class Task(BaseModel):
    id: str = str(uuid.uuid4())
    name: str
    status: str = "open"
    user: Optional[str] = None
    active_period: Optional[ActivePeriod] = None
    problems: List[Problem] = []

class Milestone(BaseModel):
    id: str = str(uuid.uuid4())
    name: str
    deadline: str
    tasks: List[str] = [] #ID's

class Project(BaseModel):
    id: str = str(uuid.uuid4())
    name: str
    deadline: str
    github_repo: str
    users: List[str] #ID's
    milestones: List[str] = [] #ID's
    
class AddUserRequest(BaseModel):
    user_id: str
class AddMilestoneRequest(BaseModel):
    milestone_id: str
class AddTaskRequest(BaseModel):
    task_id: str
class DeleteProblemRequest(BaseModel):
    problem_id: str
class PutUserRequest(BaseModel):
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

@app.post("/projects/{project_id}/milestones")
def add_milestone_to_project(project_id: str, add_req: AddMilestoneRequest):
    result = db.projects.update_one({"id": project_id}, {"$addToSet": {"milestones": add_req.milestone_id}})
    if result.matched_count == 0:
        raise HTTPException(status_code=404, detail="Project not found")
    return {"message": "Milestone added to project"}

# ======================================================================
# 
#       MILESTONES
# 
# ======================================================================

@app.post("/milestones")
def create_milestone(milestone: Milestone):
    db.milestones.insert_one(milestone.model_dump())
    return milestone

@app.get("/milestones/{milestone_id}")
def get_milestone_by_id(milestone_id: str):
    milestone = db.milestones.find_one({"id": milestone_id}, {'_id':0})
    if not milestone:
        raise HTTPException(status_code=404, detail="Milestone not found")
    return milestone

@app.get("/milestones/project/{project_id}")
def get_milestones_from_project(project_id: str):
    ids = db.projects.find_one({"id": project_id}, {'milestones':1, '_id':0})['milestones']
    milestones = db.milestones.find({"id": {"$in": ids}}, {'_id':0})
    return list(milestones)

@app.post("/milestones/{milestone_id}/tasks")
def add_task_to_milestone(milestone_id: str, add_req: AddTaskRequest):
    result = db.milestones.update_one({"id": milestone_id}, {"$addToSet": {"tasks": add_req.task_id}})
    if result.matched_count == 0:
        raise HTTPException(status_code=404, detail="Milestone not found")
    return {"message": "Milestone added to project"}

# ======================================================================
# 
#       TASKS
# 
# ======================================================================

@app.post("/tasks")
def create_task(task: Task):
    db.tasks.insert_one(task.model_dump())
    return task

@app.get("/tasks/{task_id}")
def get_task_by_id(task_id: str):
    task = db.tasks.find_one({"id": task_id}, {'_id':0})
    if not task:
        raise HTTPException(status_code=404, detail="Task not found")
    return task

@app.get("/tasks/milestone/{milestone_id}")
def get_tasks_from_milestone(milestone_id: str):
    ids = db.milestones.find_one({"id": milestone_id}, {'tasks':1, '_id':0})['tasks']
    tasks = db.tasks.find({"id": {"$in": ids}}, {'_id':0})
    return list(tasks)

@app.post("/tasks/{task_id}/problems")
def add_problem_to_task(task_id: str, problem: Problem):
    result = db.tasks.update_one({"id": task_id}, {"$addToSet": {"problems": problem.model_dump()}})
    if result.matched_count == 0:
        raise HTTPException(status_code=404, detail="Task not found")
    return {"message": "Problem added to task"}

@app.delete("/tasks/{task_id}/problems")
def remove_problem_from_task(task_id: str, delete_req: DeleteProblemRequest):
    result = db.tasks.update_one({"id": task_id}, {"$pull": {"problems": {"id": delete_req.problem_id}}})
    if result.matched_count == 0:
        raise HTTPException(status_code=404, detail="Task not found")
    return {"message": "Problem removed from task"}

@app.delete("/tasks/{task_id}/problems/all")
def remove_problem_from_task(task_id: str):
    result = db.tasks.update_one({"id": task_id}, {"$set": {"problems": []}})
    if result.matched_count == 0:
        raise HTTPException(status_code=404, detail="Task not found")
    return {"message": "All problems removed from task"}

@app.put("/tasks/{task_id}/user")
def put_user_on_task(task_id: str, put_req:PutUserRequest):
    result = db.tasks.update_one({"id": task_id}, {"$set": {"user": put_req.user_id}})
    if result.matched_count == 0:
        raise HTTPException(status_code=404, detail="Task not found")
    return {"message": "User put on task"}

@app.put("/tasks/{task_id}/active-period")
def put_user_on_task(task_id: str, period:ActivePeriod):
    result = db.tasks.update_one({"id": task_id}, {"$set": {"active_period": period.model_dump()}})
    if result.matched_count == 0:
        raise HTTPException(status_code=404, detail="Task not found")
    return {"message": "active_period put on task"}

@app.patch("/tasks/{task_id}/active-period")
def put_user_on_task(task_id: str, period:ActivePeriod):
    result = db.tasks.update_one({"id": task_id}, {"$set": {"active_period": {"end": period.end}}})
    if result.matched_count == 0:
        raise HTTPException(status_code=404, detail="Task not found")
    return {"message": "end timestamp set on active_period"}



if __name__ == "__main__":
    port = os.getenv('LISTEN_PORT')
    logging.info(f"running on http://0.0.0.0:{port}")
    uvicorn.run("main:app", host="0.0.0.0", port=int(port), reload=True)