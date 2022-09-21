package config

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strconv"

	redis "github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var ctx = context.Background()

func SetupDBConnection() *redis.Client {
	_err := godotenv.Load()
	if _err != nil {
		fmt.Println("Error loading .env file" + _err.Error())
	}

	// ttl: Number(process.env.NODE_REDIS_TTL),
	// tls: (process.env.REDIS_USE_TLS === 'true')

	db, _ := strconv.Atoi(os.Getenv("NODE_REDIS_DB_NUMBER"))
	client := redis.NewClient(&redis.Options{
		Addr:      os.Getenv("REDIS_HOSTNAME") + ":" + os.Getenv("REDIS_PORT"),
		Password:  os.Getenv("REDIS_PRIMARY_ACCESS_KEY"),
		DB:        db,
		TLSConfig: &tls.Config{},
	})

	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pong)

	return client
}

func CloseDBConnection(db *redis.Client) {
	errCloseDB := db.Close()

	if errCloseDB != nil {
		panic("Fail to close connection")
	}
}
