package lanpartyoptions

import (
	commontypes "lanyard/pkg/common_types"
	"lanyard/pkg/lan_party_options/options"
	"os"
	"testing"
	"time"
)

const testFilePath string = "test/createdTestFile.txt"

func TestCreateConfigFileSuccess(t *testing.T) {

	t.Cleanup(func() {
		os.Remove(testFilePath)
	})

	got := CreateConfigFile(&commontypes.AppConfig{LanPartyDataPath: testFilePath})
	if got != nil {
		t.Error("failed to create the test file")
	}
}

func TestNoFailureIfFileExists(t *testing.T) {
	os.OpenFile(testFilePath, os.O_RDONLY|os.O_CREATE, 0666)
	t.Cleanup(func() {
		os.Remove(testFilePath)
	})

	got := CreateConfigFile(&commontypes.AppConfig{LanPartyDataPath: testFilePath})
	if got != nil {
		t.Error("error was thrown when trying to create existing file")
	}
}

func TestWriteStartDateOption(t *testing.T) {
	testConfig := &commontypes.AppConfig{LanPartyDataPath: testFilePath}
	CreateConfigFile(testConfig)
	t.Cleanup(func() {
		os.Remove(testFilePath)
		os.Remove(testFilePath + ".old")
	})

	got := UpdateConfig(testConfig, commontypes.LanParty{Options: map[string]any{options.Options[options.StartDate]: time.Unix(1793383526, 0)}})
	if got != nil {
		t.Error(got)
	}
}

func TestConfigRestoredWhenUpdateFails(t *testing.T) {
	testConfig := &commontypes.AppConfig{LanPartyDataPath: testFilePath}
	CreateConfigFile(testConfig)
	setup := UpdateConfig(testConfig, commontypes.LanParty{Options: map[string]any{options.Options[options.StartDate]: time.Unix(1793383526, 0)}})
	if setup != nil {
		t.Error("failed to set up the test files")
	}
	t.Cleanup(func() {
		os.Remove(testFilePath)
		os.Remove(testFilePath + ".old")
	})

	failedUpdate := UpdateConfig(testConfig, commontypes.LanParty{Options: map[string]any{options.Options[options.StartDate]: "invalid"}})
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

	sd := lp.Options[options.Options[options.StartDate]]
	want := time.Unix(1793383526, 0)

	if sd != want {
		t.Error("wanted: ", want, "; got: ", sd)
	}
}
