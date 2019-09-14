package config

import (
	"encoding/json"
	"fmt"
	"github.com/john-d-pelingo/godbot/helpers"
	"io/ioutil"
)

// Config represents the structure of the config file.
type Config struct {
	Token  string `json:"token"`
	User   bool   `json:"user"`
	Prefix string `json:"prefix"`
}

// InitConfig reads the file named by configFile and returns the configuration.
func InitConfig(configFile string) *Config {
	file, err := ioutil.ReadFile(configFile)
	helpers.ErrCheck(fmt.Sprintf("Unable to open file: %s.", configFile), err)

	config := Config{}

	err = json.Unmarshal(file, &config)
	helpers.ErrCheck(fmt.Sprintf("Unable to JSON parse file: %s.", configFile), err)

	return &config
}
