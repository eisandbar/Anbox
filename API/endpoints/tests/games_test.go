package test

import (
	"eisandbar/anbox/database"
	ep "eisandbar/anbox/endpoints"
	"eisandbar/anbox/typing"
	"encoding/json"
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
)

/*
These tests confirm that endpoints/games endpoints behave as we expect.
The test suite starts a postgres container before all tests and cleans tables before each test.
*/

var game = typing.Game{
	ID:          1,
	Title:       "Runner",
	Description: "A Game",
	Url:         "https://www.studio.com/runner",
	AgeRating:   6,
	Publisher:   "studio",
}

type GamesTestSuite struct {
	TestSuite
}

func (suite *GamesTestSuite) SetupSuite() {

	// Create the router
	router := mux.NewRouter()
	router.HandleFunc("/games", ep.GamesGetAll).Methods("GET")
	router.HandleFunc("/games", ep.GamesPost).Methods("POST")
	router.HandleFunc("/games/{id:[0-9]+}", ep.GamesGetOne).Methods("GET")
	router.HandleFunc("/games/{id:[0-9]+}", ep.GamesPatch).Methods("PATCH")
	router.HandleFunc("/games/{id:[0-9]+}", ep.GamesDelete).Methods("DELETE")
	suite.TestSuite.router = router

	// Start postgres container
	suite.T().Log("Creating container")
	dbContainerStart()
	createTestTable()
}

func (suite *GamesTestSuite) TearDownSuite() {
	// Stop and remove container
	dbContainerStop()
}

func (suite *GamesTestSuite) SetupTest() {
	db, err := gorm.Open("postgres", connStr)
	capture(err)
	defer db.Close()

	// Start test with empty table
	db.DropTable(&typing.Game{})
	db.CreateTable(&typing.Game{})

	// Connects Repo to the test DB
	database.Repo.Connect(connStr)
}

func (suite *GamesTestSuite) AfterTest(_, _ string) {
	database.Repo.Close()
}

func (suite *GamesTestSuite) TestGamesGetEmpty() {
	// Test expected response
	req := httptest.NewRequest("GET", "/games", nil)
	testRequest(suite.TestSuite, req, "", []interface{}{})
}

func (suite *GamesTestSuite) TestGamesGetGood() {
	// Add data to db
	database.Repo.Post(&game)

	// Test expected response
	req := httptest.NewRequest("GET", "/games", nil)
	testRequest(suite.TestSuite, req, "", []interface{}{game})
}

func (suite *GamesTestSuite) TestGamesGetGoodFilter() {
	// Add data to db
	database.Repo.Post(&game)
	database.Repo.Post(&typing.Game{})

	// Test expected response
	req := httptest.NewRequest("GET", "/games?publisher=studio", nil)
	testRequest(suite.TestSuite, req, "", []interface{}{game})
}
func (suite *GamesTestSuite) TestGamesGetBadFilter() {
	// Can't filter searches by title
	// Add data to db
	database.Repo.Post(&game)
	database.Repo.Post(&typing.Game{})

	// Test expected response
	req := httptest.NewRequest("GET", "/games?title=Runner", nil)
	testRequest(suite.TestSuite, req, "Bad Request", nil)
}

func (suite *GamesTestSuite) TestGamesGetWrongStruct() {
	// Games has no field name
	// Add data to db
	database.Repo.Post(&game)

	// Test expected response
	req := httptest.NewRequest("GET", "/games?name=John", nil)
	testRequest(suite.TestSuite, req, "Bad Request", nil)
}

func (suite *GamesTestSuite) TestGamesPostGood() {
	// Test expected response
	reqBody, _ := json.Marshal(game)
	req := httptest.NewRequest("POST", "/games", strings.NewReader(string(reqBody)))
	testRequest(suite.TestSuite, req, "", nil)

	// Test data creation
	res := typing.Game{}
	database.Repo.GetOne(&res, 1)
	suite.Equal(game, res)
}

