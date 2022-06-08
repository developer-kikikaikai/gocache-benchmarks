package cache

import (
	"time"

	target "github.com/patrickmn/go-cache"
)

//to set same condition, we define struct
type PatrickmnGoCache struct {
	c *target.Cache
}

func NewPatrickmnGoCache() Cache {
	return &PatrickmnGoCache{
		c: target.New(3*time.Minute, 5*time.Minute),
	}
}

func (c *PatrickmnGoCache) Get(k string) (interface{}, bool) {
	return c.c.Get(k)
}

func (c *PatrickmnGoCache) Set(k string, x interface{}, d time.Duration) {
	c.c.Set(k, x, d)
}
