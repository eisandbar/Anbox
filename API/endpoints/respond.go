package endpoints

import (
	typing "eisandbar/anbox/typing"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func respond(w http.ResponseWriter, error_msg string, data interface{}) {
	var response typing.Response
	response.SetResponse(error_msg, data)

	json, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error converting to json")
	}
	io.WriteString(w, string(json))
}
