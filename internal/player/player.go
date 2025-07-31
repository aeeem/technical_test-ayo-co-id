package player

import (
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/team"
)

type Player struct {
	helper.Model
	FirstName    string     `json:"first_name" gorm:"index"`
	LastName     string     `json:"last_name" gorm:"index"`
	Height       uint       `json:"height"`
	Weight       uint       `json:"weight"`
	Position     string     `json:"position"`
	JerseyNumber int        `json:"jersey_number" gorm:"uniqueIndex:idx_player_team"`
	TeamID       uint       `json:"team_id" gorm:"uniqueIndex:idx_player_team"`
	Team         *team.Team `json:"team"`
}

type PlayerRepository interface {
	Fetch(teamID uint, lastDate string, lastId int, search string, limit int) (Player []Player, err error)
	GetById(ID int) (Player Player, err error)
	Save(Player *Player) (err error)
	Update(Player *Player) (err error)
	Delete(ID int) (err error)
}

type PlayerUsecase interface {
	Fetch(teamID uint, cursor string, search string, limit int) (Player []Player, nextCursor string, err error)
	GetById(ID int) (Player Player, err error)
	Save(Player *Player) (err error)
	Update(Player *Player) (err error)
	Delete(ID int) (err error)
}
