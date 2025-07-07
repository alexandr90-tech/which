package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		return
	}
	filesToSearch := arguments[1:]

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, file := range filesToSearch {
		found := false
		fmt.Printf("Searching for: %s\n", file)

		for _, directory := range pathSplit {
			fullPath := filepath.Join(directory, file)
			// Does it exist?
			fileInfo, err := os.Stat(fullPath)
			if err == nil {
				mode := fileInfo.Mode()
				// Is it a regular file?
				if mode.IsRegular() && mode&0111 != 0 {
					fmt.Println(" Found:", fullPath)
					found = true
				}
			}
		}

		if !found {
			fmt.Println("  Executable not found in PATH")
		}
	}
}
