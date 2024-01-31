package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/ana-flav/learning-golang.git/handlers"
    "github.com/ana-flav/learning-golang.git/repository"
    "github.com/ana-flav/learning-golang.git/service"
    "github.com/ana-flav/learning-golang.git/data"
)


func main() {

    db, err := data.DataBase()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()  

    songRepo := repository.NewSongRepository(db)

    
    router := mux.NewRouter()
	router.HandleFunc("/add-song/", AddSong).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", router))
}