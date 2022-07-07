package database

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

var Repo Repository

func (repo *Repository) Connect(args ...interface{}) {
	// Connects to DB and sets it as Repo.db
	var db *gorm.DB
	var err error

	if len(args) == 0 {
		db, err = gorm.Open("postgres", connStr)
		if err != nil {
			log.Fatalf("Error connecting to postgres: %v", err)
		}
	} else {
		db, err = gorm.Open("postgres", args...)
		if err != nil {
			log.Fatalf("Error connecting to postgres: %v", err)
		}
	}
	repo.db = db
}

func (repo *Repository) Close() {
	// Closes connection to DB
	repo.db.Close()
}

func (repo *Repository) GetAll(data, filter interface{}) (error_msg string) {
	// Find all rows that match the filter in table 'data'
	repo.db.Find(data, filter)
	return
}

func (repo *Repository) GetOne(data interface{}, id int) (error_msg string) {
	// Find first row in table 'data' that matches id
	err := repo.db.First(data, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		error_msg = "Record Not Found"
	}
	return
}

func (repo *Repository) Post(data interface{}) (error_msg string) {
	// Create new record in table 'data'
	if repo.db.Create(data).Error != nil {
		error_msg = "Error Creating Record"
	}
	return
}

func (repo *Repository) Patch(data, newData interface{}, id int) (error_msg string) {
	// Update record in table 'data' that matches id
	db := repo.db.Model(data).Where("id = ?", id).Updates(newData)
	if db.RowsAffected < 1 {
		error_msg = "Record Not Found"
	}
	if db.Error != nil {
		error_msg = "Error Updating Record"
	}
	return
}

func (repo *Repository) Delete(data interface{}, id int) (error_msg string) {
	// Delete record in table 'data' that matches id
	db := repo.db.Delete(data, id)
	if db.RowsAffected < 1 {
		error_msg = "Record Not Found"
	}
	if db.Error != nil {
		error_msg = "Error Deleting Record"
	}
	return
}
