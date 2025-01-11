package main

import (
	"context"
	pb "facade_service/protobuf_generated"
	"log"
)

type UserServiceServer struct {
	pb.UnimplementedUserServiceServer
	authClient     AuthClient
	userdbClient   UserDBClient
	notifierClient NotifierClient
}

func (s *UserServiceServer) LoginAndAuthenticate(ctx context.Context, req *pb.LoginRequest) (*pb.AuthResponse, error) {
	// login
	loginResp, ok := s.authClient.Login(req.Email, req.Password)
	if !ok {
		return &pb.AuthResponse{Valid: false}, nil
	}
	token := loginResp["token"].(string)

	// auth
	authResp, err := s.AuthenticateToken(ctx, &pb.TokenRequest{Token: token})
	if err != nil {
		return &pb.AuthResponse{Valid: false}, nil
	}
	return authResp, nil
}

func (s *UserServiceServer) AuthenticateToken(ctx context.Context, req *pb.TokenRequest) (*pb.AuthResponse, error) {
	// validate token
	resp, ok := s.authClient.ValidateToken(req.Token)
	if !ok {
		return &pb.AuthResponse{Valid: false}, nil
	}
	user_id := resp["user_id"].(string)
	// query user info
	resp, ok = s.userdbClient.QueryUser(user_id)
	if !ok {
		return &pb.AuthResponse{Valid: false}, nil
	}
	return &pb.AuthResponse{
		Valid: true,
		Token: req.Token,
		User: &pb.User{
			FirstName: resp["first_name"].(string),
			LastName:  resp["last_name"].(string),
			Id:        resp["id"].(string),
		},
	}, nil
}

func (s *UserServiceServer) CreateAccount(ctx context.Context, req *pb.RegisterRequest) (*pb.AuthResponse, error) {
	// register
	_, ok := s.authClient.Register(req.FirstName, req.LastName, req.Email, req.Password)
	if !ok {
		return &pb.AuthResponse{Valid: false}, nil
	}
	// login & auth
	authResp, err := s.LoginAndAuthenticate(ctx, &pb.LoginRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return &pb.AuthResponse{Valid: false}, nil
	}
	return authResp, nil
}

func (s *UserServiceServer) GetUserName(ctx context.Context, req *pb.UserID) (*pb.User, error) {
	// query user info
	log.Println(req.UserId)
	resp, ok := s.userdbClient.QueryUser(req.UserId)
	if !ok {
		return &pb.User{}, nil
	}
	return &pb.User{
		FirstName: resp["first_name"].(string),
		LastName:  resp["last_name"].(string),
		Id:        resp["id"].(string),
	}, nil
}

func (s *UserServiceServer) GetFriends(ctx context.Context, req *pb.UserID) (*pb.FriendsResponse, error) {
	resp, err := s.GetFriendsOrRequests(req, "friends")
	return resp, err
}

func (s *UserServiceServer) GetFriendRequests(ctx context.Context, req *pb.UserID) (*pb.FriendsResponse, error) {
	resp, err := s.GetFriendsOrRequests(req, "friendRequests")
	return resp, err
}

