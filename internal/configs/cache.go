package configs

import (
	"fmt"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func NewCache(envs Environments) *redis.Client {
	redisDB, err := strconv.Atoi(envs.REDIS_DB.String())
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", envs.REDIS_HOST.String(), envs.REDIS_PORT.String()),
		Password: envs.REDIS_PASSWORD.String(),
		DB:       redisDB,
	})
	defer redisClient.Close()

	return redisClient
}
