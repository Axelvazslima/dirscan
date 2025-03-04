package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func TreeSearching(path string, fileName string) []string {
	queue := []string{path}
	var exactMatches []string
	var possibleMatches []string

	fileName = removeFileExtension(fileName)

	for len(queue) > 0 {
		currentDir := queue[0]
		queue = queue[1:]

		entries, err := os.ReadDir(currentDir)
		if err != nil {
			fmt.Printf("Error reading directory %s: %v\n", currentDir, err)
			continue
		}

		for _, entry := range entries {
			fullPath := filepath.Join(currentDir, entry.Name())

			if entry.IsDir() {
				queue = append(queue, fullPath)
				continue
			}

			// Check file name against target
			switch nameCorrespondence(entry.Name(), fileName) {
			case 1:
				exactMatches = append(exactMatches, fullPath+" -> Exact Match")
			case 2:
				possibleMatches = append(possibleMatches, fullPath+" -> Possible Match")
			}
		}
	}

	// Return results based on priority: exact matches first, then possible
	output := make([]string, 0)
	output = append(output, exactMatches...)
	output = append(output, possibleMatches...)

	if len(output) > 0 {
		return output
	}
	return []string{"File not found"}
}

func nameCorrespondence(name string, nameToMatch string) int {
	nameToMatch = removeFileExtension(nameToMatch)
	name, nameToMatch = strings.ToLower(name), strings.ToLower(nameToMatch)
	nameLen, namenameToMatchLen := len(name), len(nameToMatch)

	if nameLen < (namenameToMatchLen-1) || nameLen > (namenameToMatchLen+1) {
		return 0
	}

	nonMatches := 0
	for i := range nameLen {
		if nameToMatch[i] == '.' {
			break
		}
		if name[i] != nameToMatch[i] {
			nonMatches++
			if nonMatches > (nameLen / 5) {
				return 0
			}
		}
	}
	if nonMatches > (nameLen / 5) {
		return 0
	}

	if nonMatches == 0 {
		return 1
	}
	return 2
}

func removeFileExtension(Name string) string {
	return strings.TrimSuffix(Name, filepath.Ext(Name))
}
