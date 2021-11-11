package services

import (
	"os"
	"strings"

	"github.com/evertonaleixo/stone-test/models"
	"github.com/go-redis/redis/v7"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Can be migrate to a postgres like
	database, err := gorm.Open(sqlite.Open("simple.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&models.Account{})
	database.AutoMigrate(&models.Transfer{})

	DB = database
}

var REDISCLIENT *redis.Client

func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	passwd := os.Getenv("REDIS_PWD")
	if len(dsn) == 0 {
		dsn = "spinyfin.redistogo.com:9986"
	}

	redisOpts := &redis.Options{
		Addr: dsn, //redis port
	}
	if strings.Compare(passwd, "") != 0 {
		redisOpts.Password = passwd
	}

	REDISCLIENT = redis.NewClient(redisOpts)
	_, err := REDISCLIENT.Ping().Result()
	if err != nil {
		panic(err)
	}
}
