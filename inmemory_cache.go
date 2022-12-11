package main

import (
	"encoding/json"
	"time"

	"github.com/allegro/bigcache/v3"
)

type cacheStorage struct {
	memcache *bigcache.BigCache
}

const (
	CACHE_TTL  = 5 * time.Minute
	CACHE_SIZE = 128
)

var CacheStorage *cacheStorage

func init() {
	if CacheStorage != nil {
		err := CacheStorage.memcache.Close()
		if err != nil {
			panic(err)
		}
		CacheStorage = nil
	}

	config := bigcache.DefaultConfig(CACHE_TTL)
	config.HardMaxCacheSize = CACHE_SIZE
	cache, err := bigcache.NewBigCache(config)
	if err != nil {
		panic(err)
	}

	CacheStorage = &cacheStorage{
		memcache: cache,
	}
}

func (c *cacheStorage) Get(key string) ([]byte, error) {
	return c.memcache.Get(key)
}

func (c *cacheStorage) Save(key string, v interface{}) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	return c.memcache.Set(key, b)
}

func (c *cacheStorage) Invalidate(key string, bin []byte) bool {
	return c.Put(key, bin)
}

func (c *cacheStorage) Put(key string, bin []byte) bool {
	err := c.memcache.Set(key, bin)
	return err == nil
}

func (c *cacheStorage) Remove(key string) bool {
	err := c.memcache.Delete(key)
	return err == nil
}
