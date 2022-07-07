package typing

import "github.com/jinzhu/gorm"

type User struct {
	ID       int    `json:"user_id" schema:"-"`
	Username string `json:"username" schema:"-"`
	Age      int    `json:"age"`
	Email    string `json:"email" schema:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = 0
	return
}

func (u *User) BeforePatch(tx *gorm.DB) (err error) {
	u.ID = 0
	return
}
