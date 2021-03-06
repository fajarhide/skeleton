package middleware

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

// BearerClaims data structure for claims
type BearerClaims struct {
	UserAuthorized bool   `json:"authorised,bool"`
	RToken         string `json:"rtoken"`
	Type           int    `json:"type"`
	jwt.StandardClaims
}

// BearerVerify function to verify token
func BearerVerify(rsaPublicKey *rsa.PublicKey, mustAuthorized bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if os.Getenv("NO_TOKEN") == "1" {
				return next(c)
			}

			req := c.Request()
			header := req.Header
			auth := header.Get("Authorization")

			if len(auth) <= 0 {
				return echo.NewHTTPError(http.StatusUnauthorized, "authorization is empty")
			}

			splitToken := strings.Split(auth, " ")
			if len(splitToken) < 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "authorization is empty")
			}

			if splitToken[0] != "Bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "authorization is invalid")
			}

			tokenStr := splitToken[1]
			token, err := jwt.ParseWithClaims(tokenStr, &BearerClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return rsaPublicKey, nil
			})

			refreshToken := c.QueryParam("refreshtoken")
			if err != nil && refreshToken == "" && !strings.Contains(c.Path(), "refreshtoken") {
				if strings.Contains(err.Error(), "token") && strings.Contains(err.Error(), "expired") {
					err = errors.New("Your session has expired. Please login again")
				}
				return echo.NewHTTPError(http.StatusExpectationFailed, err.Error())
			}

			if claims, ok := token.Claims.(*BearerClaims); err == nil && token.Valid && ok {
				if mustAuthorized {
					if claims.UserAuthorized {
						c.Set("token", token)
						return next(c)
					}
					fmt.Printf("%+v", claims)
					return echo.NewHTTPError(http.StatusUnauthorized, "Resource need an authorised user")
				}
				c.Set("token", token)
				return next(c)
			} else if ve, ok := err.(*jwt.ValidationError); ok {
				if refreshToken != "" && strings.Contains(c.Path(), "refreshtoken") {
					c.Set("token", token)
					return next(c)
				}
				var errorStr string
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					errorStr = fmt.Sprintf("Invalid token format: %s", tokenStr)
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					errorStr = "Token has been expired"
				} else {
					errorStr = fmt.Sprintf("Token Parsing Error: %s", err.Error())
				}
				return echo.NewHTTPError(http.StatusUnauthorized, errorStr)
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unknown token error")
			}
		}
	}
}
