package main

import (
	"context"
	pb "gateway_service/protobuf_generated"
	"log"
)

type ProjectServiceServer struct {
	pb.UnimplementedProjectServiceServer
	pfc ProjectFacadeClient
	nc  NotifierClient
}

type World struct {
	World string
}
type HelloWorld struct {
	HelloWorld string
}

func (s *ProjectServiceServer) Hello(ctx context.Context, req *pb.World) (*pb.HelloWorld, error) {
	rpcres := &HelloWorld{}
	rpcreq := &World{World: req.World}
	s.pfc.call("Hello", rpcreq, rpcres)
	log.Println(rpcres.HelloWorld)
	return &pb.HelloWorld{HelloWorld: rpcres.HelloWorld}, nil
}

// proj_fac: CreateProject
// notifier: send update projects list to user
func (s *ProjectServiceServer) CreateProject(ctx context.Context, req *pb.ProjectCreateRequest) (*pb.Empty, error) {
	createReq := &RPCCreateProjectRequest{
		User_id:     req.UserId,
		Name:        req.Name,
		Deadline:    req.Deadline,
		Github_repo: req.GithubRepo,
	}
	response := &RPCMinimalProject{}
	s.pfc.call("CreateProject", createReq, response)

	// Notifier: send update projects list to user
	s.nc.SendUpdate(req.UserId, "project_add", response.Id, map[string]interface{}{
		"id":       response.Id,
		"name":     response.Name,
		"deadline": response.Deadline,
	})

	return &pb.Empty{}, nil
}

// proj_fac: GetFullProjectById
func (s *ProjectServiceServer) GetFullProjectById(ctx context.Context, req *pb.ProjectID) (*pb.Project, error) {
	getReq := &RPCGetProjectByIdRequest{
		Proj_id: req.ProjectId,
	}
	response := &RPCFullProject{}
	s.pfc.call("GetFullProjectById", getReq, response)

	return &pb.Project{
		Id:         response.Id,
		Name:       response.Name,
		Deadline:   response.Deadline,
		GithubRepo: response.GithubRepo,
		Users:      convertUsersToProto(response.Users),
		Milestones: convertMilestonesToProto(response.Milestones),
	}, nil
}

// proj_fac: GetProjectsFromUser
func (s *ProjectServiceServer) GetProjectsFromUser(ctx context.Context, req *pb.UserID) (*pb.ProjectsList, error) {
	getReq := &RPCGetProjectsFromUserRequest{
		User_id: req.UserId,
	}
	response := &RPCMinimalProjects{}
	s.pfc.call("GetProjectsFromUser", getReq, response)

	return &pb.ProjectsList{
		Projects: convertMinimalProjectsToProto(response.Projects),
	}, nil
}

// proj_fac: AddUserToProject
// notifier:
// - publish update users list
// - publish notification "new member"
// - send update projects list to friend
// - send notification to friend "you are added"
func (s *ProjectServiceServer) AddUserToProject(ctx context.Context, req *pb.AddUserToProjectRequest) (*pb.Empty, error) {
	addReq := &RPCAddUserToProjectRequest{
		Proj_id: req.ProjectId,
		User_id: req.UserId,
	}
	response := &RPCUserAddToProjectResponse{}
	s.pfc.call("AddUserToProject", addReq, response)

	s.nc.PublishUpdate("project", req.ProjectId, "user_add", req.UserId, map[string]interface{}{
		"id":         response.User.Id,
		"first_name": response.User.FirstName,
		"last_name":  response.User.LastName,
	})

	s.nc.PublishNotification("projects_list", req.ProjectId, response.User.FirstName+" was added to the project: "+response.Project.Name)
	s.nc.SendUpdate(req.UserId, "new_project", req.ProjectId, map[string]interface{}{
		"id":           response.Project.Id,
		"name":         response.Project.Name,
		"deadline":     response.Project.Deadline,
		"num_of_users": response.Project.NumOfUsers,
	})
	s.nc.SendNotification(req.UserId, "You have been added to a project.")
	s.nc.Subscribe(req.UserId, "projects_list", []string{req.ProjectId})

	return &pb.Empty{}, nil
}

