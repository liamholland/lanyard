package departures

import (
	"errors"
	lp "lanyard/pkg/lan_party"
	"time"
)

func HowManyAtTime(lanparty *lp.LanParty, t *time.Time) int64 {
	var present int64 = 0

	for _, d := range lanparty.DepartureTimes {
		if t.Before(d) {
			present += 1
		}
	}

	return present
}

func AddDepartureTime(lanparty *lp.LanParty, timestamp int64) error {
	t := time.Unix(timestamp, 0)
	if t.Before(time.Now()) {
		return errors.New("departure is before current time")
	}

	lanparty.DepartureTimes = append(lanparty.DepartureTimes, t)
	return nil
}
