package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Redis Redis `json:"redis"`
	MySQL MySQL `json:"database"`
}

type MySQL struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        string `json:"port"`
	Database    string `json:"database"`
	MaxIdleConn int    `json:"maxIdleConn"`
	MaxOpenConn int    `json:"maxOpenConn"`
	MaxIdleTime int    `json:"maxIdleTime"`
	MaxLifetime int    `json:"maxLifetime"`
}

type Redis struct {
	Address    string `json:"address" yaml:"address"`
	Password   string `json:"password" yaml:"password"`
	Timeout    int    `json:"timeout" yaml:"timeout"`
	PoolSize   int    `json:"poolSize" yaml:"poolSize"`
	MaxConnAge int    `json:"maxConnAge" yaml:"maxConnAge"`
	TTL        int    `json:"ttl" yaml:"ttl"`
}

func New(pathConfig string) *Config {
	env := "local"
	if v, ok := os.LookupEnv("ENV"); ok {
		env = v
	}

	viper.AddConfigPath(pathConfig)
	viper.SetConfigName("config." + env)

	fmt.Println("make config with config." + env)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
