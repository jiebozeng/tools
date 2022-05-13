package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

var config *viper.Viper

func InitConfig()  {
	config = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./")
	config.AddConfigPath("./bin")
	err := config.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("server config reading exception %s",err))
	}
}

func GetPort() string {
	port := config.GetString("server.port")
	return port
}

func GetIsDebug() bool {
	return config.GetBool("isdebug")
}