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

you can use
```sh
ly help
```
or
```sh
ly help [command path]
```
to get more information on any commands

***

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

### Running a server

start the server with
```sh
ly server start
```

check its status with
```sh
ly server up
```

print the address of the server and web page with
```sh
> ly server where
server is running :)
server responded with:  hello gamer

local:  http://localhost:1311
LAN web page:  http://192.168.68.128:1311/lanyard
```

change settings of the config on the running server with
```sh
ly server update [setting] [value]
```

### Resolving the game list
you can resolve the game list against installed games with the cli as follows:
```sh
> ly gamelist resolve /home/liamh/lm/SteamLibrary/steamapps/common
result of game resolution:
> Straftat                      missing
> Half-Life                     ok!
> Halo 3                        missing
> Double Action                 missing
> Unreal Tournament             missing
```

the command takes multiple paths as args:
```sh
> builds/ly gamelist resolve /home/liamh/lm/SteamLibrary/steamapps/common /home/liamh/lm/SteamLibrary/steamapps/common/Halo\ The\ Master\ Chief\ Collection -p builds/gl.txt
result of game resolution:
> Unreal Tournament             missing
> Straftat                      missing
> Half-Life                     ok!
> Halo 3                        ok!
> Double Action                 missing
```

this service is also available via the server and is exposed via the web page using the file upload input type. so any user can resolve their games against the list. once they have, the admin can check who has which games installed:
```sh
> ly server users blame
Current User Game Status:

> liam:
> Unreal Tournament             missing
> Double Action                 missing
> Half-Life                     ok!
> Halo 3                        ok!
> Straftat                      missing
```

or query a specific game:
```sh
> ly server users status "Halo 3"
result of game resolution for Halo 3:
> liam                          ok!
```

the provided paths to the web page can either be the game directory or the containing directory. For example, both the paths of the Halo 3 directory itself, and the directory of Halo The Master Chief Collection will resolve Halo 3. the cli will only work with the containing directory right now, since it means shorter paths, and the execution is much faster.

### Checking player numbers

on the server, users can enter the time they will be departing the LAN party. the server tracks this and allows any user to query how many players will be available at a selected time.

```txt
there will be 1 players at 2026-11-15T22:06
```

any user can also enter the name of a game on the list (or not on the list) and query the time to play it. if the game has a specified number of ideal players, the server will tell you the last time the game could be played with that number of players.

```txt
play Halo 3 before Sunday 23:06
```

***

## Building

- clone the repo
- intall go
- add to the code
- run tests with `go test ./...`
- build with `go build -o [build path]/ly -v ./cli/...`
- view commands with `[build path]/ly --help`
