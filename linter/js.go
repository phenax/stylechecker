package linter

import (
	"path/filepath"
)

//
// NewJSLinter - Create a new linter instance for linting js
// 
// params
// -- jsRoot {string} JS source root directory
// 
// returns
// -- {*Linter}
//
func NewJSLinter(jsRoot string) (*Linter) {

	path, _ := filepath.Abs(".")
	jsRootAbsolutePath, _ := filepath.Abs(jsRoot)

	return &Linter{

		Name: "eslint",

		Args: []string{
			"--config",
			filepath.Join(path, ".eslintrc.json"),
			jsRootAbsolutePath,
			"--ignore-pattern",
			filepath.Join(jsRoot, "vendor/*"),
		},
	}
}
