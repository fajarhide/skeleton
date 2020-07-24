package presenter

import (
	"net/http"

	"github.com/fajarhide/skeleton/helper"
	"github.com/fajarhide/skeleton/modules/user/sanitizer"
	"github.com/fajarhide/skeleton/modules/user/usecase"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// userHTTPHandler - auth http handler data
type userHTTPHandler struct {
	UserUseCase usecase.UserUseCase
}

// NewUserHTTPHandler - function for initializing user http handler
func NewUserHTTPHandler(userUseCase usecase.UserUseCase) *userHTTPHandler {
	return &userHTTPHandler{
		UserUseCase: userUseCase,
	}
}

// Mount - function for mounting route
func (h *userHTTPHandler) Mount(group *echo.Group) {
	group.GET("/:userID", h.GetProfile)
}

// GetProfile - function for GetProfile by ID
func (h *userHTTPHandler) GetProfile(c echo.Context) (err error) {
	var userID int

	ctx := "UserPresenter-GetProfile"

	if userID, err = sanitizer.GetProfile(c); err != nil {
		resp := helper.NewResponse(http.StatusBadRequest, err.Error(), nil)
		return resp.WriteResponse(c)
	}

	response, err := h.UserUseCase.GetProfile(cast.ToInt(userID))
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), ctx, "GetProfile")
		resp := helper.NewResponse(http.StatusBadRequest, err.Error(), nil)
		return resp.WriteResponse(c)
	}

	return c.JSON(http.StatusOK, response)
}
