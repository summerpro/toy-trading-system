package database

import "sync"

type CacheDb struct {
	db      DB
	cacheDb map[string][]byte
	mutex   sync.Mutex
}

func (c *CacheDb) Set(key, value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cacheDb[string(key)] = value
}

func (c *CacheDb) Get(key []byte) []byte {
	res := c.cacheDb[string(key)]
	if res == nil {
		res = c.db.Get(key)
		c.Set(key, res)
	}
	return res
}

func (c *CacheDb) Commit() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for k, v := range c.cacheDb {
		c.db.Set([]byte(k), v)
	}
}
