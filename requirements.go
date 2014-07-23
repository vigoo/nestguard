package main

type Requirements map[string]string

func noRequirements() Requirements {
	return make(map[string]string)
}

func addRequirements(oldReqs Requirements, newReqs Requirements) {
	for k, v := range newReqs {
		oldReqs[k] = v
	}
}
