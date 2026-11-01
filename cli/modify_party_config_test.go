package main

import (
	common_types "lanyard/pkg/common_types"
	service "lanyard/pkg/lan_party_options"
	"lanyard/pkg/lan_party_options/options"
	"os"
	"strconv"
	"testing"
	"time"
)

const testPath string = "test/validoptions.txt"

func TestCreatePartyConfigFileWithDefaults(t *testing.T) {
	testPath := "test/create_test.txt"
	appConfig := &common_types.AppConfig{LanPartyDataPath: testPath}
	t.Cleanup(func() {
		os.Remove(testPath)
	})

	CreatePartyConfigFileWithDefaults(appConfig)

	if _, err := os.Stat(testPath); os.IsNotExist(err) {
		t.Errorf("Expected config file to be created, but it does not exist at %s", testPath)
	}
}

func TestCreatePartyConfigFileWithDefaults_ServiceFailure(t *testing.T) {
	testPath := "fake_test_dir/fail_create.txt"
	appConfig := &common_types.AppConfig{LanPartyDataPath: testPath}

	CreatePartyConfigFileWithDefaults(appConfig)

	if _, err := os.Stat(testPath); !os.IsNotExist(err) {
		t.Errorf("Expected config file to not be created, but it does exists at %s", testPath)
	}
}

func TestGetCurrentSettingForOption_ConfigNotFound(t *testing.T) {
	appConfig := &common_types.AppConfig{LanPartyDataPath: "non_existent_path"}

	GetCurrentSettingForOption(appConfig, "some_option")
}

func TestGetCurrentSettingForOption_ValidSetting(t *testing.T) {
	appConfig := &common_types.AppConfig{LanPartyDataPath: testPath}

	GetCurrentSettingForOption(appConfig, options.Options[options.StartDate])
}

func TestGetCurrentSettingForOption_InvalidOptionName(t *testing.T) {
	appConfig := &common_types.AppConfig{LanPartyDataPath: testPath}

	GetCurrentSettingForOption(appConfig, "insetting_name")
}

func TestGetCurrentSettingForOption_EmptyValue(t *testing.T) {
	appConfig := &common_types.AppConfig{LanPartyDataPath: testPath}

	GetCurrentSettingForOption(appConfig, "some_option")
}

func TestUpdateCurrentSettingForOption_Success(t *testing.T) {
	appConfig := &common_types.AppConfig{LanPartyDataPath: testPath}

	initialTime := time.Unix(1793383526, 0)
	initialParty := common_types.LanParty{Options: make(map[string]any)}
	initialParty.Options[options.Options[options.StartDate]] = initialTime
	service.UpdateConfig(appConfig, initialParty)

	newValue := time.Now().Add(time.Hour)
	UpdateCurrentSettingForOption(appConfig, options.Options[options.StartDate], strconv.FormatInt(newValue.Unix(), 10))

	updatedSettings := loadConfig(appConfig)

	if updatedSettings == nil {
		t.Fatal("Failed to load config after update.")
	}

	actualValue, ok := updatedSettings.Options[options.Options[options.StartDate]]
	if !ok {
		t.Fatal("startDate option not found in updated config.")
	}

	if newValue.Unix() != actualValue.(time.Time).Unix() {
		t.Errorf("Expected value to be %v, but got %v", newValue, actualValue)
	}
}

func TestUpdateCurrentSettingForOption_InvalidOptionName(t *testing.T) {
	appConfig := &common_types.AppConfig{LanPartyDataPath: testPath}

	service.UpdateConfig(appConfig, common_types.LanParty{Options: map[string]any{options.Options[options.StartDate]: time.Now()}})

	UpdateCurrentSettingForOption(appConfig, "invalid_setting", strconv.FormatInt(time.Now().Unix(), 10))
}

func TestUpdateCurrentSettingForOption_InvalidValue(t *testing.T) {
	appConfig := &common_types.AppConfig{LanPartyDataPath: testPath}

	service.UpdateConfig(appConfig, common_types.LanParty{Options: map[string]any{options.Options[options.StartDate]: time.Now()}})

	UpdateCurrentSettingForOption(appConfig, options.Options[options.StartDate], "not a valid time")
}

func TestUpdateCurrentSettingForOption_ServiceFailure(t *testing.T) {
	appConfig := &common_types.AppConfig{LanPartyDataPath: "fake_test_dir/settings.txt"}

	service.UpdateConfig(appConfig, common_types.LanParty{Options: map[string]any{options.Options[options.StartDate]: time.Now()}})

	UpdateCurrentSettingForOption(appConfig, options.Options[options.StartDate], strconv.FormatInt(time.Now().Add(time.Hour).Unix(), 10))
}
