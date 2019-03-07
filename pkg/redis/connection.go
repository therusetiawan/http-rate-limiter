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

// get the length of the list
func LLen(key string) (value int64, err error) {
	value, err = client.LLen(key).Result()
	return
}

// check key is exists or not
func Exists(key string) (value int64) {
	value, _ = client.Exists(key).Result()
	return
}

// insert all the specified values at the tail of the list stored at key.
// if key does not exist, it is created as empty list before performing the push operation
func RPush(key string, value interface{}) (err error) {
	err = client.RPush(key, value).Err()
	return
}

// inserts value at the tail of the list stored at key,
// only if key already exists and holds a list
func RPushX(key string, value interface{}) (err error) {
	err = client.RPushX(key, value).Err()
	return
}

// set expired time
func ExpireAt(key string, exp time.Time) (err error) {
	err = client.ExpireAt(key, exp).Err()
	return
}

// do RpushX and ExpireAt in transaction mode
func RPushAndExpireAt(key string, value interface{}, exp time.Time) error {
	pipe := client.TxPipeline()
	pipe.RPush(key, value)
	pipe.ExpireAt(key, exp)

	_, err := pipe.Exec()
	return err
}
