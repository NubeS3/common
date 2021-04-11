//+build windows linux

package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ArangoHost       string `mapstructure:"arango_host"`
	ArangoUser       string `mapstructure:"arango_user"`
	ArangoPassword   string `mapstructure:"arango_password"`
	SeaweedMasterUrl string `mapstructure:"seaweed_master_host"`
	NatsUrl          string `mapstructure:"nats_url"`
	StanUrl          string `mapstructure:"stan_url"`
	JwtSecret        string `mapstructure:"jwt_secret"`
}

var Conf Config

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	viper.ReadInConfig()

	err := viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
