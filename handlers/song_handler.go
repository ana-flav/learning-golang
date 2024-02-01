package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/ana-flav/learning-golang.git/models"
	"github.com/ana-flav/learning-golang.git/service"
)

type SongService interface {
	AddSong(http.ResponseWriter, *http.Request)
	
}
type SongHandler struct {
	songService service.SongService
}

func NewSongHandler(songService service.SongService) *SongHandler {
	return &SongHandler{
		songService: songService,
	}
}

func (sh *SongHandler) AddSong(w http.ResponseWriter, r *http.Request) {

	var newSongTaylor models.SongTaylor

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newSongTaylor)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return 
	}

	err = sh.songService.AddSong(newSongTaylor)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}