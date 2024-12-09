package domain

import (
	dbTypes "apppathway.com/pkg/db_api/internals/project/pkg/types"
	errHandling "apppathway.com/pkg/errors"
	dt "apppathway.com/pkg/user/auth/internals/project/pkg/types"
	"fmt"
	"log"
)

func (dl Logic) RegisterUser(req *dt.RegisterUserRequest) (*dt.RegisterUserResponse, error) {
	vars := map[string]string{"$attr": "email", "$value": req.User.Email}
	registered, err := dl.store.Equal(vars, []string{"uid"})
	if err != nil {
		return nil, err
	}
	if registered {
		err := fmt.Errorf(": email already registered")
		return nil, errHandling.ConflictError(err)
	}
	req.User.Role = "free-tier"
	results := &[]dt.AppCredentials{}
	spec := dbTypes.MutationSpec{Data: req.User, DBReturnSpec: dbTypes.DBReturnSpec{DType: "User", StructSlice: results, Fields: []string{"uid"}}}
	err = dl.store.CommitMutation(spec)
	if err != nil {
		return nil, err
	}
	dbCreds := *results
	if len(dbCreds) != 1 {
		err := fmt.Errorf("len(dbCreds) != 1; %+v", dbCreds)
		return nil, errHandling.FatalDatabaseErr(err)
	}
	creds := dbCreds[0]
	claims, err := User(*req.User).claimsForAccessToken()
	if err != nil {
		return nil, err
	}
	t := newAuthToken(claims)
	if creds.AccessToken, err = t.newAccessToken(); err != nil {
		return nil, err
	}
	creds.RefreshToken, err = t.newRefreshToken()
	if err != nil {
		return nil, err
	}
	spec.Data = struct {
		Token string `json:"refreshToken"`
	}{Token: creds.RefreshToken}
	spec.DType = "RefreshTokenStore"
	err = dl.store.CommitMutation(spec)
	fmt.Printf("creds: %+v \n", creds)
	createRepoAccount(*req.User)
	return nil, err
}
func (dl Logic) LoginUser(req *dt.LoginUserRequest) (*dt.LoginUserResponse, error) {
	vars := map[string]string{"$attr": "email", "$value": req.User.Email}
	returnFields := []string{"uid", "email", "role"}
	uid, valid, err := dl.store.ValidCredentials(vars, returnFields, req.User.Password)
	if err != nil {
		log.Printf("%v\n", err)
		err = fmt.Errorf("invalid credentials.")
		return nil, errHandling.AuthenticationError(err)
	}
	if !valid {
		err := fmt.Errorf("invalid credentials.")
		return nil, errHandling.AuthenticationError(err)
	}
	req.User.Role = "free-tier"
	claims, err := User(*req.User).claimsForAccessToken()
	if err != nil {
		return nil, err
	}
	t := newAuthToken(claims)
	accessToken, err := t.newAccessToken()
	if err != nil {
		return nil, err
	}
	refreshToken, err := t.newRefreshToken()
	if err != nil {
		return nil, err
	}
	spec := dbTypes.MutationSpec{}
	spec.Data = struct {
		Token string `json:"refreshToken"`
	}{Token: refreshToken}
	spec.DType = "RefreshTokenStore"
	spec.Fields = []string{"uid"}
	fmt.Println("USER: ", uid)
	resp := &dt.LoginUserResponse{User: &dt.User{ID: uid}, AccessToken: accessToken, RefreshToken: refreshToken}
	req.User.ID = uid
	return resp, dl.store.CommitMutation(spec)
}
func (dl Logic) VerifyUserAuth(req *dt.VerifyUserAuthRequest) (*dt.VerifyUserAuthResponse, error) {
	var err error
	fmt.Println("verify: ", req.AuthToken)
	if jwtToken, jwtErr := jwtTokenFromString(req.AuthToken); jwtErr != nil {
		err = errHandling.AuthorizationError(jwtErr)
		fmt.Println("verify err 1: ", err)
	} else {
		if jwtToken.Valid {
		} else {
			err = errHandling.AuthenticationError(fmt.Errorf("invalid token"))
		}
	}
	resp := &dt.VerifyUserAuthResponse{}
	if err == nil {
		resp.IsAuthorized = true
	}
	return resp, err
}
func (dl Logic) RefreshUserAuth(req *dt.RefreshUserAuthRequest) (*dt.RefreshUserAuthResponse, error) {
	resp := dt.RefreshUserAuthResponse{}
	return &resp, nil
}
