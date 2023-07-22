package cache

import "sync"

type Cache struct {
	Data  map[string]interface{}
	mutex sync.Mutex ///sincronize el access to Cache data
}

/**
 * Funcion set SimpleCache to value
**/
func (c *Cache) Set(key string, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.Data[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	value, ok := c.Data[key]
	return value, ok
}

func (c *Cache) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	delete(c.Data, key)
}
