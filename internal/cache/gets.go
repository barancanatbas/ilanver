package cache

import (
	"encoding/json"
	"fmt"
	"ilanver/internal/config"

	"github.com/gomodule/redigo/redis"
)

func GetFromCache(key string, list interface{}) bool { // testler bittikten sonra buradaki fmt'ler kaldırılacak !
	// cache de var mı ?
	var client = config.Pool.Get()
	defer client.Close()
	value, err := redis.String(client.Do("get", key))
	if err != nil {
		return false
	}
	err = json.Unmarshal([]byte(value), &list)
	if err != nil {
		fmt.Println("->json hatası var")
		return false
	}

	return true
}

func GetHashCache(key string) (map[string]string, bool) {
	fmt.Println("girdi")
	var client = config.Pool.Get()
	defer client.Close()
	value, err := redis.StringMap(client.Do("HGETALL", key))

	if err != nil {
		return make(map[string]string, 0), false
	}
	fmt.Println("value : ", value)
	return value, true
}

func GetFromCacheString(key string) (string, bool) {
	// cache de var mı ?
	var client = config.Pool.Get()
	defer client.Close()
	value, err := redis.String(client.Do("get", key))
	if err != nil {
		return "", false
	}

	return string(value), true
}

func GetFromCacheInt(key string) (int, bool) {
	var client = config.Pool.Get()
	defer client.Close()
	value, err := redis.Int(client.Do("get", key))
	if err != nil {
		return 0, false
	}
	return value, true
}