// proj facacde: CreateMilestoneInProject
// notifier: publish update milestones list
func (s *ProjectServiceServer) CreateMilestoneInProject(ctx context.Context, req *pb.MilestoneCreateRequest) (*pb.Empty, error) {
	createReq := &RPCCreateMilestoneInProjectRequest{
		Proj_id:  req.ProjectId,
		Name:     req.Name,
		Deadline: req.Deadline,
	}
	response := &RPCMilestone{}
	s.pfc.call("CreateMilestoneInProject", createReq, response)

	// Notifier: publish update milestones list
	s.nc.PublishUpdate("project", req.ProjectId, "milestone_add", response.Id, map[string]interface{}{
		"id":                    response.Id,
		"name":                  response.Name,
		"deadline":              response.Deadline,
		"tasks":                 response.Tasks,
		"num_of_tasks":          response.NumOfTasks,
		"num_of_finished_tasks": response.NumOfFinishedTasks,
		"num_of_problems":       response.NumOfProblems,
	})

	return &pb.Empty{}, nil
}

// proj facade: CreateTaskInMilestone
// notifier: publish udpate tasks list
func (s *ProjectServiceServer) CreateTaskInMilestone(ctx context.Context, req *pb.TaskCreateRequest) (*pb.Empty, error) {
	createReq := &RPCCreateTaskInMilestoneRequest{
		Milestone_id: req.MilestoneId,
		Name:         req.Name,
	}
	response := &RPCTask{}
	s.pfc.call("CreateTaskInMilestone", createReq, response)

	// Notifier: publish update tasks list
	s.nc.PublishUpdate("project", req.ProjectId, "new_task_in_milestne", req.MilestoneId, map[string]interface{}{
		"id":              response.Id,
		"name":            response.Name,
		"status":          response.Status,
		"num_of_problems": response.NumOfProblems,
		"is_assigned":     response.IsAssigned,
	})

	return &pb.Empty{}, nil
}

// proj facade: AddProblemToTask
// notifier:
// - publish udpate problems list
// - publish notification "new problem"
func (s *ProjectServiceServer) AddProblemToTask(ctx context.Context, req *pb.ProblemAddRequest) (*pb.Empty, error) {
	addReq := &RPCAddProblemToTaskRequest{
		Task_id:      req.TaskId,
		Problem_name: req.Problem.Name,
	}
	response := &RPCProblem{}
	s.pfc.call("AddProblemToTask", addReq, response)

	// Notifier: publish update problems list
	s.nc.PublishUpdate("project", req.ProjectId, "new_problem_in_task", req.TaskId, map[string]interface{}{
		"id":        response.Id,
		"name":      response.Name,
		"posted_at": response.PostedAt,
	})

	// Notifier: publish notification "new problem"
	s.nc.PublishNotification("projects_list", req.ProjectId, "A new problem: "+response.Name+" was added to a task.")

	return &pb.Empty{}, nil
}

// proj facade: ResolveProblem
// notifier:
// - publish udpate problems list
// - publish notification "problem solved"
func (s *ProjectServiceServer) ResolveProblem(ctx context.Context, req *pb.ResolveProblemRequest) (*pb.Empty, error) {
	resolveReq := &RPCResolveProblemRequest{
		Task_id:    req.TaskId,
		Problem_id: req.ProblemId,
	}
	response := &RPCProblem{}
	s.pfc.call("ResolveProblem", resolveReq, response)

	// Notifier: publish update problems list
	s.nc.PublishUpdate("project", req.ProjectId, "problem_resolve_in_task", req.TaskId, map[string]interface{}{
		"problem_id": req.ProblemId,
	})

	// Notifier: publish notification "problem solved"
	s.nc.PublishNotification("projects_list", req.TaskId, "The problem: "+response.Name+"was resolved.")

	return &pb.Empty{}, nil
}

