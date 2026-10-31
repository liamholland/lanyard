package lanpartyoptions

import (
	commontypes "lanyard/pkg/common_types"
	"os"
	"testing"
)

const testFilePath string = "./test/createdTestFile.txt"

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
