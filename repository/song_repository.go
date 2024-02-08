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

func (sr *SongRepository) GetAllSongs() ([]models.SongTaylor, error){
	results, err := sr.db.Query("SELECT * FROM songs")

	if err != nil{
		return nil, err
	}
	defer results.Close()

	var songs = make([]models.SongTaylor, 0)

	for results.Next(){
		var song models.SongTaylor
		err = results.Scan(&song.ID, &song.LinkSong, &song.Name)
		if err != nil{
			return nil, err
		}

		songs = append(songs, song)
	}
	return songs, nil
}