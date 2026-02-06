package config

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/adrg/xdg"
)

type Config struct {
	HighlightColor string
	NoHighlight    bool
	WordsPerMinute int
}

var configPath string = "sread/config.toml"

func defaultConfig() Config {
	return Config{HighlightColor: "#98FF98", NoHighlight: false, WordsPerMinute: 300}
}

func parseConfig(rawConfig string) Config {
	var config Config
	_, err := toml.Decode(rawConfig, &config)
	if err != nil {
		config := defaultConfig()
		if err = writeConfig(config); err != nil {
			log.Fatal(err)
		}
	}

	return config
}

func writeConfig(config Config) error {
	configFilePath, err := xdg.ConfigFile(configPath)
	if err != nil {
		return err
	}

	file, err := os.Create(configFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	if err = encoder.Encode(config); err != nil {
		return err
	}

	return nil
}

func ReadConfig() (Config, error) {
	configFilePath, err := xdg.ConfigFile(configPath)
	if err != nil {
		return Config{}, err
	}

	bytes, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, err
	}

	return parseConfig(string(bytes)), nil
}
