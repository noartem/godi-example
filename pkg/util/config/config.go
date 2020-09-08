package config

import (
	"github.com/noartem/godi"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

// IPath path to config file
type IPath string

// Path default config path
const Path IPath = "./cmd/api/config.yml"

// Config application config file
type Config struct {
	Port uint `json:"port" yaml:"port"`

	SwaggerPath string `json:"swagger_path" yaml:"swagger_path"`

	DB struct {
		Host     string `json:"host" yaml:"host"`
		Port     uint   `json:"port" yaml:"port"`
		User     string `json:"user" yaml:"user"`
		Name     string `json:"name" yaml:"name"`
		Password string `json:"password" yaml:"password"`
	} `json:"db" yaml:"db"`

	JWT struct {
		Algorithm  string `json:"algo" yaml:"algo"`
		Secret     string `json:"secret" yaml:"secret"`
		TTL        uint   `json:"ttl" yaml:"ttl"`
		RefreshTTL uint   `json:"refresh_ttl" yaml:"refresh_ttl"`
	} `json:"jwt" yaml:"jwt"`
}

// NewConfig read and parse config file
func NewConfig(path IPath) (*Config, *godi.BeanOptions, error) {
	file, err := ioutil.ReadFile(string(path))
	if err != nil {
		return nil, nil, err
	}

	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		return nil, nil, err
	}

	options := &godi.BeanOptions{
		Type: godi.Singleton,
	}

	return &config, options, nil
}

func (config *Config) Get() *Config {
	return config
}
