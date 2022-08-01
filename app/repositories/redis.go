package repositories

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

var redisClient *redis.Client

func init() {
	_err := godotenv.Load()
	if _err != nil {
		fmt.Println("Error loading .env file" + _err.Error())
	}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pong)

	redisClient = client
}

func SetClient(client *redis.Client){
	redisClient = client
}