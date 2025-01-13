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

func (p *ProjectService) CreateProject(user_id string, name string, deadline string, github_repo string) (map[string]interface{}, bool) {
	data := map[string]interface{}{
		"name":        name,
		"deadline":    deadline,
		"github_repo": github_repo,
		"users":       []string{user_id},
	}
	return p.db.POST("/projects", data)
}
func (p *ProjectService) GetProjectById(proj_id string) (map[string]interface{}, bool) {
	project, ok := p.db.GET("/projects/" + proj_id)
	if !ok {
		return map[string]interface{}{}, false
	}
	milestones, ok := p.db.GET("/milestones/project/" + proj_id)
	if !ok {
		return map[string]interface{}{}, false
	}
	project["milestones"] = milestones
	return project, true
}
func (p *ProjectService) GetProjectsFromUser(user_id string) (map[string]interface{}, bool) {
	return p.db.GET("/projects/user/" + user_id)
}
func (p *ProjectService) AddUserToProject(proj_id string, user_id string) (map[string]interface{}, bool) {
	data := map[string]interface{}{
		"user_id": user_id,
	}
	return p.db.POST("/projects/"+proj_id+"/users", data)
}
func (p *ProjectService) CreateMilestoneInProject(proj_id string, name string, deadline string) (map[string]interface{}, bool) {
	data := map[string]interface{}{
		"name":     name,
		"deadline": deadline,
	}
	milestone, ok := p.db.POST("/milestones", data)
	if !ok {
		return map[string]interface{}{}, false
	}
	data = map[string]interface{}{
		"milestone_id": milestone["id"],
	}
	return p.db.POST("/projects/"+proj_id+"/milestones", data)
}
func (p *ProjectService) GetMilestoneById(milestone_id string) (map[string]interface{}, bool) {
	milestone, ok := p.db.GET("/milestones/" + milestone_id)
	if !ok {
		return map[string]interface{}{}, false
	}
	tasks, ok := p.db.GET("/tasks/milestone/" + milestone_id)
	if !ok {
		return map[string]interface{}{}, false
	}
	milestone["tasks"] = tasks
	return milestone, true
}
func (p *ProjectService) CreateTaskInMilestone(milestone_id string, name string) (map[string]interface{}, bool) {
	data := map[string]interface{}{
		"name":   name,
		"status": "open",
	}
	task, ok := p.db.POST("/tasks", data)
	if !ok {
		return map[string]interface{}{}, false
	}
	data = map[string]interface{}{
		"task_id": task["id"],
	}
	return p.db.POST("/milestones/"+milestone_id+"/tasks", data)
}
func (p *ProjectService) GetTaskById(task_id string) (map[string]interface{}, bool) {
	return p.db.GET("/tasks/" + task_id)
}
func (p *ProjectService) AddProblemToTask(task_id string, problem string) (map[string]interface{}, bool) {
	currentTime := time.Now()
	data := map[string]interface{}{
		"problem":   problem,
		"posted_at": currentTime.Format("2006-01-02"), // GEEN IDEE WAAROM, maar go MOET deze exacte datum hebben als format
	}
	return p.db.POST("/tasks/"+task_id+"/problems", data)
}
func (p *ProjectService) ResolveProblem(task_id string, problem_id string) (map[string]interface{}, bool) {
	data := map[string]interface{}{
		"problem_id": problem_id,
	}
	return p.db.DELETE_WITH_BODY("/tasks/"+task_id+"/problems", data)
}
func (p *ProjectService) AssignTask(task_id string, user_id string) (map[string]interface{}, bool) {
	data := map[string]interface{}{
		"user_id": user_id,
	}
	_, ok := p.db.PUT("/tasks/"+task_id+"/user", data)
	if !ok {
		return map[string]interface{}{}, false
	}
	currentTime := time.Now()
	data = map[string]interface{}{
		"start": currentTime.Format("2006-01-02"),
	}
	_, ok = p.db.PUT("/tasks/"+task_id+"/active-period", data)
	if !ok {
		return map[string]interface{}{}, false
	}
	data = map[string]interface{}{
		"status": "active",
	}
	return p.db.PATCH("/tasks/"+task_id+"/status", data)
}
func (p *ProjectService) CompleteTask(task_id string) (map[string]interface{}, bool) {
	_, ok := p.db.DELETE("/tasks/" + task_id + "/problems/all")
	if !ok {
		return map[string]interface{}{}, false
	}
	currentTime := time.Now()
	data := map[string]interface{}{
		"end": currentTime.Format("2006-01-02"),
	}
	_, ok = p.db.PATCH("/tasks/"+task_id+"/active-period", data)
	if !ok {
		return map[string]interface{}{}, false
	}
	data = map[string]interface{}{
		"status": "closed",
	}
	return p.db.PATCH("/tasks/"+task_id+"/status", data)
}
