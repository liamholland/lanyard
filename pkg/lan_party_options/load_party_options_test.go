package lanpartyoptions

import (
	"errors"
	commontypes "lanyard/pkg/common_types"
	"os"
	"testing"
	"time"
)

func TestConfigPathNotFound(t *testing.T) {
	_, got := LoadConfig(&commontypes.AppConfig{LanPartyDataPath: "invalidpath"})
	if !errors.Is(got, os.ErrNotExist) {
		t.Errorf("did not return corrent error type: %T", got)
	}
}

func TestLoadConfigValidOptions(t *testing.T) {
	got, _ := LoadConfig(&commontypes.AppConfig{LanPartyDataPath: "./test/validoptions.txt"})
	want := commontypes.LanParty{StartDate: time.Unix(1793383526, 0)}
	if *got != want {
		t.Error()
	}
}

func TestLoadConfigInvalidOptions(t *testing.T) {
	got, _ := LoadConfig(&commontypes.AppConfig{LanPartyDataPath: "./test/invalidoptions.txt"})
	want := commontypes.LanParty{}
	if *got != want {
		t.Error()
	}
}
