package options

import types "lanyard/pkg/common_types"

type readFunction func(*types.LanParty, any) error
type writeFunction func(any, *LanPartyOption) (string, error)

type LanPartyOption struct {
	FileDeclaration string
	OnRead          readFunction
	OnWrite         writeFunction
}
