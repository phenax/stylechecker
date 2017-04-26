package linter

import (
	"os/exec"
	"os"
	"strings"
)

type Linter struct {
	Name string
	Args []string
}

func (linter *Linter) String() string {
	return "Command: " + linter.Name + " " + strings.Join(linter.Args[:], " ")
}


func (linter *Linter) Lint() ([]byte, error) {

	command := exec.Command(linter.Name, linter.Args...)
	command.Stdout = os.Stdout;

	err := command.Run()

	return nil, err
}

