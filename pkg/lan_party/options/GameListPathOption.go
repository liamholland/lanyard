package options

import (
	"errors"
	"fmt"
	lp "lanyard/pkg/lan_party"
	"log"
	"path/filepath"
)

var GameListPathOption LanPartyOption = LanPartyOption{
	Options[GameListPath],

	func(lp *lp.LanParty, v any) error {
		if v == nil {
			return errors.ErrUnsupported
		}

		fp := filepath.Clean(v.(string))
		lp.Options[Options[GameListPath]] = fp
		return nil
	},

	func(v any, self *LanPartyOption) (string, error) {
		if v == nil {
			return "", errors.ErrUnsupported
		}

		p, ok := v.(string)
		if !ok {
			return "", errors.ErrUnsupported
		}

		if !filepath.IsAbs(p) {
			log.Print("gamelist file path must be absolute")
			return "", filepath.ErrBadPattern
		}

		return fmt.Sprintf("%s:%s", Options[GameListPath], p), nil
	},
}
