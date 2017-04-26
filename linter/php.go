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

	phpRootAbsolutePath, _ := filepath.Abs(phpRoot)

	phpcsPath, _ := filepath.Join(phpRootAbsolutePath, "../", "vendor/bin/phpcs")

	return &Linter{

		Name: phpcsPath,

		Args: []string{
			phpRootAbsolutePath,
		},
	}
}
