package repository

import (
	"database/sql"
	"fmt"

	"github.com/ana-flav/learning-golang.git/models"
	"github.com/google/uuid"
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
	_, err := sr.db.Exec("INSERT INTO song (name, link) VALUES ($1, $2)", song.Name, song.LinkSong)
	return err
}
func (sr *SongRepository) GetAllSongs() ([]models.SongTaylor, error) {
    results, err := sr.db.Query("SELECT * FROM song")
	fmt.Println(err);
    if err != nil {
        return nil, err
    }
    defer results.Close()

    var songs []models.SongTaylor

    for results.Next() {
        var song models.SongTaylor
        err := results.Scan(&song.ID, &song.LinkSong, &song.Name)
		fmt.Println("teste", err); 
        if err != nil {
            return nil, err
        }
		fmt.Println(song);
        songs = append(songs, song)
    }

    if err := results.Err(); err != nil {
        return nil, err
    }

    return songs, nil
}

func (sr *SongRepository) DeleteSong(IdSong uuid.UUID) error {
	fmt.Println("teste")
	_, err := sr.db.Query("DELETE FROM song WHERE id_song = $1", IdSong);
	fmt.Println(err)
	return err
}


func (sr *SongRepository) UpdateSong(idSong uuid.UUID, updatedSong models.SongTaylor) error {
	fmt.Println(updatedSong.Name, updatedSong.LinkSong, idSong)
	_, err := sr.db.Exec("UPDATE song SET name = $1, link = $2 WHERE id_song= $3", updatedSong.Name, updatedSong.LinkSong, idSong)
	fmt.Println(err)
	return err
}

func (sr *SongRepository) GetSongById(idSong uuid.UUID) (models.SongTaylor, error) {
    var song models.SongTaylor

   
    err := sr.db.QueryRow("SELECT id_song, name, link FROM song WHERE id_song = $1", idSong).
        Scan(&song.ID, &song.Name, &song.LinkSong)

    if err != nil {
        return models.SongTaylor{}, err
    }

    return song, nil
}

