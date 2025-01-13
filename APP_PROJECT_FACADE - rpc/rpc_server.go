package main

import (
	"time"
)

type ProjectService struct {
	db *ProjectDBClient
}

func NewProjectService(db *ProjectDBClient) *ProjectService {
	return &ProjectService{
		db: db,
	}
}

func (p *ProjectService) CreateProject(req *CreateProjectRequest, res *JSONResponse) error {
	data := map[string]interface{}{
		"name":        req.Name,
		"deadline":    req.Deadline,
		"github_repo": req.Github_repo,
		"users":       []string{req.User_id},
	}
	jsonRes, _ := p.db.POST("/projects", data)
	res.Data = jsonRes
	return nil
}
func (p *ProjectService) GetProjectById(req *GetProjectByIdRequest, res *JSONResponse) error {
	project, _ := p.db.GET("/projects/" + req.Proj_id)
	milestones, _ := p.db.GET("/milestones/project/" + req.Proj_id)
	project["milestones"] = milestones
	res.Data = project
	return nil
}
func (p *ProjectService) GetProjectsFromUser(req *GetProjectsFromUserRequest, res *JSONResponse) error {
	jsonRes, _ := p.db.GET("/projects/user/" + req.User_id)
	res.Data = jsonRes
	return nil
}
func (p *ProjectService) AddUserToProject(req *AddUserToProjectRequest, res *EmptyResponse) error {
	data := map[string]interface{}{
		"user_id": req.User_id,
	}
	p.db.POST("/projects/"+req.Proj_id+"/users", data)
	return nil
}
func (p *ProjectService) CreateMilestoneInProject(req *CreateMilestoneInProjectRequest, res *JSONResponse) error {
	data := map[string]interface{}{
		"name":     req.Name,
		"deadline": req.Deadline,
	}
	milestone, _ := p.db.POST("/milestones", data)
	data = map[string]interface{}{
		"milestone_id": milestone["id"],
	}
	p.db.POST("/projects/"+req.Proj_id+"/milestones", data)
	res.Data = milestone
	return nil
}
func (p *ProjectService) GetMilestoneById(req *GetMilestoneByIdRequest, res *JSONResponse) error {
	milestone, _ := p.db.GET("/milestones/" + req.Milestone_id)
	tasks, _ := p.db.GET("/tasks/milestone/" + req.Milestone_id)
	milestone["tasks"] = tasks
	res.Data = milestone
	return nil
}
func (p *ProjectService) CreateTaskInMilestone(req *CreateTaskInMilestoneRequest, res *JSONResponse) error {
	data := map[string]interface{}{
		"name":   req.Name,
		"status": "open",
	}
	task, _ := p.db.POST("/tasks", data)
	data = map[string]interface{}{
		"task_id": task["id"],
	}
	p.db.POST("/milestones/"+req.Milestone_id+"/tasks", data)
	res.Data = task
	return nil
}
func (p *ProjectService) GetTaskById(req *GetTaskByIdRequest, res *JSONResponse) error {
	jsonRes, _ := p.db.GET("/tasks/" + req.Task_id)
	res.Data = jsonRes
	return nil
}
func (p *ProjectService) AddProblemToTask(req *AddProblemToTaskRequest, res *EmptyResponse) error {
	currentTime := time.Now()
	data := map[string]interface{}{
		"name":      req.Problem_name,
		"posted_at": currentTime.Format("2006-01-02"), // GEEN IDEE WAAROM, maar go MOET deze exacte datum hebben als format
	}
	p.db.POST("/tasks/"+req.Task_id+"/problems", data)
	return nil
}
func (p *ProjectService) ResolveProblem(req *ResolveProblemRequest, res *EmptyResponse) error {
	data := map[string]interface{}{
		"problem_id": req.Problem_id,
	}
	p.db.DELETE_WITH_BODY("/tasks/"+req.Task_id+"/problems", data)
	return nil
}
func (p *ProjectService) AssignTask(req *AssignTaskRequest, res *EmptyResponse) error {
	data := map[string]interface{}{
		"user_id": req.User_id,
	}
	p.db.PUT("/tasks/"+req.Task_id+"/user", data)
	currentTime := time.Now()
	data = map[string]interface{}{
		"start": currentTime.Format("2006-01-02"),
	}
	p.db.PUT("/tasks/"+req.Task_id+"/active-period", data)
	data = map[string]interface{}{
		"status": "active",
	}
	p.db.PATCH("/tasks/"+req.Task_id+"/status", data)
	return nil
}
func (p *ProjectService) CompleteTask(req *CompleteTaskRequest, res *EmptyResponse) error {
	p.db.DELETE("/tasks/" + req.Task_id + "/problems/all")
	currentTime := time.Now()
	data := map[string]interface{}{
		"end": currentTime.Format("2006-01-02"),
	}
	p.db.PATCH("/tasks/"+req.Task_id+"/active-period", data)
	data = map[string]interface{}{
		"status": "closed",
	}
	p.db.PATCH("/tasks/"+req.Task_id+"/status", data)
	return nil
}
