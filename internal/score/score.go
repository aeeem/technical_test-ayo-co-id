package score

import (
	"technical_test-ayo-co-id/internal/helper"
)

type Score struct {
	helper.Model
	MatchID   uint   `json:"match_id" gorm:"index;foreignKey:MatchID;references:ID"`
	TeamID    uint   `json:"team_id"`
	PlayerID  uint   `json:"player_id"`
	ScoreTime string `json:"score_time"`
}

type ScoreRepository interface {
	FetchScoreByMatchID(MatchID int) (Score []Score, err error)
	GetByID(ID uint) (Score Score, err error)
	FetchScoreTeamByMatchID(MatchID int, teamID uint) (Score []Score, err error)
	Save(score *Score, isHome bool) (err error)
	Update(score *Score) (err error)
	Delete(ID int) (err error)
}

type ScoreUsecase interface {
	FetchScoreByMatchID(MatchID int) (Score []Score, err error)
	FetchScoreTeamByMatchID(MatchID int, teamID uint) (Score []Score, err error)
	Save(score *Score) (err error)
	Update(score *Score) (err error)
	Delete(ID int) (err error)
}
