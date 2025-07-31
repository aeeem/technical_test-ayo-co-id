package repository

import (
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/score"

	"gorm.io/gorm"
)

type scoreRepository struct {
	DB *gorm.DB
}

func NewScoreRepository(db *gorm.DB) score.ScoreRepository {
	return &scoreRepository{
		DB: db,
	}
}

func (r *scoreRepository) FetchScoreByMatchID(MatchID int) (Score []score.Score, err error) {
	err = r.DB.Where("match_id = ?", MatchID).Find(&Score).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}

func (r *scoreRepository) GetByID(ID uint) (Score score.Score, err error) {
	err = r.DB.Where("id = ?", ID).First(&Score).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}

func (r *scoreRepository) FetchScoreTeamByMatchID(MatchID int, teamID uint) (Score []score.Score, err error) {
	err = r.DB.Where("match_id = ? AND team_id = ?", MatchID, teamID).Find(&Score).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}
func (r *scoreRepository) Save(Score *score.Score) (err error) {
	err = r.DB.Model(&Score).Create(&Score).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return

}
func (r *scoreRepository) Update(Score *score.Score) (err error) {
	err = r.DB.Model(&Score).Updates(&Score).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}
func (r *scoreRepository) Delete(ID int) (err error) {
	err = r.DB.Model(&score.Score{}).Delete("id = ?", ID).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}
