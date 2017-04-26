package linter

import (
	"os"
	"os/exec"
	"strings"
)

type Linter struct {
	Name string
	Args []string
}

func (linter *Linter) String() string {
	return "Command: " + linter.Name + " " + strings.Join(linter.Args[:], " ")
}

func (linter *Linter) Lint() error {

	command := exec.Command(linter.Name, linter.Args...)
	command.Stdout = os.Stdout;

	return command.Run()
}

