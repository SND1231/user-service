package main

import (
	"context"
	"log"
	"net"

	pb "github.com/SND1231/user-service/proto"
	user_app_service "github.com/SND1231/user-service/user_app_service"
	"google.golang.org/grpc"
)

const (
	port = ":9001"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedUserServiceServer
}

// GET User
func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var id = in.Id
	var user, err = user_app_service.GetUser(id)
	return &pb.GetUserResponse{User: &user}, err
}

// Login
func (s *server) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	var request = *in
	var id, token, err = user_app_service.LoginUser(request)
	if err == nil {
		return &pb.LoginResponse{Id: id, Token: token}, nil
	} else {
		return &pb.LoginResponse{}, err
	}
}

// GET Users
func (s *server) GetUsers(ctx context.Context, in *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	var users, err = user_app_service.GetUsers(*in)
	return &pb.GetUsersResponse{Users: users}, err
}

// Create User
func (s *server) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	var request = *in
	var id, token, err = user_app_service.CreateUser(request)
	if err == nil {
		return &pb.CreateUserResponse{Id: id, Token: token}, nil
	} else {
		return &pb.CreateUserResponse{}, err
	}
}

// Update User
func (s *server) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	var request = *in
	var id, err = user_app_service.UpdateUser(request)
	if err == nil {
		return &pb.UpdateUserResponse{Id: id}, nil
	} else {
		return &pb.UpdateUserResponse{}, err
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
