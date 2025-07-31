package match_http

type MatchRequest struct {
	MatchDate   string `json:"match_date" validate:"required" example:"1994-12-15"`
	MatchTime   string `json:"match_time" validate:"required" example:"15:45:00"`
	MatchStatus string `json:"match_status" validate:"required,oneof=ongoing finished upcoming" example:"upcoming"`
	HomeTeam    uint   `json:"home_team" gorm:"index" validate:"required"`
	AwayTeam    uint   `json:"away_team" gorm:"index" validate:"required"`
}

type MatchRequestUpdate struct {
	MatchRequest
	ID               uint   `json:"id" validate:"required" example:"1"`
	TotalScoreHome   uint   `json:"total_score_home"`
	TotalScoreAway   uint   `json:"total_score_away"`
	Winner           uint   `json:"winner" gorm:"index" validate:"numeric"`
	TeamWinnerName   string `json:"team_winner_name" gorm:"index" validate:"ascii"`
	PlayerMvpID      uint   `json:"player_mvp_id" gorm:"index" validate:"numeric"`
	MatchDescription string `json:"match_description" validate:"ascii"`
}
