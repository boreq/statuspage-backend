// Package config holds the global config.
package config

import (
	"encoding/json"
	"io/ioutil"
)

type ConfigStruct struct {
	Debug              bool
	ServeAddress       string
	ScriptsDirectory   string
	UpdateEverySeconds int
}

// Config points to the current config struct used by the other parts of the
// program.
var Config *ConfigStruct = Default()

// Default returns the default config.
func Default() *ConfigStruct {
	conf := &ConfigStruct{
		Debug:              false,
		ServeAddress:       "127.0.0.1:8118",
		ScriptsDirectory:   "/path/to/scripts/directory",
		UpdateEverySeconds: 30,
	}
	return conf
}

// Load loads the config from the specified json file. If certain keys are not
// present in the loaded config file the current values of the config struct
// are preserved.
func Load(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, Config)
}
