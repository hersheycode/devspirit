package authdriver

import (
	"context"
	"fmt"

	errHandling "apppathway.com/pkg/errors"
	apitools "apppathway.com/pkg/net/grpc/tools"
	authPb "apppathway.com/pkg/user/auth/internals/project/api/v1"
	authTypes "apppathway.com/pkg/user/auth/internals/project/pkg/types"
	"github.com/gogf/gf/util/gconv"
)

func (au AuthPathway) LoginUser(ctx context.Context, req authTypes.LoginUserRequest) (authTypes.LoginUserResponse, error) {
	pbReq := &authPb.LoginUserRequest{}
	gconv.Struct(req, pbReq)
	resp, err := au.api.LoginUser(ctx, pbReq)
	if err != nil {
		err := fmt.Errorf("auth err: could not login: %v", err)
		return authTypes.LoginUserResponse{}, errHandling.AuthenticationError(err)
	}
	apitools.Close(au.conn)
	domainResp := authTypes.LoginUserResponse{}
	gconv.Struct(resp, &domainResp)
	apitools.SaveCreds(domainResp.User.ID, domainResp.AccessToken)
	return domainResp, nil
}
