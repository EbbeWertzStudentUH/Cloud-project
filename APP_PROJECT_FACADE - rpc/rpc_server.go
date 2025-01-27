package main

import "time"

type ProjectService struct {
	db  *ProjectDBClient
	udb *UserDBClient
}

func NewProjectService(db *ProjectDBClient, udb *UserDBClient) *ProjectService {
	return &ProjectService{
		db:  db,
		udb: udb,
	}
}

type World struct {
	World string
}
type HelloWorld struct {
	HelloWorld string
}

func (p *ProjectService) Hello(req *World, res *HelloWorld) error {
	res.HelloWorld = "hello " + req.World
	return nil
}

func (p *ProjectService) CreateProject(req *CreateProjectRequest, res *MinimalProject) error {
	data := map[string]interface{}{
		"name":        req.Name,
		"deadline":    req.Deadline,
		"github_repo": req.Github_repo,
		"users":       []string{req.User_id},
	}
	jsonRes, _ := p.db.POST("/projects", data)
	res.Id = jsonRes["id"].(string)
	res.Name = jsonRes["name"].(string)
	res.Deadline = jsonRes["deadline"].(string)
	res.NumOfUsers = len(jsonRes["users"].([]interface{}))
	return nil
}

func (p *ProjectService) GetFullProjectById(req *GetProjectByIdRequest, res *FullProject) error {
	projectJSON, _ := p.db.GET("/projects/" + req.Proj_id)
	milestonesJSON, _ := p.db.GETMULTI("/milestones/project/" + req.Proj_id)
	milestones := []Milestone{}
	for _, milestoneJSON := range milestonesJSON {
		milestone := p.milestoneJSONToMilestone(milestoneJSON)
		milestones = append(milestones, milestone)
	}
	user_ids_interface := projectJSON["users"].([]interface{})
	user_ids := []string{}
	for _, user_id := range user_ids_interface {
		user_ids = append(user_ids, user_id.(string))
	}
	users := p.getUsers(user_ids)
	res.Id = projectJSON["id"].(string)
	res.Name = projectJSON["name"].(string)
	res.Deadline = projectJSON["deadline"].(string)
	res.GithubRepo = projectJSON["github_repo"].(string)
	res.Users = users
	res.Milestones = milestones
	return nil
}

func (p *ProjectService) GetProjectsFromUser(req *GetProjectsFromUserRequest, res *MinimalProjects) error {
	projectJSONs, _ := p.db.GETMULTI("/projects/user/" + req.User_id)
	projects := []MinimalProject{}
	for _, projectJSON := range projectJSONs {
		projects = append(projects, MinimalProject{
			Id:         projectJSON["id"].(string),
			Name:       projectJSON["name"].(string),
			Deadline:   projectJSON["deadline"].(string),
			NumOfUsers: len(projectJSON["users"].([]interface{})),
		})
	}
	res.Projects = projects
	return nil
}

func (p *ProjectService) AddUserToProject(req *AddUserToProjectRequest, res *UserAddToProjectResponse) error {
	data := map[string]interface{}{
		"user_id": req.User_id,
	}
	p.db.POST("/projects/"+req.Proj_id+"/users", data)
	userJSOn, _ := p.udb.QueryUsers([]string{req.User_id})
	projectJSON, _ := p.db.GET("/projects/" + req.Proj_id)
	res.Project = MinimalProject{
		Id:         projectJSON["id"].(string),
		Name:       projectJSON["name"].(string),
		Deadline:   projectJSON["deadline"].(string),
		NumOfUsers: len(projectJSON["users"].([]interface{})),
	}
	res.User = User{
		Id:        userJSOn[0]["id"].(string),
		FirstName: userJSOn[0]["first_name"].(string),
		LastName:  userJSOn[0]["last_name"].(string),
	}
	return nil
}

func (p *ProjectService) CreateMilestoneInProject(req *CreateMilestoneInProjectRequest, res *Milestone) error {
	data := map[string]interface{}{
		"name":     req.Name,
		"deadline": req.Deadline,
	}
	milestoneJSON, _ := p.db.POST("/milestones", data)
	milestone := p.milestoneJSONToMilestone(milestoneJSON)
	data = map[string]interface{}{
		"milestone_id": milestoneJSON["id"],
	}
	p.db.POST("/projects/"+req.Proj_id+"/milestones", data)
	*res = milestone
	return nil
}

