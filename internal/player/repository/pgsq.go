package repostory

import (
	"technical_test-ayo-co-id/internal/helper"
	"technical_test-ayo-co-id/internal/player"

	"gorm.io/gorm"
)

type playerRepository struct {
	DB *gorm.DB
}

func NewPlayerRepository(DB *gorm.DB) player.PlayerRepository {
	return &playerRepository{
		DB: DB,
	}
}

func (r *playerRepository) Fetch(teamID uint, lastDate string, lastID int, search string, limit int) (Player []player.Player, err error) {
	q := r.DB
	if teamID != 0 {
		q = q.Where("team_id = ?", teamID)
	}

	if lastDate != "" {
		q = q.Where("created_at > ?", lastDate)
	}

	if lastID != 0 {
		q = q.Where("id > ?", lastID)
	}
	err = q.Limit(limit).Find(&Player).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}

func (r *playerRepository) GetById(ID int) (Player player.Player, err error) {
	err = r.DB.Where("id = ?", ID).First(&Player).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}

func (r *playerRepository) Save(Player *player.Player) (err error) {
	err = r.DB.Model(&Player).Create(&Player).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}

func (r *playerRepository) Update(Player *player.Player) (err error) {
	err = r.DB.Model(&Player).Updates(&Player).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}

func (r *playerRepository) Delete(ID int) (err error) {
	err = r.DB.Model(&player.Player{}).Delete("id = ?", ID).Error
	if err != nil {
		err = helper.CheckIfErrFromDbToStatusCode(err)
		return
	}
	return
}
