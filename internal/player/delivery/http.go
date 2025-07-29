package delivery

import (
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/player"
	"technical_test-ayo-co-id/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type PlayerHandler struct {
	Validator     *validator.XValidator
	PlayerUsecase player.PlayerUsecase
}

func NewTeamHandler(f *fiber.App, validator *validator.XValidator, PlayerUsecase player.PlayerUsecase) {
	TeamHandler := PlayerHandler{
		PlayerUsecase: PlayerUsecase,
		Validator:     validator,
	}
	teamRoute := f.Group("player")
	//todo : add middleware for auth
	teamRoute.Get("/", TeamHandler.Fetch)        //handling Fetch Data
	teamRoute.Get("/:id", TeamHandler.GetById)   //handling GetById Data
	teamRoute.Post("/", TeamHandler.Save)        //handling Save Data
	teamRoute.Put("/", TeamHandler.Update)       //handling Update Data
	teamRoute.Delete("/:id", TeamHandler.Delete) //handling Delete Data
}

// @Router			/player	[get]
// @Summary			fetch player data
// @Param			team_id	query	integer	false	not required but if filled will filter by team id
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
// @Tags			Player
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

// @Router			/player	[post]
// @Summary			Insert data into databases
// @Tags			Player
// @Accept			json
// @Param			id	body	team_http.TeamRequest	true "team post request"
// @Success			200	{object}	helper.StdResponse{data=player.Player}
func (h PlayerHandler) Save(c *fiber.Ctx) (err error) {
	return
}

func (h PlayerHandler) Update(c *fiber.Ctx) (err error) {
	return
}

func (h PlayerHandler) Delete(c *fiber.Ctx) (err error) {
	return
}
