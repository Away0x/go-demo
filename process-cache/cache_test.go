package cache_test

import (
	"log"
	"sync"
	"testing"

	"cache"
	"cache/algorithm/lru"

	"github.com/matryer/is"
)

func TestCacheGet(t *testing.T) {
	db := map[string]string{
		"key1": "val1",
		"key2": "val2",
		"key3": "val3",
		"key4": "val4",
	}
	getter := cache.GetFunc(func(key string) interface{} {
		log.Println("[From DB] find key", key)

		if val, ok := db[key]; ok {
			return val
		}
		return nil
	})
	c := cache.NewCache(getter, lru.New(0, nil))

	is := is.New(t)

	var wg sync.WaitGroup

	for k, v := range db {
		wg.Add(1)
		go func(k, v string) {
			defer wg.Done()
			is.Equal(c.Get(k), v)

			is.Equal(c.Get(k), v)
		}(k, v)
	}
	wg.Wait()

	is.Equal(c.Get("unknown"), nil)
	is.Equal(c.Get("unknown"), nil)

	is.Equal(c.Stat().NGet, 10)
	is.Equal(c.Stat().NHit, 4)
}
