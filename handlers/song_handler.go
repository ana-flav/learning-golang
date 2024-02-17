package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ana-flav/learning-golang.git/models"
	"github.com/ana-flav/learning-golang.git/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type SongService interface {
	AddSongHandler(http.ResponseWriter, *http.Request)
	GetSongHandler(http.ResponseWriter, *http.Request)
	DeleteSongHandler(http.ResponseWriter, *http.Request)
	UpdateSongHandler(http.ResponseWriter, *http.Request)
	
}
type SongHandler struct {
	songService service.SongService
}

func NewSongHandler(songService service.SongService) *SongHandler {
	return &SongHandler{
		songService: songService,
	}
}

func (sh *SongHandler) AddSongHandler(w http.ResponseWriter, r *http.Request) {

	var newSongTaylor models.SongTaylor

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newSongTaylor)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return 
	}
	fmt.Println("New song: ", newSongTaylor)


	
	err = sh.songService.AddSongService(newSongTaylor)
	if err != nil {
		fmt.Println("Error adding song: ", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Song added successfully")
	fmt.Println("Song added successfully: ", http.StatusCreated	)
	w.WriteHeader(http.StatusCreated)

}

func (sh *SongHandler) GetSongHandler(w http.ResponseWriter, r *http.Request){
	var allSongsTaylor = make([]models.SongTaylor, 0)
	var err error
	allSongsTaylor, err = sh.songService.GetAllSongsService()

	fmt.Println(allSongsTaylor);
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

func (sh *SongHandler) DeleteSongHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idString, ok := vars["IdSong"]
    if !ok {
        http.Error(w, "ID not provided in URL", http.StatusBadRequest)
        return
    }

    idSong, err := uuid.Parse(idString)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    err = sh.songService.DeleteSongService(idSong)
    if err != nil {
        http.Error(w, "Error deleting the song", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func (sh *SongHandler) UpdateSongHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idString, ok := vars["IdSong"]
    if !ok {
        http.Error(w, "ID not provided in URL", http.StatusBadRequest)
        return
    }

    idSong, err := uuid.Parse(idString)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    var updatedSong models.SongTaylor
    decoder := json.NewDecoder(r.Body)
    err = decoder.Decode(&updatedSong)
    if err != nil {
        http.Error(w, "Error decoding JSON", http.StatusBadRequest)
        return
    }

    err = sh.songService.UpdateSongService(idSong, updatedSong)
	fmt.Println(err)
    if err != nil {
        http.Error(w, "Error updating the song", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func (sh *SongHandler) GetSongByIdHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idString, ok := vars["IdSong"]
    if !ok {
        http.Error(w, "ID not provided in URL", http.StatusBadRequest)
        return
    }

    idSong, err := uuid.Parse(idString)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    song, err := sh.songService.GetSongByIdService(idSong)
    if err != nil {
        http.Error(w, "Error retrieving the song", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(song)
}












