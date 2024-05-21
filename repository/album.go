package repository

import (
	"github.com/pulse227/server-recruit-challenge-sample/model"
	"gorm.io/gorm"
)

type IAlbumRepository interface {
	GetAll() ([]*model.Album, error)
	Get(id model.AlbumID) (*model.Album, error)
	Add(album *model.Album) error
	Delete(id model.AlbumID) error
}

type albumRepository struct {
	db gorm.DB
}

func NewAlbumRepisitory(db gorm.DB) *albumRepository  {
		return &albumRepository{
			db: db,
		}
}

func (sr *albumRepository) GetAll() ([]*model.Album, error) {
	albums := []*model.Album{}
	if err := sr.db.Preload("Singer").Find(&albums).Error; err != nil {
		return []*model.Album{}, err
	}
	return albums, nil
}

func (sr *albumRepository) Get(id model.AlbumID) (*model.Album, error) {
	album := &model.Album{}
	if err := sr.db.Preload("Singer").First(&album, id).Error; err != nil {
		return &model.Album{}, err
	}
	return album, nil
}

func (sr *albumRepository) Add(album *model.Album) error {
	if err := sr.db.Create(album).Error; err != nil {
		return err
	}
	return nil
}

func (sr *albumRepository) Delete(id model.AlbumID) error {
	album := &model.Album{}
	if err := sr.db.Where("id = ?", id).Delete(&album).Error; err != nil {
		return err
	}
	return nil
}