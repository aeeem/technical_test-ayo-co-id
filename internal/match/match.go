package match

import (
	"database/sql"
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/player"
	"technical_test-ayo-co-id/internal/score"
	"technical_test-ayo-co-id/internal/team"
)

type Match struct {
	helper.Model
	MatchDate        string         `json:"match_date" gorm:"type:date;index"`
	MatchTime        string         `json:"match_time" gorm:"type:varchar(200);index"`
	HomeTeam         uint           `json:"home_team" gorm:"index"`
	MatchStatus      string         `json:"match_status" gorm:"type:varchar(200);index"`
	AwayTeam         uint           `json:"away_team" gorm:"index"`
	MatchDescription sql.NullString `json:"match_description" gorm:"type:varchar(200);default:null" swaggertype:"string"`
	TotalScoreHome   sql.NullInt32  `json:"total_score_home" gorm:"index;default:null" swaggertype:"integer"`
	TotalScoreAway   sql.NullInt32  `json:"total_score_away" gorm:"index;default:null" swaggertype:"integer"`
	Winner           sql.NullInt32  `json:"winner" gorm:"index;default:null" swaggertype:"integer"`
	TeamWinnerName   sql.NullString `json:"team_winner_name" gorm:"type:varchar(255);index;default:null" swaggertype:"string"`
	PlayerMvpID      sql.NullInt32  `json:"player_mvp_id" gorm:"index;default:null" swaggertype:"integer"`
	//this type of model should be handled differently only inside Match details usecase
	HomeTeamDetail *team.Team     `json:"home_team_details" gorm:"-"`
	AwayTeamDetail *team.Team     `json:"away_team_details" gorm:"-"`
	PlayerMVP      *player.Player `json:"player_mvp" gorm:"-"`
	ScoreHome      []score.Score  `json:"score_home" gorm:"-"`
	ScoreAway      []score.Score  `json:"score_away" gorm:"-"`
	MatchScore     []score.Score  `json:"match_score" gorm:"-"`
	PreviousMatch  []Match        `json:"previous_match" gorm:"-"`
}

type MatchRepository interface {
	Fetch(AwayTeamID uint, HomeTeamID uint, mvp uint, lastDate string, lastId int, search string, limit int) (match []Match, err error)
	GetById(ID int) (match Match, err error)
	Save(match *Match) (err error)
	Update(match *Match) (err error)
	Delete(ID int) (err error)
}

type MatchUsecase interface {
	Fetch(AwayTeamID uint, HomeTeamID uint, mvp uint, Cusror string, search string, limit int) (match []Match, nextCursor string, err error)
	GetById(ID int) (match Match, err error)
	Save(match *Match) (err error)
	Update(match *Match) (err error)
	Delete(ID int) (err error)
}
