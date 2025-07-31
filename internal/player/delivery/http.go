package player_http

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/player"
	"technical_test-ayo-co-id/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type PlayerHandler struct {
	Validator     *validator.XValidator
	PlayerUsecase player.PlayerUsecase
}

func NewPlayerHandler(f *fiber.App, validator *validator.XValidator, PlayerUsecase player.PlayerUsecase, middleware fiber.Handler) {
	PlayerHandler := PlayerHandler{
		PlayerUsecase: PlayerUsecase,
		Validator:     validator,
	}
	playerRoutes := f.Group("player")
	//todo: add middleware for auth
	playerRoutes.Get("/", PlayerHandler.Fetch)      //handling Fetch Data
	playerRoutes.Get("/:id", PlayerHandler.GetById) //handling GetById Data

	admin := playerRoutes.Group("/admin")
	admin.Use(middleware)
	admin.Post("/", PlayerHandler.Save)        //handling Save Data
	admin.Put("/", PlayerHandler.Update)       //handling Update Data
	admin.Delete("/:id", PlayerHandler.Delete) //handling Delete Data
}

// @Router			/player	[get]
// @Summary			fetch player data
// @Param			team_id	query	integer	false	"not required but if filled will filter by team id"
// @Param			cursor	query	string	false	"base64 encoded"
// @Param			search	query	string	false "Team abcd"
// @Param			limit	query	integer	true "10"
// @Tags			player
// @Success			200	{object}	helper.ListResponse{data=[]player.Player}
func (h PlayerHandler) Fetch(c *fiber.Ctx) (err error) {
	teamID := c.QueryInt("team_id")
	cursor := c.Query("cursor")
	search := c.Query("search")
	limit := c.QueryInt("limit", 10)

	res, nextcursor, err := h.PlayerUsecase.Fetch(uint(teamID), cursor, search, limit)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonListResponseSuccess(c, nextcursor, res)
}

// @Router			/player/{id}	[get]
// @Summary			get detail player data
// @Param			id	path	integer	true	"user id"
// @Tags			player
// @Success			200	{object}	helper.StdResponse{data=player.Player}
func (h PlayerHandler) GetById(c *fiber.Ctx) (err error) {
	ID, err := c.ParamsInt("id")
	if err != nil {
		err = helper.BadRequest
		return helper.JsonErrorResponse(c, err)
	}
	res, err := h.PlayerUsecase.GetById(ID)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseSuccess(c, res)
}

// @Router			/player/admin	[post]
// @Summary			Insert data into databases
// @Security ApiKeyAuth
// @Tags			player
// @Accept			json
// @Param			id	body	player_http.PlayerRequest	true "team post request"
// @Success			200	{object}	helper.StdResponse{data=player.Player}
func (h PlayerHandler) Save(c *fiber.Ctx) (err error) {
	//validation
	PlayerRequest := PlayerRequest{}
	if err := json.Unmarshal(c.Body(), &PlayerRequest); err != nil {
		return helper.JsonErrorResponseValidation(c, err)
	}

	if errs := h.Validator.Validate(PlayerRequest); len(errs) > 0 && errs[0].Error {
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
	//end of validation
	PlayerModel := player.Player{
		FirstName:    PlayerRequest.FirstName,
		LastName:     PlayerRequest.LastName,
		Height:       PlayerRequest.Height,
		Weight:       PlayerRequest.Weight,
		Position:     PlayerRequest.Position,
		JerseyNumber: PlayerRequest.JerseyNumber,
		TeamID:       PlayerRequest.TeamID,
	}
	err = h.PlayerUsecase.Save(&PlayerModel)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseCreated(c, PlayerModel)
}

// @Router			/player/admin	[put]
// @Summary			Update data from databases
// @Security ApiKeyAuth
// @Tags			player
// @Accept			json
// @Param			id	body	player_http.PlayerRequestUpdate	true "Player put request"
// @Success			200	{object}	helper.StdResponse{data=player.Player}
func (h PlayerHandler) Update(c *fiber.Ctx) (err error) {
	// validation
	PlayerRequest := PlayerRequestUpdate{}
	if err := json.Unmarshal(c.Body(), &PlayerRequest); err != nil {
		return helper.JsonErrorResponseValidation(c, err)
	}

	if errs := h.Validator.Validate(PlayerRequest); len(errs) > 0 && errs[0].Error {
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
	// end of validation
	PlayerModel := player.Player{
		Model: helper.Model{
			ID: PlayerRequest.ID,
		},
		FirstName:    PlayerRequest.FirstName,
		LastName:     PlayerRequest.LastName,
		Height:       PlayerRequest.Height,
		Weight:       PlayerRequest.Weight,
		Position:     PlayerRequest.Position,
		JerseyNumber: PlayerRequest.JerseyNumber,
		TeamID:       PlayerRequest.TeamID,
	}
	err = h.PlayerUsecase.Update(&PlayerModel)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}

	return helper.JsonStandardResponseUpdated(c, PlayerModel)
}

// @Router			/player/admin/{id}	[delete]
// @Summary			Delete data from databases
// @Tags			player
// @Accept			json
// @Security ApiKeyAuth
// @Param			id	path	integer	true "Player id"
// @Success			200	{object}	helper.StdResponse{data=nil}
func (h PlayerHandler) Delete(c *fiber.Ctx) (err error) {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return helper.JsonErrorResponseValidation(c, err)
	}

	err = h.PlayerUsecase.Delete(ID)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}

	return helper.JsonStandardResponseDeleted(c)
}
