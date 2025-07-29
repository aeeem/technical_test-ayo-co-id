package team

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	TeamName string `json:"team_name" `
	Logo     string `json:"logo" `
	Founded  string `json:"founded" `
	Address  string `json:"address" `
	City     string `json:"city" `
}

type TeamRepository interface {
	Fetch(lastDate string, lastId int, search string) (team []Team, err error)
	GetById(ID int) (team Team, err error)
	Save(Team *Team) (err error)
	Update(Team *Team) (err error)
	Delete(ID int) (err error)
}

type TeamUsecase interface {
	Fetch(cursor string, search string) (team []Team, nextCursor string, err error)
	GetById(ID int) (team Team, err error)
	Save(Team *Team) (err error)
	Update(Team *Team) (err error)
	Delete(ID int) (err error)
}
