package endpoints

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"

	database "eisandbar/anbox/database"
	typing "eisandbar/anbox/typing"
)

func GamesGetAll(w http.ResponseWriter, r *http.Request) {
	// Get all game records that match filter
	var games []typing.Game
	var filter typing.Game

	err := schema.NewDecoder().Decode(&filter, r.URL.Query())
	if err != nil {
		respond(w, "Bad Request", nil)
		return
	}

	error_msg := database.Repo.GetAll(&games, filter)
	respond(w, error_msg, games)
}

func GamesPost(w http.ResponseWriter, r *http.Request) {
	// Create new game record
	var game typing.Game

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&game)
	if err != nil {
		respond(w, "Bad Request", nil)
		return
	}

	error_msg := database.Repo.Post(&game)
	respond(w, error_msg, nil)
}

func GamesGetOne(w http.ResponseWriter, r *http.Request) {
	// Find game record by id
	var game typing.Game

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	error_msg := database.Repo.GetOne(&game, id)
	respond(w, error_msg, game)
}

func GamesPatch(w http.ResponseWriter, r *http.Request) {
	// Update game record by id
	var game typing.Game

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&game)
	if err != nil {
		respond(w, "Bad Request", nil)
		return
	}

	error_msg := database.Repo.Patch(&typing.Game{}, game, id)
	respond(w, error_msg, nil)
}

func GamesDelete(w http.ResponseWriter, r *http.Request) {
	// Delete game record by id
	var game typing.Game

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	error_msg := database.Repo.Delete(&game, id)
	respond(w, error_msg, nil)
}
