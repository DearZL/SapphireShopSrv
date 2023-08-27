package global

import (
	"SapphireShop/SapphireShop_srv/email_srv/common/config"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	Config *viper.Viper
	Logger *zap.Logger
	Redis  *redis.Client
)

func InitConfig() {
	Config = config.InitConfig()
}

func InitLogConfig() {
	if Config.Get("app.mode") == "debug" {
		Logger = config.ZapConfig(true)
	}
}

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.host" + ":" + viper.GetString("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       0,
	})
	_, err := Redis.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("redis初始化出错%v", err))
		return
	}
}
