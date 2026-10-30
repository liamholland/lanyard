package lanpartyoptions

import (
	commontypes "lanyard/pkg/common_types"
	"testing"
	"time"
)

func TestConfigPathNotFound(t *testing.T) {
	got := *LoadConfig(&commontypes.AppConfig{LanPartyDataPath: "invalidpath"})
	want := commontypes.LanParty{}
	if got != want {
		t.Error()
	}
}

func TestLoadConfigValidOptions(t *testing.T) {
	got := *LoadConfig(&commontypes.AppConfig{LanPartyDataPath: "./test/validoptions.txt"})
	want := commontypes.LanParty{StartDate: time.Unix(1793383526, 0)}
	if got != want {
		t.Error()
	}
}

func TestLoadConfigInvalidOptions(t *testing.T) {
	got := *LoadConfig(&commontypes.AppConfig{LanPartyDataPath: "./test/invalidoptions.txt"})
	want := commontypes.LanParty{}
	if got != want {
		t.Error()
	}
}
