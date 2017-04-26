package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/phenax/stylechecker/config"
	"github.com/phenax/stylechecker/linter"
)


func main() {

	args := os.Args[:]
	var configFile string;

	if(len(args) >= 2) {

		switch(args[1]) {
			case "init": {
				configInit()
				return;
			}
		}

		configFile, _ = filepath.Abs(args[1])
	} else {
		configFile, _ = filepath.Abs("./stupidsid.json")
	}

	conf, err := config.LoadConfigFile(configFile)
	if err != nil { panic(err) }

	jsStatus := useLinter(linter.NewJSLinter(conf.Paths["js"]))

	if(jsStatus) {
		colorPrint(32, "You are good to go. Push away my friend.\n");
	}
}

func useLinter(linter *linter.Linter) bool {

	fmt.Println("\n", linter, "\n")

	err := linter.Lint()

	if(err != nil) {
		colorPrint(31, "Not so fast. Fix all of this before the push\n\n")
		return false
	}

	return true
}


func colorPrint(color int, str string) {
	fmt.Printf("\033[%dm%s\033[0m", color, str)
}


func configInit() {

	path, _ := filepath.Abs("./stupidsid.json")

	err := config.Create(path, &config.StupidsidConfig{
		Name: "foo",
		Description: "baar",
		Paths: map[string]string {
			"root": ".",
			"js": ".",
		},
	})

	if err != nil { panic(err) }

}
