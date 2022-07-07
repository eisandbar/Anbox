# Anbox
A demo api for anbox

## Setup

To start the containers
1) cd into the Anbox dir
2) run command: 
```
docker-compose up --build
```

Now you can access the api at localhost:3333

## Games API

Only the games API was implemented.
You can send requests to it with curl:
```
curl localhost:3333/games | json_pp -json_opt pretty,canonical
curl -d '{"title":"Runner", "publisher":"studio"}' -X POST localhost:3333/games | json_pp -json_opt pretty,canonical
curl localhost:3333/games/1 | json_pp -json_opt pretty,canonical
curl -d '{"age_rating":6}' -X PATCH localhost:3333/games/1 | json_pp -json_opt pretty,canonical
curl -X DELETE localhost:3333/games/1 | json_pp -json_opt pretty,canonical
```

## Testing

To run the tests
1) cd into the API dir
2) run command:
```
go test ./... -v
```
This will run all the tests in the project