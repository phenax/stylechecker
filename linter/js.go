package linter

import (
	"path/filepath"
)

func NewJSLinter(jsRoot string) (*Linter) {

	path, _ := filepath.Abs(".")
	jsRootAbsolutePath, _ := filepath.Abs(jsRoot)

	return &Linter{

		Name: "eslint",

		Args: []string{
			"--config",
			filepath.Join(path, ".eslintrc.json"),
			jsRootAbsolutePath,
		},
	}
}
