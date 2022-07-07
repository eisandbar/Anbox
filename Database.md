## Database

Our database will have 3 main resources: games, users and links.

games:
```json
{
    "game_id": 1,
    "title": "Runner",
    "description": "A Game",
    "url": "https://www.studio.com/runner",
    "age_rating": 6,
    "publisher": "studio"
}
```

users:
```json
 {
    "user_id": 1,
    "username": "john",
    "age": 19,
    "email": "john@mail.com"
}
```

links:
```json
{
    "link_id": 1,
    "game_id": 1,
    "user_id": 1,
    "title": "Runner",
    "username": "john",
    "hours_played": 18
}
```

I think it is best to use an SQL database because we are working with structured relational data.

As a bonus if we ever need to analyze the data it will be easier to do so in SQL.

I chose to separate links into a separate tables because this way we can query both by game and user id. It also makes it easier to update and to add additional columns (e.g. achievements, last played, etc.).

Additionally, in case of increased load, we can use a redis cache for common queries, especially for games, since the list won't be updated often and will be queried by all users.