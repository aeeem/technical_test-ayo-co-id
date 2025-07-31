package user

import "technical_test-ayo-co-id/internal/helper"

type User struct {
	helper.Model
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
}

type UserRepository interface {
	Save(user *User) (err error)
	GetByEmail(email string) (user User, err error)
}

type UserUsecase interface {
	Register(user *User) (err error)
	GetByEmail(email string) (user User, err error)
}
