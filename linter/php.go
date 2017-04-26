package linter

import (
	"path/filepath"
)

func NewPHPLinter(phpRoot string) (*Linter) {

	path, _ := filepath.Abs(".")
	jsRootAbsolutePath, _ := filepath.Abs(phpRoot)

	return &Linter{

		Name: "eslint",

		Args: []string{
			"--config",
			filepath.Join(path, ".eslintrc.json"),
			jsRootAbsolutePath,
		},
	}
}
