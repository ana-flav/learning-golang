package service

import (
	"errors"
	"fmt"

	"github.com/ana-flav/learning-golang.git/models"
	"github.com/ana-flav/learning-golang.git/repository"
	"github.com/google/uuid"
)

type SongService interface {
    AddSongService(song models.SongTaylor) error
    GetAllSongsService() ([]models.SongTaylor, error)  
	DeleteSongService(IdSong uuid.UUID ) error
	UpdateSongService(idSong uuid.UUID, updatedSong models.SongTaylor) error
	GetSongByIdService(idSong uuid.UUID) (models.SongTaylor, error)
}


type songService struct {
	songRepository repository.SongRepository 
}

func NewSongService(songRepository repository.SongRepository) *songService {
	return &songService{
		songRepository: songRepository,
	}
}

func (s *songService) AddSongService(song models.SongTaylor) error {
	fmt.Println("Song to be added: ", song)
	if song.Name == "" || song.LinkSong == "" {
		return errors.New("Name and LinkSong cannot be empty")
	} else if (len(song.Name) >= 150 || len(song.Name) <= 5) || (len(song.LinkSong) >= 300 || len(song.LinkSong) <= 5) {
		return errors.New("Name may be between 150 and 5 characters and link may be between 300 and 3 characters.")
	}

	fmt.Println("Song to be added: ", song.LinkSong, song.Name)
	return s.songRepository.InsertSong(song)
}

func (s *songService) GetAllSongsService() ([]models.SongTaylor, error) {
	fmt.Println("teste");
    return s.songRepository.GetAllSongs()
}

func (s *songService) DeleteSongService(IdSong uuid.UUID ) error{
	return s.songRepository.DeleteSong(IdSong)
}

func (s *songService) UpdateSongService(idSong uuid.UUID, updatedSong models.SongTaylor) error {

	if updatedSong.Name == "" || updatedSong.LinkSong == "" {
		return errors.New("Name and LinkupdatedSong cannot be empty")
	} else if (len(updatedSong.Name) >= 150 || len(updatedSong.Name) <= 5) || (len(updatedSong.LinkSong) >= 300 || len(updatedSong.LinkSong) <= 5) {
		return errors.New("Name may be between 150 and 5 characters and link may be between 300 and 3 characters.")
	}
	return s.songRepository.UpdateSong(idSong, updatedSong)
}

func (s *songService) GetSongByIdService(idSong uuid.UUID) (models.SongTaylor, error){
	return s.songRepository.GetSongById(idSong)
}


