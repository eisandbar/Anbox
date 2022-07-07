package test

import (
	"eisandbar/anbox/typing"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os/exec"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

type Data struct {
	ID   int
	name string
}

type TestSuite struct {
	suite.Suite
	router *mux.Router
}

var connStr = "host=localhost port=5432 user=postgres dbname=test_db password='password' sslmode=disable"

func dbContainerStart() {
	// Starts a postgres container on port 5432 of localhost. We need sleep 2 so that the container can finish startup before we attempt to connect
	cmdStr := "docker run -d -p 5432:5432 --name test_db --rm -e POSTGRES_PASSWORD=password postgres; do sleep 2"
	out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)

}

func dbContainerStop() {
	// Terminates the container to remove traces
	cmdStr := "docker rm -f test_db"
	out, _ := exec.Command("/bin/sh", "-c", cmdStr).Output()
	fmt.Printf("%s", out)
}

func createTestTable() {
	// First we connect to the default database before creating a test_db
	// If we fail to create the test_db we check if it already exists.
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password='password' sslmode=disable")
	capture(err)
	db = db.Exec("CREATE DATABASE test_db;")
	if db.Error != nil {
		fmt.Println("Unable to create DB test_db, attempting to connect assuming it exists...")
		db, err = gorm.Open("postgres", connStr)
		if err != nil {
			fmt.Println("Unable to connect to test_db")
			capture(err)
		}
	}
	defer db.Close()
	capture(err)
}

func testRequest(suite TestSuite, req *http.Request, error_msg string, data interface{}) {
	// Using the router in TestSuite we send request req and confirm the response matches data
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req) // send request

	resp := w.Result()
	respBody, _ := io.ReadAll(resp.Body) // response body

	var response typing.Response
	response.SetResponse(error_msg, data)
	json, _ := json.Marshal(response) // what we expect the response to be

	suite.Equal(json, respBody)
}

func capture(err error) {
	if err != nil {
		log.Fatalf("%s", err)
	}
}
