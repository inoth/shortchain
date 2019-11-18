package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

var (
	once sync.Once
	conf *Config
)

type Config struct {
	DB struct {
		Redis string `yaml:"Redis"`
		Mongo []struct {
			Name       string `yaml:"Name"`
			ConnectStr string `yaml:"ConnectStr"`
		} `yaml:"Mongo"`
	} `yaml:"DB"`
	ServerPort string `yaml:"ServerPort"`
	Domain     string `yaml:"Domain"`
	Token      string `yaml:"Token"`
}

func Instance() *Config {
	once.Do(func() {
		conf = &Config{}
	})
	return conf
}

func (c *Config) Init() error {
	yamlFile, err := ioutil.ReadFile("conf.yml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}
	return nil
}
