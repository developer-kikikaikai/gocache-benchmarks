package cache

import (
	"time"

	target "github.com/akyoto/cache"
)

type AkyotoCache struct {
	c *target.Cache
}

func NewAkyotoCache() Cache {
	return &AkyotoCache{
		c: target.New(5 * time.Minute),
	}
}

func (c *AkyotoCache) Get(k string) (interface{}, bool) {
	return c.c.Get(k)
}

func (c *AkyotoCache) Set(k string, x interface{}, d time.Duration) {
	c.c.Set(k, x, d)
}
