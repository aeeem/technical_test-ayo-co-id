package auth_http

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"technical_test-ayo-co-id/internal/auth"
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/user"
	"technical_test-ayo-co-id/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	AuthUsecase auth.AuthUsecase
	UserUsecase user.UserUsecase
	Validator   *validator.XValidator
}

func NewAuthHandler(f *fiber.App, Validator *validator.XValidator, Userusecase user.UserUsecase, authUsecase auth.AuthUsecase) (middleware fiber.Handler) {
	AuthHandler := AuthHandler{
		AuthUsecase: authUsecase,
		UserUsecase: Userusecase,
		Validator:   Validator,
	}
	userRoutes := f.Group("auth")
	//todo: add middleware for auth
	userRoutes.Post("/register", AuthHandler.Register) //handling Fetch Data
	userRoutes.Post("/login", AuthHandler.Login)       //handling GetById Data
	auth := userRoutes.Group("/logout").Use(New(Config{
		AuthHandler: AuthHandler,
	}))
	auth.Post("/", AuthHandler.Logout) //handling GetById Data

	return New(Config{
		AuthHandler: AuthHandler,
	})
}

// @Router			/auth/register	[post]
// @Summary			register auth data into databases
// @Tags			Auth
// @Accept			json
// @Param			id	body	auth_http.Register	true "auth register post request"
// @Success			200	{object}	helper.StdResponse{}
func (h AuthHandler) Register(c *fiber.Ctx) (err error) {
	Register := Register{}
	if err := json.Unmarshal(c.Body(), &Register); err != nil {
		return helper.JsonErrorResponseValidation(c, err)
	}
	if errs := h.Validator.Validate(Register); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}
		err = errors.New(strings.Join(errMsgs, "/n"))
		return helper.JsonErrorResponseValidation(c, err)
	}
	User := user.User{
		Email:    Register.Email,
		Password: Register.Password,
	}
	err = h.UserUsecase.Register(&User)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseCreated(c, nil)
}

// @Router			/auth/login	[post]
// @Summary			login auth data into databases
// @Tags			Auth
// @Accept			json
// @Param			id	body	auth_http.Login	true "auth login post request"
// @Success			200	{object}	helper.StdResponse{data=auth.Auth}
func (h AuthHandler) Login(c *fiber.Ctx) (err error) {
	login := Login{}
	if err := json.Unmarshal(c.Body(), &login); err != nil {
		return helper.JsonErrorResponseValidation(c, err)
	}
	if errs := h.Validator.Validate(login); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}
		err = errors.New(strings.Join(errMsgs, "/n"))
		return helper.JsonErrorResponseValidation(c, err)
	}

	Auth, err := h.AuthUsecase.Login(login.Email, login.Password)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseCreated(c, Auth)
}

// @Router			/auth/logout	[post]
// @Summary			logout auth data into databases
// @Tags			Auth
// @Accept			json
// @Param			id	body	auth_http.Logout	true "auth register post request"
// @Success			200	{object}	helper.StdResponse{}
func (h AuthHandler) Logout(c *fiber.Ctx) (err error) {
	logout := Logout{}
	if err := json.Unmarshal(c.Body(), &logout); err != nil {
		return helper.JsonErrorResponseValidation(c, err)
	}
	if errs := h.Validator.Validate(logout); len(errs) > 0 && errs[0].Error {
		errMsgs := make([]string, 0)

		for _, err := range errs {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: '%v' | Needs to implement '%s'",
				err.FailedField,
				err.Value,
				err.Tag,
			))
		}
		err = errors.New(strings.Join(errMsgs, "/n"))
		return helper.JsonErrorResponseValidation(c, err)
	}

	err = h.AuthUsecase.Logout(logout.Token)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseSuccess(c, nil)
}
