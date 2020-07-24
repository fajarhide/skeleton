package presenter

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/fajarhide/skeleton/helper"
	"github.com/fajarhide/skeleton/modules/user/usecase"
	pb "github.com/fajarhide/skeleton/protogo/user"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// UserGRPCHandler data structure
type UserGRPCHandler struct {
	UserUseCase usecase.UserUseCase
}

// NewGRPCHandler function for initializing grpc handler object
func NewGRPCHandler(userUseCase usecase.UserUseCase) *UserGRPCHandler {
	return &UserGRPCHandler{
		UserUseCase: userUseCase,
	}
}

// GetProfile - function for getting detail user by id via grpc
func (h *UserGRPCHandler) GetProfile(c context.Context, arg *pb.UserQuery) (*pb.ResponseMessage, error) {
	ctx := "UserPresenter-GetProfile"

	userID := arg.ID
	user, err := h.UserUseCase.GetProfile(cast.ToInt(userID))
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), ctx, "get_detail_user")
		return nil, status.Error(codes.Internal, err.Error())
	}

	msg := pb.ResponseMessage{
		Message: "success to proceed data",
		Email:   user.Email,
		Name:    user.Name,
	}

	return &msg, nil
}
