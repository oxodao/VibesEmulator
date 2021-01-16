package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ListeningAddress string
	WebrootURL       string
}

func Load() (*Config, error) {
	cfg := Config{}

	ctt, err := ioutil.ReadFile("config.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(ctt, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}