package config

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"encoding/json"
	"bytes"
)

type StupidsidConfig struct {
	Name string;
	Description string;
	Paths map[string]string;
}


func (conf *StupidsidConfig) String() string {
	format := "{\n\tName: %s,\n\tDescription: %s,\n\tPaths: %+v\n}"
	return fmt.Sprintf(format, conf.Name, conf.Description, conf.Paths);
}


func LoadConfigFile(pathname string) (*StupidsidConfig, error) {

	content, err := ioutil.ReadFile(pathname)

	if err != nil { return nil, err }

	conf := &StupidsidConfig{}
	err = json.Unmarshal(content, conf)

	if err != nil { return nil, err }

	conf.Paths["root"], _ = filepath.Abs(conf.Paths["root"])

	for key, val := range conf.Paths {
		if key != "root" {
			conf.Paths[key] = filepath.Join(conf.Paths["root"], val)
		}
	}

	return conf, nil
}


func CompileConfigTemplate(conf *StupidsidConfig) []byte {

	content, _ := json.Marshal(conf)

	var indentedJSON bytes.Buffer;

	json.Indent(&indentedJSON, content, "", "    ")

	return indentedJSON.Bytes();
}


func Create(pathname string, conf *StupidsidConfig) error {

	content := CompileConfigTemplate(conf)

	err := ioutil.WriteFile(pathname, content, 0777)

	return err
}

