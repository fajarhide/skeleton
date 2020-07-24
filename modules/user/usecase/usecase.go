package usecase

import "github.com/fajarhide/skeleton/modules/user/model"

// UserUseCase - user use case interface abstraction
type UserUseCase interface {
	GetProfile(id int) (*model.User, error)
}
