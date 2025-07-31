package team

import (
	"technical_test-ayo-co-id/internal/helper"
)

type Team struct {
	helper.Model
	TeamName    string `json:"team_name" example:"team_name"`
	Logo        string `json:"logo" example:"logo"`
	YearFounded string `json:"year_founded" example:"1994"`
	Address     string `json:"address" example:"address"`
	City        string `json:"city" example:"city"`
}

type TeamRepository interface {
	Fetch(lastDate string, lastId int, search string, limit int) (team []Team, err error)
	GetById(ID int) (team Team, err error)
	Save(Team *Team) (err error)
	Update(Team *Team) (err error)
	Delete(ID int) (err error)
}

type TeamUsecase interface {
	Fetch(cursor string, search string, limit int) (team []Team, nextCursor string, err error)
	GetById(ID int) (team Team, err error)
	Save(Team *Team) (err error)
	Update(Team *Team) (err error)
	Delete(ID int) (err error)
}
