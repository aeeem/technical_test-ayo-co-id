package repository

import (
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/team"

	"gorm.io/gorm"
)

type TeamPostgreRepository struct {
	DB *gorm.DB
}

func NewTeamPostgreRepository(db *gorm.DB) team.TeamRepository {
	return &TeamPostgreRepository{
		DB: db,
	}
}

func (r *TeamPostgreRepository) Fetch(lastDate string, lastID int, search string, limit int) (team []team.Team, err error) {
	q := r.DB

	if search != "" {
		q = q.Where("ILIKE %?%", search)
	}
	if lastDate != "" {
		q = q.Where("created_at > ?", lastDate)
	}
	if lastID != 0 {
		q = q.Where("id > ?", lastID)
	}
	err = q.Limit(limit).Find(&team).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}

func (r *TeamPostgreRepository) GetById(ID int) (team team.Team, err error) {
	err = r.DB.Where("id = ?", ID).First(&team).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}

func (r *TeamPostgreRepository) Save(Team *team.Team) (err error) {
	err = r.DB.Model(&Team).Create(Team).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}

func (r *TeamPostgreRepository) Update(Team *team.Team) (err error) {
	err = r.DB.Model(&Team).Updates(&Team).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}
func (r *TeamPostgreRepository) Delete(ID int) (err error) {
	err = r.DB.Model(&team.Team{}).Delete("id = ?", ID).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}
