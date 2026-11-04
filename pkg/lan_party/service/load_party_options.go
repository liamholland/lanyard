package service

import (
	"bufio"
	cfg "lanyard/pkg/config"
	lp "lanyard/pkg/lan_party"
	opts "lanyard/pkg/lan_party/options"
	"log"
	"os"
	"strings"
)

func LoadConfig(config *cfg.AppConfig) (*lp.LanParty, error) {
	if _, err := os.Stat(config.LanPartyDataPath); err == nil {
		return readConfigFromFile(config.LanPartyDataPath), nil
	} else if os.IsNotExist(err) {
		log.Print("the options file at the path '", config.LanPartyDataPath, "' does not exist")
		return nil, err
	} else {
		return nil, err
	}
}

func readConfigFromFile(path string) *lp.LanParty {
	config := lp.LanParty{Options: make(map[string]any)}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("failed to open file with path '%s': %s\n", path, err.Error())
		return &config
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		readOption(s.Text(), &config)
	}

	if err := s.Err(); err != nil {
		log.Print("error while scanning tokens: ", err.Error())
	}

	return &config
}

func readOption(line string, lp *lp.LanParty) {
	input := strings.Split(line, ":")
	if opts.IsValidOptionName(input[0]) {
		option := opts.GetOptionFromName(input[0])
		err := option.OnRead(lp, input[1])
		if err != nil {
			log.Print(option.FileDeclaration, " has a malformed value declaration ", input[1])
		}
	}
}
