package types

import (
	errHandling "apppathway.com/pkg/errors"
	pb "apppathway.com/pkg/user/auth/internals/project/api/v1"
)

type User struct {
	ID        string `json:"uid"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	PhoneNum  string `json:"phoneNum"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	ProfileId string `json:"profileId"`
	AccountId string `json:"accountId"`
}
type RegisterUserRequest struct{ User *User }
type RegisterUserResponse struct{ User *User }
type LoginUserRequest struct{ User *User }
type AppCredentials struct {
	UID          string `json:"uid"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
type LoginUserResponse struct {
	User         *User
	AccessToken  string
	RefreshToken string
}
type RefreshUserAuthRequest struct {
	AccessToken  string
	RefreshToken string
}
type RefreshUserAuthResponse struct{ AuthToken string }
type VerifyUserAuthRequest struct{ AuthToken string }
type VerifyUserAuthResponse struct{ IsAuthorized bool }
type IsAuthorizedRequest struct {
	Token     string
	RouteName string
	Vars      map[string]string
}
type IsAuthorizedResponse struct {
	Status bool `json:"status"`
}
type UserService interface {
	RegisterUser(*RegisterUserRequest) (*RegisterUserResponse, error)
	LoginUser(*LoginUserRequest) (*LoginUserResponse, error)
	RefreshUserAuth(*RefreshUserAuthRequest) (*RefreshUserAuthResponse, error)
	VerifyUserAuth(*VerifyUserAuthRequest) (*VerifyUserAuthResponse, error)
}
type UserClientService interface{}
type UserGRPCPort interface {
	RegisterUser(*pb.RegisterUserRequest) (*pb.RegisterUserResponse, error)
	LoginUser(*pb.LoginUserRequest) (*pb.LoginUserResponse, error)
	VerifyUserAuth(*pb.VerifyUserAuthRequest) (*pb.VerifyUserAuthResponse, error)
	RefreshUserAuth(*pb.RefreshUserAuthRequest) (*pb.RefreshUserAuthResponse, error)
}
type UserRestPort interface {
	RegisterUser(RegisterUserRequest) (*RegisterUserResponse, *errHandling.AppError)
	LoginUser(LoginUserRequest) (*LoginUserResponse, *errHandling.AppError)
	RefreshUserAuth(RefreshUserAuthRequest) (*RefreshUserAuthResponse, *errHandling.AppError)
	VerifyUserAuth(VerifyUserAuthRequest) (*VerifyUserAuthResponse, *errHandling.AppError)
}
type UserClientGRPCPort interface{}
type DomainServices interface{ UserService }
type ClientServices interface{ UserClientService }
type MiddlewareServices interface{ AuthMiddlewareService }
type DomainPorts interface{ ServerGRPCPorts }
type ClientPorts interface{ ClientGRPCPorts }
type ServerGRPCPorts interface{ UserGRPCPort }
type ServerRestPorts interface{ UserRestPort }
type ClientGRPCPorts interface{ UserClientGRPCPort }
type AuthMiddlewareService interface {
	IsAuthorized(*IsAuthorizedRequest) (*IsAuthorizedResponse, error)
}
