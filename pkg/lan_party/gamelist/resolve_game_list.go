package gamelist

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const minimumFileSizeToBeGame = 50 * 1000 * 1000 // 50MB

func ResolveGameListFromPaths(gl *GameList, paths []string) *GameListResolution {
	resolution := make(GameListResolution)

	for _, p := range paths {
		unresolvedGames := func() []string {
			l := make([]string, len(*gl))
			for _, g := range *gl {
				if r, ok := resolution[g]; !ok || !r {
					l = append(l, g)
				}
			}
			return l
		}()

		for _, g := range unresolvedGames {
			resolution[g] = resolveGameAtPath(g, p)
		}
	}

	return &resolution
}

func ResolveGameListFromMap(gl *GameList, games map[string]int64, resolvedGames map[string]bool) *GameListResolution {
	log.Print("starting game resolution for game list")

	resolution := make(GameListResolution)
	namesForMatching := func() map[string]string {
		m := make(map[string]string)
		for g := range games {
			m[encodedGameName(g)] = g
		}
		return m
	}()

	for _, g := range *gl {
		resolution[g] = func() bool {
			if r, ok := resolvedGames[g]; ok && r {
				return true
			}
			return false
		}()

		if resolution[g] {
			continue
		}

		encodedG := encodedGameName(g)
		if n, ok := namesForMatching[encodedG]; !ok {
			continue
		} else if games[n] < minimumFileSizeToBeGame {
			continue
		}

		resolution[g] = true
	}

	return &resolution
}

func resolveGameAtPath(game string, path string) bool {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Print("error occurred while reading directories: ", path, err)
		return false
	}

	for _, d := range entries {
		if !d.IsDir() {
			continue
		}

		if !strings.EqualFold(encodedGameName(game), encodedGameName(d.Name())) {
			continue
		}

		investigated, suberr := investigateGameDirectory(path, d.Name())
		if suberr != nil {
			log.Print("error occurred when investigating directory: ", d.Name(), suberr)
			return false
		}
		return investigated
	}

	return false
}

func investigateGameDirectory(path string, dirName string) (bool, error) {
	if entries, err := os.ReadDir(filepath.Join(path, dirName)); err != nil {
		log.Print("failed to read game directory: ", dirName, err)
	} else if len(entries) == 0 {
		return false, nil
	} else if len(entries) >= 50 {
		return true, nil
	} else {
		return guessBasedOnDirectorySize(filepath.Join(path, dirName), minimumFileSizeToBeGame)
	}

	return false, nil
}

func guessBasedOnDirectorySize(path string, minSize int64) (bool, error) {
	var size int64 = 0
	err := filepath.Walk(path, func(path string, info fs.FileInfo, fileReadErr error) error {
		if fileReadErr != nil {
			return fileReadErr
		}

		if !info.IsDir() {
			size += info.Size()
		}
		return fileReadErr
	})

	if err != nil {
		log.Print("error occurred while determining size of game directory: ", path, err)
		return false, err
	}

	log.Printf("read %d bytes in game directory %s\n", size, path)

	if size >= minSize {
		return true, nil
	}

	return false, nil
}

func encodedGameName(game string) string {
	return strings.ToLower(strings.ReplaceAll(game, " ", ""))
}
