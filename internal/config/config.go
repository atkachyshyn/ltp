package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config stores the configuration settings retrieved from the config.yaml file.
type Config struct {
	ServerAddress string `yaml:"serverAddress"`
	KrakenAPIURL  string `yaml:"krakenAPIURL"`
}

// LoadConfig reads the specified config file and decodes the YAML configuration.
func LoadConfig(path string) (*Config, error) {
	configFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
