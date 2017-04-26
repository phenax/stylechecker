package linter

import (
	"os"
	"os/exec"
	"strings"
)


//
// Linter - Linter instance
//
type Linter struct {
	// Name - Linter name/path to linter
	Name string
	// Args - command line arguements
	Args []string
}


//
// String - toString fn
//
func (linter *Linter) String() string {
	return "Command: " + linter.Name + " " + strings.Join(linter.Args[:], " ")
}

//
// Lint - Run the linter
//
func (linter *Linter) Lint() error {
	return linter.Exec()
}

//
// Excute the command given in the linter config
//
func (linter *Linter) Exec() error {
	command := exec.Command(linter.Name, linter.Args...)
	command.Stdout = os.Stdout;

	return command.Run()
}
