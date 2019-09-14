package config

import (
	"encoding/json"
	"fmt"
	"github.com/john-d-pelingo/godbot/helpers"
	"io/ioutil"
)

type config struct {
	Token  string `json:"token"`
	User   bool   `json:"user"`
	Prefix string `json:"prefix"`
}

func InitConfig(configFile string) *config {
	file, err := ioutil.ReadFile(configFile)
	helpers.ErrCheck(fmt.Sprintf("Unable to open file: %s.", configFile), err)

	config := config{}

	err = json.Unmarshal(file, &config)
	helpers.ErrCheck(fmt.Sprintf("Unable to JSON parse file: %s.", configFile), err)

	return &config
}
