package query

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/fajarhide/skeleton/helper"
	"github.com/fajarhide/skeleton/modules/user/model"
	log "github.com/sirupsen/logrus"
)

// userQueryMysql - user query implementation
// it is not exposed to outer package
type userQueryMysql struct {
	dbRead     *sqlx.DB
	dbWrite    *sqlx.DB
}

// NewUserQueryMysql - function for initializing user query
func NewUserQueryMysql(dbRead *sqlx.DB, dbWrite *sqlx.DB) *userQueryMysql {
	return &userQueryMysql{dbRead: dbRead, dbWrite: dbWrite}
}

// FindUserByID - function for getting detail user by ID
func (uq *userQueryMysql) FindUserByID(ID int) (*model.User, error) {
	ctx := "UserQuery-FindUserByID"

	// initiate variables
	var (
		user model.User
	)

	sq := `SELECT id, email, password, name FROM users WHERE id = ?`

	err := uq.dbRead.Get(&user, sq, ID)
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), ctx, "query_database")
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("data not found")
	}
	return &user, nil
}


// FindUserByEmail - function for getting detail user by Email
func (uq *userQueryMysql) FindUserByEmail(email string) (*model.User, error) {
	ctx := "UserQuery-FindUserByID"

	// initiate variables
	var (
		user model.User
	)

	sq := `SELECT id, email, password, name FROM users WHERE email = ?`

	err := uq.dbRead.Get(&user, sq, email)
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), ctx, "query_database")
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("data not found")
	}
	return &user, nil
}
