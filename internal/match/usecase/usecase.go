package usecase

import (
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/match"
	"technical_test-ayo-co-id/internal/player"
	"technical_test-ayo-co-id/internal/score"
	"technical_test-ayo-co-id/internal/team"

	"github.com/rs/zerolog/log"
)

type matchUsecase struct {
	matchRepository  match.MatchRepository
	playerRepository player.PlayerRepository
	teamUsecase      team.TeamUsecase
	ScoreUsecase     score.ScoreUsecase
}

func NewMatchUsecase(matchrepository match.MatchRepository, playerRepository player.PlayerRepository, teamUsecase team.TeamUsecase, ScoreUsecase score.ScoreUsecase) match.MatchUsecase {
	return &matchUsecase{
		playerRepository: playerRepository,
		ScoreUsecase:     ScoreUsecase,
		teamUsecase:      teamUsecase,
		matchRepository:  matchrepository,
	}
}

func (u *matchUsecase) Fetch(AwayTeamID uint, HomeTeamID uint,
	mvp uint,
	cursor string,
	search string,
	limit int) (match []match.Match, nextCursor string, err error) {
	date, ID, err := helper.CursorToDateAndID(cursor)
	if err != nil {
		//handling error
		return
	}

	match, err = u.matchRepository.Fetch(AwayTeamID, HomeTeamID, mvp, date, ID, search, limit)
	if err != nil {
		return
	}

	if len(match) > 0 && len(match) == limit {
		lastDataCreatedAt := match[len(match)-1].CreatedAt.String()
		lastDataID := match[len(match)-1].ID
		//encode last data of this response into last cursor data
		nextCursor, err = helper.DateAndIDToCursor(lastDataCreatedAt, int(lastDataID))
		if err != nil {
			return
		}

	}
	return
}

func (u *matchUsecase) GetById(ID int) (match match.Match, err error) {
	match, err = u.matchRepository.GetById(ID)
	if err != nil {
		//handling error
		return
	}

	//get away team details
	AwayTeamDetails, err := u.teamUsecase.GetById(int(match.AwayTeam))
	if err != nil {
		//handling error
		return
	}
	if AwayTeamDetails.ID != 0 {
		match.AwayTeamDetail = &AwayTeamDetails
	}

	//get home team details
	HomeTeamDetail, err := u.teamUsecase.GetById(int(match.HomeTeam))
	if err != nil {
		//handling error
		return
	}
	if HomeTeamDetail.ID != 0 {
		match.HomeTeamDetail = &HomeTeamDetail
	}

	//get match score
	match.MatchScore, err = u.ScoreUsecase.FetchScoreByMatchID(int(match.ID))
	if err != nil {
		return
	}
	//get away score
	match.ScoreAway, err = u.ScoreUsecase.FetchScoreTeamByMatchID(int(match.ID), match.AwayTeam)
	if err != nil {
		return
	}
	//get home score
	match.ScoreHome, err = u.ScoreUsecase.FetchScoreTeamByMatchID(int(match.ID), match.HomeTeam)
	if err != nil {
		return
	}

	TotalScoreHome := len(match.ScoreHome)
	if !match.TotalScoreHome.Valid {
		match.TotalScoreHome.Scan(int32(TotalScoreHome))
	}

	TotalScoreAway := len(match.ScoreAway)
	if !match.TotalScoreAway.Valid {
		match.TotalScoreAway.Scan(int32(TotalScoreAway))
	}

	if !match.Winner.Valid {
		if TotalScoreAway < TotalScoreHome {
			err = match.Winner.Scan(int32(match.HomeTeam))
			if err != nil {
				log.Info().Err(err).Msg("Error decide winner for home team")
				err = helper.InternalServerErr
				return
			}
		} else if TotalScoreAway > TotalScoreHome {
			err = match.Winner.Scan(int32(match.AwayTeam))
			if err != nil {
				log.Info().Err(err).Msg("Error decide winner for away team")
				err = helper.InternalServerErr
				return
			}
		} else {
			err = match.Winner.Scan(int32(0))
			if err != nil {
				log.Info().Err(err).Msg("Error decide winner for draw")
				err = helper.InternalServerErr
				return
			}
		}
	}

	//get mvp
	if match.PlayerMvpID.Valid {

		PlayerMVP, err := u.playerRepository.GetById(int(match.PlayerMvpID.Int32))
		if err != nil {
			err = helper.InternalServerErr
			return match, err
		}
		match.PlayerMVP = &PlayerMVP
	}
	return
}
func (u *matchUsecase) Save(match *match.Match) (err error) {
	err = u.matchRepository.Save(match)
	if err != nil {
		return
	}
	return
}
func (u *matchUsecase) Update(Match *match.Match) (err error) {
	//check if team is already created
	Oldmatch, err := u.GetById(int(Match.ID))
	if err != nil {
		return
	}
	//in case gorm doesnt return error not found
	if Oldmatch.ID == 0 {
		return
	}
	err = u.matchRepository.Update(Match)
	if err != nil {
		//handle error
		return
	}
	return
}
func (u *matchUsecase) Delete(ID int) (err error) {
	OldPlayer, err := u.GetById(ID)
	if err != nil {
		return
	}
	//in case gorm doesnt return error not found
	if OldPlayer.ID == 0 {
		return
	}
	err = u.matchRepository.Delete(ID)
	if err != nil {
		return
	}
	return
}
