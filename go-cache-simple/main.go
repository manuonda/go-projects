package main

import (
	"fmt"
	"time"

	"github.com/manuonda/go-projects/go-cache-simple/cache"
)

func main() {

	fmt.Println("Simple Cache Service")
	simpleCache := cache.Cache{
		Data: make(map[string]interface{}),
	}

	simpleCache.Set("key1", "value1")
	simpleCache.Set("key2", "value2")
	value, ok := simpleCache.Get("key2")
	if ok {
		fmt.Println("Value :", value)
	} else {
		fmt.Println("Key not Foundt")
	}
	simpleCache.Delete("key1")

	fmt.Println("-- Cache Time ----")
	timeCache := cache.NewCache()
	timeCache.Set("key3", "value3", 5*time.Second)
	if value, found := timeCache.Get("key3"); found {
		fmt.Println("Value3 get : ", value)
	} else {
		fmt.Println("Not found value of key3")
	}
	time.Sleep(6 * time.Second)

	if value, found := timeCache.Get("key3"); found {
		fmt.Println("Value3 get : ", value)
	} else {
		fmt.Println("Not found value of key3")
	}

}
