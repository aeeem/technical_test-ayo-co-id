package usecase

import (
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/player"
	"technical_test-ayo-co-id/internal/team"
)

type PlayerUsecase struct {
	playerRepository player.PlayerRepository
	TeamUsecase      team.TeamUsecase
}

func NewPlayerUsecase(playerRepository player.PlayerRepository, TeamUsecase team.TeamUsecase) player.PlayerUsecase {
	return &PlayerUsecase{
		TeamUsecase:      TeamUsecase,
		playerRepository: playerRepository,
	}
}

func (u *PlayerUsecase) Fetch(teamID uint, cursor string, search string, limit int) (Player []player.Player, nextCursor string, err error) {
	// get cursor and decode it
	date, ID, err := helper.CursorToDateAndID(cursor)
	if err != nil {
		//handling error
		return
	}

	Player, err = u.playerRepository.Fetch(teamID, date, ID, search, limit)
	if err != nil {
		//handling error
		return
	}

	if len(Player) > 0 && len(Player) == limit {
		lastDataCreatedAt := Player[len(Player)-1].CreatedAt.String()
		lastDataID := Player[len(Player)-1].ID
		//encode last data of this response into last cursor data
		nextCursor, err = helper.DateAndIDToCursor(lastDataCreatedAt, int(lastDataID))
		if err != nil {
			return
		}

	}

	return
}

func (u PlayerUsecase) GetById(ID int) (player player.Player, err error) {
	player, err = u.playerRepository.GetById(ID)
	if err != nil {
		//handling error
		return
	}

	return
}

func (u PlayerUsecase) Save(Player *player.Player) (err error) {
	//check if team is available
	Team, err := u.TeamUsecase.GetById(int(Player.TeamID))
	if err != nil {
		if err == helper.ErrNotFound {
			return helper.BadRequest
		}
		return
	}
	if Team.ID == 0 {
		return helper.BadRequest
	}
	err = u.playerRepository.Save(Player)
	if err != nil {
		return
	}
	return
}

func (u PlayerUsecase) Update(Player *player.Player) (err error) {
	//check if team is already created
	OldPlayer, err := u.GetById(int(Player.ID))
	if err != nil {
		return
	}
	//in case gorm doesnt return error not found
	if OldPlayer.ID == 0 {
		return
	}
	err = u.playerRepository.Update(Player)
	if err != nil {
		//handle error
		return
	}
	return
}

func (u PlayerUsecase) Delete(ID int) (err error) {
	//check team
	OldPlayer, err := u.GetById(ID)
	if err != nil {
		return
	}
	//in case gorm doesnt return error not found
	if OldPlayer.ID == 0 {
		return
	}
	err = u.playerRepository.Delete(ID)
	if err != nil {
		return
	}
	return
}
