package main

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

// DBConfig db config
type DBConfig struct {
	Host     string `toml:"host"`
	Port     uint16 `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Name     string `toml:"name"`
}

// Config migration config
type Config struct {
	DB *DBConfig `toml:"database"`
}

// NewConfig creates new config
func NewConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open config file")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read config file")
	}
	var config Config
	if err := toml.Unmarshal(buf, &config); err != nil {
		return nil, errors.Wrap(err, "failed to create Config from file")
	}
	return &config, nil
}
