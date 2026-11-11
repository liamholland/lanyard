package departures

import (
	"errors"
	lp "lanyard/pkg/lan_party"
	"sort"
	"time"
)

func HowManyAtTime(lanparty *lp.LanParty, t *time.Time) int64 {
	var present int64 = 0

	for id, d := range lanparty.DepartureTimes {
		if id == 0 {
			continue
		} else if t.Before(d) {
			present += 1
		}
	}

	return present
}

func AddDepartureTime(lanparty *lp.LanParty, timestamp int64, id int64) error {
	t := time.Unix(timestamp, 0)
	if t.Before(time.Now()) {
		return errors.New("departure is before current time")
	}

	lanparty.DepartureTimes[id] = t
	return nil
}

func GetTimeWithNPeople(lanparty *lp.LanParty, n int64) time.Time {
	deps := func() []time.Time {
		s := make([]time.Time, 0)
		for id, t := range lanparty.DepartureTimes {
			if id == 0 {
				continue
			}

			s = append(s, t)
		}
		return s
	}()

	sort.Slice(deps, func(i, j int) bool {
		return deps[i].After(deps[j])
	})

	maxP := func() int64 {
		lD := int64(len(deps))
		if lD < n {
			return lD
		} else {
			return n
		}
	}()

	return deps[maxP-1]
}
