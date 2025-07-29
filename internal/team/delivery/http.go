package http

import (
	"technical_test-ayo-co-id/internal/team"
	"technical_test-ayo-co-id/internal/validator"

	"github.com/gofiber/fiber/v2"
)

type TeamHandler struct {
	Validator   *validator.XValidator
	TeamUsecase team.TeamUsecase
}

func NewTeamHandler(f *fiber.App, validator validator.XValidator, TeamUsecase team.TeamUsecase) {
	TeamHandler := TeamHandler{
		TeamUsecase: TeamUsecase,
		Validator:   &validator,
	}
	teamRoute := f.Group("team")
	//todo : add middleware for auth
	teamRoute.Get("/", TeamHandler.Fetch)      //handling Fetch Data
	teamRoute.Get("/:id", TeamHandler.GetById) //handling GetById Data
	teamRoute.Post("/", TeamHandler.Save)      //handling Save Data
	teamRoute.Put("/", TeamHandler.Update)     //handling Update Data
	teamRoute.Delete("/", TeamHandler.Update)  //handling Update Data

}

// Getting all data for admins
func (h *TeamHandler) Fetch(c *fiber.Ctx) (err error) {

	return

}

func (h *TeamHandler) GetById(c *fiber.Ctx) (err error) {
	return
}
func (h *TeamHandler) Save(c *fiber.Ctx) (err error) {
	return
}
func (h *TeamHandler) Update(c *fiber.Ctx) (err error) {
	return
}
func (h *TeamHandler) Delete(c *fiber.Ctx) (err error) {
	return
}
