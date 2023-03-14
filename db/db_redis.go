package db

import (
	"fmt"
	"github.com/go-redis/redis"
)

var (
	redisdb *redis.Client
)

func InitRedis() (err error) {
	redisdb = redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		})
	_, err = redisdb.Ping().Result()
	return
}

func Save(key string, value string) error {
	res, err := redisdb.Set(key, value, 0).Result()
	if err != nil {
		return fmt.Errorf("save error, %s, %s", err.Error(), res)
	}
	return nil
}

func Get(key string) (string, error) {
	return redisdb.Get(key).Result()
}

func Del(key string) {
	redisdb.Del(key).Result()
}
