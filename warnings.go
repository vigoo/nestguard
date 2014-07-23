package main

import (
	"fmt"
)

type Warning struct {
	message string
	repo    string
}

type Warnings []Warning

// Creates an empty list of warnings
func noWarnings() Warnings {
	return make([]Warning, 0)
}

// Checks if there is any warnings
func hasWarnings(warnings Warnings) bool {
	return len(warnings) > 0
}

// Prints the list of warnings in human readable format
func printWarnings(warnings Warnings) {
	fmt.Printf("Warnings:\n")

	for _, warning := range warnings {
		fmt.Printf("%s: %s\n", warning.repo, warning.message)
	}
}
