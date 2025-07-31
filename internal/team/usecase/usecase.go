package usecase

import (
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/team"
)

type teamUsecase struct {
	teamRepository team.TeamRepository
}

func NewTeamUsecase(teamRepository team.TeamRepository) team.TeamUsecase {
	return &teamUsecase{
		teamRepository: teamRepository,
	}
}

func (u *teamUsecase) Fetch(cursor string, search string, limit int) (team []team.Team, nextCursor string, err error) {
	//get cursor and decode it
	date, ID, err := helper.CursorToDateAndID(cursor)
	if err != nil {
		//handling error
		return
	}

	team, err = u.teamRepository.Fetch(date, ID, search, limit)
	if err != nil {
		//handling error
		return
	}

	if len(team) > 0 && len(team) == limit {
		lastDataCreatedAt := team[len(team)-1].CreatedAt.String()
		lastDataID := team[len(team)-1].ID
		//encode last data of this response into last cursor data
		nextCursor, err = helper.DateAndIDToCursor(lastDataCreatedAt, int(lastDataID))
		if err != nil {
			return
		}

	}

	return
}

func (u *teamUsecase) GetById(ID int) (team team.Team, err error) {
	team, err = u.teamRepository.GetById(ID)
	if err != nil {
		//handling error
		return
	}

	return
}

func (u *teamUsecase) Save(Team *team.Team) (err error) {
	err = u.teamRepository.Save(Team)
	if err != nil {
		return
	}
	return
}

func (u *teamUsecase) Update(Team *team.Team) (err error) {
	//check if team is already created
	OldTeams, err := u.GetById(int(Team.ID))
	if err != nil {
		return
	}
	//in case gorm doesnt return error not found
	if OldTeams.ID == 0 {
		return
	}
	err = u.teamRepository.Update(Team)
	if err != nil {
		//handle error
		return
	}
	return
}

func (u *teamUsecase) Delete(ID int) (err error) {
	//check team
	OldTeams, err := u.GetById(ID)
	if err != nil {
		return
	}
	//in case gorm doesnt return error not found
	if OldTeams.ID == 0 {
		return
	}
	err = u.teamRepository.Delete(ID)
	if err != nil {
		return
	}
	return
}
