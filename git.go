package main

import (
	"log"
	"os/exec"
	"strings"
)

// Creates a list of warnings for uncommitted changes in the
// git repository pointed by path. The repoName parameter is
// stored in the generated warning structures.
func uncommittedChanges(path string, repoName string) []Warning {
	git_status := exec.Command("git", "status", "--porcelain")
	git_status.Dir = path
	out, err := git_status.Output()

	if err != nil {
		log.Fatal(err)
		return make([]Warning, 0)
	} else {
		lines := strings.Split(string(out), "\n")
		result := make([]Warning, 0, len(lines))
		for i := 0; i < len(lines); i++ {
			trimmed := strings.TrimSpace(lines[i])
			if !strings.HasPrefix(trimmed, "??") &&
				len(trimmed) > 0 {
				result = append(result, Warning{
					message: trimmed,
					repo:    repoName})
			}
		}

		return result
	}
}

// Gets the HEAD commit's hash of the git repository pointed by path,
// or returns ??? if it could not read it.
func headCommitHash(path string) string {
	git_log := exec.Command("git", "log", "-n", "1", "--pretty=oneline")
	git_log.Dir = path
	out, err := git_log.Output()

	if err != nil {
		log.Fatal(err)
		return "???"
	} else {
		fields := strings.Fields(string(out))
		return fields[0]
	}
}

// Gets the current branch of the git repository pointed by path, or returns ???
// if it could not read it.
func currentBranch(path string) string {
	git_rp := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	git_rp.Dir = path
	out, err := git_rp.Output()

	if err == nil {
		return strings.Split(string(out), "\n")[0]
	} else {
		return "???"
	}
}
