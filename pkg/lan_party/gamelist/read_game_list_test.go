package gamelist

import (
	"os"
	"testing"
)

const testFilePath string = "test/games.txt"

func TestReadGameList_Success(t *testing.T) {
	list, err := ReadGameList(testFilePath)
	if err != nil {
		t.Fatalf("ReadGameList returned an unexpected error: %v", err)
	}

	expected := []string{
		"Titanfall 2",
		"Half-Life",
		"Halo 3",
	}

	if len(*list) != len(expected) {
		t.Errorf("Expected %d games, got %d", len(expected), len(*list))
	}

	for i, expectedGame := range expected {
		if (*list)[i] != expectedGame {
			t.Errorf("Index %d: Expected %q, got %q", i, expectedGame, (*list)[i])
		}
	}
}

func TestReadGameList_FileNotFound(t *testing.T) {
	nonExistentPath := "test/missing.txt"

	list, err := ReadGameList(nonExistentPath)

	if list != nil {
		t.Errorf("Expected nil list when file not found, got list of length %d", len(*list))
	}

	if err == nil {
		t.Fatal("Expected an error when file does not exist, got nil")
	}

	// Check if the error indicates that the file does not exist
	if !os.IsNotExist(err) {
		t.Errorf("Expected error type os.IsNotExist, got %v", err)
	}
}

func TestReadIdealNumbers_Success(t *testing.T) {
	ideal, err := ReadIdealPlayers(testFilePath)
	if err != nil {
		t.Fatalf("GetIdealPlayersForGame returned an unexpected error: %v", err)
	}

	if len(*ideal) != 2 {
		t.Fatalf("Expected 2 games with ideal numbers, got %d", len(*ideal))
	}

	if val, ok := (*ideal)[encodedGameName("Titanfall 2")]; !ok || val != 4 {
		t.Errorf("Expected Titanfall to have 4 players, got %v", val)
	}

	if val, ok := (*ideal)[encodedGameName("Halo 3")]; !ok || val != 6 {
		t.Errorf("Expected Halo to have 6 players, got %v", val)
	}
}

func TestGetIdealPlayersForGame_FileNotFound(t *testing.T) {
	nonExistentPath := "test/missing_ideal.txt"

	ideal, err := ReadIdealPlayers(nonExistentPath)

	if ideal != nil {
		t.Errorf("Expected nil ideal map when file not found, got %v", ideal)
	}

	if err == nil {
		t.Fatal("Expected an error when file does not exist, got nil")
	}

	if !os.IsNotExist(err) {
		t.Errorf("Expected error type os.IsNotExist, got %v", err)
	}
}

func TestGetIdealPlayersForGame_Success(t *testing.T) {
	if val, _ := GetIdealPlayersForGame("Titanfall 2", testFilePath); val != 4 {
		t.Errorf("Expected Titanfall to have 4 players, got %v", val)
	}

	if val, _ := GetIdealPlayersForGame("Halo 3", testFilePath); val != 6 {
		t.Errorf("Expected Halo to have 6 players, got %v", val)
	}

	if val, _ := GetIdealPlayersForGame("Half-Life", testFilePath); val != -1 {
		t.Errorf("Expected not to get a value for Half-Life, got %v", val)
	}

	if val, _ := GetIdealPlayersForGame("Fake Game", testFilePath); val != -1 {
		t.Errorf("Expected not to get a value for non-existent game, got %v", val)
	}
}
