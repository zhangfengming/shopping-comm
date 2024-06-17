package common

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetConsulConfig(url, fileKey string) (*viper.Viper, error) {
	viper := viper.New()
	viper.AddRemoteProvider("consul", url, fileKey)

	viper.SetConfigType("json")
	err := viper.ReadRemoteConfig()
	if err != nil {
		fmt.Println(err)
	}
	return viper, err
}

// type MysqlConfig struct {
// 	Host     string `json:"host"`
// 	Port     string `json:"port"`
// 	User     string `json:"user"`
// 	Pwd      string `json:"pwd"`
// 	Database string `json:"database"`
// }

// 获取mysql 配置
func GetMysqlFromCousul(viper *viper.Viper, path ...string) (db *gorm.DB, err error) {
	str := viper.GetString("user") + ":" + viper.GetString("pwd") + "@tcp(" + viper.GetString("host") + ":" + viper.GetString("port") + ")/" + viper.GetString("database") + "?charset=utf8mb4&parseTime=True&loc=Local"
	// DB, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/user-center?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	db, errr := gorm.Open(mysql.Open(str), &gorm.Config{Logger: newLogger})

	if errr != nil {
		log.Println(err)
	}

	return db, nil
}

/**
{
	"user" : "root",
	"pwd" : "",
	"DB" : 0,
	"address" : "127.0.0.1",
	"port": 6379,
	"poolSize" : 30,
	"minIdeConn": 30
}
*/
// 获取redis 配置
func GetRedisFromCousul(viper *viper.Viper, path ...string) (red *redis.Client, err error) {
	redis := redis.NewClient(&redis.Options{
		Addr:         viper.GetString("address") + ":" + viper.GetString("port"),
		Password:     viper.GetString("pwd"),
		DB:           viper.GetInt("DB"),
		PoolSize:     viper.GetInt("poolSize"),
		MinIdleConns: viper.GetInt("minIdeConn"),
	})
	//cluster
	// clusterClients := redis.NewClusterClient(&redis.clusteroptions{
	// 	Addrs: []string{"192.168.100,131:6380", "192.168.100.131:6381", "192.168.100.131:6382"},
	// })
	// fmt.Println(clusterClients)
	return redis, nil
}

// 设置用户登录信息
func SetUserToken(red *redis.Client, key string, val []byte, timeTTL time.Duration) {
	red.Set(key, val, timeTTL)
}

// 获取用户登录信息
func GetUserToken(red *redis.Client, key string) string {
	res, err := red.Get(key).Result()
	if err != nil {
		fmt.Println("get redis token error", err)
	}
	return res
}
