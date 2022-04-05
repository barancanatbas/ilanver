package cache

import (
	"encoding/json"
	"fmt"
	"ilanver/internal/config"

	"github.com/gomodule/redigo/redis"
)

func SetHashCache(key string, hash map[string]string) {
	var client = config.Pool.Get()
	defer client.Close()

	for k, v := range hash {
		client.Do("HSET", key, k, v)
	}
}

func SetFromCache(key string, list interface{}, ex uint64) {
	var client = config.Pool.Get()
	defer client.Close()
	jsondata, err := json.Marshal(list)
	if err != nil {
		fmt.Println("json converter")
	}
	val, err := redis.Int64(client.Do("set", key, string(jsondata), "ex", ex))
	if err != nil && val <= 0 {
		fmt.Println("set redis")
	}
}

func SetFromCacheNoEx(key string, list interface{}) {
	var client = config.Pool.Get()
	defer client.Close()
	jsondata, err := json.Marshal(list)
	if err != nil {
		fmt.Println("json converter")
	}
	val, err := redis.Int64(client.Do("set", key, string(jsondata)))
	if err != nil && val <= 0 {
		fmt.Println("set redis")
	}
}
