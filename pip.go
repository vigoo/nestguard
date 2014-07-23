package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var ghReqRegexp = regexp.MustCompile(`-e git\+git@github\.com\:[0-9A-Za-z_\-]+/[0-9A-Za-z_\-]+\.git@([0-9A-Za-z]*)\#egg=([0-9A-Za-z_\-]+)`)

// Parse a line of the requirements file. Only supports github repositories currently.
func parseRequirement(line string) (string, string, bool) {

	if len(line) > 0 {
		if !strings.HasPrefix(line, "#") {
			result := ghReqRegexp.FindStringSubmatch(line)

			if len(result) == 3 {
				return strings.Replace(result[2], "_", "-", -1), result[1], true
			}
		}
	}

	return "", "", false
}

// Readln returns a single line (without the ending \n)
// from the input buffered reader.
// An error is returned if there is an error with the
// buffered reader.
func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

// Read a PIP requirements file and returns the map of requirements
// pointing to github repositories
func readRequirements(path string) Requirements {
	result := noRequirements()

	f, err := os.Open(path)
	if err == nil {
		r := bufio.NewReader(f)
		s, e := Readln(r)

		for e == nil {
			dep, ver, res := parseRequirement(s)
			if res {
				result[dep] = ver
			}
			s, e = Readln(r)
		}
	} else {
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
	}

	return result
}

// Reads two pip requirement files (for example one for development, one for
// production), and produces warnings if they differ in github repo dependencies
func readMultipleRequirements(path1 string, path2 string) (Requirements, Warnings) {
	warnings := noWarnings()
	req1 := readRequirements(path1)
	req2 := readRequirements(path2)

	if len(req1) == 0 {
		return req2, warnings
	} else if len(req2) == 0 {
		return req1, warnings
	} else {
		for k, v2 := range req2 {
			if v1, ok := req1[k]; ok {
				// k is also in req1
				if v1 != v2 {
					warnings = append(warnings, Warning{
						message: fmt.Sprintf("Dependency %s differs in %s and %s", k, filepath.Base(path1), filepath.Base(path2)),
						repo:    "root"})
				}
			} else {
				req1[k] = v2
			}
		}

		return req1, warnings
	}
}
