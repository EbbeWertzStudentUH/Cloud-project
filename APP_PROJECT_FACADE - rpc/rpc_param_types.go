package main

type EmptyResponse struct {
}

type CreateProjectRequest struct {
	User_id     string
	Name        string
	Deadline    string
	Github_repo string
}
type MinimalProject struct {
	Id         string
	Name       string
	Deadline   string
	NumOfUsers int
}
type UserAddToProjectResponse struct {
	Project MinimalProject
	User    User
}
type MinimalProjects struct {
	Projects []MinimalProject
}
type User struct {
	Id        string
	FirstName string
	LastName  string
}
type Problem struct {
	Id       string
	Name     string
	PostedAt string
}
type Task struct {
	Id              string
	Name            string
	Status          string
	User            *User
	ActiveStartDate *string
	ActiveEndDate   *string
	NumOfProblems   int
	IsAssigned      bool
	Problems        []Problem
}
type Milestone struct {
	Id                 string
	Name               string
	Deadline           string
	NumOfProblems      int
	NumOfTasks         int
	NumOfFinishedTasks int
	Tasks              []Task
}
type GetProjectByIdRequest struct {
	Proj_id string
}
type FullProject struct {
	Id         string
	Name       string
	Deadline   string
	GithubRepo string
	Users      []User
	Milestones []Milestone
}

//
//
//

type GetProjectsFromUserRequest struct {
	User_id string
}
type AddUserToProjectRequest struct {
	Proj_id string
	User_id string
}
type CreateMilestoneInProjectRequest struct {
	Proj_id  string
	Name     string
	Deadline string
}
type GetMilestoneByIdRequest struct {
	Milestone_id string
}
type CreateTaskInMilestoneRequest struct {
	Milestone_id string
	Name         string
}
type GetTaskByIdRequest struct {
	Task_id string
}
type AddProblemToTaskRequest struct {
	Task_id      string
	Problem_name string
}
type ResolveProblemRequest struct {
	Task_id    string
	Problem_id string
}
type AssignTaskRequest struct {
	Task_id string
	User_id string
}
type CompleteTaskRequest struct {
	Task_id string
}
