package player_http

type PlayerRequest struct {
	FirstName    string `json:"first_name" validate:"required,alphaunicode" example:"Arif"`
	LastName     string `json:"last_name" validate:"required,ascii" example:"Maulana"`
	Height       uint   `json:"height" validate:"required,numeric,gt=0" example:"173"`
	Weight       uint   `json:"weight" validate:"required,numeric,gt=0" example:"200"`
	Position     string `json:"position" validate:"required,oneof=attacker midfielder defender goalkeeper" example:"midfielder"`
	JerseyNumber int    `json:"jersey_number" validate:"numeric,gte=0" example:"11"`
	TeamID       uint   `json:"team_id" validate:"required" example:"1"`
}
type PlayerRequestUpdate struct {
	PlayerRequest
	ID uint `json:"id" validate:"required" example:"1"`
}
