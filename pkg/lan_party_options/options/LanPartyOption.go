package options

import types "lanyard/pkg/common_types"

type loadFunction func(*types.LanParty, any) error

type LanPartyOption struct {
	FileDeclaration string
	LoadFunction    loadFunction
}
