package configs

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Port           int `yaml:"port"`
	Authentication struct {
		Hosts []string `yaml:"hosts"`
	} `yaml:"authentication"`
	Post struct {
		Hosts []string `yaml:"hosts"`
	}	`yaml:"posts"`
}

type appConfig struct {
	AppConfig *AppConfig `yaml:"app_config"`
}

func getConfig(path string) (*appConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open App config (path=%s) error: %s", path, err)
	}
	defer f.Close()

	bs, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read App config (path=%s) error: %s", path, err)
	}

	conf := &appConfig{}
	if err := yaml.Unmarshal(bs, conf); err != nil {
		return nil, fmt.Errorf("unmarshal App config (path=%s) error: %s", path, err)
	}
	return &appConfig{}, nil
}

func GetAppConfig(path string) (*AppConfig, error) {
	conf, err := getConfig(path)
	if err != nil {
		return nil, err
	}
	return conf.AppConfig, nil
}


