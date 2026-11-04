package options

import (
	"fmt"
	lp "lanyard/pkg/lan_party"
)

type readFunction func(*lp.LanParty, any) error
type writeFunction func(any) (string, error)

type LanPartyOption struct {
	FileDeclaration string
	OnRead          readFunction
	OnWrite         writeFunction
}

func (lp LanPartyOption) GetFileString(v any) (string, error) {
	stringValue, err := lp.OnWrite(v)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%s\n", lp.FileDeclaration, stringValue), nil

}
