package main

import (
	"context"
	pb "gateway_service/protobuf_generated"
)

type NotificationServiceServer struct {
	pb.UnimplementedNotificationServiceServer
	notifierClient NotifierClient
	userService    *UserServiceServer
	projectService *ProjectServiceServer
}

func (s *NotificationServiceServer) SubscribeFriendList(ctx context.Context, req *pb.UserID) (*pb.Empty, error) {
	resp, _ := s.userService.GetFriendsOrRequests(req, "friends")
	var friend_ids []string
	for _, user := range resp.Users {
		friend_ids = append(friend_ids, user.Id)
	}
	s.notifierClient.Subscribe(req.UserId, "friends", friend_ids)
	return &pb.Empty{}, nil
}

func (s *NotificationServiceServer) SubscribeProjectsList(ctx context.Context, req *pb.UserID) (*pb.Empty, error) {
	resp, _ := s.projectService.GetProjectsFromUser(ctx, req)
	var project_ids []string
	for _, project := range resp.Projects {
		project_ids = append(project_ids, project.Id)
	}
	s.notifierClient.Subscribe(req.UserId, "projects_list", project_ids)
	return &pb.Empty{}, nil
}

func (s *NotificationServiceServer) SwitchProjectSubscription(ctx context.Context, req *pb.ProjectSubscribeRequest) (*pb.Empty, error) {
	if req.UnsubscribeProject != nil {
		s.notifierClient.UnSubscribe(req.UserId, "project", []string{*req.UnsubscribeProject})
	}
	s.notifierClient.Subscribe(req.UserId, "project", []string{req.SubscribeProject})
	return &pb.Empty{}, nil
}
