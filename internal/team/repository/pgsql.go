package repository

import (
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

func (r *TeamPostgreRepository) Fetch(lastDate string, lastID int, search string) (team []team.Team, err error) {
	q := r.DB.Where("created_at > ? and id > ", lastDate, lastID)

	if search != "" {
		q = q.Where("ILIKE %?%", search)
	}

	err = q.Find(&team).Error
	if err != nil {
		return
	}
	return
}

func (r *TeamPostgreRepository) GetById(ID int) (team team.Team, err error) {
	err = r.DB.Where("id = ?").First(&team).Error
	if err != nil {
		return
	}
	return
}

func (r *TeamPostgreRepository) Save(Team *team.Team) (err error) {
	err = r.DB.Model(&Team).Create(Team).Error
	if err != nil {
		return
	}
	return
}

func (r *TeamPostgreRepository) Update(Team *team.Team) (err error) {
	err = r.DB.Model(&Team).Updates(&Team).Error
	if err != nil {
		return
	}
	return
}
func (r *TeamPostgreRepository) Delete(ID int) (err error) {
	err = r.DB.Delete("id = ?", ID).Error
	return
}
