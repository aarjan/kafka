package config

import (
	"io/ioutil"

	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

var (
	conf AppConfig
)

// Load load all the config for the application
func Load() AppConfig {
	conf, err := readEnv()
	if err != nil {
		log.Fatal("Error While Loading config from Env - " + err.Error())
	}
	return conf
}

// Read Config
func Read() AppConfig {
	return conf
}

// Read info file to read build release file
func readInfoFile() string {
	data, err := ioutil.ReadFile("data/info.json") // just pass the file name
	if err != nil {
		log.Fatal("Error While opening data/info.json file - " + err.Error())
		return ""
	}
	return string(data)
}

// Read Env on App start
func readEnv() (AppConfig, error) {
	defaults()
	err := envconfig.Process("event", &conf)
	if err != nil {
		return AppConfig{}, err
	}
	return conf, nil
}

// defaults add the default value to the Config
func defaults() {
	conf.ESHost = "localhost"
	conf.ESPort = "9200"
	conf.ListenHost = "localhost"
	conf.ListenPort = 8080
}
