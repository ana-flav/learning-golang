package service

import (
	"errors"

	"github.com/ana-flav/learning-golang.git/models"
	"github.com/ana-flav/learning-golang.git/repository"
)

type SongService interface {
	AddSong(song models.SongTaylor) error
}

type songService struct {
	songRepository repository.SongRepository 
}

func NewSongService(songRepository repository.SongRepository) *songService {
	return &songService{
		songRepository: songRepository,
	}
}

func (s *songService) AddSong(song models.SongTaylor) error {
	if song.Name == "" || song.LinkSong == "" {
		return errors.New("Name and LinkSong cannot be empty")
	} else if (len(song.Name) >= 150 || len(song.Name) <= 5) || (len(song.LinkSong) >= 300 || len(song.LinkSong) <= 5) {
		return errors.New("Name may be between 150 and 5 characters and link may be between 300 and 3 characters.")
	}

	return s.songRepository.InsertSong(song)
}
