# User Experience

Users should be able to work with three resources: games, users and links

## Games
Games can be created with the command
```
games create /path/to/file/
```	

Games can be listed with the command 
```
games ls
```
You can filter your search with the filter flag that accepts a key-value pair.\
Accepted keys:
| Name       | Value                   |
| ---------- | ----------------------- |
| publisher  | Name of the publisher   |
| age rating | The age rating          |

Example :
```
games ls --filter genre=action
```

Games can be updated with the command
```
games update <game ID> /path/to/file/
```

And deleted with
```
games delete <game ID>
```

To get information about a specific game use command
```
games info <game ID>
```

## Users
Users can be created with the command
```
users create /path/to/file/
```

Users can be listed with the command
```
users ls
```
You can filter your search with the filter flag that accepts a key-value pair.\
Accepted keys:
| Name | Value           |
| ---- | --------------- |
| age  | Age of the user |

Example :
```
users ls --filter age=21
```

Users can be updated with the command
```
users update <user ID> /path/to/file
```

And deleted with 
```
users delete <user ID>
```

To get information about a specific user use command
```
users info <user ID>
```

## Links

You can link users to the games they play using the command
```
links create /path/to/file/
```

To see what games a user plays you can use the command
```
links ls --filter user=<user ID>
```

To see who plays a certain game you can use the command
```
links ls --filter game=<game ID>
```

Links can be updated with the command
```
links update <link ID> /path/to/file/
```

And deleted with
```
links delete <link ID>
```

To get information about a specific link use command
```
links info <link ID>
```