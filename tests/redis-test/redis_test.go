package redistest

import (
	"ilanver/internal/cache"
	"ilanver/internal/config"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestRedis(t *testing.T) {
	config.Pool = config.NewPool()

	key := "test"
	var value string

	cache.SetFromCache(key, "test", 6)

	err := cache.GetFromCache(key, &value)

	assert.Equal(t, err, true)

	time.Sleep(time.Second * 6)

	err = cache.GetFromCache(key, &value)

	assert.Equal(t, err, false)
}

func TestRedisHmap(t *testing.T) {
	config.Pool = config.NewPool()

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
	config.Pool = config.NewPool()

	key := "test"

	cache.SetFromCache(key, "test", 6)

	exists := cache.Exists(key)

	assert.Equal(t, exists, true)

	cache.DelFromCache(key)

	exists = cache.Exists(key)

	assert.Equal(t, exists, false)
}
