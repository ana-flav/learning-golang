package handlers	

import (
	"encoding/json"
	"net/http"
	"service"
)

func AddSong(response http.ResponseWriter, request *http.Request) {

	// jsonResponse := service.AddSong("nome", "LinkSong")

	// if jsonResponse == nil {
	// 	returnError(response, http.StatusInternalServerError, "Internal Server Error")
	// } else {	
	// 	returnResponse(response, http.StatusOK, jsonResponse)
	// }

}