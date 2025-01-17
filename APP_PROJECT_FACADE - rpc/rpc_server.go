package main

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

func (p *ProjectService) CreateProject(req *CreateProjectRequest, res *MinimalProjectResponse) error {
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

// func (p *ProjectService) GetProjectsFromUser(req *GetProjectsFromUserRequest, res *JSONResponse) error {
// 	jsonRes, _ := p.db.GET("/projects/user/" + req.User_id)
// 	res.Data = map[string]interface{}{
// 		"projects": jsonRes,
// 	}
// 	return nil
// }
// func (p *ProjectService) AddUserToProject(req *AddUserToProjectRequest, res *EmptyResponse) error {
// 	data := map[string]interface{}{
// 		"user_id": req.User_id,
// 	}
// 	p.db.POST("/projects/"+req.Proj_id+"/users", data)
// 	return nil
// }
// func (p *ProjectService) CreateMilestoneInProject(req *CreateMilestoneInProjectRequest, res *JSONResponse) error {
// 	data := map[string]interface{}{
// 		"name":     req.Name,
// 		"deadline": req.Deadline,
// 	}
// 	milestone, _ := p.db.POST("/milestones", data)
// 	data = map[string]interface{}{
// 		"milestone_id": milestone["id"],
// 	}
// 	p.db.POST("/projects/"+req.Proj_id+"/milestones", data)
// 	res.Data = milestone
// 	return nil
// }
// func (p *ProjectService) GetMilestoneById(req *GetMilestoneByIdRequest, res *JSONResponse) error {
// 	milestone, _ := p.db.GET("/milestones/" + req.Milestone_id)
// 	tasks, _ := p.db.GET("/tasks/milestone/" + req.Milestone_id)
// 	milestone["tasks"] = tasks
// 	res.Data = milestone
// 	return nil
// }
// func (p *ProjectService) CreateTaskInMilestone(req *CreateTaskInMilestoneRequest, res *JSONResponse) error {
// 	data := map[string]interface{}{
// 		"name":   req.Name,
// 		"status": "open",
// 	}
// 	task, _ := p.db.POST("/tasks", data)
// 	data = map[string]interface{}{
// 		"task_id": task["id"],
// 	}
// 	p.db.POST("/milestones/"+req.Milestone_id+"/tasks", data)
// 	res.Data = task
// 	return nil
// }
// func (p *ProjectService) GetTaskById(req *GetTaskByIdRequest, res *JSONResponse) error {
// 	jsonRes, _ := p.db.GET("/tasks/" + req.Task_id)
// 	res.Data = jsonRes
// 	return nil
// }
// func (p *ProjectService) AddProblemToTask(req *AddProblemToTaskRequest, res *EmptyResponse) error {
// 	currentTime := time.Now()
// 	data := map[string]interface{}{
// 		"name":      req.Problem_name,
// 		"posted_at": currentTime.Format("2006-01-02"), // GEEN IDEE WAAROM, maar go MOET deze exacte datum hebben als format
// 	}
// 	p.db.POST("/tasks/"+req.Task_id+"/problems", data)
// 	return nil
// }
// func (p *ProjectService) ResolveProblem(req *ResolveProblemRequest, res *EmptyResponse) error {
// 	data := map[string]interface{}{
// 		"problem_id": req.Problem_id,
// 	}
// 	p.db.DELETE_WITH_BODY("/tasks/"+req.Task_id+"/problems", data)
// 	return nil
// }
// func (p *ProjectService) AssignTask(req *AssignTaskRequest, res *EmptyResponse) error {
// 	data := map[string]interface{}{
// 		"user_id": req.User_id,
// 	}
// 	p.db.PUT("/tasks/"+req.Task_id+"/user", data)
// 	currentTime := time.Now()
// 	data = map[string]interface{}{
// 		"start": currentTime.Format("2006-01-02"),
// 	}
// 	p.db.PUT("/tasks/"+req.Task_id+"/active-period", data)
// 	data = map[string]interface{}{
// 		"status": "active",
// 	}
// 	p.db.PATCH("/tasks/"+req.Task_id+"/status", data)
// 	return nil
// }
// func (p *ProjectService) CompleteTask(req *CompleteTaskRequest, res *EmptyResponse) error {
// 	p.db.DELETE("/tasks/" + req.Task_id + "/problems/all")
// 	currentTime := time.Now()
// 	data := map[string]interface{}{
// 		"end": currentTime.Format("2006-01-02"),
// 	}
// 	p.db.PATCH("/tasks/"+req.Task_id+"/active-period", data)
// 	data = map[string]interface{}{
// 		"status": "closed",
// 	}
// 	p.db.PATCH("/tasks/"+req.Task_id+"/status", data)
// 	return nil
// }
