package repository

import (
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/match"

	"gorm.io/gorm"
)

type matchRepository struct {
	DB *gorm.DB
}

func NewMatchRepository(DB *gorm.DB) match.MatchRepository {
	return &matchRepository{DB: DB}
}

func (r *matchRepository) Fetch(AwayTeamID uint,
	HomeTeamID uint,
	mvp uint,
	lastDate string,
	lastId int,
	search string, limit int) (Match []match.Match, err error) {
	q := r.DB.Model(&match.Match{})
	if HomeTeamID != 0 {
		q = r.DB.Where("home_team_id = ?", HomeTeamID)
	}

	if AwayTeamID != 0 {
		q = r.DB.Where("away_team_id = ?", AwayTeamID)
	}

	if mvp != 0 {
		q = r.DB.Where("player_mvp_id = ?", mvp)
	}

	if lastDate != "" {
		q = r.DB.Where("created_at > ?", lastDate)
	}

	if lastId != 0 {
		q = r.DB.Where("id > ?", lastId)
	}

	if search != "" {
		q = q.Where("team_winner_name ILIKE %?%", search)
	}

	err = q.Limit(limit).
		Find(&Match).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}

func (r *matchRepository) GetById(ID int) (match match.Match, err error) {
	err = r.DB.Where("id = ?", ID).First(&match).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}

func (r *matchRepository) Save(Match *match.Match) (err error) {
	err = r.DB.Model(&match.Match{}).Create(&Match).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}

func (r *matchRepository) Update(match *match.Match) (err error) {
	err = r.DB.Model(&match).Updates(&match).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}

func (r *matchRepository) Delete(ID int) (err error) {
	err = r.DB.Model(&match.Match{}).Delete("id = ?", ID).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}
