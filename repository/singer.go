package repository

import (
	"github.com/pulse227/server-recruit-challenge-sample/model"
	"gorm.io/gorm"
)

type ISingerRepository interface {
	GetAll() ([]*model.Singer, error)
	Get(id model.SingerID) (*model.Singer, error)
	Add(singer *model.Singer) error
	Delete(id model.SingerID) error
}

type singerRepository struct {
	db gorm.DB
}

func NewSingerRepisitory(db gorm.DB) *singerRepository  {
		return &singerRepository{
			db: db,
		}
}

func (sr *singerRepository) GetAll() ([]*model.Singer, error) {
	singers := []*model.Singer{}
	if err := sr.db.Find(&singers).Error; err != nil {
		return []*model.Singer{}, err
	}
	return singers, nil
}

func (sr *singerRepository) Get(id model.SingerID) (*model.Singer, error) {
	singer := &model.Singer{}
	if err := sr.db.First(&singer, id).Error; err != nil {
		return &model.Singer{}, err
	}
	return singer, nil
}

func (sr *singerRepository) Add(singer *model.Singer) error {
	if err := sr.db.Create(singer).Error; err != nil {
		return err
	}
	return nil
}

func (sr *singerRepository) Delete(id model.SingerID) error {
	singer := &model.Singer{}
	if err := sr.db.Where("id = ?", id).Delete(&singer).Error; err != nil {
		return err
	}
	return nil
}