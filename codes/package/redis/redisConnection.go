package redis

import (
	"log"

	"github.com/go-redis/redis"
)

func RedisConnection() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "host.docker.internal:6379",
		Password: "redispassword",
		DB:       0,
	})

	err := client.Ping().Err()
	if err != nil {
		log.Printf("Errors %s while pinging", err)
		return nil, err
	}

	return client, nil
}
