package presenter

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/cast"
	"net/http"

	"github.com/labstack/echo"
	"github.com/fajarhide/skeleton/helper"
	"github.com/fajarhide/skeleton/modules/auth/sanitizer"
	"github.com/fajarhide/skeleton/modules/auth/usecase"
	log "github.com/sirupsen/logrus"
	"github.com/fajarhide/skeleton/keys"
	"github.com/fajarhide/skeleton/middleware"
)

// userHTTPHandler - auth http handler data
type userHTTPHandler struct {
	AuthUseCase usecase.AuthUseCase
}

// NewUserHTTPHandler - function for initializing user http handler
func NewUserHTTPHandler(authUseCase usecase.AuthUseCase) *userHTTPHandler {
	return &userHTTPHandler{
		AuthUseCase: authUseCase,
	}
}

// Mount - function for mounting route
func (h *userHTTPHandler) Mount(group *echo.Group) {
	publicKey, _ := keys.InitPublicKey()
	group.POST("/login", h.Login)
	group.GET("/me", h.GetProfile, middleware.BearerVerify(publicKey, true))
}

// Login - function for Login
func (h *userHTTPHandler) Login(c echo.Context) (err error) {
	ctx := "UserPresenter-GetProfile"
	request, err := sanitizer.Login(c);
	if  err != nil {
		resp := helper.NewResponse(http.StatusBadRequest, err.Error(), nil)
		return resp.WriteResponse(c)
	}

	response, err := h.AuthUseCase.Login(request.Email, request.Password)
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), ctx, "GetProfile")
		resp := helper.NewResponse(http.StatusBadRequest, err.Error(), nil)
		return resp.WriteResponse(c)
	}

	return c.JSON(http.StatusOK, response)
}



// Login - function for Login
func (h *userHTTPHandler) GetProfile(c echo.Context) (err error) {
	ctx := "UserPresenter-GetProfile"
	token, _ := c.Get("token").(*jwt.Token)
	claims, _ := token.Claims.(*middleware.BearerClaims)
	userID := cast.ToInt(claims.Audience)
	response, err := h.AuthUseCase.GetProfile(userID)
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), ctx, "GetProfile")
		resp := helper.NewResponse(http.StatusBadRequest, err.Error(), nil)
		return resp.WriteResponse(c)
	}
	return c.JSON(http.StatusOK, response)
}
