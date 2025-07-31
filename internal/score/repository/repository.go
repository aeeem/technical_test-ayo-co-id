package repository

import (
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/match"
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
func (r *scoreRepository) Save(Score *score.Score, ishome bool) (err error) {
	err = r.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(&Score).Create(&Score).Error
		if err != nil {
			tx.Rollback()
			err = helper.CheckIfErrFromDbToStatusCode(err)
			return err
		}
		//checking if score for away or for home team
		if ishome {
			err = tx.Model(&match.Match{}).Where("id = ?", Score.MatchID).
				Update("total_score_away", gorm.Expr("Coalesce(total_score_away,0) + ?", 1)).Error
			if err != nil {
				tx.Rollback()
				err = helper.CheckIfErrFromDbToStatusCode(err)
				return err
			}
		} else {
			err = tx.Model(&match.Match{}).Where("id = ?", Score.MatchID).
				Update("total_score_home", gorm.Expr("Coalesce(total_score_home,0) + ?", 1)).Error
			if err != nil {
				tx.Rollback()
				err = helper.CheckIfErrFromDbToStatusCode(err)
				return err
			}
		}
		return err
	})
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
