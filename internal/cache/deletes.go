package cache

import (
	"fmt"
	"ilanver/internal/config"

	"github.com/gomodule/redigo/redis"
)

func DeleteKeysFromPattern(pattern string) {
	var client = config.Pool.Get()
	defer client.Close()

	val, err := redis.Strings(client.Do("keys", pattern))
	if err == nil {
		for _, item := range val {
			client.Do("del", item)
		}
	}
}

func DelFromCache(key string) {
	var client = config.Pool.Get()
	client.Do("del", key)
	defer client.Close()
}

// eğer rediste verilen key değerine göre bir değer varsa true döner, yoksa false.
func Exists(key string) bool {
	var client = config.Pool.Get()
	defer client.Close()
	val, err := redis.Int64(client.Do("exists", key))
	if err != nil {
		return false
	}

	if val <= 0 {
		return false
	}

	fmt.Println(val)

	return true
}