func (s *UserServiceServer) AcceptFriendRequest(ctx context.Context, req *pb.FriendEditRequest) (*pb.FriendsResponse, error) {
	// remove friend request bij user (resultaat bijhouden want dit is updated nieuwe friend requests lijst)
	resp, err := s.RemoveOrAddFriendsOrRequests(req.UserId, req.FriendId, "removeFriendRequest")
	// add friend bij user
	s.RemoveOrAddFriendsOrRequests(req.UserId, req.FriendId, "addFriend")
	// add user bij friend
	s.RemoveOrAddFriendsOrRequests(req.FriendId, req.UserId, "addFriend")

	user, ok := s.userdbClient.QueryUser(req.UserId)
	if !ok {
		return &pb.FriendsResponse{Users: []*pb.User{}}, nil
	}
	// stuur bericht
	message := user["first_name"].(string) + " is nu je vriend!"
	s.notifierClient.SendNotification(req.FriendId, message)
	// breng friend op de hoogte van nieuwe friends lijst
	s.notifierClient.SendUpdate(req.FriendId, "new_friend", req.UserId, user)
	// subscribe friend om updates te krijgen over jouw
	s.notifierClient.Subscribe(req.FriendId, "friends", []string{req.UserId})
	// subscribe jezelf om updates te krijgen over friend
	s.notifierClient.Subscribe(req.UserId, "friends", []string{req.FriendId})
	return resp, err
}
func (s *UserServiceServer) RemoveFriend(ctx context.Context, req *pb.FriendEditRequest) (*pb.FriendsResponse, error) {
	// remove user bij friend
	s.RemoveOrAddFriendsOrRequests(req.FriendId, req.UserId, "removeFriend")
	// remove friend bij user
	resp, err := s.RemoveOrAddFriendsOrRequests(req.UserId, req.FriendId, "removeFriend")

	user, ok := s.userdbClient.QueryUser(req.UserId)
	if !ok {
		return &pb.FriendsResponse{Users: []*pb.User{}}, nil
	}
	// breng friend op de hoogte van nieuwe friends lijst
	s.notifierClient.SendUpdate(req.FriendId, "removed_friend", req.UserId, user)
	// unsubscribe friend van updates over jouw
	s.notifierClient.UnSubscribe(req.FriendId, "friends", []string{req.UserId})
	// unsubscribe jezelf van updates over friend
	s.notifierClient.UnSubscribe(req.UserId, "friends", []string{req.FriendId})
	return resp, err
}
func (s *UserServiceServer) SendFriendRequest(ctx context.Context, req *pb.FriendEditRequest) (*pb.Empty, error) {
	// voeg bij friend een request toe met als friend user
	s.RemoveOrAddFriendsOrRequests(req.FriendId, req.UserId, "addFriendRequest")
	// haal naam van user
	user, ok := s.userdbClient.QueryUser(req.UserId)
	if !ok {
		return &pb.Empty{}, nil
	}
	// stuur bericht naar friend
	message := user["first_name"].(string) + " " + user["last_name"].(string) + " heeft je een vriendverzoek gestuurd!"
	s.notifierClient.SendNotification(req.FriendId, message)
	s.notifierClient.SendUpdate(req.FriendId, "new_friend_request", req.UserId, user)
	return &pb.Empty{}, nil
}
func (s *UserServiceServer) RejectFriendRequest(ctx context.Context, req *pb.FriendEditRequest) (*pb.FriendsResponse, error) {
	resp, err := s.RemoveOrAddFriendsOrRequests(req.UserId, req.FriendId, "removeFriendRequest")
	return resp, err
}

func (s *UserServiceServer) GetFriendsOrRequests(req *pb.UserID, graphqltype string) (*pb.FriendsResponse, error) {
	resp, ok := s.userdbClient.QueryFriendsOrRequests(req.UserId, graphqltype)
	if !ok {
		return &pb.FriendsResponse{Users: []*pb.User{}}, nil
	}
	var users []*pb.User
	for _, userMap := range resp {
		user := &pb.User{
			FirstName: userMap["first_name"].(string),
			LastName:  userMap["last_name"].(string),
			Id:        userMap["id"].(string),
		}
		users = append(users, user)
	}
	return &pb.FriendsResponse{
		Users: users,
	}, nil
}

func (s *UserServiceServer) RemoveOrAddFriendsOrRequests(userID string, friendID string, graphqltype string) (*pb.FriendsResponse, error) {
	resp, ok := s.userdbClient.RemoveOrAddFriendsOrRequests(userID, friendID, graphqltype)
	if !ok {
		return &pb.FriendsResponse{Users: []*pb.User{}}, nil
	}
	var users []*pb.User
	for _, userMap := range resp {
		user := &pb.User{
			FirstName: userMap["first_name"].(string),
			LastName:  userMap["last_name"].(string),
			Id:        userMap["id"].(string),
		}
		users = append(users, user)
	}
	return &pb.FriendsResponse{
		Users: users,
	}, nil
}
