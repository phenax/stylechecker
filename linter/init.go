package linter

import (
	"os/exec"
)

type Linter struct {
	name string
	args []string
}


func (linter *Linter) Lint() (string, error) {

	stdout, err := exec.
		Command(linter.name, linter.args...).
		Output()

	return string(stdout), err
}

