package query

import (
	"github.com/fajarhide/skeleton/modules/user/model"
)

// UserQuery - user query interface abstraction
type UserQuery interface {
	FindUserByID(ID int) (*model.User, error)
	FindUserByEmail(email string) (*model.User, error)
}
