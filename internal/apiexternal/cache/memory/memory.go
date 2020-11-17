package memory

import (
	"errors"
	"sync"
	"time"
)

type Storage struct {
	sync.RWMutex
	items             map[string]Item
	defaultExpiration time.Duration
	cleanupInterval   time.Duration
}

type Item struct {
	Value      interface{}
	Expiration int64
	Created    time.Time
}

func InitCash(defaultExpiration, cleanupInterval time.Duration) *Storage {

	items := make(map[string]Item)

	cache := Storage{
		items:             items,
		defaultExpiration: defaultExpiration,
		cleanupInterval:   cleanupInterval,
	}

	if cleanupInterval > 0 {
		//move data to database
		cache.StartGC()
	}
	return &cache
}

// Set setting a cache by key
func (c *Storage) Set(key string, value interface{}, duration time.Duration) {

	var expiration int64

	if duration == 0 {
		duration = c.defaultExpiration
	}
	if duration > 0 {
		expiration = time.Now().Add(duration).UnixNano()
	}

	c.Lock()
	defer c.Unlock()

	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
		Created:    time.Now(),
	}
}

func (c *Storage) Get(key string) (interface{}, bool) {

	c.RLock()
	defer c.RUnlock()

	item, found := c.items[key]

	if !found {
		return nil, false
	}

	if item.Expiration > 0 {

		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}
	}

	return item.Value, true
}

func (c *Storage) Delete(key string) error {

	c.Lock()
	defer c.Unlock()

	if _, found := c.items[key]; !found {
		return errors.New("NullPointer Exception: key not found")
	}

	delete(c.items, key)
	return nil
}

func (c *Storage) StartGC() {
	go c.GC()
}

func (c *Storage) GC() {

	for {
		<-time.After(c.cleanupInterval)

		if c.items == nil {
			return
		}

		if keys := c.expiredKeys(); len(keys) != 0 {
			c.clearItems(keys)

		}
	}
}

func (c *Storage) expiredKeys() (keys []string) {

	c.RLock()
	defer c.RUnlock()

	for k, i := range c.items {
		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
			keys = append(keys, k)
		}
	}
	return
}

func (c *Storage) clearItems(keys []string) {

	c.Lock()
	defer c.Unlock()

	for _, k := range keys {
		delete(c.items, k)
	}
}
