package typing

type Link struct {
	ID          int    `json:"link_id" schema:"-"`
	GameId      int    `json:"game_id"`
	UserId      int    `json:"user_id"`
	Title       string `json:"title" schema:"-"`
	Username    string `json:"username" schema:"-"`
	HoursPlayed int    `json:"hours_played" schema:"-"`
}
