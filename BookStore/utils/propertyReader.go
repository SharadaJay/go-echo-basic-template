package utils

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type Config struct {
	Port string `yaml:"port"`
}

var Cfg map[string]Config

func init() {
	readPropertyFile(&Cfg)
}

func readPropertyFile(cfg *map[string]Config) {
	path := os.ExpandEnv("properties/${ENV}.yml")
	byteArray, err := ioutil.ReadFile(path)
	if err != nil {
		handleError(err)
	}

	err = yaml.Unmarshal(byteArray, cfg)
	if err != nil {
		handleError(err)
	}
}

func handleError(err error) {
	log.Fatalf("FATAL_ERROR : %v", err)
}
