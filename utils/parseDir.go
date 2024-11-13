// SPDX-License-Identifier: GPL-2.0-or-later
package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// ParseDirectory walks through the directory and recursively parses .nix files.
// It launches a goroutine for each .nix file found, calling ParseFile.
func ParseDirectory(rootDir, dirname string, ch chan<- Option, wg *sync.WaitGroup) {
	filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		// Handle any error accessing a path
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}

		// Check if the current file has a .nix extension and is not a directory
		if !info.IsDir() && filepath.Ext(path) == ".nix" {
			wg.Add(1)
			go ParseFile(rootDir, path, ch, wg)
		}
		return nil
	})
}
