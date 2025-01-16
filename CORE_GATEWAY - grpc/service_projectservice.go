package main

import (
	"context"
	pb "gateway_service/protobuf_generated"
)

type ProjectServiceServer struct {
	pb.UnimplementedProjectServiceServer
	projectFacadeClient ProjectFacadeClient
	notifierClient      NotifierClient
}

func (s *ProjectServiceServer) InitialiseProject(ctx context.Context, req *pb.Project) (*pb.Empty, error) {
	// TODO: Call projectFacadeClient to create the project
	// TODO: Notify using notifierClient
	return &pb.Empty{}, nil
}

func (s *ProjectServiceServer) GetProjectById(ctx context.Context, req *pb.ProjectID) (*pb.Project, error) {
	// TODO: Call projectFacadeClient to fetch the project by ID
	return &pb.Project{}, nil
}

func (s *ProjectServiceServer) AddUserToProject(ctx context.Context, req *pb.AddUserToProjectRequest) (*pb.Empty, error) {
	// TODO: Call projectFacadeClient to add the user to the project
	// TODO: Notify using notifierClient (e.g., publish updates, send notifications)
	return &pb.Empty{}, nil
}

func (s *ProjectServiceServer) GetProjectsFromUser(ctx context.Context, req *pb.UserID) (*pb.ProjectsList, error) {
	// TODO: Call projectFacadeClient to fetch projects for the user
	return &pb.ProjectsList{}, nil
}

func (s *ProjectServiceServer) CreateMilestoneInProject(ctx context.Context, req *pb.MilestoneAddRequest) (*pb.Empty, error) {
	// TODO: Call projectFacadeClient to create the milestone
	// TODO: Notify using notifierClient (e.g., publish milestone updates)
	return &pb.Empty{}, nil
}

func (s *ProjectServiceServer) GetMilestoneById(ctx context.Context, req *pb.MilestoneID) (*pb.Milestone, error) {
	// TODO: Call projectFacadeClient to fetch the milestone by ID
	return &pb.Milestone{}, nil
}

func (s *ProjectServiceServer) CreateTaskInMilestone(ctx context.Context, req *pb.TaskAddRequest) (*pb.Empty, error) {
	// TODO: Call projectFacadeClient to create the task
	// TODO: Notify using notifierClient (e.g., publish task updates)
	return &pb.Empty{}, nil
}

func (s *ProjectServiceServer) GetTaskById(ctx context.Context, req *pb.TaskID) (*pb.Task, error) {
	// TODO: Call projectFacadeClient to fetch the task by ID
	return &pb.Task{}, nil
}

func (s *ProjectServiceServer) AddProblemToTask(ctx context.Context, req *pb.ProblemAddRequest) (*pb.Empty, error) {
	// TODO: Call projectFacadeClient to add the problem to the task
	// TODO: Notify using notifierClient (e.g., publish problem updates, send notifications)
	return &pb.Empty{}, nil
}

func (s *ProjectServiceServer) ResolveProblem(ctx context.Context, req *pb.ProblemID) (*pb.Empty, error) {
	// TODO: Call projectFacadeClient to resolve the problem
	// TODO: Notify using notifierClient (e.g., publish problem updates, send notifications)
	return &pb.Empty{}, nil
}

func (s *ProjectServiceServer) AssignTask(ctx context.Context, req *pb.TaskAssignRequest) (*pb.Empty, error) {
	// TODO: Call projectFacadeClient to assign the task
	// TODO: Notify using notifierClient (e.g., publish task updates)
	return &pb.Empty{}, nil
}

func (s *ProjectServiceServer) CompleteTask(ctx context.Context, req *pb.TaskID) (*pb.Empty, error) {
	// TODO: Call projectFacadeClient to mark the task as complete
	// TODO: Notify using notifierClient (e.g., publish task updates)
	return &pb.Empty{}, nil
}
