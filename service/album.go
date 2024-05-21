package service

import (
	"context"

	"github.com/pulse227/server-recruit-challenge-sample/model"
	"github.com/pulse227/server-recruit-challenge-sample/repository"
)

type AlbumService interface {
	GetAlbumListService(ctx context.Context) ([]*model.Album, error)
	GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.Album, error)
	PostAlbumService(ctx context.Context, album *model.Album) error
	DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error
}

type albumService struct {
	albumRepository repository.IAlbumRepository
	singerRepository repository.ISingerRepository
}

var _ AlbumService = (*albumService)(nil)

func NewAlbumService(albumRepository repository.IAlbumRepository, singerRepository repository.ISingerRepository) *albumService {
	return &albumService{
		albumRepository: albumRepository,
		singerRepository: singerRepository,
	}		
}

func (s *albumService) GetAlbumListService(ctx context.Context) ([]*model.Album, error) {
	albums, err := s.albumRepository.GetAll()
	if err != nil {
		return nil, err
	}
	return albums, nil
}

func (s *albumService) GetAlbumService(ctx context.Context, albumID model.AlbumID) (*model.Album, error) {
	album, err := s.albumRepository.Get(albumID)
	if err != nil {
		return nil, err
	}
	return album, nil
}

func (s *albumService) PostAlbumService(ctx context.Context, album *model.Album) error {
	var singer *model.Singer
	var err error
	if singer, err = s.singerRepository.Get(album.SingerID); err != nil {
		return err
	}

	album.Singer = *singer
	if err := s.albumRepository.Add(album); err != nil {
		return err
	}
	return nil
}

func (s *albumService) DeleteAlbumService(ctx context.Context, albumID model.AlbumID) error {
	if err := s.albumRepository.Delete(albumID); err != nil {
		return err
	}
	return nil
}
