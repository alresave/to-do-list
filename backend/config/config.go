package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DbUri     string
	RabbitUri string
	RedisUri  string
	RedisPwd  string
	RedisDb   int
	ServerUrl string
	JwtKey    string
}

func InitConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	config := Config{
		DbUri:     viper.GetString("DATABASE_URI"),
		RabbitUri: viper.GetString("RABBIT_URI"),
		RedisUri:  viper.GetString("127.0.0.1:6379"),
		RedisPwd:  viper.GetString("REDIS_PASSWORD"),
		RedisDb:   viper.GetInt("REDIS_DBNUMBER"),
		ServerUrl: viper.GetString("SERVER_URL"),
		JwtKey:    viper.GetString("JWT_KEY"),
	}
	return &config, nil
}
