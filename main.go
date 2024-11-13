// SPDX-License-Identifier: GPL-2.0-or-later
package main

import (
	"flag"
	"fmt"
	"nixmkdoc/utils"
	"sync"
)

const version = "0.0.1"

func main() {
	// Define command-line flags
	fileFlag := flag.String("file", "", "Path to a single .nix file")
	dirFlag := flag.String("dir", "", "Path to a directory containing .nix files")
	flag.StringVar(fileFlag, "f", "", "Alias for -file")
	flag.StringVar(dirFlag, "d", "", "Alias for -dir")

	// Custom help message
	flag.Usage = func() {
		fmt.Printf(`nixMkDoc - Version %s
A tool to generate Markdown documentation from Nix files.

Usage:
  nixMkDoc [options]

Options:
  --file, -f <path>   Path to a single .nix file
  --dir, -d <path>    Path to a directory containing .nix files
  --help, -h          Show this help message

Example:
  nixMkDoc -file path/to/file.nix
  nixMkDoc -dir path/to/directory
`, version)
	}

	flag.Parse()

	// Show help if no file or directory flag is provided
	if *fileFlag == "" && *dirFlag == "" {
		flag.Usage()
		return
	}

	var wg sync.WaitGroup
	ch := make(chan utils.Option)
	var options []utils.Option

	// Process a single file or directory based on flags
	if *fileFlag != "" {
		utils.ParseFile(*fileFlag, *fileFlag, ch, &wg)
	} else if *dirFlag != "" {
		utils.ParseDirectory(*dirFlag, *dirFlag, ch, &wg)
	}

	// Close the channel once all goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	// Read options from the channel and append them to `options`
	for option := range ch {
		options = append(options, option)
	}

	// Write all options to the Markdown file
	fmt.Printf("Writing %d options to Markdown file.\n", len(options))
	utils.WriteMarkdown("options.md", options)
}
