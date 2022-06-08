package cache

import (
	"time"

	"github.com/allegro/bigcache"
	target "github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
)

type GoCache struct {
	c *target.Cache
}

func NewGoCache() Cache {
	bigcacheClient, _ := bigcache.NewBigCache(bigcache.DefaultConfig(5 * time.Minute))
	bigcacheStore := store.NewBigcache(bigcacheClient, nil) // No otions provided (as second argument)
	return &GoCache{
		c: target.New(bigcacheStore),
	}
}

func (c *GoCache) Get(k string) (interface{}, bool) {
	v, err := c.c.Get(k)
	return v, (err == nil)
}

func (c *GoCache) Set(k string, x interface{}, d time.Duration) {
	c.c.Set(k, x, nil)
}
