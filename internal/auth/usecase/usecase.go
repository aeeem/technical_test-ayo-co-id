package usecase

import (
	"fmt"
	"technical_test-ayo-co-id/internal/auth"
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/user"
	"time"

	"github.com/rs/zerolog/log"
)

type authUsecase struct {
	UserUsecase    user.UserUsecase
	AuthRepository auth.AuthRepository
}

func NewAuthusecase(authRepository auth.AuthRepository, userUsecase user.UserUsecase) auth.AuthUsecase {
	return &authUsecase{
		AuthRepository: authRepository,
		UserUsecase:    userUsecase,
	}
}

func (u *authUsecase) Login(email string, password string) (auth auth.Auth, err error) {
	//check if user is exist
	user, err := u.UserUsecase.GetByEmail(email)
	if err != nil {
		err = helper.MakeUsecaseLevelErr(400, "user not found")
		return
	}

	if !helper.CompareHashPassword(user.Password, password) {
		err = helper.MakeUsecaseLevelErr(403, "wrong password")
		return
	}
	//create tokens
	tokenStructure := fmt.Sprintf("%s=%d=%d", user.Email, user.ID, time.Now().Unix())
	auth.Token, err = helper.GenerateHashPassword(tokenStructure)
	if err != nil {
		log.Print(err)
		err = helper.MakeUsecaseLevelErr(500, "error generating token")
		return
	}
	auth.UserID = user.ID
	auth.ExpiredIn = time.Now().Add(time.Hour * 24).Format(time.RFC3339)
	//save session
	err = u.AuthRepository.Save(&auth)
	if err != nil {
		return
	}
	return
}
func (u *authUsecase) Logout(token string) (err error) {
	//get token
	auth, err := u.AuthRepository.GetAuthByToken(token)
	if err != nil {
		return
	}
	err = u.AuthRepository.Delete(int(auth.ID))
	if err != nil {
		return
	}
	return
}

func (u *authUsecase) CheckToken(token string) (auth auth.Auth, err error) {
	auth, err = u.AuthRepository.GetAuthByToken(token)
	if err != nil {
		if err == helper.ErrNotFound {
			err = helper.MakeUsecaseLevelErr(403, "Unauthorized")
			return
		}
		return
	}
	TokenExpired, err := time.Parse(time.RFC3339, auth.ExpiredIn)
	if err != nil {
		log.Print(err)
		err = helper.MakeUsecaseLevelErr(500, "error parsing expired")
		return
	}
	if time.Now().After(TokenExpired) {
		err = helper.MakeUsecaseLevelErr(401, "token expired")
		return
	}
	return
}
