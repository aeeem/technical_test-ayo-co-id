package team_http

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/team"
	"technical_test-ayo-co-id/internal/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type TeamHandler struct {
	Validator   *validator.XValidator
	TeamUsecase team.TeamUsecase
}

func NewTeamHandler(f *fiber.App, validator *validator.XValidator, TeamUsecase team.TeamUsecase, middleware fiber.Handler) {
	TeamHandler := TeamHandler{
		TeamUsecase: TeamUsecase,
		Validator:   validator,
	}
	teamRoute := f.Group("team")
	//todo : add middleware for auth
	teamRoute.Get("/", TeamHandler.Fetch)      //handling Fetch Data
	teamRoute.Get("/:id", TeamHandler.GetById) //handling GetById Data
	admin := teamRoute.Group("/admin")
	admin.Use(middleware)
	admin.Post("/", TeamHandler.Save)        //handling Save Data
	admin.Put("/", TeamHandler.Update)       //handling Update Data
	admin.Delete("/:id", TeamHandler.Delete) //handling Delete Data
}

// @Router			/team	[get]
// @Summary			fetch team data
// @Param			cursor	query	string	false	"base64 encoded"
// @Param			search	query	string	false "Team abcd"
// @Param			limit	query	int	true "10"
// @Tags			Team
// @Success			200	{object}	helper.ListResponse{data=[]team.Team}
func (h *TeamHandler) Fetch(c *fiber.Ctx) (err error) {
	cursor := c.Query("cursor")
	search := c.Query("search")
	limit := c.QueryInt("limit", 10)

	res, nextCursor, err := h.TeamUsecase.Fetch(cursor, search, limit)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonListResponseSuccess(c, nextCursor, res)
}

// @Router			/team/{id}	[get]
// @Summary			get detail team data
// @Param			id	path	integer	true	"user id"
// @Tags			Team
// @Success			200	{object}	helper.StdResponse{data=team.Team}
func (h *TeamHandler) GetById(c *fiber.Ctx) (err error) {
	ID, err := c.ParamsInt("id")
	if err != nil {
		err = helper.BadRequest
		return helper.JsonErrorResponse(c, err)
	}
	res, err := h.TeamUsecase.GetById(ID)
	if err != nil {
		log.Info().Any("err", err).Msg("log from get by id")
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseSuccess(c, res)
}

// @Router			/team/admin	[post]
// @Summary			Insert data into databases
// @Security ApiKeyAuth
// @Tags			Team
// @Accept			json
// @Param			id	body	team_http.TeamRequest	true "team post request"
// @Success			200	{object}	helper.StdResponse{data=team.Team}
func (h *TeamHandler) Save(c *fiber.Ctx) (err error) {
	TeamRequest := TeamRequest{}
	if err := json.Unmarshal(c.Body(), &TeamRequest); err != nil {
		return helper.JsonErrorResponseValidation(c, err)
	}
	if errs := h.Validator.Validate(TeamRequest); len(errs) > 0 && errs[0].Error {
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
	TeamModel := team.Team{
		TeamName:    TeamRequest.TeamName,
		Logo:        TeamRequest.Logo,
		YearFounded: TeamRequest.YearFounded,
		Address:     TeamRequest.Address,
		City:        TeamRequest.City,
	}
	err = h.TeamUsecase.Save(&TeamModel)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseCreated(c, TeamModel)
}

// @Router			/team/admin	[put]
// @Summary			Update data from databases
// @Tags			Team
// @Security ApiKeyAuth
// @Accept			json
// @Param			id	body	team_http.TeamRequestUpdate	true "team put request"
// @Success			200	{object}	helper.StdResponse{data=team.Team}
func (h *TeamHandler) Update(c *fiber.Ctx) (err error) {
	TeamRequest := TeamRequestUpdate{}
	if err := json.Unmarshal(c.Body(), &TeamRequest); err != nil {
		return helper.JsonErrorResponseValidation(c, err)
	}
	if errs := h.Validator.Validate(TeamRequest); len(errs) > 0 && errs[0].Error {
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

	TeamModel := team.Team{
		Model: helper.Model{
			ID: TeamRequest.ID,
		},
		TeamName:    TeamRequest.TeamName,
		Logo:        TeamRequest.Logo,
		YearFounded: TeamRequest.YearFounded,
		Address:     TeamRequest.Address,
		City:        TeamRequest.City,
	}
	err = h.TeamUsecase.Update(&TeamModel)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseSuccess(c, TeamModel)
}

// @Router			/team/admin/{id}	[delete]
// @Summary			Delete data from databases
// @Tags			Team
// @Accept			json
// @Security ApiKeyAuth
// @Param			id	path	integer	true "user id"
// @Success			200	{object}	helper.StdResponse{data=nil}
func (h *TeamHandler) Delete(c *fiber.Ctx) (err error) {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	err = h.TeamUsecase.Delete(ID)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseDeleted(c)
}
