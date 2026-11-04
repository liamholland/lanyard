package service

import (
	"errors"
	cfg "lanyard/pkg/config"
	opts "lanyard/pkg/lan_party/options"
	"os"
	"testing"
	"time"
)

func TestConfigPathNotFound(t *testing.T) {
	_, got := LoadConfig(&cfg.AppConfig{LanPartyDataPath: "invalidpath"})
	if !errors.Is(got, os.ErrNotExist) {
		t.Errorf("did not return corrent error type: %T", got)
	}
}

func TestLoadConfigValidOptions(t *testing.T) {
	got, _ := LoadConfig(&cfg.AppConfig{LanPartyDataPath: "./test/validoptions.txt"})
	want := time.Unix(1793383526, 0)
	if got.Options[opts.Options[opts.StartDate]] != want {
		t.Error()
	}
}

func TestLoadConfigInvalidOptions(t *testing.T) {
	got, _ := LoadConfig(&cfg.AppConfig{LanPartyDataPath: "./test/invalidoptions.txt"})
	if len(got.Options) != 0 {
		t.Error(got)
	}
}
