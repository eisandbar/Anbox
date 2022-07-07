# API Structure
- [/](#/)
- [/games](#games)
    - [/games/\<id\>](#gamesid)
- [/users](#users)
    - [/users/\<id\>](#usersid)
- [/links](#links)
    - [/links/\<id\>](#linksid)

## /games

### GET
Description: Get a list of games

Parameters:
You can use query parameters to filter what games are returned
| Name       | Type     |
| ---------- | -------- |
| publisher  | string   |
| age rating | int      |

Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": [
        {
            "game_id": 1,
            "title": "Runner",
            "description": "A Game",
            "url": "https://www.studio.com/runner",
            "age_rating": 6,
            "publisher": "studio"
            
        },
        {
            "game_id": 2,
            "title": "Swimmer",
            "description": "Sequel to Runner",
            "url": "https://www.studio.com/swimmer",
            "age_rating": 6,
            "publisher": "studio"
        }
    ]
}
```

### POST
Description: Create a new game

Payload:
```json
{
    "title": "Runner",
    "description": "A Game",
    "url": "https://www.studio.com/runner",
    "age_rating": 6,
    "publisher": "studio"
}
```

Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": null
}
```


## /games/\<id\>

### GET
Description: Get information about a game by ID

Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": {
        "game_id": 1,
        "title": "Runner",
        "description": "A Game",
        "url": "https://www.studio.com/runner",
        "age_rating": 6,
        "publisher": "studio"
    }
}
```

### PATCH
Description: Update an existing game

Payload:
```json
{
    "url": "https://www.new_studio.com/runner",
    "publisher": "new_studio"
}
```
Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": null
}
```
 
### DELETE
Description: Delete an existing game

Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": null
}
```

## /users

### GET
Description: Get a list of users

Parameters:
You can use query parameters to filter what users are returned
| Name      | Type     |
| --------- | -------- |
| age       | int      |

Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": [
        {
            "user_id": 1,
            "username": "john",
            "age": 19,
            "email": "john@mail.com"
        },
        {
            "user_id": 2,
            "username": "joe",
            "age": 22,
            "email": "joe@mail.com"
        }
    ]
}
```

### POST
Description: Create a new user

Payload:
```json
{
    "username": "john",
    "age": 19,
    "email": "john@mail.com"
}
```

Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": null
}
```

## /users/\<id\>

### GET
Description: Get information about a user by ID

Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": {
        "user_id": 1,
        "username": "john",
        "age": 19,
        "email": "john@mail.com"
    }
}
```

### PATCH
Description: Update an existing user

Payload:
```json
{
    "age": 20
}
```

Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": null
}
```

### DELETE
Description: Delete an existing user

Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": null
}
```

## /links
 
### GET
Description: Get a list of user-game links

Parameters:
You can use query parameters to filter the links either by game_id or user_id
| Name      | Type     |
| --------- | -------- |
| user      | int      |
| game      | int      |

Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": [
        {
            "link_id": 1,
            "game_id": 1,
            "user_id": 1,
            "title": "Runner",
            "username": "john",
            "hours_played": 18
        },
        {
            "link_id": 2,
            "game_id": 2,
            "user_id": 1,
            "title": "Swimmer",
            "username": "john",
            "hours_played": 2
        }
    ]
}
```

### POST
Description: create a new link

Payload:
```json
{
    "game_id": 1,
    "user_id": 1,
    "title": "Runner",
    "username": "john",
    "hours_played": 18
}
```

Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": null
}
```

## /links/\<id\>

### GET
Description: Get information about a link by ID

Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": {
        "link_id": 1,
        "game_id": 1,
        "user_id": 1,
        "title": "Runner",
        "username": "john",
        "hours_played": 18
    }
}
```

### PATCH
Description: Update an existing link

Payload:
```json
{
    "hours_played": 22
}
```

Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": null
}
```

### DELETE
Description: Delete an existing link

Output:
```json
{
    "status": "Success",
    "status_code": 200,
    "error_code": 0,
    "error_msg": "",
    "metadata": null
}
```
