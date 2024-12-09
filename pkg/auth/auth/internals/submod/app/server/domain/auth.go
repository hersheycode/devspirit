package domain

import (
	"errors"
	"fmt"
	"time"

	errHandling "apppathway.com/pkg/errors"
	dt "apppathway.com/pkg/user/auth/internals/project/pkg/types"
	"github.com/dgrijalva/jwt-go"
)

const HMAC_SAMPLE_SECRET = "hmacSampleSecret"
const ACCESS_TOKEN_DURATION = time.Hour
const REFRESH_TOKEN_DURATION = time.Hour * 24 * 30

type User dt.User
type refreshTokenClaims struct{ jwt.StandardClaims }
type accessTokenClaims struct{ jwt.StandardClaims }
type authToken struct{ *jwt.Token }

func (u User) claimsForAccessToken() (accessTokenClaims, error) {
	if u.Role == "free-tier" {
		return u.claimsForFreeTierUser(), nil
	}
	fmt.Printf("[DEBUG] Invalid Role '%v' or Not Found \n", u.Role)
	err := fmt.Errorf("invalid credentials: agent name is not registered and/or invalid api key and/or role")
	return accessTokenClaims{}, errHandling.ConflictError(err)
}
func (u User) claimsForFreeTierUser() accessTokenClaims {
	return accessTokenClaims{StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix()}}
}
func newAuthToken(claims accessTokenClaims) authToken {
	return authToken{Token: jwt.NewWithClaims(jwt.SigningMethodHS256, claims)}
}
func (t authToken) newAccessToken() (string, error) {
	signedString, err := t.Token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		fmt.Printf("Failed while signing access token: %v \n", err)
		err := fmt.Errorf("cannot generate access token \n")
		return "", errHandling.UnexpectedError(err)
	}
	return signedString, nil
}
func (t authToken) newRefreshToken() (string, error) {
	c := t.Token.Claims.(accessTokenClaims)
	refreshClaims := c.newClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedString, err := token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		fmt.Printf("[DEBUG] Failed while signing refresh token: %v", err.Error())
		err = fmt.Errorf("cannot generate refresh token \n")
		return "", errHandling.UnexpectedError(err)
	}
	return signedString, nil
}
func (c accessTokenClaims) newClaims() refreshTokenClaims {
	return refreshTokenClaims{StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(REFRESH_TOKEN_DURATION).Unix()}}
}
func (c refreshTokenClaims) newClaims() accessTokenClaims {
	return accessTokenClaims{StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(ACCESS_TOKEN_DURATION).Unix()}}
}
func jwtTokenFromString(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &accessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		err = fmt.Errorf("failed while parsing token: %v", err.Error())
		return nil, errHandling.UnexpectedError(err)
	}
	return token, nil
}
func isAccessTokenValid(toBeValidated string) *jwt.ValidationError {
	_, err := jwt.Parse(toBeValidated, func(token *jwt.Token) (interface{}, error) {
		return []byte(HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		var vErr *jwt.ValidationError
		if errors.As(err, &vErr) {
			return vErr
		}
	}
	return nil
}
func newAccessTokenFromRefreshToken(refreshToken string) (string, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &refreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		err := fmt.Errorf("invalid or expired refresh token")
		return "", errHandling.AuthenticationError(err)
	}
	r := token.Claims.(*refreshTokenClaims)
	accessTokenClaims := r.newClaims()
	authToken := newAuthToken(accessTokenClaims)
	return authToken.newAccessToken()
}
