package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"encoding/json"
	"bytes"
)


const (
	// The indentation for the config file
	ConfigIndentation = "    "

	// All config file permissions
	ConfigFilePermission = 0777
)


//
// StupidsidConfig - Stupidsid configuration structure
//
type StupidsidConfig struct {

	// Name - The project name
	Name string

	// Description - Some information about the project
	Description string

	// Paths - Relative paths to the root field
	// { root, php, js, css }
	Paths map[string]string
}


//
// String - toString formatting
//
func (conf *StupidsidConfig) String() string {
	format := "{\n\tName: %s,\n\tDescription: %s,\n\tPaths: %+v\n}"
	return fmt.Sprintf(format, conf.Name, conf.Description, conf.Paths)
}


//
// LoadConfigFile - Load the stupidsid.json config file
// 
// params
// -- pathname {string}  The configuration file path
// 
// returns
// -- {*StupidsidConfig}
// -- {error}
//
func LoadConfigFile(pathname string) (*StupidsidConfig, error) {

	content, err := ioutil.ReadFile(pathname)

	if err != nil { return nil, err }

	conf := &StupidsidConfig{}

	// Config json to config instance
	err = json.Unmarshal(content, conf)

	if err != nil { return nil, err }

	conf.Paths["root"], _ = filepath.Abs(conf.Paths["root"])

	// All Paths inside config
	for key, val := range conf.Paths {
		if key != "root" {
			conf.Paths[key + "_absolute"] = filepath.Join(conf.Paths["root"], val)
		}
	}

	return conf, nil
}



//
// CompileConfigTemplate - Compile stupidsid config file from the instance
// 
// params
// -- conf {*StupidsidConfig}  Stupidsid config instance
// 
// returns
// -- {[]byte}
//
func CompileConfigTemplate(conf *StupidsidConfig) []byte {

	var indentedJSON bytes.Buffer

	content, _ := json.Marshal(conf)

	// Add indentation to config
	json.Indent(&indentedJSON, content, "", ConfigIndentation)

	return indentedJSON.Bytes()
}


//
// Create - Create a new config file from config instance
// 
// params
// -- pathname {string}  The path to the new file to create
// -- conf {*StupidsidConfig} Config instance
// 
// returns
// -- {error}
//
func Create(pathname string, conf *StupidsidConfig) error {

	content := CompileConfigTemplate(conf)

	err := ioutil.WriteFile(pathname, content, ConfigFilePermission)

	return err
}



//
// GenerateConfigFile - Generate a given config file
// 
// params
// -- filename {string}  The config filename
// -- dest {string}  The destination directory
// 
// returns
// -- {error}
//
func GenerateConfigFile(filename string, dest string) error {

	compile := GetTemplateFn(filename)
	if compile == nil { return nil }

	// Compile the template
	content := compile()

	fmt.Println(filepath.Join(dest, filename))
	fmt.Println([]byte(content))

	// Write the compiled template out to the destination
	err := ioutil.WriteFile(
		filepath.Join(dest, filename),
		[]byte(content),
		ConfigFilePermission,
	)

	return err
}

