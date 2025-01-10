package main

import (
	"context"
	pb "facade_service/protobuf_generated"
)

type NotificationServiceServer struct {
	pb.UnimplementedNotificationServiceServer
	notifierClient NotifierClient
	userService    *UserServiceServer
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
