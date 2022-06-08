package cache

import (
	"time"

	"github.com/bluele/gcache"
)

const (
	cacheSize = 1_000_000
	cacheTTL  = 1 * time.Hour // default expiration
)

type GCache struct {
	c gcache.Cache
}

func NewGCache() Cache {
	return &GCache{
		c: gcache.New(cacheSize).Simple().Expiration(cacheTTL).Build(),
	}
}
func NewGCacheLRU() Cache {
	return &GCache{
		c: gcache.New(cacheSize).LRU().Expiration(cacheTTL).Build(),
	}
}
func NewGCacheLFU() Cache {
	return &GCache{
		c: gcache.New(cacheSize).LFU().Expiration(cacheTTL).Build(),
	}
}
func NewGCacheARC() Cache {
	return &GCache{
		c: gcache.New(cacheSize).ARC().Expiration(cacheTTL).Build(),
	}
}

func (c *GCache) Get(k string) (interface{}, bool) {
	v, err := c.c.Get(k)
	return v, (err == nil)
}

func (c *GCache) Set(k string, x interface{}, d time.Duration) {
	if d == 0 {
		c.c.Set(k, x)
	} else {
		c.c.SetWithExpire(k, x, d)
	}
}
