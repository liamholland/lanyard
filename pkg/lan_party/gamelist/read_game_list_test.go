package gamelist

import (
	"os"
	"testing"
)

func TestReadGameList_Success(t *testing.T) {
	testFilePath := "test/games.txt"
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
