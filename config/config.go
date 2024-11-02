package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Directory Directory `yaml:"directory"`
	Files     []File    `yaml:"files"`
}

type Directory struct {
	Name string `yaml:"name"`
	Path string `yaml:"path"`
}

type File struct {
	File string `yaml:"file"`
	Name string `yaml:"name"`
}

type ConfigLoader struct {
	path string
}

func NewConfigLoader(path string) *ConfigLoader {
	return &ConfigLoader{path: path}
}

func (cl *ConfigLoader) Load() (*Config, error) {
	data, err := os.ReadFile(cl.path)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