// proj facade: AssignTask
// notifier: publish udpate task
func (s *ProjectServiceServer) AssignTask(ctx context.Context, req *pb.TaskAssignRequest) (*pb.Empty, error) {
	assignReq := &RPCAssignTaskRequest{
		Task_id: req.TaskId,
		User_id: req.UserId,
	}
	response := &RPCTask{}
	s.pfc.call("AssignTask", assignReq, response)

	// Notifier: publish update task
	s.nc.PublishUpdate("project", req.ProjectId, "task_update", req.TaskId, map[string]interface{}{
		"status":              response.Status,
		"active_period_start": response.ActiveStartDate,
		"is_assigned":         response.IsAssigned,
		"user": map[string]interface{}{
			"id":         response.User.Id,
			"first_name": response.User.FirstName,
			"last_name":  response.User.LastName,
		},
	})

	return &pb.Empty{}, nil
}

// proj facade: CompleteTask
// notifier:
// - publish udpate task
// - publish notification "task completed"
func (s *ProjectServiceServer) CompleteTask(ctx context.Context, req *pb.TaskCompleteRequest) (*pb.Empty, error) {
	completeReq := &RPCCompleteTaskRequest{
		Task_id: req.TaskId,
	}
	response := &RPCTask{}
	s.pfc.call("CompleteTask", completeReq, response)

	// Notifier: publish update task
	s.nc.PublishUpdate("project", req.ProjectId, "task_update", req.TaskId, map[string]interface{}{
		"status":            response.Status,
		"active_period_end": response.ActiveStartDate,
	})

	// Notifier: publish notification "task completed"
	s.nc.PublishNotification("projects_list", req.ProjectId, "A task was completed.")

	return &pb.Empty{}, nil
}

func convertUsersToProto(users []RPCUser) []*pb.User {
	protoUsers := make([]*pb.User, len(users))
	for i, user := range users {
		protoUsers[i] = &pb.User{
			Id:        user.Id,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
	}
	return protoUsers
}

func convertMilestonesToProto(milestones []RPCMilestone) []*pb.Milestone {
	protoMilestones := make([]*pb.Milestone, len(milestones))
	for i, milestone := range milestones {
		protoMilestones[i] = &pb.Milestone{
			Id:                 milestone.Id,
			Name:               milestone.Name,
			Deadline:           milestone.Deadline,
			NumOfProblems:      int32(milestone.NumOfProblems),
			NumOfTasks:         int32(milestone.NumOfTasks),
			NumOfFinishedTasks: int32(milestone.NumOfFinishedTasks),
			Tasks:              convertTasksToProto(milestone.Tasks),
		}
	}
	return protoMilestones
}

func convertTasksToProto(tasks []RPCTask) []*pb.Task {

	protoTasks := make([]*pb.Task, len(tasks))
	for i, task := range tasks {
		var user *pb.User = nil
		if task.User != nil {
			user = &pb.User{
				Id:        task.User.Id,
				FirstName: task.User.FirstName,
				LastName:  task.User.LastName,
			}
		}
		protoTasks[i] = &pb.Task{
			Id:                task.Id,
			Name:              task.Name,
			Status:            task.Status,
			NumOfProblems:     int32(task.NumOfProblems),
			IsAssigned:        task.IsAssigned,
			User:              user,
			ActivePeriodStart: task.ActiveStartDate,
			ActivePeriodEnd:   task.ActiveEndDate,
			Problems:          convertProblemsToProto(task.Problems),
		}
	}
	return protoTasks
}

func convertProblemsToProto(problems []RPCProblem) []*pb.Problem {
	protoProblems := make([]*pb.Problem, len(problems))
	for i, problem := range problems {
		protoProblems[i] = &pb.Problem{
			Id:       &problem.Id,
			Name:     problem.Name,
			PostedAt: problem.PostedAt,
		}
	}
	return protoProblems
}

func convertMinimalProjectsToProto(projects []RPCMinimalProject) []*pb.Project {
	protoProjects := make([]*pb.Project, len(projects))
	for i, project := range projects {
		protoProjects[i] = &pb.Project{
			Id:       project.Id,
			Name:     project.Name,
			Deadline: project.Deadline,
		}
	}
	return protoProjects
}
