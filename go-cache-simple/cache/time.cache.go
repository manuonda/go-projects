package cache

import (
	"sync"
	"time"
)

// struct save data
type cacheEntry struct {
	value      interface{}
	expiration time.Time
}

type TimeCache struct {
	Data  map[string]cacheEntry
	mutex sync.Mutex
}

func NewCache() *TimeCache {
	return &TimeCache{
		Data: make(map[string]cacheEntry),
	}
}

func (tc *TimeCache) Set(key string, value interface{}, expiration time.Duration) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	tc.Data[key] = cacheEntry{
		value:      value,
		expiration: time.Now().Add(expiration),
	}

}

func (tc *TimeCache) Get(key string) (interface{}, bool) {
	tc.mutex.Lock()
	defer tc.mutex.Unlock()

	entry, found := tc.Data[key]

	if !found {
		return nil, false
	}

	if entry.expiration.Before(time.Now()) {
		delete(tc.Data, key)
		return nil, false
	}

	return entry.value, true
}
