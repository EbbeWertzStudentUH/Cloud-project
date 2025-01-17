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
	res.NumOfUsers = len(jsonRes["users"].([]string))
	return nil
}

func (p *ProjectService) GetFullProjectById(req *GetProjectByIdRequest, res *FullProject) error {
	projectJSON, _ := p.db.GET("/projects/" + req.Proj_id)
	milestonesJSON, _ := p.db.GETMULTI("/milestones/project/" + projectJSON["id"].(string))
	milestones := []Milestone{}
	for _, milestoneJSON := range milestonesJSON {
		milestone := p.milestoneJSONToMilestone(milestoneJSON)
		milestones = append(milestones, milestone)
	}
	users := p.getUsers(projectJSON["users"].([]string))
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
			NumOfUsers: len(projectJSON["users"].([]string)),
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
		NumOfUsers: len(projectJSON["users"].([]string)),
	}
	res.User = User{
		Id:        userJSOn[0]["id"].(string),
		FirstName: userJSOn[0]["id"].(string),
		LastName:  userJSOn[0]["id"].(string),
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
	tasksJSON, _ := p.db.GET("/tasks/milestone/" + milestoneJSON["id"].(string))
	tasks := []Task{}
	numOfProblems := 0
	NumOfFinishedTasks := 0
	for _, taskJSON := range tasksJSON {
		task := p.taskJSONTOTask(taskJSON.(map[string]interface{}))
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
	if ok {
		userPtr = &(p.getUsers([]string{user_id.(string)})[0])
	}
	var activeStartPtr *string = nil
	var activeEndPtr *string = nil
	activePeriod, ok := taskJSON["active_period"]
	if ok {
		activeStart, ok := activePeriod.(map[string]interface{})["start"].(string)
		if ok {
			activeStartPtr = &activeStart
		}
		activeEnd, ok := activePeriod.(map[string]interface{})["end"].(string)
		if ok {
			activeEndPtr = &activeEnd
		}
	}
	problems := []Problem{}
	for _, problemJSON := range taskJSON["problems"].([]map[string]string) {
		problems = append(problems, Problem{
			Id:       problemJSON["id"],
			Name:     problemJSON["name"],
			PostedAt: problemJSON["posted_at"],
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
