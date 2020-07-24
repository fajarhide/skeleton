package token

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/fajarhide/skeleton/helper"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

//BearerClaims data structure
type BearerClaims struct {
	UserAuthorized bool `json:"authorised,bool"`
	jwt.StandardClaims
}

// VerifyTokenIgnoreExpiration function for verifying token and ignore expiration
func VerifyTokenIgnoreExpiration(rsaPublicKey *rsa.PublicKey, oldAccessToken string) (*BearerClaims, error) {
	ctx := "Token-VerifyTokenIgnoreExpiration"

	var errorStr error

	token, err := jwt.ParseWithClaims(oldAccessToken, &BearerClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			errorStr = fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			helper.Log(log.ErrorLevel, errorStr.Error(), ctx, "parse_token")
			return nil, errorStr
		}
		return rsaPublicKey, nil
	})

	if claims, ok := token.Claims.(*BearerClaims); ok {
		return claims, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {

		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			errorStr = fmt.Errorf("invalid token format: %s", oldAccessToken)
		} else {
			errorStr = fmt.Errorf("token parsing error: %s", err.Error())
		}

		helper.Log(log.ErrorLevel, err.Error(), ctx, "validate_token")
		return nil, errorStr
	}

	errorStr = errors.New("unknown errors")
	helper.Log(log.ErrorLevel, errorStr.Error(), ctx, "validate_token")
	return nil, errorStr
}
