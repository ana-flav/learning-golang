package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ana-flav/learning-golang.git/models"
	"github.com/ana-flav/learning-golang.git/service"
)

type SongService interface {
	AddSong(http.ResponseWriter, *http.Request)
	GetSong(http.ResponseWriter, *http.Request)
	
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
	fmt.Println("New song: ", newSongTaylor)


	
	err = sh.songService.AddSong(newSongTaylor)
	if err != nil {
		fmt.Println("Error adding song: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Song added successfully")
	fmt.Println("Song added successfully: ", http.StatusCreated	)
	w.WriteHeader(http.StatusCreated)

}

func (sh *SongHandler) GetSong(w http.ResponseWriter, r *http.Request){
	var allSongsTaylor = make([]models.SongTaylor, 0)
	var err error
	allSongsTaylor, err = sh.songService.GetAllSongs()

	if err != nil {
        http.Error(w, "Erro ao obter todas as músicas", http.StatusInternalServerError)
        return
    }
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(allSongsTaylor); err != nil {
        http.Error(w, "Erro ao codificar a resposta JSON", http.StatusInternalServerError)
        return
    }
}