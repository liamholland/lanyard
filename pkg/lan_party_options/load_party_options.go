package lanpartyoptions

import (
	"bufio"
	"errors"
	types "lanyard/pkg/common_types"
	options "lanyard/pkg/lan_party_options/options"
	"log"
	"os"
	"strings"
)

func LoadConfig(config *types.AppConfig) *types.LanParty {
	if _, err := os.Stat(config.LanPartyDataPath); err == nil {
		return readConfigFromFile(config.LanPartyDataPath)
	} else if errors.Is(err, os.ErrNotExist) {
		return &types.LanParty{}
	}

	return &types.LanParty{}
}

func readConfigFromFile(path string) *types.LanParty {
	config := types.LanParty{}

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

func readOption(line string, lp *types.LanParty) {
	input := strings.Split(line, ":")
	if options.IsValidOptionName(input[0]) {
		option := options.GetOptionFromName(input[0])
		err := option.LoadFunction(lp, input[1])
		if err != nil {
			log.Print(option.FileDeclaration, " has a malformed value declaration ", input[1])
		}
	}
}
