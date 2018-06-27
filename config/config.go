package config

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

type ClientConfig struct {
	EndpointURL string `yaml:"endpointUrl"`
	AuthToken   string `yaml:"authToken"`
}

func Create(fileName string) (*ClientConfig, error) {
	return getFileConfig(fileName)
}

func getFileConfig(fileName string) (*ClientConfig, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	config := new(ClientConfig)
	marshalErr := yaml.Unmarshal(data, config)

	if marshalErr != nil {
		return nil, marshalErr
	}

	return config, nil
}
