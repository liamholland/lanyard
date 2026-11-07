package gamelist

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadGameList(fp string) (*GameList, error) {
	if _, err := os.Stat(fp); err == nil {
		return readGameListFromFile(fp), nil
	} else if os.IsNotExist(err) {
		log.Print("no list exists at the location '", fp, "'")
		return nil, err
	} else {
		log.Print("error while reading games: ", err)
		return nil, err
	}
}

func GetIdealPlayersForGame(game string, fp string) (int64, error) {
	ideal, err := ReadIdealPlayers(fp)
	if err != nil {
		return 0, err
	}

	if numPlayers, ok := (*ideal)[encodedGameName(game)]; !ok || numPlayers == 0 {
		return -1, nil
	} else {
		return numPlayers, nil
	}
}

func ReadIdealPlayers(fp string) (*IdealNumbers, error) {
	if _, err := os.Stat(fp); err == nil {
		return readIdealNumbersFromFile(fp), nil
	} else if os.IsNotExist(err) {
		log.Print("no list exists at the location '", fp, "'")
		return nil, err
	} else {
		log.Print("error while reading ideal players: ", err)
		return nil, err
	}
}

func readGameListFromFile(fp string) *GameList {
	gl := make(GameList, 0, 15)

	file, err := os.Open(fp)
	if err != nil {
		log.Printf("failed to open file with path '%s': %s\n", fp, err.Error())
		return nil
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		if s.Text() != "" {
			gameName, _, _ := strings.Cut(s.Text(), "#")
			gl = append(gl, strings.TrimSpace(gameName))
		}
	}

	if err := s.Err(); err != nil {
		log.Print("error while reading games: ", err.Error())
	}

	return &gl
}

func readIdealNumbersFromFile(fp string) *IdealNumbers {
	ideal := make(IdealNumbers)

	file, err := os.Open(fp)
	if err != nil {
		log.Printf("failed to open file with path '%s': %s\n", fp, err.Error())
		return nil
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		if s.Text() != "" {
			gameName, idealNumber, present := strings.Cut(s.Text(), "#")
			if !present {
				continue
			}

			v, parseErr := strconv.ParseInt(idealNumber, 10, 16)
			if parseErr != nil {
				log.Print("error while parsing ideal number of players for ", gameName, " value may not be valid")
				continue
			}

			ideal[encodedGameName(gameName)] = v
		}
	}

	if err := s.Err(); err != nil {
		log.Print("error while reading games: ", err.Error())
	}

	return &ideal
}
