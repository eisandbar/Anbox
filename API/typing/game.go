package typing

import (
	"github.com/jinzhu/gorm"
)

type Game struct {
	ID          int    `json:"game_id" schema:"-"`
	Title       string `json:"title" schema:"-"`
	Description string `json:"description" schema:"-"`
	Url         string `json:"url" schema:"-"`
	AgeRating   int    `json:"age_rating" schema:"age_rating"`
	Publisher   string `json:"publisher"`
}

func (g *Game) BeforeCreate(tx *gorm.DB) (err error) {
	g.ID = 0
	return
}

func (g *Game) BeforePatch(tx *gorm.DB) (err error) {
	g.ID = 0
	return
}
