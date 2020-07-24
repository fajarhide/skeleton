package usecase

import (
	"github.com/fajarhide/skeleton/modules/auth/token"
	"github.com/fajarhide/skeleton/helper"
	"github.com/fajarhide/skeleton/modules/auth/model"
	userModel "github.com/fajarhide/skeleton/modules/user/model"
	"github.com/fajarhide/skeleton/modules/user/query"
	log "github.com/sirupsen/logrus"
	"strconv"

)

// userUseCaseImpl - user use case implementation
type userUseCaseImpl struct {
	userQuery      query.UserQuery
	tokenGenerator token.AccessTokenGenerator
}

// NewAuthUseCase - function for initializing user use case implementation
func NewAuthUseCase(
	userQuery query.UserQuery,
	tokenGenerator token.AccessTokenGenerator,
) AuthUseCase {
	return &userUseCaseImpl{
		userQuery:      userQuery,
		tokenGenerator: tokenGenerator,
	}
}

// Login - function for login
func (au *userUseCaseImpl) Login(email, password string) (*model.LoginResponse, error) {
	ctx := "AuthUseCase-GetProfile"

	user, err := au.userQuery.FindUserByEmail(email)
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), ctx, "find_detail_user")
		return nil, err
	}
	// put user validation here

	// set token
	claims := token.Claim{
		Issuer:     token.Issuer,
		Audience:   strconv.Itoa(user.ID),
		Subject:    user.Email,
		Authorised: true,
	}

	tokenResult := <-au.tokenGenerator.GenerateAccessToken(claims)
	// set response
	resp := model.LoginResponse{
		Email:        user.Email,
		TokenID:      tokenResult.AccessToken.AccessToken,
		DisplayName:  user.Name,
		Registered:   "true",
		LocalID:      helper.RandomStringBase64(28),
		Kind:         "identitytoolkit#VerifyPasswordResponse",
	}

	return &resp, nil
}



// GetProfile - function for getting user profile
func (au *userUseCaseImpl) GetProfile(ID int) (*userModel.User, error) {
	ctx := "AuthUseCase-GetProfile"

	user, err := au.userQuery.FindUserByID(ID)
	if err != nil {
		helper.Log(log.ErrorLevel, err.Error(), ctx, "find_detail_user")
		return nil, err
	}

	return user, nil
}