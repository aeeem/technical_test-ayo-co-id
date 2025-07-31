package auth

import "technical_test-ayo-co-id/internal/helper"

type Auth struct {
	helper.Model
	Token     string
	ExpiredIn string
	UserID    uint
}

type AuthRepository interface {
	Save(auth *Auth) (err error)
	GetAuthByToken(token string) (auth Auth, err error)
	GetAuthByUserID(userID uint) (auth Auth, err error)
	Delete(ID int) (err error)
}

type AuthUsecase interface {
	Login(email string, password string) (auth Auth, err error)
	Logout(token string) (err error)
	CheckToken(token string) (auth Auth, err error)
}
