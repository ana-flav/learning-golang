package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)


func main() {
    router := mux.NewRouter()
	router.HandleFunc("/add-song/", AddSong).Methods("POST")
    log.Fatal(http.ListenAndServe(":8000", router))
}