package config

import (
	"fmt"
	"github.com/spf13/viper"
)

// 读取配置
func InitConfig()  {
	viper.SetConfigName("application")  //配置文件名称
	viper.SetConfigType("yaml")// 读取yaml配置文件
	viper.AddConfigPath(".")      // 设置配置文件和可执行二进制文件在用一个目录
	err := viper.ReadInConfig() // 搜索路径，并读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}