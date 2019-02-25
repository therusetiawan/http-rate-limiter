package redis

import (
	"time"

	"github.com/go-redis/redis"
)

var client *redis.Client

func (c *Config) NewConnection() error {
	client = redis.NewClient(&redis.Options{
		Addr:     c.Addr,
		Password: c.Password,
		DB:       c.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}

func GetConnection() *redis.Client {
	return client
}

func LLen(key string) (value int64, err error) {
	value, err = client.LLen(key).Result()
	return
}

func Exists(key string) (value int64) {
	value, _ = client.Exists(key).Result()
	return
}

func RPush(key string, value interface{}) (err error) {
	err = client.RPush(key, value).Err()
	return
}

func RPushX(key string, value interface{}) (err error) {
	err = client.RPushX(key, value).Err()
	return
}

func ExpireAt(key string, exp time.Time) (err error) {
	err = client.ExpireAt(key, exp).Err()
	return
}
