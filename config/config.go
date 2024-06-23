package config

import (
	"embed"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Domain string `yaml:"domain"`
	Port   string `yaml:"port"`
	Auth   struct {
		Token string `yaml:"token"`
	} `yaml:"auth"`
}

//go:embed config.yaml
var EmbeddedConfig embed.FS

var Data Config

func init() {
	var bytes []byte
	var err error
	bytes, err = os.ReadFile("config/config.yaml")
	if err != nil {
		fmt.Println("读取外部配置失败")
		bytes, err = EmbeddedConfig.ReadFile("config.yaml")
		if err != nil {
			log.Fatalf("Error reading embedded config file: %v", err)
		}
	}
	err = yaml.Unmarshal(bytes, &Data)
	if err != nil {
		log.Fatalf("Error parsing config file: %v", err)
	}
}
