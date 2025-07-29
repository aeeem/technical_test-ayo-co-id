package team_http

type TeamRequest struct {
	TeamName string `json:"team_name" validate:"required"`
	Logo     string `json:"logo" validate:"required"`
	Founded  string `json:"founded" validate:"required"`
	Address  string `json:"address" validate:"required"`
	City     string `json:"city" validate:"required"`
}
type TeamRequestUpdate struct {
	TeamRequest
	ID uint `json:"id" validate:"required"`
}
