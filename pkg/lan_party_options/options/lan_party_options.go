package options

import commontypes "lanyard/pkg/common_types"

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

func GetOptionValueFromName(partyConfig *commontypes.LanParty, option string) any {
	return GetOptionValue(partyConfig, optionNames[option])
}

func GetOptionValue(partyConfig *commontypes.LanParty, option ValidLanPartyOptionName) any {
	return partyConfig.Options[Options[option]]
}
