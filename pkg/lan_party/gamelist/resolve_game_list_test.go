package gamelist

import (
	"path/filepath"
	"testing"
)

const deletedGame = "just leftover data"
const noGame = "nogame"
const game = "game name"
const testDir = "test"

func TestGuessBasedOnDirectorySize(t *testing.T) {
	var minSize int64 = 1000

	t.Run("SizeBelowThreshold", func(t *testing.T) {
		result, err := guessBasedOnDirectorySize(filepath.Join(testDir, deletedGame), minSize)
		if err != nil {
			t.Fatalf("guessBasedOnDirectorySize failed unexpectedly: %v", err)
		}
		if result != false {
			t.Error("Expected false for size below threshold, got true")
		}
	})

	t.Run("SizeAtOrAboveThreshold", func(t *testing.T) {
		result, err := guessBasedOnDirectorySize(filepath.Join(testDir, game), minSize)
		if err != nil {
			t.Fatalf("guessBasedOnDirectorySize failed unexpectedly: %v", err)
		}
		if result != true {
			t.Error("Expected true for size above threshold, got false")
		}
	})
}

func TestInvestigateGameDirectory(t *testing.T) {
	t.Run("DirectoryIsEmpty", func(t *testing.T) {
		result, err := investigateGameDirectory(filepath.Join(testDir, noGame), noGame)
		if err != nil {
			t.Fatalf("investigateGameDirectory failed unexpectedly: %v", err)
		}
		if result != false {
			t.Error("Expected false when directory is empty, got true")
		}
	})

	t.Run("DirectoryDoesNotExist", func(t *testing.T) {
		result, err := investigateGameDirectory(filepath.Join(testDir, "missing_game"), "missing_game")
		if err != nil {
			t.Fatalf("investigateGameDirectory failed unexpectedly: %v", err)
		}
		if result != false {
			t.Error("Expected false when directory does not exist, got true")
		}
	})
}

func TestResolveGameAtPath(t *testing.T) {
	t.Run("FoundInvalidGame", func(t *testing.T) {
		result := resolveGameAtPath(deletedGame, testDir)
		if result != false {
			t.Error("Expected false when a game directory is found but size fails, got true")
		}
	})

	t.Run("NoGameDirectoryFound", func(t *testing.T) {
		result := resolveGameAtPath("NonExistentDir", testDir)
		if result != false {
			t.Error("Expected false when no game directory is found, got true")
		}
	})
}

func TestResolveGameList(t *testing.T) {
	games := GameList{noGame, deletedGame, game}

	resolution := *ResolveGameList(&games, []string{testDir})

	if len(resolution) < len(games) {
		t.Error("not all games were returned in resolution")
	}

	for g, r := range resolution {
		if r {
			t.Errorf("game %s was resolved but all test data is invalid", g)
		}
	}

}
