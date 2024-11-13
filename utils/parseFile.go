// SPDX-License-Identifier: GPL-2.0-or-later
package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// ParseFile reads a single .nix file and sends the options found to a channel.
// It uses a context stack to manage nested option scopes and extracts option details.
func ParseFile(rootDir, filename string, ch chan<- Option, wg *sync.WaitGroup) {
	defer wg.Done()

	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return
	}
	defer file.Close()

	// Get the relative path for file identification in outputs.
	relPath, err := filepath.Rel(rootDir, filename)
	if err != nil {
		relPath = filepath.Base(filename)
	}

	scanner := bufio.NewScanner(file)
	var context []string // Stack to manage nested option context
	inOption := false    // Flag indicating start of a new option
	var option Option    // Reusable variable for each option

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Handle the start of a new context block (indicated by '{')
		if strings.HasSuffix(line, "{") && !strings.Contains(line, "lib.mkOption {") {
			contextName := ExtractContextName(line)
			if contextName != "" {
				context = append(context, contextName)
			}
			continue
		} else if line == "};" && inOption { // End of the active option
			// Complete the option and send it to the channel
			if option.Name != "" {
				ch <- option
			}
			inOption = false
			continue
		} else if line == "};" && len(context) > 0 { // Handle end of a context block
			context = context[:len(context)-1]
			continue
		}

		// Detect the start of a new option block `lib.mkOption`
		if strings.Contains(line, "lib.mkOption {") {
			inOption = true
			option = Option{
				FileName: relPath,
			}
			option.Name = BuildFullOptionName(
				context,
				line,
			)
			option.Name = strings.Trim(option.Name, ".")
		}

		// Parse the attributes of the active option
		if inOption {
			if strings.Contains(line, "type =") {
				option.Type = ExtractFieldValue(line)
				if strings.HasPrefix(option.Type, "lib.types.") {
					option.Type = strings.TrimPrefix(option.Type, "lib.types.")
				}
			} else if strings.Contains(line, "default =") {
				option.Default = ExtractFieldValue(line)
			} else if strings.Contains(line, "description = ''") {
				option.Description = ExtractMultilineValue(scanner)
			} else if strings.Contains(line, "example =") {
				option.Example = ExtractFieldValue(line)
			}
		}
	}

	// Handle any scanning error
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
	}
}
