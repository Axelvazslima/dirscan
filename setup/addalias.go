package setup

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func AddAlias() string {
	return addAlias()
}

func addAlias() string {
	// Get the user's home directory.
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "Error getting home directory:" + err.Error()
	}

	// Construct the full path to ~/.bashrc.
	bashrcPath := filepath.Join(homeDir, ".bashrc")

	// Define the alias we want to insert.
	aliasLine := "alias godirscan='go run ~/dirscan/main.go'"

	// Define a marker in your .bashrc that indicates where snippets start.
	// The alias will be inserted immediately before the line that contains this marker.
	// Adjust this marker string to fit your .bashrc file.
	marker := "# Snippets start"

	// Open .bashrc for reading.
	file, err := os.Open(bashrcPath)
	if err != nil {
		return "Error opening .bashrc:" + err.Error()
	}
	defer file.Close()

	// Read the file line by line.
	scanner := bufio.NewScanner(file)
	var lines []string
	aliasExists := false
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		// Check if the alias already exists.
		if strings.Contains(line, "alias godirscan=") {
			aliasExists = true
		}
	}
	if err := scanner.Err(); err != nil {
		return "Error reading .bashrc:" + err.Error()
	}

	if aliasExists {
		return "Alias 'godirscan' already exists in .bashrc"
	}

	// Find the insertion index based on the marker.
	insertIndex := -1
	for i, line := range lines {
		if strings.Contains(line, marker) {
			insertIndex = i
			break
		}
	}

	// If the marker wasn't found, choose a safe fallback (for example, insert in the middle).
	if insertIndex == -1 {
		insertIndex = len(lines) / 2
		fmt.Println("Marker not found; inserting alias at line", insertIndex)
	}

	// Insert the aliasLine into the slice of lines.
	newLines := append(lines[:insertIndex], append([]string{aliasLine}, lines[insertIndex:]...)...)

	// Write the modified content back to .bashrc.
	output := strings.Join(newLines, "\n") + "\n"
	if err := os.WriteFile(bashrcPath, []byte(output), 0644); err != nil {
		return "Error writing to .bashrc:" + err.Error()
	}

	return "Alias 'godirscan' added successfully to .bashrc"
}
