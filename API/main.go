package main

import (
	database "eisandbar/anbox/database"
	ep "eisandbar/anbox/endpoints"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

func main() {
	// Connecting to DB and initializing tables
	database.Repo.Connect()
	database.Repo.InitDB()
	defer database.Repo.Close()

	router := mux.NewRouter()
	setRoutes(router)

	handler := cors.Default().Handler(router)

	fmt.Println("Listening on port:", CONN_PORT)

	log.Fatal(http.ListenAndServe(":"+CONN_PORT, handler))
}

func setRoutes(router *mux.Router) {
	// games
	router.HandleFunc("/games", ep.GamesGetAll).Methods("GET")
	router.HandleFunc("/games", ep.GamesPost).Methods("POST")
	router.HandleFunc("/games/{id:[0-9]+}", ep.GamesGetOne).Methods("GET")
	router.HandleFunc("/games/{id:[0-9]+}", ep.GamesPatch).Methods("PATCH")
	router.HandleFunc("/games/{id:[0-9]+}", ep.GamesDelete).Methods("DELETE")

	// // users
	// router.HandleFunc("/users", ep.UsersGetAll).Methods("GET")
	// router.HandleFunc("/users", ep.UsersPost).Methods("POST")
	// router.HandleFunc("/users/{id:[0-9]+}", ep.UsersGetOne).Methods("GET")
	// router.HandleFunc("/users/{id:[0-9]+}", ep.UsersPatch).Methods("PATCH")
	// router.HandleFunc("/users/{id:[0-9]+}", ep.UsersDelete).Methods("DELETE")

	// //links
	// router.HandleFunc("/links", ep.LinksGetAll).Methods("GET")
	// router.HandleFunc("/links", ep.LinksPost).Methods("POST")
	// router.HandleFunc("/links/{id:[0-9]+}", ep.LinksGetOne).Methods("GET")
	// router.HandleFunc("/links/{id:[0-9]+}", ep.LinksPatch).Methods("PATCH")
	// router.HandleFunc("/links/{id:[0-9]+}", ep.LinksDelete).Methods("DELETE")
}
