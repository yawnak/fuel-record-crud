package config

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Config struct {
	DBConn DBConnConfig
	Server ServerConfig
}

type DBConnConfig struct {
	Host   string `config:"POSTGRES_HOST"`
	Port   string `config:"POSTGRES_PORT"`
	User   string `config:"POSTGRES_USER"`
	DbName string `config:"POSTGRES_DB"`
}

type ServerConfig struct {
	Port string `config:"PORT"`
}

func BindConfig(conf any, paths ...string) error {
	defer config.ClearAll()
	if reflect.ValueOf(conf).Type().Kind() != reflect.Ptr {
		return errors.New("conf is not a pointer")
	}
	config.WithOptions(config.ParseEnv, func(opt *config.Options) {
		opt.DecoderConfig.TagName = "config"
	})
	// add driver for support yaml content
	config.AddDriver(yaml.Driver)

	err := config.LoadFiles(paths...)
	if err != nil {
		return fmt.Errorf("error loading config files: %w", err)
	}

	err = config.Decode(conf)
	return err
}

func MustGetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("env variable %s not found", key))
	}
	return val
}

func NewDBConnConfig(filepath string) (*DBConnConfig, error) {
	conf := &DBConnConfig{}
	err := BindConfig(conf, filepath)
	return conf, err
}

func NewServerConfig(filepath string) (*ServerConfig, error) {
	conf := &ServerConfig{}
	err := BindConfig(conf, filepath)
	return conf, err
}
