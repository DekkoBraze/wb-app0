package cache

import (
	"errors"
	"sync"
	"log"

	dbPkg "app0/internal/database"
)

type Cache struct {
	sync.RWMutex
	items map[string]dbPkg.Order
}

func New() *Cache {
	items := make(map[string]dbPkg.Order)

	cache := Cache{
		items: items,
	}

	return &cache
}

func (c *Cache) Set(orders dbPkg.Order) {
	c.Lock()
	defer c.Unlock()

	c.items[orders.Order_uid] = orders
	log.Print("data in cache!")
}

func (c *Cache) Get(key string) (item dbPkg.Order, isFound bool) {
	c.RLock()
	defer c.RUnlock()

	item, isFound = c.items[key]
	return
}

func (c *Cache) Delete(key string) (err error) {
	c.Lock()
	defer c.Unlock()

	if _, found := c.items[key]; !found {
		return errors.New("key not found")
	}

	delete(c.items, key)

	return
}
