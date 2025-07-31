package repository

import (
	"technical_test-ayo-co-id/internal/auth"
	"technical_test-ayo-co-id/internal/helper"

	"gorm.io/gorm"
)

type authRepository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) auth.AuthRepository {
	return &authRepository{
		DB: db,
	}
}

func (r *authRepository) Save(auth *auth.Auth) (err error) {
	err = r.DB.Model(&auth).Create(auth).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}
func (r *authRepository) GetAuthByToken(token string) (auth auth.Auth, err error) {
	err = r.DB.Where("token = ?", token).First(&auth).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}
func (r *authRepository) Delete(ID int) (err error) {
	err = r.DB.Model(&auth.Auth{}).Delete("id = ?", ID).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}

func (r *authRepository) GetAuthByUserID(userID uint) (auth auth.Auth, err error) {
	err = r.DB.Where("user_id = ?", userID).First(&auth).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}
