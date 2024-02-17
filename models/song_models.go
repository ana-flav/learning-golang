package models 
import (
	"github.com/google/uuid"
)

type SongTaylor struct {
	ID 		 uuid.UUID
	Name 	 string
	LinkSong string
}