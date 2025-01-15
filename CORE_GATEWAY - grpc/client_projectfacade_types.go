package main

type RPCJSONResponse struct {
	Data map[string]interface{}
}
type RPCEmptyResponse struct {
}

type RPCCreateProjectRequest struct {
	User_id     string
	Name        string
	Deadline    string
	Github_repo string
}
type RPCGetProjectByIdRequest struct {
	Proj_id string
}
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
