package main

import (
	"io/ioutil"

	"gitlab.com/kompu/twitter-saver/core"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Db             core.DbConfig
	Twitter        core.TwitterConfig
	Json           core.JsonConfig
	AutoDeleteDays *int `yaml:"autoDeleteDays"`
}

func readConfig(path *string) (*Config, error) {
	data, err := ioutil.ReadFile(*path)
	if err != nil {
		return nil, err
	}

	config := Config{}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
