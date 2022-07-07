package typing

type User struct {
	ID       int    `json:"user_id" schema:"-"`
	Username string `json:"username" schema:"-"`
	Age      int    `json:"age"`
	Email    string `json:"email" schema:"-"`
}
