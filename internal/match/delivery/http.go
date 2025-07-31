package match_http

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/match"

	"technical_test-ayo-co-id/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type MatchHandler struct {
	Validator    *validator.XValidator
	MatchUsecase match.MatchUsecase
}

func NewMatchHandler(f *fiber.App, validator *validator.XValidator, Matchusecase match.MatchUsecase) {
	MatchHandler := MatchHandler{
		Validator:    validator,
		MatchUsecase: Matchusecase,
	}
	matchRoutes := f.Group("match")
	//todo: add middleware for auth
	matchRoutes.Get("/", MatchHandler.Fetch)        //handling Fetch Data
	matchRoutes.Get("/:id", MatchHandler.GetById)   //handling GetById Data
	matchRoutes.Post("/", MatchHandler.Save)        //handling Save Data
	matchRoutes.Put("/", MatchHandler.Update)       //handling Update Data
	matchRoutes.Delete("/:id", MatchHandler.Delete) //handling Delete Data
}

// @Router			/match	[get]
// @Summary			fetch match data
// @Param			home_id	query	integer	false	"not required but if filled it will filter match history between teams"
// @Param			away_id	query	integer	false	"not required but if filled it will filter match history between teams"
// @Param			player_id	query	integer	false "not required but if filled it will return match history where player become MVP"
// @Param			cursor	query	string	false	"base64 encoded"
// @Param			search	query	string	false "Team abcd"
// @Param			limit	query	integer	true "10"
// @Tags			match
// @Success			200	{object}	helper.ListResponse{data=[]match.Match}
func (h *MatchHandler) Fetch(c *fiber.Ctx) (err error) {
	home_id := c.QueryInt("home_id")
	away_id := c.QueryInt("away_id")
	player_id := c.QueryInt("player_id")
	cursor := c.Query("cursor")
	search := c.Query("search")
	limit := c.QueryInt("limit", 10)

	res, nextcursor, err := h.MatchUsecase.Fetch(uint(home_id), uint(away_id), uint(player_id), cursor, search, limit)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonListResponseSuccess(c, nextcursor, TransformArrToJson(res))
}

// @Router			/match/{id}	[get]
// @Summary			get detail match data
// @Param			id	path	integer	true	"user id"
// @Tags			match
// @Success			200	{object}	helper.StdResponse{data=match.Match}
func (h *MatchHandler) GetById(c *fiber.Ctx) (err error) {
	ID, err := c.ParamsInt("id")
	if err != nil {
		err = helper.BadRequest
		return helper.JsonErrorResponse(c, err)
	}
	res, err := h.MatchUsecase.GetById(ID)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseSuccess(c, TransformIntoJson(res))
}

// @Router			/match	[post]
// @Summary			Insert data into databases
// @Tags			match
// @Accept			json
// @Param			id	body	match_http.MatchRequest	true "team post request"
// @Success			200	{object}	helper.StdResponse{data=match.Match}
func (h *MatchHandler) Save(c *fiber.Ctx) (err error) {
	MatchRequest := MatchRequest{}
	if err := json.Unmarshal(c.Body(), &MatchRequest); err != nil {
		return helper.JsonErrorResponseValidation(c, err)
	}

	if errs := h.Validator.Validate(MatchRequest); len(errs) > 0 && errs[0].Error {
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
	MatchModel := match.Match{
		MatchDate:   MatchRequest.MatchDate,
		MatchTime:   MatchRequest.MatchTime,
		HomeTeam:    MatchRequest.HomeTeam,
		MatchStatus: MatchRequest.MatchStatus,
		AwayTeam:    MatchRequest.AwayTeam,
	}

	err = h.MatchUsecase.Save(&MatchModel)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseSuccess(c, TransformIntoJson(MatchModel))
}

// @Router			/match	[put]
// @Summary			Update data from databases
// @Description		Updating data after match finished
// @Tags			match
// @Accept			json
// @Param			id	body	match_http.MatchRequestUpdate	true "match put request"
// @Success			200	{object}	helper.StdResponse{data=match.Match}
func (h *MatchHandler) Update(c *fiber.Ctx) (err error) {
	MatchRequestUpdate := MatchRequestUpdate{}
	if err := json.Unmarshal(c.Body(), &MatchRequestUpdate); err != nil {
		return helper.JsonErrorResponseValidation(c, err)
	}

	if errs := h.Validator.Validate(MatchRequestUpdate); len(errs) > 0 && errs[0].Error {
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
	MatchModel := match.Match{
		Model: helper.Model{
			ID: MatchRequestUpdate.ID,
		},
		MatchDate: MatchRequestUpdate.MatchDate,
		MatchTime: MatchRequestUpdate.MatchTime,
		HomeTeam:  MatchRequestUpdate.HomeTeam,
		AwayTeam:  MatchRequestUpdate.AwayTeam,
	}

	//use update for updating match data after the match ended
	MatchModel.Winner.Int32 = int32(MatchRequestUpdate.Winner)
	MatchModel.TotalScoreHome.Int32 = int32(MatchRequestUpdate.TotalScoreHome)
	MatchModel.TotalScoreAway.Int32 = int32(MatchRequestUpdate.TotalScoreAway)
	MatchModel.TeamWinnerName.String = MatchRequestUpdate.TeamWinnerName
	MatchModel.PlayerMvpID.Int32 = int32(MatchRequestUpdate.PlayerMvpID)
	MatchModel.MatchDescription.String = MatchRequestUpdate.MatchDescription
	err = h.MatchUsecase.Update(&MatchModel)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseSuccess(c, TransformIntoJson(MatchModel))
}

// @Router			/match/{id}	[delete]
// @Summary			Delete data from databases
// @Tags			match
// @Accept			json
// @Param			id	path	integer	true "match id"
// @Success			200	{object}	helper.StdResponse{data=nil}
func (h *MatchHandler) Delete(c *fiber.Ctx) (err error) {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return helper.JsonErrorResponseValidation(c, err)
	}

	err = h.MatchUsecase.Delete(ID)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}

	return helper.JsonStandardResponseDeleted(c)
}
