package utils

import (
	"bufio"
	"regexp"
	"strings"
)

// buildFullOptionName constructs the full name of an option based on the context stack.
func BuildFullOptionName(context []string, line string) string {
	var fullNameParts []string
	nameRegex := regexp.MustCompile(`([\w\.]+)\s*=\s*lib\.mkOption`)
	matches := nameRegex.FindStringSubmatch(line)

	// Append the matched option name to the full name parts, if available.
	if len(matches) > 1 {
		fullNameParts = append(fullNameParts, matches[1])
	}

	// Traverse the context stack in reverse to build the option's full name.
	for i := len(context) - 1; i >= 0; i-- {
		if context[i] == "options" {
			break
		}
		fullNameParts = append([]string{context[i]}, fullNameParts...)
	}

	// Join all parts to form the complete option name.
	return strings.Join(fullNameParts, ".")
}

// ExtractContextName retrieves the context name from a line that starts a new block.
func ExtractContextName(line string) string {
	parts := strings.Split(line, "=")
	if len(parts) > 1 {
		return strings.TrimSpace(parts[0])
	}
	return ""
}

// ExtractFieldValue retrieves the value of a field from a single line.
func ExtractFieldValue(line string) string {
	parts := strings.SplitN(line, "=", 2)
	if len(parts) > 1 {
		value := strings.TrimSpace(parts[1])
		return strings.Trim(value, `";`)
	}
	return ""
}

// ExtractMultilineValue reads a multiline value enclosed by " from a scanner.
func ExtractMultilineValue(scanner *bufio.Scanner) string {
	var description strings.Builder
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "'';" {
			break
		}
		description.WriteString(line + "\n")
	}
	return strings.TrimSpace(description.String())
}
