package config

import (
	"context"
	"crypto/tls"
	"log"
	"os"
	"session-service-v2/app/logger"
	"strconv"

	redis "github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var logg *logger.FpayLogger = logger.GetLogger()

func SetupDBConnection() *redis.Client {

	db, _ := strconv.Atoi(os.Getenv("NODE_REDIS_DB_NUMBER"))
	client := redis.NewClient(&redis.Options{
		Addr:      os.Getenv("REDIS_HOSTNAME") + ":" + os.Getenv("REDIS_PORT"),
		Password:  os.Getenv("REDIS_PRIMARY_ACCESS_KEY"),
		DB:        db,
		TLSConfig: &tls.Config{},
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		logg.Error("error configuring redis...")
		log.Fatal(err)
	}

	logg.Info("Redis configuration set successfully")

	return client
}

func CloseDBConnection(db *redis.Client) {
	errCloseDB := db.Close()

	if errCloseDB != nil {
		panic("Fail to close connection")
	}
}
