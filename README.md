```
	
 __         ______     __   __     __  __     ______     ______     _____    
/\ \       /\  __ \   /\ "-.\ \   /\ \_\ \   /\  __ \   /\  == \   /\  __-.  
\ \ \____  \ \  __ \  \ \ \-.  \  \ \____ \  \ \  __ \  \ \  __<   \ \ \/\ \ 
 \ \_____\  \ \_\ \_\  \ \_\\"\_\  \/\_____\  \ \_\ \_\  \ \_\ \_\  \ \____- 
  \/_____/   \/_/\/_/   \/_/ \/_/   \/_____/   \/_/\/_/   \/_/ /_/   \/____/ 
                                                                             
```

a suite of tools for running LAN parties

**lanyard** allows users to configure settings for a LAN party, such as the start date and the list of games to play via a cli, then start a server which allows every present player to view and contribute to that information

**lanyard** can do things like tell you how many players will be availble on a given day, tell players which games from the games list they have yet to install, and suggest when the best time to play a game is based on the ideal player numbers for that game

***

## Usage

create a party config
```sh
ly party create
```

get options
```sh
ly party get [option]
```

set options
```sh
ly party set [option] [value]
```

resolve installed games against the games list
```sh
ly gamelist resolve [paths]
```

### Configuring the game list
the game list can be any text file. you can configure multiple lists and swap them out on the fly by pointing the server or cli at a different one.

example game list:
```txt
Halo 3 #3
Double Action
Unreal Tournament
Straftat #4
Half-Life #8
```
the ideal player numbers can optionally be specified by adding `#[number]` next to the game name.


***

## Building

- clone the repo
- intall go
- add to the code
- run tests with `go test ./...`
- build with `go build -o [build path]/ly -v ./cli/...`
- view commands with `[build path]/ly --help`
