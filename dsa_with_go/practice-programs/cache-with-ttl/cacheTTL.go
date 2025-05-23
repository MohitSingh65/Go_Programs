package main

import (
	"fmt"
	"sync"
	"time"
)

type CacheItem struct {
	value  interface{}
	expiry time.Time
}

type Cache struct {
	data map[string]CacheItem
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]CacheItem),
	}
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = CacheItem{
		value:  value,
		expiry: time.Now().Add(ttl),
	}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	item, ok := c.data[key]
	if !ok {
		return nil, false
	}

	if item.expiry.Before(time.Now()) {
		delete(c.data, key)
		return nil, false
	}
	return item.value, true
}

func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.data, key)
}

func (c *Cache) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data = make(map[string]CacheItem)
}

func main() {
	cache := NewCache()

	// Adding data to the cache with a TTL of 2 seconds
	cache.Set("name", "mohit", 2*time.Second)
	cache.Set("weight", 75, 5*time.Second)

	// Retrieving data from the cache
	if val, ok := cache.Get("name"); ok {
		fmt.Println("Value for name:", val)
	}

	// Wait for some time to see expiry in action
	time.Sleep(3 * time.Second)

	// Retrieving expired data from the cache
	if _, ok := cache.Get("name"); !ok {
		fmt.Println("Name key has expired")
	}

	// Retrieving data before expiry
	if val, ok := cache.Get("weight"); ok {
		fmt.Println("Value for weight before expiry:", val)
	}

	// Wait for some time to see expiry in action
	time.Sleep(3 * time.Second)

	// Retrieving expired data from the cache
	if _, ok := cache.Get("weight"); !ok {
		fmt.Println("Weight key has expired")
	}

	// Deleting data from the cache
	cache.Set("key", "val", 2*time.Second)
	cache.Delete("key")

	// Clearing the cache
	cache.Clear()

	time.Sleep(time.Second) // Sleep to allow cache operations to complete
}
