package cache

import (
	"encoding/json"
	"errors"
	"fmt"
	"sync"

	dbPkg "app0/internal/database"
	"app0/internal/structs"
)

type Cache struct {
	sync.RWMutex
	items map[string]structs.Order
}

// Инициализация
func New() *Cache {
	items := make(map[string]structs.Order)

	cache := Cache{
		items: items,
	}

	return &cache
}

// Добавление данных
func (c *Cache) Set(order structs.Order) {
	c.Lock()
	defer c.Unlock()
	c.items[order.Order_uid] = order
}

// Взятие данных
func (c *Cache) Get(key string) (item structs.Order, isFound bool) {
	c.RLock()
	defer c.RUnlock()

	item, isFound = c.items[key]
	return
}

// Удаление данных
func (c *Cache) Delete(key string) (err error) {
	c.Lock()
	defer c.Unlock()

	if _, found := c.items[key]; !found {
		return errors.New("key not found")
	}

	delete(c.items, key)

	return
}

// Вывод данных в консоль
func (c *Cache) PrintItems() {
	fmt.Println("Cached items:")
	for i, item := range c.items {
		fmt.Println("-----------------------------")
		fmt.Println("Item ", i)
		fmt.Println(item)
	}
}

// Восстановление данных из бд
func (c *Cache) Recover(db dbPkg.Database) (err error) {
	var order structs.Order
	dbOrders, err := db.SelectOrders()
	if err != nil {
		return
	}
	for _, byteOrder := range dbOrders {
		json.Unmarshal(byteOrder, &order)
		c.Set(order)
	}
	return
}
