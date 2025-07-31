package repository

import (
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/user"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (r *userRepository) Save(user *user.User) (err error) {
	err = r.DB.Model(&user).Create(user).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}
func (r *userRepository) GetByEmail(email string) (user user.User, err error) {

	err = r.DB.Model(&user).Where("email = ?", email).First(&user).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}
