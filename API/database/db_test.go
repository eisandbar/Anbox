package database

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

/*
These tests confirm that all the SQL behavior in dataase/repository.go
is as we expect using a DB mock instead of a real DB.
Not all of these tests need asserts as mock.Expect is already an assert
*/

type DBTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
}

type Data struct {
	ID   int
	Name string
}

var testData = Data{
	ID:   1,
	Name: "John",
}

func (suite *DBTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mock = mock
	Repo.Connect(db)
}

func (suite *DBTestSuite) AfterTest(_, _ string) {
	require.NoError(suite.T(), suite.mock.ExpectationsWereMet())
}

func (suite *DBTestSuite) TestGetAll() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "data"`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name"}).AddRow(testData.ID, testData.Name))

	var res []Data
	Repo.GetAll(&res, Data{})
	assert.Equal(suite.T(), []Data{testData}, res)
}

func (suite *DBTestSuite) TestPost() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "data" ("name") 
		 VALUES ($1) RETURNING "data"."id"`)).
		WithArgs(testData.Name).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(testData.ID))
	suite.mock.ExpectCommit()

	Repo.Post(&Data{Name: "John"})
}

func (suite *DBTestSuite) TestGetOne() {
	suite.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "data"
		WHERE ("data"."id" = 1)
		ORDER BY "data"."id" ASC LIMIT 1`)).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name"}).AddRow(testData.ID, testData.Name))

	var res Data
	Repo.GetOne(&res, 1)
	assert.Equal(suite.T(), testData, res)
}

func (suite *DBTestSuite) TestPatch() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(regexp.QuoteMeta(
		`UPDATE "data" SET "name" = $1 WHERE (id = $2)`)).
		WithArgs(testData.Name, testData.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	Repo.Patch(&Data{}, &Data{Name: "John"}, 1)
}

func (suite *DBTestSuite) TestDelete() {
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec(regexp.QuoteMeta(
		`DELETE FROM "data"  WHERE ("data"."id" = 1)`)).
		WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()

	Repo.Delete(&Data{}, 1)
}

func TestDBTestSuite(t *testing.T) {
	suite.Run(t, new(DBTestSuite))
}
