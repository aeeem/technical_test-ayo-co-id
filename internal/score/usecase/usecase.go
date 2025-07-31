package usecase

import (
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/match"
	"technical_test-ayo-co-id/internal/score"

	"github.com/rs/zerolog/log"
)

type scoreUsecase struct {
	Matchrepository match.MatchRepository
	ScoreRepository score.ScoreRepository
}

func NewScoreUsecase(Matchrepository match.MatchRepository, ScoreRepository score.ScoreRepository) score.ScoreUsecase {
	return &scoreUsecase{
		Matchrepository: Matchrepository,
		ScoreRepository: ScoreRepository,
	}
}

func (u *scoreUsecase) FetchScoreByMatchID(MatchID int) (Score []score.Score, err error) {
	Score, err = u.ScoreRepository.FetchScoreByMatchID(MatchID)
	if err != nil {
		return
	}
	log.Info().Any("Score", Score).Msg("score")
	return
}
func (u *scoreUsecase) FetchScoreTeamByMatchID(MatchID int, teamID uint) (Score []score.Score, err error) {
	Score, err = u.ScoreRepository.FetchScoreTeamByMatchID(MatchID, teamID)
	if err != nil {
		return
	}
	log.Info().Any("Score", Score).Msg("score")
	return
}
func (u *scoreUsecase) Save(score *score.Score) (err error) {
	//check if match is available
	Match, err := u.Matchrepository.GetById(int(score.MatchID))
	if err != nil {
		return
	}

	//check if player is available
	Player, err := u.Matchrepository.GetById(int(score.PlayerID))
	if err != nil {
		return
	}
	if Player.ID != 0 && Match.ID != 0 {
		err = u.ScoreRepository.Save(score)
		if err != nil {
			return
		}
	} else {
		err = helper.ErrNotFound
		return
	}
	return
}
func (u *scoreUsecase) Update(score *score.Score) (err error) {
	// check if match is available
	Match, err := u.Matchrepository.GetById(int(score.MatchID))
	if err != nil {
		return
	}
	//check if score is already created
	OldScore, err := u.ScoreRepository.GetByID(score.ID)
	if err != nil {
		return
	}
	// check if player is available
	Player, err := u.Matchrepository.GetById(int(score.PlayerID))
	if err != nil {
		return
	}
	if Player.ID != 0 && Match.ID != 0 && OldScore.ID != 0 {
		err = u.ScoreRepository.Update(score)
		if err != nil {
			return
		}
	} else {
		err = helper.ErrNotFound
		return
	}
	return
}

func (u *scoreUsecase) Delete(ID int) (err error) {
	err = u.ScoreRepository.Delete(ID)
	if err != nil {
		return
	}
	return
}
