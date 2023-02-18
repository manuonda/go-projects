package settings

import "gopkg.in/yaml.v2"

var settingsFile []byte

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type Settings struct {
	Port string         `yaml:"port"`
	DB   DatabaseConfig `yaml:"database"`
}

func New() (*Settings, error) {
	var s Settings
	err := yaml.Unmarshal(settingsFile, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}