func (p *ProjectService) CreateTaskInMilestone(req *CreateTaskInMilestoneRequest, res *Task) error {
	data := map[string]interface{}{
		"name":   req.Name,
		"status": "open",
	}
	taskJSOn, _ := p.db.POST("/tasks", data)
	data = map[string]interface{}{
		"task_id": taskJSOn["id"],
	}
	p.db.POST("/milestones/"+req.Milestone_id+"/tasks", data)
	task := p.taskJSONTOTask(taskJSOn)
	*res = task
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

func (p *ProjectService) ResolveProblem(req *ResolveProblemRequest, res *Problem) error {
	data := map[string]interface{}{
		"problem_id": req.Problem_id,
	}
	problemJSON, _ := p.db.GET("/tasks/" + req.Task_id + "/problems/" + req.Problem_id)
	p.db.DELETE_WITH_BODY("/tasks/"+req.Task_id+"/problems", data)
	res.Id = problemJSON["id"].(string)
	res.Name = problemJSON["name"].(string)
	res.PostedAt = problemJSON["posted_at"].(string)
	return nil
}
func (p *ProjectService) AssignTask(req *AssignTaskRequest, res *Task) error {
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
	taskJSON, _ := p.db.GET("/tasks/" + req.Task_id)
	*res = p.taskJSONTOTask(taskJSON)
	return nil
}
func (p *ProjectService) CompleteTask(req *CompleteTaskRequest, res *Task) error {
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
	taskJSON, _ := p.db.GET("/tasks/" + req.Task_id)
	*res = p.taskJSONTOTask(taskJSON)
	return nil
}

// ====================================
// 	PRIVATE HELPERS
// ====================================

func (p *ProjectService) getUsers(userID []string) []User {
	usersJSONs, _ := p.udb.QueryUsers(userID)
	users := []User{}
	for _, userJSON := range usersJSONs {
		users = append(users, User{
			Id:        userJSON["id"].(string),
			FirstName: userJSON["first_name"].(string),
			LastName:  userJSON["last_name"].(string),
		})
	}
	return users
}

func (p *ProjectService) milestoneJSONToMilestone(milestoneJSON map[string]interface{}) Milestone {
	tasksJSON, _ := p.db.GETMULTI("/tasks/milestone/" + milestoneJSON["id"].(string))
	tasks := []Task{}
	numOfProblems := 0
	NumOfFinishedTasks := 0
	for _, taskJSON := range tasksJSON {
		task := p.taskJSONTOTask(taskJSON)
		tasks = append(tasks, task)
		numOfProblems += task.NumOfProblems
		if task.Status == "closed" {
			NumOfFinishedTasks++
		}
	}
	return Milestone{
		Id:                 milestoneJSON["id"].(string),
		Name:               milestoneJSON["name"].(string),
		Deadline:           milestoneJSON["deadline"].(string),
		NumOfProblems:      numOfProblems,
		NumOfTasks:         len(tasks),
		NumOfFinishedTasks: NumOfFinishedTasks,
		Tasks:              tasks,
	}
}

func (p *ProjectService) taskJSONTOTask(taskJSON map[string]interface{}) Task {
	user_id, ok := taskJSON["user"]
	var userPtr *User = nil
	if ok && user_id != nil {
		userPtr = &(p.getUsers([]string{user_id.(string)})[0])
	}
	var activeStartPtr *string = nil
	var activeEndPtr *string = nil
	activePeriod, ok := taskJSON["active_period"]
	if ok && activePeriod != nil {
		activeStart, ok := activePeriod.(map[string]interface{})["start"]
		if ok && activeStart != nil {
			activeStartStr := activeStart.(string)
			activeStartPtr = &activeStartStr
		}
		activeEnd, ok := activePeriod.(map[string]interface{})["end"]
		if ok && activeEnd != nil {
			activeEndStr := activeEnd.(string)
			activeEndPtr = &activeEndStr
		}
	}
	problems := []Problem{}
	for _, problemJSON := range taskJSON["problems"].([]interface{}) {
		problems = append(problems, Problem{
			Id:       problemJSON.(map[string]interface{})["id"].(string),
			Name:     problemJSON.(map[string]interface{})["name"].(string),
			PostedAt: problemJSON.(map[string]interface{})["posted_at"].(string),
		})
	}

	return Task{
		Id:              taskJSON["id"].(string),
		Name:            taskJSON["name"].(string),
		Status:          taskJSON["status"].(string),
		User:            userPtr,
		ActiveStartDate: activeStartPtr,
		ActiveEndDate:   activeEndPtr,
		NumOfProblems:   len(problems),
		IsAssigned:      userPtr != nil,
		Problems:        problems,
	}
}
