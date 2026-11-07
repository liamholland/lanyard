package options

import (
	"errors"
	"fmt"
	lp "lanyard/pkg/lan_party"
	"log"
	"strconv"
	"time"
)

var StartDateOption LanPartyOption = LanPartyOption{
	Options[StartDate],

	func(lanp *lp.LanParty, v any) error {
		if v == nil {
			return errors.ErrUnsupported
		}

		t, err := parseValidPOSIXTimestamp(v)
		if err != nil {
			return err
		}

		lanp.Options[Options[StartDate]] = t
		return nil
	},

	func(v any) (string, error) {
		if v == nil {
			return "", errors.ErrUnsupported
		}

		t, ok := v.(time.Time)
		if !ok {
			return "", errors.ErrUnsupported
		}

		return fmt.Sprintf("%d", t.Unix()), nil
	},
}

func parseValidPOSIXTimestamp(v any) (time.Time, error) {
	i, err := strconv.ParseInt(v.(string), 10, 64)
	if err != nil {
		log.Print(v, " is not a valid POSIX timestamp:", err)
		return time.Time{}, errors.ErrUnsupported
	}
	return time.Unix(i, 0), nil
}
