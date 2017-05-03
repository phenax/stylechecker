package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/phenax/stylechecker/config"
	"github.com/phenax/stylechecker/linter"
)


const (

	// ErrorMessage - Color for error messages
	ErrorMessage = 31

	// SuccessMessage - Color for success messages
	SuccessMessage = 32

	// InfoMessage - Color for info
	InfoMessage = 38
)



// Main fn
func main() {

	args := os.Args[:]
	var configFile string

	// Make sense of the arguements
	if(len(args) >= 2) {
		switch(args[1]) {
			case "init": {
				configInit()
				return
			}
			default: configFile, _ = filepath.Abs(args[1])
		}
	} else {
		// By default look for the config in the current directory
		configFile, _ = filepath.Abs("./stupidsid.json")
	}

	// Load the config file and create the instance
	conf, err := config.LoadConfigFile(configFile)

	if err != nil {
		colorPrint(ErrorMessage, err.Error() + "\n")
		return
	}

	// Language specific linting goes here
	isPHPSafe := useLinter(linter.NewPHPLinter(conf.Paths["php"]))
	isJSSafe := useLinter(linter.NewJSLinter(conf.Paths["js"]))
	isCSSSafe := useLinter(linter.NewCSSLinter(conf.Paths["css"]))

	// Error message
	if isPHPSafe && isJSSafe && isCSSSafe {
		// TODO: git push option(if a flag is passed, push automatically)
		colorPrint(SuccessMessage, "You are good to go. Push away my friend.\n")
		return
	}


	// Show what went wrong
	errorLangs := ""

	if(!isPHPSafe) { errorLangs+= ", PHP" }
	if(!isJSSafe) { errorLangs+= ", JS" }
	if(!isCSSSafe) { errorLangs+= ", CSS" }

	colorPrint(ErrorMessage, "\n\nNot so fast. Fix all of this before the push\n")
	colorPrint(InfoMessage, "Error(s) in " + errorLangs + "\n")
}


//
// useLinter - Use the linter passed
// 
// params
// -- linter {*linter.Linter}  The linter instance
// 
// returns
// -- {bool} Operation status
//
func useLinter(linter *linter.Linter) bool {

	fmt.Println("\n", linter, "\n")

	// Execute the linter
	err := linter.Lint()

	return err == nil
}


//
// colorPrint - Print a text with color
// 
// params
// -- color {int}
// -- str {string}
//
func colorPrint(color int, str string) {
	fmt.Printf("\033[%dm%s\033[0m", color, str)
}



//
// configInit - Initialize the config files for the root directory
//
func configInit() {

	absoluteRoot, _ := filepath.Abs(".")
	configPath := filepath.Join(absoluteRoot, "stupidsid.json")

	// Create a default stupidsid.json file
	err := config.Create(configPath, &config.StupidsidConfig{
		Name: "project-name",
		Description: "Some information about the project",
		Paths: map[string]string {
			"root": ".",
			"js": "./static/js",
			"css": "./static/css",
			"php": "./src",
		},
	})

	if err != nil { panic(err) }

	// ESLint config file generation
	_, err = exec.LookPath("eslint")

	if err != nil {
		colorPrint(ErrorMessage, "You need to install eslint\n")
	}

	// Generate the sass config file
	errCSS := config.GenerateConfigFile(
		config.SCSSConfigFilename,
		absoluteRoot,
	)

	// Generate the js config file
	errJS := config.GenerateConfigFile(
		config.JSConfigFilename,
		absoluteRoot,
	)

	// Error check
	if(errCSS != nil || errCSS != nil) {
		fmt.Println(errCSS, errJS)
		return
	}
}
