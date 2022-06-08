package cache

import (
	"time"

	"github.com/dgraph-io/ristretto"
)

type Ristretto struct {
	c *ristretto.Cache
}

func NewRistretto() Cache {
	c, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 50,
		MaxCost:     1 * 1024 * 1024 * 1024,
		BufferItems: 64,
	})

	return &Ristretto{
		c: c,
	}
}

func (c *Ristretto) Get(k string) (interface{}, bool) {
	return c.c.Get(k)
}

func (c *Ristretto) Set(k string, x interface{}, d time.Duration) {
	c.c.Set(k, x, 0)
	c.c.Wait()
}
