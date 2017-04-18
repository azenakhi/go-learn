package common

import (
	"encoding/json"
	"log"
	"os"
)

type configuration struct {
	SecretKey  string `json:"secret_key"`
	JenkinsURL string `json:"jenkins_url"`
}

var AppConfig configuration

func init() {
	loadConfig()
}

func loadConfig() {
	file, err := os.Open("common/config.json")
	if err != nil {
		log.Fatalf("[loadConfig] %s\n", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	AppConfig = configuration{}
	err = decoder.Decode(&AppConfig)
	if err != nil {
		log.Fatalf("[loadAppConfig] %s\n", err)
	}
}
