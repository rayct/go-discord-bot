package config

import (
	"encoding/json"
	"fmt"
	"os"
)

var (
	Token     string
	BotPrefix string

	config *configStruct
)

type configStruct struct {
	Token     string
	BotPrefix string
}

func ReadConfig() error {

	file, err := os.ReadFile("config/config.json")

	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(file, &config)

	if err != nil {
		fmt.Println(err)
	}

	Token = config.Token
	BotPrefix = config.BotPrefix

	return nil
}
