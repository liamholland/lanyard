package options

import (
	lp "lanyard/pkg/lan_party"
	"maps"
)

type ValidLanPartyOptionName int

const (
	StartDate ValidLanPartyOptionName = iota
)

var Options = map[ValidLanPartyOptionName]string{
	StartDate: "start",
}

var optionNames = func(m map[ValidLanPartyOptionName]string) map[string]ValidLanPartyOptionName {
	var on = make(map[string]ValidLanPartyOptionName)
	for k, v := range m {
		on[v] = k
	}
	return on
}(Options)

func IsValidOptionName(optionName string) bool {
	_, ok := optionNames[optionName]
	return ok
}

func GetOptionFromName(optionName string) *LanPartyOption {
	if IsValidOptionName(optionName) {
		return GetOption(optionNames[optionName])
	}

	return nil
}

func GetOption(optionDeclaration ValidLanPartyOptionName) *LanPartyOption {
	switch optionDeclaration {
	case StartDate:
		return &StartDateOption
	default:
		return nil
	}
}

func GetOptionValueFromName(partyConfig *lp.LanParty, option string) any {
	return GetOptionValue(partyConfig, optionNames[option])
}

func GetOptionValue(partyConfig *lp.LanParty, option ValidLanPartyOptionName) any {
	return partyConfig.Options[Options[option]]
}

func Copy(original *lp.LanParty) *lp.LanParty {
	copy := lp.LanParty{Options: make(map[string]any)}
	maps.Copy(copy.Options, original.Options)
	return &copy
}
