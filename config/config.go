package config


import (
	"fmt"
	"github.com/spf13/viper"
)

var Config *viper.Viper

func InitConfig(mode string) {
	Config = viper.GetViper()
	Config.AddConfigPath("config")
	Config.SetConfigType("toml")
	Config.SetConfigName("application")
	Config.ReadInConfig()
	Config.SetConfigName("application-" + mode)
	Config.MergeInConfig()
	fmt.Println(Config.GetString("database.mysql.host"))
}
