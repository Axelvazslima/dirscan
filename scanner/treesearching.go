package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func TreeSearching(path string, fileName string) []string {
	queue := []string{path}
	var output []string

	for len(queue) > 0 {
		curDir := queue[0]
		queue = queue[1:]

		ls, err := os.ReadDir(curDir)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			continue
		}
		for _, entry := range ls {
			fullPath := filepath.Join(curDir, entry.Name())

			if !entry.IsDir() {
				switch nameCorrespondence(entry.Name(), fileName) {
				case 1:
					return []string{fullPath + " -> Exact Match"}
				case 2:
					output = append(output, fullPath+" -> Possible Match")
				}
			} else {
				queue = append(queue, fullPath)
			}
		}
	}

	if len(output) > 0 {
		return output
	}
	return []string{"File not found"}
}

func nameCorrespondence(name string, nameToMatch string) int {
	name, nameToMatch = strings.ToLower(name), strings.ToLower(nameToMatch)
	nameLen, namenameToMatchLen := len(name), len(nameToMatch)

	if nameLen < (namenameToMatchLen-1) || nameLen > (namenameToMatchLen+1) {
		return 0
	}

	nonMatches := 0
	for i := range nameLen {
		if name[i] != nameToMatch[i] {
			nonMatches++
			if nonMatches > (nameLen / 5) {
				return 0
			}
		}
	}

	if nonMatches == 0 {
		return 1
	}
	return 2
}
