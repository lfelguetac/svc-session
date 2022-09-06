package config

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"os"

	redis "github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
)

var ctx = context.Background()

func SetupDBConnection() *redis.Client {
	_err := godotenv.Load()
	if _err != nil {
		fmt.Println("Error loading .env file" + _err.Error())
	}

	client := redis.NewClient(&redis.Options{
		Addr:      os.Getenv("REDIS_URI"),
		Password:  os.Getenv("REDIS_PWD"),
		DB:        0,
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
	log.Println("errCloseDB : ", db)
	errCloseDB := db.Close()
	log.Println("errCloseDB : ", errCloseDB)
	if errCloseDB != nil {
		panic("Fail to close connection")
	}
}
