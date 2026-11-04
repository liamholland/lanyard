package options

import lp "lanyard/pkg/lan_party"

type readFunction func(*lp.LanParty, any) error
type writeFunction func(any, *LanPartyOption) (string, error)

type LanPartyOption struct {
	FileDeclaration string
	OnRead          readFunction
	OnWrite         writeFunction
}
