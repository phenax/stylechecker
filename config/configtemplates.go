package config

import (
)

const (

	// scss-lint config file name
	SCSSConfigFilename = ".scss-lint.yml"

	// eslint config file name
	JSConfigFilename = ".eslintrc.json"
)



//
// GetTemplateFn - Get the template function for the file to generate
// 
// params
// -- filename {string}
// 
// returns
// -- {func() string}
//
func GetTemplateFn(filename string) (func() string) {

	switch(filename) {
		case SCSSConfigFilename: return scssConfigTemplate
		case JSConfigFilename: return jsConfigTemplate
	}

	return nil
}

//
// Sass config template fn
//
func scssConfigTemplate() string {
	return `
scss_files: 'static/css/**/*.scss'

exclude: 'static/css/plugins/**'

linters:
  BorderZero:
    enabled: false

  Indentation:
    enabled: true
    character: tab
    allow_non_nested_indentation: true
`
}


//
// js config template fn
//
func jsConfigTemplate() string {
	return `
{
    "env": {
        "browser": true,
        "commonjs": true,
        "es6": true,
        "node": true
    },
    "extends": "eslint:recommended",
    "parserOptions": {
        "sourceType": "module"
    },
    "rules": {
        "indent": [
            "error",
            "tab"
        ],
        "linebreak-style": [
            "error",
            "unix"
        ],
        "quotes": [
            "error",
            "single"
        ],
        "semi": [
            "error",
            "always"
        ]
    }
}
`
}


