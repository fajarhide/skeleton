package router

import (
	"github.com/fajarhide/skeleton/config"
	"os"
	"time"

	authUseCase "github.com/fajarhide/skeleton/modules/auth/usecase"
	authToken "github.com/fajarhide/skeleton/modules/auth/token"
	userQuery "github.com/fajarhide/skeleton/modules/user/query"
	userUseCase "github.com/fajarhide/skeleton/modules/user/usecase"
	"github.com/fajarhide/skeleton/helper"
	"github.com/fajarhide/skeleton/keys"
	log "github.com/sirupsen/logrus"
)

// Service handler data structure
type Service struct {
	UserUseCase userUseCase.UserUseCase
	AuthUseCase authUseCase.AuthUseCase
}

// MakeHandler - function for initializing handler of the services
func MakeHandler() *Service {
	// initiate database connection here
	readDB := config.ReadMysqlDB()
	writeDB := config.WriteMysqlDB()

	privateKey, err := keys.InitPrivateKey()
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), "http_handler", "private_key")
		os.Exit(1)
	}

	tokenAge, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_AGE"))
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), "http_handler", "parse_token_age")
		os.Exit(1)
	}
	refreshTokenAge, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_AGE"))
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), "http_handler", "parse_refresh_token_age")
		os.Exit(1)
	}

	accessTokenGenerator := authToken.NewJwtGenerator(privateKey, tokenAge, refreshTokenAge)

	userQuery := userQuery.NewUserQueryMysql(readDB, writeDB)
	uUseCase := userUseCase.NewUserUseCase(userQuery)
	aUseCase := authUseCase.NewAuthUseCase(userQuery, accessTokenGenerator)

	return &Service{
		UserUseCase: uUseCase,
		AuthUseCase:aUseCase,
	}
}
