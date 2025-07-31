package match_http

import (
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/match"
	"technical_test-ayo-co-id/internal/player"
	"technical_test-ayo-co-id/internal/score"
	"technical_test-ayo-co-id/internal/team"
)

type MatchResponse struct {
	helper.Model
	MatchDate        string  `json:"match_date" gorm:"type:date;index"`
	MatchTime        string  `json:"match_time" gorm:"type:time;index"`
	HomeTeam         uint    `json:"home_team" gorm:"index"`
	MatchStatus      string  `json:"match_status" gorm:"type:varchar(200);index"`
	AwayTeam         uint    `json:"away_team" gorm:"index"`
	MatchDescription string  `json:"match_description" gorm:"type:varchar(200);default:null" swaggertype:"string"`
	TotalScoreHome   *int32  `json:"total_score_home" gorm:"index;default:null" swaggertype:"integer"`
	TotalScoreAway   *int32  `json:"total_score_away" gorm:"index;default:null" swaggertype:"integer"`
	Winner           *int32  `json:"winner" gorm:"index;default:null" swaggertype:"integer"`
	TeamWinnerName   *string `json:"team_winner_name" gorm:"type:character(255);index;default:null" swaggertype:"string"`
	PlayerMvpID      *int32  `json:"player_mvp_id" gorm:"index;default:null" swaggertype:"integer"`
	//preloaded model should be handled differently
	HomeTeamDetail *team.Team     `json:"home_team_details" gorm:"foreignKey:HomeTeam;references:ID"`
	AwayTeamDetail *team.Team     `json:"away_team_details" gorm:"foreignKey:AwayTeam;references:ID"`
	PlayerMVP      *player.Player `json:"player_mvp"`
	ScoreHome      []score.Score  `json:"score_home" gorm:"foreignKey:MatchID;references:ID"`
	ScoreAway      []score.Score  `json:"score_away" gorm:"foreignKey:MatchID;references:ID"`
	MatchScore     []score.Score  `json:"match_score" gorm:"foreignKey:MatchID;references:ID"`
}

func TransformArrToJson(matches []match.Match) []MatchResponse {
	var matchResponse []MatchResponse
	for _, match := range matches {
		matchResponse = append(matchResponse, TransformIntoJson(match))
	}
	return matchResponse
}

func TransformIntoJson(match match.Match) MatchResponse {
	matchResponse := MatchResponse{
		Model:       match.Model,
		MatchDate:   match.MatchDate,
		MatchTime:   match.MatchTime,
		HomeTeam:    match.HomeTeam,
		AwayTeam:    match.AwayTeam,
		MatchStatus: match.MatchStatus,
		MatchScore:  match.MatchScore,
		ScoreHome:   match.ScoreHome,
		ScoreAway:   match.ScoreAway,
	}

	if match.TotalScoreHome.Valid {
		matchResponse.TotalScoreHome = &match.TotalScoreHome.Int32
	}

	if match.TotalScoreAway.Valid {

		matchResponse.TotalScoreAway = &match.TotalScoreAway.Int32
	}

	if match.Winner.Valid {

		matchResponse.Winner = &match.Winner.Int32
	}

	if match.TeamWinnerName.Valid {

		matchResponse.TeamWinnerName = &match.TeamWinnerName.String
	}

	if match.PlayerMvpID.Valid {

		matchResponse.PlayerMvpID = &match.PlayerMvpID.Int32
	}

	if match.MatchDescription.Valid {

		matchResponse.MatchDescription = match.MatchDescription.String
	}

	if match.HomeTeamDetail != nil {

		matchResponse.HomeTeamDetail = match.HomeTeamDetail
	}

	if match.AwayTeamDetail != nil {

		matchResponse.AwayTeamDetail = match.AwayTeamDetail
	}

	if match.PlayerMVP != nil {
		matchResponse.PlayerMVP = match.PlayerMVP
	}

	return matchResponse
}
