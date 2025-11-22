package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port int `mapstructure:"port"`
}

type MySQLConfig struct {
	Host      string `mapstructure:"host"`
	Port      int    `mapstructure:"port"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	Database  string `mapstructure:"database"`
	Charset   string `mapstructure:"charset"`
	ParseTime bool   `mapstructure:"parseTime"`
	Loc       string `mapstructure:"loc"`
}

type RedisConfig struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type RabbitConfig struct {
	URL string `mapstructure:"url"`
}

type LoggingConfig struct {
	Level string `mapstructure:"level"`
}

type Config struct {
	Server   ServerConfig `mapstructure:"server"`
	MySQL    MySQLConfig  `mapstructure:"mysql"`
	Redis    RedisConfig  `mapstructure:"redis"`
	RabbitMQ RabbitConfig `mapstructure:"rabbitmq"`
	Logging  LoggingConfig `mapstructure:"logging"`
}

func loadConfig() (*Config, error) {
	v := viper.New()

	v.SetConfigName("configs")
	v.SetConfigType("yaml")
	v.AddConfigPath("configs") // folder

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &cfg, nil
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Config error: %v", err)
	}

	fmt.Println("Server will run on port:", cfg.Server.Port)
	fmt.Println("MySQL host:", cfg.MySQL.Host)
}
