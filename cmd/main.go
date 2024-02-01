package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ana-flav/learning-golang.git/data"
	"github.com/ana-flav/learning-golang.git/handlers"
	"github.com/ana-flav/learning-golang.git/repository"
	"github.com/ana-flav/learning-golang.git/service"
	"github.com/gorilla/mux"
)


func main() {

    db, err := data.DataBase()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()  

    fmt.Println("Successfully connected to the database")
    songRepo := repository.NewSongRepository(db)
    songService := service.NewSongService(*songRepo)
    songHandler := handlers.NewSongHandler(songService)
   

    
    router := mux.NewRouter()
    fmt.Println("Server is running on port 8000")   
	router.HandleFunc("/add-song", songHandler.AddSong).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", router))
}