func (suite *GamesTestSuite) TestGamesPostBadRequest() {
	// Creating bad struct
	data := Data{name: "John"}

	// Test expected response
	reqBody, _ := json.Marshal(data)
	req := httptest.NewRequest("POST", "/games", strings.NewReader(string(reqBody)))
	testRequest(suite.TestSuite, req, "Bad Request", nil)

	// Test no data creation
	res := typing.Game{}
	database.Repo.GetOne(&res, 1)
	suite.Equal(typing.Game{}, res)
}

func (suite *GamesTestSuite) TestGamesPostNoTable() {
	// Deleting table
	db, err := gorm.Open("postgres", connStr)
	capture(err)
	defer db.Close()
	db.DropTable(&typing.Game{})

	// Test expected response
	reqBody, _ := json.Marshal(game)
	req := httptest.NewRequest("POST", "/games", strings.NewReader(string(reqBody)))
	testRequest(suite.TestSuite, req, "Error Creating Record", nil)
}

func (suite *GamesTestSuite) TestGamesPostIdIncrement() {
	// Making sure you can't set Id with post
	// Creating bad data
	data := game
	data.ID = 2

	// Test expected response
	reqBody, _ := json.Marshal(data)
	req := httptest.NewRequest("POST", "/games", strings.NewReader(string(reqBody)))
	testRequest(suite.TestSuite, req, "", nil)

	// Test data creation
	res := typing.Game{}
	database.Repo.GetOne(&res, 1)
	suite.Equal(game, res)
}

func (suite *GamesTestSuite) TestGamesIdGetGood() {
	// Add data to db
	database.Repo.Post(&game)

	// Test expected response
	req := httptest.NewRequest("GET", "/games/1", nil)
	testRequest(suite.TestSuite, req, "", game)
}

func (suite *GamesTestSuite) TestGamesIdGetNotFound() {
	// Test expected response
	req := httptest.NewRequest("GET", "/games/1", nil)
	testRequest(suite.TestSuite, req, "Record Not Found", game)
}

func (suite *GamesTestSuite) TestGamesIdPatchGood() {
	// Add data to db
	database.Repo.Post(&typing.Game{})

	// Test expected response
	reqBody, _ := json.Marshal(game)
	req := httptest.NewRequest("PATCH", "/games/1", strings.NewReader(string(reqBody)))
	testRequest(suite.TestSuite, req, "", nil)

	// Test data update
	res := typing.Game{}
	database.Repo.GetOne(&res, 1)
	suite.Equal(game, res)
}

func (suite *GamesTestSuite) TestGamesIdPatchBadRequest() {
	// Add data to db
	database.Repo.Post(&typing.Game{})

	// Creating bad struct
	data := Data{name: "John"}

	// Test expected response
	reqBody, _ := json.Marshal(data)
	req := httptest.NewRequest("PATCH", "/games/1", strings.NewReader(string(reqBody)))
	testRequest(suite.TestSuite, req, "Bad Request", nil)

	// Test data update
	res := typing.Game{}
	database.Repo.GetOne(&res, 1)
	suite.Equal(typing.Game{ID: 1}, res)
}

func (suite *GamesTestSuite) TestGamesIdPatchNotFound() {
	// Test expected response
	reqBody, _ := json.Marshal(game)
	req := httptest.NewRequest("PATCH", "/games/1", strings.NewReader(string(reqBody)))
	testRequest(suite.TestSuite, req, "Record Not Found", nil)
}

func (suite *GamesTestSuite) TestGamesIdDeleteGood() {
	// Add data to db
	database.Repo.Post(&typing.Game{})

	// Test expected response
	req := httptest.NewRequest("DELETE", "/games/1", nil)
	testRequest(suite.TestSuite, req, "", nil)

	// Test data delete
	res := typing.Game{}
	database.Repo.GetOne(&res, 1)
	suite.Equal(typing.Game{}, res)
}

func (suite *GamesTestSuite) TestGamesIdDeleteNotFound() {
	// Test expected response
	req := httptest.NewRequest("DELETE", "/games/1", nil)
	testRequest(suite.TestSuite, req, "Record Not Found", nil)
}

func (suite *GamesTestSuite) TestSimple() {
	suite.Equal(true, true)
}

func TestGamesTestSuite(t *testing.T) {
	suite.Run(t, new(GamesTestSuite))
}
