package service

import (
	"errors"
	"fmt"

	"github.com/ana-flav/learning-golang.git/models"
	"github.com/ana-flav/learning-golang.git/repository"
)

type SongService interface {
    AddSong(song models.SongTaylor) error
    GetAllSongs() ([]models.SongTaylor, error)  
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
	fmt.Println("Song to be added: ", song)
	if song.Name == "" || song.LinkSong == "" {
		return errors.New("Name and LinkSong cannot be empty")
	} else if (len(song.Name) >= 150 || len(song.Name) <= 5) || (len(song.LinkSong) >= 300 || len(song.LinkSong) <= 5) {
		return errors.New("Name may be between 150 and 5 characters and link may be between 300 and 3 characters.")
	}

	fmt.Println("Song to be added: ", song.LinkSong, song.Name)
	return s.songRepository.InsertSong(song)
}

func (s *songService) GetAllSongs() ([]models.SongTaylor, error) {
    return s.songRepository.GetAllSongs()
}


