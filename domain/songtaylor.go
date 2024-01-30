package domain 

import "time"

type SongTaylor struct {
	ID int
	Name string
	LinkSong string
	CreatedAt *time.Time
}