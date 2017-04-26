package linter

import (
	"path/filepath"
)

//
// NewPHPLinter - Create a new linter instance for linting php
// 
// params
// -- phpRoot {string} PHP source root directory
// 
// returns
// -- {*Linter}
//
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
