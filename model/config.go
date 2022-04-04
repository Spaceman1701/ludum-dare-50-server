package model

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var config *Config

type ShrineConfig struct {
	Radius               int `yaml:"radius"`
	Buffer               int `yaml:"buffer"`
	BaseSacrificeContrib int `yaml:"baseSacrificeContrib"`
	BaseDeathContrib     int `yaml:"baseDeathContrib"`
	UpgradeThreshold     int `yaml:"upgradeThreshold"`
	DestroyThreshold     int `yaml:"destroyThreshold"`
	DeclineRate          int `yaml:"declineRate"`
	SpawnPower           int `yaml:"spawnPower"`
	UsageCost            int `yaml:"usageCost"`
}

type Config struct {
	Shrine ShrineConfig `yaml:"shrine"`
}

func loadConfig(path string) Config {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var c Config
	err = yaml.UnmarshalStrict(file, &c)
	if err != nil {
		panic(err)
	}
	return c
}

func GetConfig() Config {
	if config == nil {
		fmt.Printf("config is nil, loading\n")
		c := loadConfig("config.yaml")
		config = &c
	}
	return *config
}
