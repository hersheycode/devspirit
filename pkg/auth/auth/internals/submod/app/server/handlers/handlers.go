package handlers

import (
	pb "apppathway.com/pkg/user/auth/internals/project/api/v1"
	"context"
	"fmt"
	"log"
)

func (s *server) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	fmt.Printf("1) RegisterUser: %+v\n", req)
	var resp *pb.RegisterUserResponse
	var err error
	u := req.GetUser()
	if u.GetEmail() != "" && u.GetPassword() != "" {
		resp, err = s.Services.RegisterUser(req)
		if err != nil {
			log.Println(err)
			return &pb.RegisterUserResponse{}, err
		}
		fmt.Printf("2) RegisterUser: %+v\n", req)
	}
	return resp, nil
}
func (s *server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	fmt.Printf("1) app req: %+v\n", req)
	var resp *pb.LoginUserResponse
	var err error
	u := req.GetUser()
	if u.GetEmail() != "" && u.GetPassword() != "" {
		resp, err = s.Services.LoginUser(req)
		if err != nil {
			log.Println(err)
			return &pb.LoginUserResponse{}, err
		}
		fmt.Printf("2) app req: %+v\n", req)
	} else {
		log.Println("email or password was empty...")
		return &pb.LoginUserResponse{}, err
	}
	return resp, nil
}
func (s *server) RefreshUserAuth(ctx context.Context, req *pb.RefreshUserAuthRequest) (*pb.RefreshUserAuthResponse, error) {
	fmt.Printf("1) app req: %+v\n", req)
	var resp *pb.RefreshUserAuthResponse
	var err error
	if req.GetAccessToken() != "" && req.GetRefreshToken() != "" {
		resp, err = s.Services.RefreshUserAuth(req)
		if err != nil {
			fmt.Println(err)
			return &pb.RefreshUserAuthResponse{}, err
		}
		fmt.Printf("2) app req: %+v\n", req)
	} else {
		err := fmt.Errorf("access or refresh token empty...")
		return &pb.RefreshUserAuthResponse{}, err
	}
	return resp, nil
}
func (s *server) VerifyUserAuth(ctx context.Context, req *pb.VerifyUserAuthRequest) (*pb.VerifyUserAuthResponse, error) {
	var resp *pb.VerifyUserAuthResponse
	var err error
	if req.GetAuthToken() != "" {
		resp, err = s.Services.VerifyUserAuth(req)
		if err != nil {
			return &pb.VerifyUserAuthResponse{}, err
		}
	} else {
		err := fmt.Errorf("auth token empty...")
		return &pb.VerifyUserAuthResponse{}, err
	}
	fmt.Printf("Verify Response: %+v \n", resp)
	return resp, nil
}
