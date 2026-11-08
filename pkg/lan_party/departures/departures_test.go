package departures

import (
	lanparty "lanyard/pkg/lan_party"
	"testing"
	"time"
)

func TestInvalidOption(t *testing.T) {
	testLP := lanparty.LanParty{DepartureTimes: []time.Time{
		time.Now().Add(time.Hour * 3),
		time.Now().Add(time.Hour * 2),
		time.Now().Add(time.Hour * 1),
	}}

	now := time.Now()
	if numNow := HowManyAtTime(&testLP, &now); numNow != 3 {
		t.Errorf("Expected 3 players now, got %v", numNow)
	}

	oneHour := now.Add(time.Hour)
	if numHour := HowManyAtTime(&testLP, &oneHour); numHour != 2 {
		t.Errorf("Expected 2 players in an hour, got %v", numHour)
	}

	fourHours := now.Add(time.Hour * 4)
	if numFourHours := HowManyAtTime(&testLP, &fourHours); numFourHours != 0 {
		t.Errorf("Expected no players in four hours, got %v", numFourHours)
	}
}
