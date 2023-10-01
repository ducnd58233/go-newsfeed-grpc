package configs

import (
	"fmt"
	"io"
	"os"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gopkg.in/yaml.v2"
)

type AuthConfig struct {
	Port  int           `yaml:"port"`
	MySQL mysql.Config  `yaml:"mysql"`
	Redis redis.Options `yaml:"redis"`
}

type authConfig struct {
	AuthConfig *AuthConfig `yaml:"auth_config"`
}

func getConfig(path string) (*authConfig, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open App config (path=%s) error: %s", path, err)
	}
	defer f.Close()

	bs, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("read App config (path=%s) error: %s", path, err)
	}

	conf := &authConfig{}
	if err := yaml.Unmarshal(bs, conf); err != nil {
		return nil, fmt.Errorf("unmarshal App config (path=%s) error: %s", path, err)
	}
	return &authConfig{}, nil
}

func GetAuthConfig(path string) (*AuthConfig, error) {
	conf, err := getConfig(path)
	if err != nil {
		return nil, err
	}
	return conf.AuthConfig, nil
}
