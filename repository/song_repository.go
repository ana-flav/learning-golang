package repository

import (
	"database/sql"
	"github.com/ana-flav/learning-golang.git/models"
)

type SongRepository struct {
	db *sql.DB
}


func NewSongRepository(db *sql.DB) *SongRepository {
	// Construtor que cria instancia do tipo SongRepository e inicializa o campo db como o objeto
	// sql.DB
	return &SongRepository{
		db: db,
	}
}

func (sr *SongRepository) InsertSong(song models.SongTaylor) error {
	_, err := sr.db.Exec("INSERT INTO songs (name, link) VALUES ($1, $2)", song.Name, song.LinkSong)
	return err
}