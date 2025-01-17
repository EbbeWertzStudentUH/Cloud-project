package main

type RPCEmptyResponse struct {
}

type RPCCreateProjectRequest struct {
	User_id     string
	Name        string
	Deadline    string
	Github_repo string
}
type RPCMinimalProject struct {
	Id         string
	Name       string
	Deadline   string
	NumOfUsers int
}
type RPCUserAddToProjectResponse struct {
	Project RPCMinimalProject
	User    RPCUser
}
type RPCMinimalProjects struct {
	Projects []RPCMinimalProject
}
type RPCUser struct {
	Id        string
	FirstName string
	LastName  string
}
type RPCProblem struct {
	Id       string
	Name     string
	PostedAt string
}
type RPCTask struct {
	Id              string
	Name            string
	Status          string
	User            *RPCUser
	ActiveStartDate *string
	ActiveEndDate   *string
	NumOfProblems   int
	IsAssigned      bool
	Problems        []RPCProblem
}
type RPCMilestone struct {
	Id                 string
	Name               string
	Deadline           string
	NumOfProblems      int
	NumOfTasks         int
	NumOfFinishedTasks int
	Tasks              []RPCTask
}
type RPCGetProjectByIdRequest struct {
	Proj_id string
}
type RPCFullProject struct {
	Id         string
	Name       string
	Deadline   string
	GithubRepo string
	Users      []RPCUser
	Milestones []RPCMilestone
}

//
//
//

type RPCGetProjectsFromUserRequest struct {
	User_id string
}
type RPCAddUserToProjectRequest struct {
	Proj_id string
	User_id string
}
type RPCCreateMilestoneInProjectRequest struct {
	Proj_id  string
	Name     string
	Deadline string
}
type RPCGetMilestoneByIdRequest struct {
	Milestone_id string
}
type RPCCreateTaskInMilestoneRequest struct {
	Milestone_id string
	Name         string
}
type RPCGetTaskByIdRequest struct {
	Task_id string
}
type RPCAddProblemToTaskRequest struct {
	Task_id      string
	Problem_name string
}
type RPCResolveProblemRequest struct {
	Task_id    string
	Problem_id string
}
type RPCAssignTaskRequest struct {
	Task_id string
	User_id string
}
type RPCCompleteTaskRequest struct {
	Task_id string
}
