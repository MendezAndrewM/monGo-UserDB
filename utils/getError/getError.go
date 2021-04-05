package geterror

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mendezandrewm/monGo-UserDB/models"
)

func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = models.ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
