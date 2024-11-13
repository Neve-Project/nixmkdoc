package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

// WriteMarkdown generates a Markdown file containing the options found in .nix files.
// It organizes options by directory and file, outputting each with its metadata.
func WriteMarkdown(filename string, options []Option) {
	outputFile, err := os.Create(filename)
	if err != nil {
		fmt.Printf("Error creating markdown file: %v\n", err)
		return
	}
	defer outputFile.Close()

	// Group options by directory and file for structured output
	optionsByDir := make(map[string]map[string][]Option)
	for _, opt := range options {
		dir := filepath.Dir(opt.FileName)
		if _, ok := optionsByDir[dir]; !ok {
			optionsByDir[dir] = make(map[string][]Option)
		}
		optionsByDir[dir][opt.FileName] = append(optionsByDir[dir][opt.FileName], opt)
	}

	// Sort directories alphabetically for consistent output order
	dirs := make([]string, 0, len(optionsByDir))
	for dir := range optionsByDir {
		dirs = append(dirs, dir)
	}
	sort.Strings(dirs)

	// Iterate through sorted directories and files to write options
	for _, dir := range dirs {
		files := make([]string, 0, len(optionsByDir[dir]))
		for file := range optionsByDir[dir] {
			files = append(files, file)
		}
		sort.Strings(files)

		for _, file := range files {
			// Write file section header
			fmt.Fprintf(outputFile, "## %s\n\n", file)

			// Write each option's details
			for _, opt := range optionsByDir[dir][file] {
				fmt.Fprintf(outputFile, "### %s\n", opt.Name)
				fmt.Fprintf(outputFile, "- **type**: %s\n", opt.Type)
				fmt.Fprintf(outputFile, "- **default**: %s\n", opt.Default)
				fmt.Fprintf(outputFile, "- **description**: %s\n", opt.Description)
				if opt.Example != "" {
					fmt.Fprintf(outputFile, "- **example**: %s\n", opt.Example)
				}
				fmt.Fprintln(outputFile)
			}
		}
	}
}
