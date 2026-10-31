package options

import (
	"errors"
	"fmt"
	types "lanyard/pkg/common_types"
	"log"
	"strconv"
	"time"
)

var StartDateOption LanPartyOption = LanPartyOption{
	Options[StartDate],

	func(lp *types.LanParty, v any) error {
		if v == nil {
			return errors.ErrUnsupported
		}

		t, err := parseValidPOSIXTimestamp(v)
		if err != nil {
			return err
		}

		lp.Options[Options[StartDate]] = t
		return nil
	},

	func(v any, self *LanPartyOption) (string, error) {
		if v == nil {
			return "", errors.ErrUnsupported
		}

		t, ok := v.(time.Time)
		if !ok {
			return "", errors.ErrUnsupported
		}

		return fmt.Sprintf("%s:%d", self.FileDeclaration, t.Unix()), nil
	},
}

func parseValidPOSIXTimestamp(v any) (time.Time, error) {
	i, err := strconv.ParseInt(v.(string), 10, 64)
	if err != nil {
		log.Print(v, " is not a valid POSIX timestamp")
		return time.Time{}, errors.ErrUnsupported
	}
	return time.Unix(i, 0), nil
}
