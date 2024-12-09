package serverports

import (
	pb "apppathway.com/pkg/user/auth/internals/project/api/v1"
	domainTypes "apppathway.com/pkg/user/auth/internals/project/pkg/types"
	"apppathway.com/pkg/user/auth/internals/submod/app/server/domain"
	"fmt"
	"github.com/gogf/gf/util/gconv"
)

type Services domainTypes.DomainPorts
type ServerPorts struct{ repo domain.Services }

func New(repo domain.Services) ServerPorts {
	return ServerPorts{repo}
}
func (s ServerPorts) RegisterUser(req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	domainReq := &domainTypes.RegisterUserRequest{}
	gconv.Struct(req, domainReq)
	domainResp, err := s.repo.RegisterUser(domainReq)
	if err != nil {
		return nil, err
	}
	resp := &pb.RegisterUserResponse{}
	gconv.Struct(domainResp, resp)
	return resp, err
}
func (s ServerPorts) LoginUser(req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	domainReq := &domainTypes.LoginUserRequest{}
	gconv.Struct(req, domainReq)
	domainResp, err := s.repo.LoginUser(domainReq)
	if err != nil {
		return nil, err
	}
	domainUsr := domainResp.User
	user := &pb.User{ID: domainUsr.ID}
	resp := &pb.LoginUserResponse{}
	gconv.Struct(domainResp, resp)
	resp.User = user
	fmt.Printf("DOMAIN RESPONSE AFTER: %+v \n", resp)
	return resp, err
}
func (s ServerPorts) VerifyUserAuth(req *pb.VerifyUserAuthRequest) (*pb.VerifyUserAuthResponse, error) {
	domainReq := &domainTypes.VerifyUserAuthRequest{}
	gconv.Struct(req, domainReq)
	domainResp, err := s.repo.VerifyUserAuth(domainReq)
	if err != nil {
		return nil, err
	}
	resp := &pb.VerifyUserAuthResponse{}
	gconv.Struct(domainResp, resp)
	return resp, err
}
func (s ServerPorts) RefreshUserAuth(req *pb.RefreshUserAuthRequest) (*pb.RefreshUserAuthResponse, error) {
	domainReq := &domainTypes.RefreshUserAuthRequest{}
	gconv.Struct(req, domainReq)
	domainResp, err := s.repo.RefreshUserAuth(domainReq)
	if err != nil {
		return nil, err
	}
	resp := &pb.RefreshUserAuthResponse{}
	gconv.Struct(domainResp, resp)
	return resp, err
}
