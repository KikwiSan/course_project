<<<<<<< HEAD
package main

import (
	"context"

	pb "github.com/KikwiSan/course_project/users"
	"github.com/iqubb/src/users/api"
	"golang.org/x/crypto/bcrypt"
)

type UserServer struct {
	repository api.Repository
}

func (s *UserServer) GetUser(ctx context.Context, usReq *pb.Request) (*pb.Response, error) {
	result, _ := s.repository.GetUser(ctx, usReq.Id)
	user := api.UnconvertingUser(result)
	return &pb.Response{
		User: user,
	}, nil
}

func (s *UserServer) CreateUser(ctx context.Context, newReq *pb.User) (*pb.Response, error) {
	pass, _ := bcrypt.GenerateFromPassword([]byte(newReq.Password), bcrypt.DefaultCost)
	newReq.Password = string(pass)
	_ = s.repository.CreateUser(ctx, api.ConvertingUser(newReq))
	return &pb.Response{
		User: newReq,
	}, nil
}

func (s *UserServer) GetAllUsers(ctx context.Context, allReq *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	results, _ := s.repository.GetAllUsers(ctx)
	users := api.UnconvertingAllUsers(results)
	return &pb.GetAllUsersResponse{
		AllUsers: users,
	}, nil
}
=======
package main

import (
	"context"

	pb "github.com/KikwiSan/course_project/users"
	"github.com/iqubb/src/users/api"
	"golang.org/x/crypto/bcrypt"
)

type UserServer struct {
	repository api.Repository
}

func (s *UserServer) GetUser(ctx context.Context, usReq *pb.Request) (*pb.Response, error) {
	result, _ := s.repository.GetUser(ctx, usReq.Id)
	user := api.UnconvertingUser(result)
	return &pb.Response{
		User: user,
	}, nil
}

func (s *UserServer) CreateUser(ctx context.Context, newReq *pb.User) (*pb.Response, error) {
	pass, _ := bcrypt.GenerateFromPassword([]byte(newReq.Password), bcrypt.DefaultCost)
	newReq.Password = string(pass)
	_ = s.repository.CreateUser(ctx, api.ConvertingUser(newReq))
	return &pb.Response{
		User: newReq,
	}, nil
}

func (s *UserServer) GetAllUsers(ctx context.Context, allReq *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	results, _ := s.repository.GetAllUsers(ctx)
	users := api.UnconvertingAllUsers(results)
	return &pb.GetAllUsersResponse{
		AllUsers: users,
	}, nil
}
>>>>>>> 9006d60a1538c3867b12bcf075dde1053ad2a85e
