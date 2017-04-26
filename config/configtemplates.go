package config

import (
)

const (
	SCSSConfigFilename = ".scss-lint.yml"
	JSConfigFilename = ".eslintrc.json"
)


func scssConfigTemplate() string {
	return ""
}

func jsConfigTemplate() string {
	return ""
}


func GetTemplateFn(filename string) (func() string) {

	switch(filename) {
		case SCSSConfigFilename: return scssConfigTemplate
		case JSConfigFilename: return jsConfigTemplate
	}

	return nil
}

