package score_http

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/score"
	"technical_test-ayo-co-id/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type ScoreHandler struct {
	Validator    *validator.XValidator
	ScoreUsecase score.ScoreUsecase
}

func NewScoreHandler(f *fiber.App, Validator *validator.XValidator, ScoreUsecase score.ScoreUsecase) {
	ScoreHandler := ScoreHandler{
		Validator:    Validator,
		ScoreUsecase: ScoreUsecase,
	}

	scoreGroup := f.Group("score")
	//todo : add middleware
	scoreGroup.Get("/:match_id", ScoreHandler.FetchScoreByMatchID)
	scoreGroup.Get("/:match_id/:team_id", ScoreHandler.FetchScoreTeamByMatchID)
	scoreGroup.Post("/", ScoreHandler.Save)
	scoreGroup.Put("/", ScoreHandler.Update)
	scoreGroup.Delete("/:id", ScoreHandler.Delete)

}

// @Router			/score/{match_id}	[get]
// @Summary			fetch team data
// @Param			match_id	path	integer	false	"match id"
// @Tags			Score
// @Success			200	{object}	helper.ListResponse{data=[]score.Score}
func (h *ScoreHandler) FetchScoreByMatchID(c *fiber.Ctx) (err error) {
	matchID, err := c.ParamsInt("match_id")
	if err != nil {
		err = errors.New("matchid should be integer")
		return helper.JsonErrorResponseValidation(c, err)
	}

	res, err := h.ScoreUsecase.FetchScoreByMatchID(matchID)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonListResponseSuccess(c, "", res)
}

// @Router			/score/{match_id}/{team_id}	[get]
// @Summary			fetch team data
// @Param			match_id	path	integer	false	"match id"
// @Param			team_id	path	integer	false	"team id"
// @Tags			Score
// @Success			200	{object}	helper.ListResponse{data=[]score.Score}
func (h *ScoreHandler) FetchScoreTeamByMatchID(c *fiber.Ctx) (err error) {
	matchID, err := c.ParamsInt("match_id")
	if err != nil {
		err = errors.New("match should be integer")
		return helper.JsonErrorResponseValidation(c, err)
	}
	teamID, err := c.ParamsInt("team_id")
	if err != nil {
		err = errors.New("team should be integer")
		return helper.JsonErrorResponseValidation(c, err)
	}
	res, err := h.ScoreUsecase.FetchScoreTeamByMatchID(matchID, uint(teamID))
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonListResponseSuccess(c, "", res)
}

// @Router			/score	[post]
// @Summary			Insert data into databases
// @Tags			Score
// @Accept			json
// @Param			id	body	score_http.ScoreRequest	true "score post request"
// @Success			200	{object}	helper.StdResponse{data=score.Score}
func (h *ScoreHandler) Save(c *fiber.Ctx) (err error) {
	ScoreRequest := ScoreRequest{}
	if err := json.Unmarshal(c.Body(), &ScoreRequest); err != nil {
		return helper.JsonErrorResponseValidation(c, err)
	}
	if errs := h.Validator.Validate(ScoreRequest); len(errs) > 0 && errs[0].Error {
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
	ScoreModel := score.Score{
		TeamID:    ScoreRequest.TeamID,
		MatchID:   ScoreRequest.MatchID,
		PlayerID:  ScoreRequest.PlayerID,
		ScoreTime: ScoreRequest.ScoreTime,
	}
	err = h.ScoreUsecase.Save(&ScoreModel)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseCreated(c, ScoreModel)
}

// @Router			/score	[put]
// @Summary			Insert data into databases
// @Tags			Score
// @Accept			json
// @Param			id	body	score_http.ScoreUpdateRequest	true "score post request"
// @Success			200	{object}	helper.StdResponse{data=score.Score}
func (h *ScoreHandler) Update(c *fiber.Ctx) (err error) {
	ScoreRequest := ScoreUpdateRequest{}
	if err := json.Unmarshal(c.Body(), &ScoreRequest); err != nil {
		return helper.JsonErrorResponseValidation(c, err)
	}
	if errs := h.Validator.Validate(ScoreRequest); len(errs) > 0 && errs[0].Error {
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
	ScoreModel := score.Score{
		Model:     helper.Model{ID: ScoreRequest.ID},
		TeamID:    ScoreRequest.TeamID,
		MatchID:   ScoreRequest.MatchID,
		PlayerID:  ScoreRequest.PlayerID,
		ScoreTime: ScoreRequest.ScoreTime,
	}
	err = h.ScoreUsecase.Update(&ScoreModel)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseCreated(c, ScoreModel)
}
func (h *ScoreHandler) Delete(c *fiber.Ctx) (err error) {
	ID, err := c.ParamsInt("id")
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	err = h.ScoreUsecase.Delete(ID)
	if err != nil {
		return helper.JsonErrorResponse(c, err)
	}
	return helper.JsonStandardResponseDeleted(c)
}
