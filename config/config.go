package config

import (
	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
)

type Config struct {
	HighlightColor string
	NoHighlight    bool
	WordsPerMinute int
}

func defaultConfig() Config {
	return Config{HighlightColor: "#98FF98", NoHighlight: false, WordsPerMinute: 300}
}

func parseConfig(rawConfig string) Config {
	var config Config
	_, err := toml.Decode(rawConfig, &config)
	if err != nil {
		return defaultConfig()
	}

	return config
}

func ReadConfig() Config {
	configString, err := xdg.ConfigFile("sread/config.toml")
	if err != nil {
		return defaultConfig()
	}

	return parseConfig(configString)
}
