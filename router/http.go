package router

import (
	"fmt"
	"os"
	"strconv"


	"github.com/fajarhide/skeleton/middleware"
	userPresenter "github.com/fajarhide/skeleton/modules/user/presenter"
	authPresenter "github.com/fajarhide/skeleton/modules/auth/presenter"
	"github.com/labstack/echo"
	em "github.com/labstack/echo/middleware"
)

const DefaultPort = 8080

// HTTPServerMain - function for initializing main HTTP server
func (s *Service) HTTPServerMain() {

	e := echo.New()
	e.Use(middleware.Logger)
	e.Use(em.CORSWithConfig(em.CORSConfig{
		AllowMethods: []string{echo.OPTIONS, echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	if os.Getenv("DEBUG") == "1" {
		e.Debug = true
	}

	userHandler := userPresenter.NewUserHTTPHandler(s.UserUseCase)
	userGroup := e.Group("/user")
	userHandler.Mount(userGroup)

	authHandler := authPresenter.NewUserHTTPHandler(s.AuthUseCase)
	authGroup := e.Group("/auth")
	authHandler.Mount(authGroup)

	// set REST port
	var port uint16
	if portEnv, ok := os.LookupEnv("PORT"); ok {
		portInt, err := strconv.Atoi(portEnv)
		if err != nil {
			port = DefaultPort
		} else {
			port = uint16(portInt)
		}
	} else {
		port = DefaultPort
	}

	listenerPort := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(listenerPort))

}
