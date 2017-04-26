package main

import (
	"fmt"
	"os"
	// "io/ioutil"
	"github.com/phenax/stylechecker/linter"
)


func main() {

	args := os.Args[:]

	fmt.Println(args);

	// ioutil


	eslint := linter.NewJSLinter("/var/www/html/stupidsid/static/js")
	fmt.Println(eslint);

	_, err := eslint.Lint()

	if err != nil {
		// Error found
	}
}
