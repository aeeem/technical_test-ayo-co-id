package usecase

import (
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/user"
)

type userUsecase struct {
	userRepository user.UserRepository
}

func NewUserUsecase(userRepository user.UserRepository) user.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) Register(user *user.User) (err error) {
	user.Password, err = helper.GenerateHashPassword(user.Password)
	if err != nil {
		err = helper.MakeUsecaseLevelErr(500, "error generating password hash")
		return
	}
	err = u.userRepository.Save(user)
	if err != nil {
		return
	}
	return
}
func (u *userUsecase) GetByEmail(email string) (user user.User, err error) {
	user, err = u.userRepository.GetByEmail(email)
	if err != nil {
		return
	}
	return
}
