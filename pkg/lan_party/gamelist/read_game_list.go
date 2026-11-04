package gamelist

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type GameList []string

func ReadGameList(fp string) (*GameList, error) {
	if _, err := os.Stat(fp); err == nil {
		return readConfigFromFile(fp), nil
	} else if os.IsNotExist(err) {
		log.Print("no list exists at the location '", fp, "'")
		return nil, err
	} else {
		return nil, err
	}
}

func readConfigFromFile(fp string) *GameList {
	gl := make(GameList, 0, 15)

	file, err := os.Open(fp)
	if err != nil {
		log.Printf("failed to open file with path '%s': %s\n", fp, err.Error())
		return &gl
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		if s.Text() != "" {
			gl = append(gl, strings.TrimSpace(s.Text()))
		}
	}

	if err := s.Err(); err != nil {
		log.Print("error while readign games: ", err.Error())
	}

	return &gl
}
