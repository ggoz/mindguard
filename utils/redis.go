package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
)

type RedisStore struct {
	client *redis.Client
}

var RediStore *RedisStore

func InitRedisStore(redisConfig RedisConfig) *RedisStore {
	fmt.Println("rd", redisConfig)
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr + ":" + redisConfig.Port,
		Password: redisConfig.Password,
		DB:       redisConfig.Db,
	})

	RediStore = &RedisStore{client: client}

	return RediStore
}

// 存入redis
func (rs *RedisStore) Set(id string, value string) error {
	// 五分钟过期
	err := RediStore.client.Set(id, value, time.Minute*60*24).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// 从redis获取值
func (rs *RedisStore) Get(id string) string {
	val, err := RediStore.client.Get(id).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return val
}

// 从redis删除数据
func (rs *RedisStore) Del(id string) error {
	err := RediStore.client.Del(id).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
