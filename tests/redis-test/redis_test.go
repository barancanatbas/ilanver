package redistest

import (
	"ilanver/internal/cache"
	"ilanver/internal/config"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-playground/assert/v2"
	"github.com/gomodule/redigo/redis"
)

func setup() *redis.Pool {
	redisServer := mockRedis()
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", redisServer.Addr())
			if err != nil {
				return nil, err
			}
			return c, err
		},
	}
}

func mockRedis() *miniredis.Miniredis {
	s, err := miniredis.Run()

	if err != nil {
		panic(err)
	}

	return s
}

func TestSetRedis(t *testing.T) {
	config.Pool = setup()
	defer config.Pool.Close()

	key := "test"

	cache.SetFromCache(key, "test", 6)

	exists := cache.Exists(key)

	assert.Equal(t, exists, true)
}

func TestRedisGet(t *testing.T) {
	config.Pool = setup()
	defer config.Pool.Close()

	key := "test"
	var value string

	cache.SetFromCache(key, "test", 6)

	err := cache.GetFromCache(key, &value)

	assert.Equal(t, err, true)
}

func TestRedisHmap(t *testing.T) {
	config.Pool = setup()
	defer config.Pool.Close()

	key := "test"

	data := map[string]string{
		"test1": "test1",
		"test2": "test2",
	}

	cache.SetHashCache(key, data)

	exists := cache.Exists(key)

	assert.Equal(t, exists, true)

	goal := cache.GetHashCache(key)

	assert.Equal(t, goal["test1"], data["test1"])
	assert.Equal(t, goal["test2"], data["test2"])
}

func TestDelete(t *testing.T) {
	config.Pool = setup()
	defer config.Pool.Close()

	key := "test"

	cache.SetFromCache(key, "test", 6)

	exists := cache.Exists(key)

	assert.Equal(t, exists, true)

	cache.DelFromCache(key)

	exists = cache.Exists(key)

	assert.Equal(t, exists, false)
}
