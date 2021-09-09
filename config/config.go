package config

import (
	"fmt"
	"github.com/api_base/tool/database"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

const (
	filePathFormat = "%s/config/%s.yml"
)

type Config struct {
	Database database.Config `yaml:"database"`
}

func NewConfig() Config {
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatalf("get wd error: %v ", err)
	}
	if basePath == "/" {
		basePath = ""
	}
	filePath := fmt.Sprintf(filePathFormat, basePath, "local")

	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("read yaml file error: %v ", err)
	}

	configuration := Config{}
	if err := yaml.Unmarshal(configFile, &configuration); err != nil {
		log.Fatalf("error: %v", err)
	}
	return configuration
}
