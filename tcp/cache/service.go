package cache

import (
	"sync"
	"errors"
)

type Service struct {
	Map    map[string][]byte
	Locker *sync.RWMutex
}

func NewService() *Service {
	return &Service {
		Map: make(map[string][]byte),
		Locker: new(sync.RWMutex),
	}
}

func (service *Service) Get(key string) ([]byte, error) {
	val, existed := service.Map[key]
	if !existed {
		return nil, errors.New("Key: "+key+", not found!")
	}
    return val, nil
}

func (service *Service) Set(key string, val []byte) error {
	service.Locker.Lock()
	defer service.Locker.Unlock()

	service.Map[key] = val
	return nil
}

func (service *Service) Del(key string) error {
	_, existed := service.Map[key]
	if !existed {
		return errors.New("Key: "+key+", not found!")
	}
	service.Locker.Lock()
	defer service.Locker.Unlock()

	delete(service.Map, key)
	return nil
} 
