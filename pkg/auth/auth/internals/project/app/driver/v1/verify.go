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

func (au AuthPathway) VerifyUserAuth(ctx context.Context, req authTypes.VerifyUserAuthRequest) (authTypes.VerifyUserAuthResponse, error) {
	pbReq := &authPb.VerifyUserAuthRequest{}
	gconv.Struct(req, pbReq)
	resp, err := au.api.VerifyUserAuth(ctx, pbReq)
	if err != nil {
		err := fmt.Errorf("auth err: could not Register: %v", err)
		return authTypes.VerifyUserAuthResponse{}, errHandling.AuthenticationError(err)
	}
	apitools.Close(au.conn)
	domainResp := authTypes.VerifyUserAuthResponse{}
	gconv.Struct(resp, &domainResp)
	return domainResp, nil
}
