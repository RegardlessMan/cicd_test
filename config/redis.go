/**
 * @Author QG
 * @Date  2024/12/1 15:30
 * @description
**/

package config

import (
	"cicd_test/global"
	"github.com/go-redis/redis"
	"log"
)

func initRedis() {

	newClient := redis.NewClient(&redis.Options{
		Addr:     AppConfig.Redis.Addr,
		Password: AppConfig.Redis.Password,
		DB:       AppConfig.Redis.DB,
	})

	_, err := newClient.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis, got error: %v", err)
	}
	global.RedisDb = newClient
}
