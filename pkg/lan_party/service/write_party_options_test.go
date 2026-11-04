package service

import (
	cfg "lanyard/pkg/config"
	lp "lanyard/pkg/lan_party"
	opts "lanyard/pkg/lan_party/options"
	"os"
	"path/filepath"
	"testing"
	"time"
)

const testFilePath string = "test/createdTestFile.txt"

func TestCreateConfigFileSuccess(t *testing.T) {

	t.Cleanup(func() {
		os.Remove(testFilePath)
	})

	got := CreateConfigFile(&cfg.AppConfig{LanPartyDataPath: testFilePath})
	if got != nil {
		t.Error("failed to create the test file")
	}
}

func TestNoFailureIfFileExists(t *testing.T) {
	os.OpenFile(testFilePath, os.O_RDONLY|os.O_CREATE, 0666)
	t.Cleanup(func() {
		os.Remove(testFilePath)
	})

	got := CreateConfigFile(&cfg.AppConfig{LanPartyDataPath: testFilePath})
	if got != nil {
		t.Error("error was thrown when trying to create existing file")
	}
}

func TestWriteStartDateOption(t *testing.T) {
	testConfig := &cfg.AppConfig{LanPartyDataPath: testFilePath}
	CreateConfigFile(testConfig)
	t.Cleanup(func() {
		os.Remove(testFilePath)
		os.Remove(testFilePath + ".old")
	})

	got := UpdateConfig(testConfig, lp.LanParty{Options: map[string]any{opts.Options[opts.StartDate]: time.Unix(1793383526, 0)}})
	if got != nil {
		t.Error(got)
	}
}

func TestConfigRestoredWhenUpdateFails(t *testing.T) {
	testConfig := &cfg.AppConfig{LanPartyDataPath: testFilePath}
	CreateConfigFile(testConfig)
	setup := UpdateConfig(testConfig, lp.LanParty{Options: map[string]any{opts.Options[opts.StartDate]: time.Unix(1793383526, 0)}})
	if setup != nil {
		t.Error("failed to set up the test files")
	}
	t.Cleanup(func() {
		os.Remove(testFilePath)
		os.Remove(testFilePath + ".old")
	})

	failedUpdate := UpdateConfig(testConfig, lp.LanParty{Options: map[string]any{opts.Options[opts.StartDate]: "invalid"}})
	if failedUpdate == nil {
		t.Error(failedUpdate)
	}

	if _, err := os.Stat(testConfig.LanPartyDataPath); err != nil {
		t.Error(err.Error())
	}

	lp, err := LoadConfig(testConfig)
	if err != nil {
		t.Error(err)
	}

	sd := lp.Options[opts.Options[opts.StartDate]]
	want := time.Unix(1793383526, 0)

	if sd != want {
		t.Error("wanted: ", want, "; got: ", sd)
	}
}

func TestWriteGameListPathOption(t *testing.T) {
	testConfig := &cfg.AppConfig{LanPartyDataPath: testFilePath}
	CreateConfigFile(testConfig)
	t.Cleanup(func() {
		os.Remove(testFilePath)
		os.Remove(testFilePath + ".old")
	})

	absolutePath := "/tmp/test_game_list"

	got := UpdateConfig(testConfig, lp.LanParty{
		Options: map[string]any{opts.Options[opts.GameListPath]: absolutePath},
	})
	if got != nil {
		t.Error(got)
	}

	lp, err := LoadConfig(testConfig)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	want := absolutePath
	if gotVal, ok := lp.Options[opts.Options[opts.GameListPath]]; !ok || gotVal != want {
		t.Errorf("wanted: %v; got: %v", want, gotVal)
	}
}

func TestWriteRelativeGameListPathOptionFails(t *testing.T) {
	testConfig := &cfg.AppConfig{LanPartyDataPath: testFilePath}
	CreateConfigFile(testConfig)
	t.Cleanup(func() {
		os.Remove(testFilePath)
		os.Remove(testFilePath + ".old")
	})

	relativePath := "./game_list"

	got := UpdateConfig(testConfig, lp.LanParty{
		Options: map[string]any{opts.Options[opts.GameListPath]: relativePath},
	})

	if got == nil {
		t.Fatal("expected an error when writing a relative path, got nil")
	}

	if got != filepath.ErrBadPattern {
		t.Errorf("expected error %v, got %v", filepath.ErrBadPattern, got)
	}
}
