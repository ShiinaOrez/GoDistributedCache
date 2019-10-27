package cache

import (
	"errors"
	"sync"
)

type CacheService struct {
	Cache  map[string][]byte
	Locker *sync.RWMutex
}

func NewCacheService() *CacheService {
	return &CacheService{
		Cache:  make(map[string][]byte),
		Locker: new(sync.RWMutex),
	}
}

func (cache *CacheService) Get(key string) ([]byte, error) {
	value, existed := cache.Cache[key]
	if existed {
		return value, nil
	}
	return []byte{}, errors.New("Key: " + key + " Not Found!")
}

func (cache *CacheService) Set(key string, value []byte) error {
	cache.Locker.Lock()
	defer cache.Locker.Unlock()

	cache.Cache[key] = value
	return nil
}

func (cache *CacheService) Del(key string) error {
	if _, existed := cache.Cache[key]; !existed {
		return errors.New("Key: " + key + " Not Found!")
	}
	cache.Locker.Lock()
	defer cache.Locker.Unlock()

	delete(cache.Cache, key)
	return nil
}
