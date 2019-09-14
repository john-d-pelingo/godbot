package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {
	Token  string `json:"token"`
	User   bool   `json:"user"`
	Prefix string `json:"prefix"`
}

func initConfig(configFile string) *config {
	file, err := ioutil.ReadFile(configFile)
	errCheck(fmt.Sprintf("Unable to open file: %s.", configFile), err)

	config := config{}

	err = json.Unmarshal(file, &config)
	errCheck(fmt.Sprintf("Unable to JSON parse file: %s.", configFile), err)

	return &config
}
