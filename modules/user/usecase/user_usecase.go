package usecase

import (
	"github.com/fajarhide/skeleton/helper"
	"github.com/fajarhide/skeleton/modules/user/model"
	"github.com/fajarhide/skeleton/modules/user/query"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
)

// userUseCaseImpl - user use case implementation
type userUseCaseImpl struct {
	userQuery query.UserQuery
}

// NewUserUseCase - function for initializing user use case implementation
func NewUserUseCase(userQuery query.UserQuery) UserUseCase {
	return &userUseCaseImpl{
		userQuery: userQuery,
	}
}

// GetDetailUser - function for getting detail user
func (uu *userUseCaseImpl) GetProfile(id int) (*model.User, error) {
	ctx := "UserUseCase-GetProfile"

	user, err := uu.userQuery.FindUserByID(cast.ToInt(id))
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), ctx, "find_detail_user")
		return nil, err
	}

	return user, nil
}
