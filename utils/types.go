// SPDX-License-Identifier: GPL-2.0-or-later
package utils

// Option represents an extracted option from a .nix file.
// Each Option contains metadata about its location, type, default value, description, and example usage.
type Option struct {
	FileName    string // The relative file path of the .nix file where the option is defined
	Name        string // The full name of the option, including context hierarchy
	Type        string // The data type of the option (e.g., bool, int, str)
	Default     string // The default value assigned to the option
	Description string // A description explaining the purpose of the option
	Example     string // An example of how to set or use the option
}
