package options

import (
	"errors"
	types "lanyard/pkg/common_types"
	"log"
	"strconv"
	"time"
)

var StartDateOption LanPartyOption = LanPartyOption{
	Options[StartDate],

	func(lp *types.LanParty, v any) error {
		if v != nil {
			i, err := strconv.ParseInt(v.(string), 10, 64)
			if err != nil {
				log.Print(v, " is not a valid POSIX timestamp")
				return errors.ErrUnsupported
			}
			lp.StartDate = time.Unix(i, 0)
			return nil
		}

		return errors.ErrUnsupported
	},
}
