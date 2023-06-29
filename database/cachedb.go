package database

import (
	"fmt"
	"github.com/summerpro/toy-trading-system/types"
	"sync"
)

type CacheDb struct {
	db      DB
	cacheDb map[types.Address]types.Account
	mutex   sync.Mutex
}

func NewCacheDb(initSize int, db DB) *CacheDb {
	return &CacheDb{
		db:      db,
		cacheDb: make(map[types.Address]types.Account, initSize),
		mutex:   sync.Mutex{},
	}
}

func (c *CacheDb) Set(key types.Address, value types.Account) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.cacheDb[key] = value
}

func (c *CacheDb) Get(key types.Address) types.Account {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	res, ok := c.cacheDb[key]
	if !ok {
		resBytes := c.db.Get([]byte(key))
		if resBytes == nil {
			return types.Account{
				Addr:   key,
				Amount: types.ZeroAmount,
			}
		}
		res = types.UnSerializeAccount(resBytes)
		c.cacheDb[key] = res
	}
	return res
}

func (c *CacheDb) Commit() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	for k, v := range c.cacheDb {
		c.db.Set([]byte(k), v.Serialize())
	}
}

func (c *CacheDb) RangeCache() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	fmt.Println("Account List:")
	fmt.Println("------------------")
	for k, v := range c.cacheDb {
		fmt.Println(k, v.Amount)
	}
	fmt.Println("------------------")
	fmt.Println()
}
