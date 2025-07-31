package score_http

type ScoreRequest struct {
	TeamID    uint   `json:"team_id" validate:"required" example:"1"`
	MatchID   uint   `json:"match_id" validate:"required" example:"1"`
	PlayerID  uint   `json:"player_id" validate:"required" example:"1"`
	ScoreTime string `json:"score_time" validate:"required" example:"12:30:00"`
}

type ScoreUpdateRequest struct {
	ScoreRequest
	ID uint `json:"id" validate:"required" example:"1"`
}
