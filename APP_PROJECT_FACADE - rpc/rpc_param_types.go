package main

type JSONResponse struct {
	Data map[string]interface{}
}
type EmptyResponse struct {
}

type CreateProjectRequest struct {
	User_id     string
	Name        string
	Deadline    string
	Github_repo string
}
type GetProjectByIdRequest struct {
	Proj_id string
}
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
