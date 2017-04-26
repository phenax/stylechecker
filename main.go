package main

import (
	"fmt"
	"github.com/phenax/stylechecker/linter"
)


func main() {

	eslint := linter.NewJSLinter()

	output, err := eslint.Lint()

	if err != nil {
		fmt.Println("error", err.Error());
	}

	fmt.Println(string(output));

}
