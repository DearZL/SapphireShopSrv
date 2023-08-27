package global

import (
	"SapphireShop/SapphireShop_srv/user_srv/common/config"
	"SapphireShop/SapphireShop_srv/user_srv/model"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	DB     *gorm.DB
	Redis  *redis.Client
	dbErr  error
	Config *viper.Viper
	Logger *zap.Logger
)

func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8&parseTime=true&loc=Local",
		Config.GetString("userDB.username"),
		Config.GetString("userDB.password"),
		Config.GetString("userDB.host"),
		Config.GetString("userDB.port"),
		Config.GetString("userDB.name"),
	)
	DB, dbErr = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if dbErr != nil {
		panic(dbErr.Error())
		return
	}
	dbErr = DB.AutoMigrate(
		&model.User{},
	)
	if dbErr != nil {
		panic(dbErr.Error())
		return
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

func InitConfig() {
	Config = config.InitConfig()
}

func InitLogConfig() {
	if Config.Get("app.mode") == "debug" {
		Logger = config.ZapConfig(true)
	}

}
