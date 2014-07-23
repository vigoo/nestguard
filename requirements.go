package main

type Requirements map[string]string

// Creates a new, empty requirement map
func noRequirements() Requirements {
	return make(map[string]string)
}

// Merge a set of requirements into an existing set of requirements
func addRequirements(oldReqs Requirements, newReqs Requirements) {
	for k, v := range newReqs {
		oldReqs[k] = v
	}
}
