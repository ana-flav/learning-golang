package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "data"
    "service"
)


func main() {

    db, err := data.DataBase()
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()  

    
    router := mux.NewRouter()
	router.HandleFunc("/add-song/", AddSong).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", router))
}