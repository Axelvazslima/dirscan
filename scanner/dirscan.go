package scanner

import "os"

func DirScan(fileName string) []string {
	curDir, err := os.Getwd()
	if err != nil {
		return []string{"Error getting current directory:" + err.Error()}
	}
	return TreeSearching(curDir, fileName)
}
