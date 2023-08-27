package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() *viper.Viper {
	config := viper.New()
	config.SetConfigName("config") // 设置配置文件名（不包含扩展名）
	config.SetConfigType("yaml")   // 设置配置文件类型，可选
	config.AddConfigPath("../../") // 添加配置文件搜索路径
	err := config.ReadInConfig()   // 加载配置文件
	if err != nil {
		panic(fmt.Errorf("无法加载配置文件: %s", err))
	}
	return config
}
