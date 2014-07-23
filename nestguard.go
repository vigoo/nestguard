package main

import (
	"os"
	"path/filepath"
)

// Process a nested git repository and produces a set of warnings
func processRepo(path string, repoName string, reqs Requirements) []Warning {

	// Uncommitted chaneges are treated a warnings
	result := uncommittedChanges(path, repoName)

	// If we have a requirement definition for this repository and it does
	// not match the head commit or current branch name, it is also a warning
	head := headCommitHash(path)
	branch := currentBranch(path)

	req, exists := reqs[repoName]
	if exists {
		if req != head && req != branch {
			result = append(result, Warning{
				message: "HEAD revision differs from requirements",
				repo:    repoName})
		}
	}

	return result
}

func main() {

	warnings := noWarnings()
	reqs := noRequirements()

	// PIP
	pipReqs, pipWarnings := readMultipleRequirements("requirements.txt", "requirements-production.txt")
	addRequirements(reqs, pipReqs)
	warnings = append(warnings, pipWarnings...)

	// ...add other dpeendency management implementations here...

	// Looking for nested git repos
	filepath.Walk(".", func(path string, _ os.FileInfo, _ error) error {
		if filepath.Base(path) == ".git" {
			var nestedRepo = filepath.Dir(path)
			var nestedRepoName = filepath.Base(nestedRepo)

			if nestedRepoName != "." {
				repoWarns := processRepo(nestedRepo, nestedRepoName, reqs)
				warnings = append(warnings, repoWarns...)
			}
		}

		return nil
	})

	// Printing the results
	if hasWarnings(warnings) {
		printWarnings(warnings)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
