package usecase

import (
	"github.com/fajarhide/skeleton/modules/auth/model"
	userModel "github.com/fajarhide/skeleton/modules/user/model"
)

// AuthUseCase - user use case interface abstraction
type AuthUseCase interface {
	Login(enail, password string) (*model.LoginResponse, error)
	GetProfile(ID int) (*userModel.User, error)
}
