package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig()  {
	viper.SetConfigName("application") // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".") // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	_ = viper.BindEnv("todo.select.port", "TODO_SELECT_PORT")
	_ = viper.BindEnv("todo.select.uri", "TODO_SELECT_URI")
	_ = viper.BindEnv("todo.write.port", "TODO_WRITE_PORT")
	_ = viper.BindEnv("todo.write.uri", "TODO_WRITE_URI")
	_ = viper.BindEnv("web.port", "WEB_PORT")
	_ = viper.BindEnv("check.token.url", "CHECK_TOKEN_URL")
}
