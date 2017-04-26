package linter

import (
	"path/filepath"
)

//
// NewCSSLinter - Create a new linter instance for linting scss
// 
// params
// -- scssRoot {string} SCSS source root directory
// 
// returns
// -- {*Linter}
//
func NewCSSLinter(scssRoot string) (*Linter) {

	scssRootAbsolutePath, _ := filepath.Abs(scssRoot)

	return &Linter{

		Name: "scss-lint",

		Args: []string{
			scssRootAbsolutePath,
		},
	}
}
