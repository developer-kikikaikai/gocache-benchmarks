package cache

import (
	"testing"
	"time"
)

var goCache Cache

func init() {
	goCache = NewGoCache()
}

func Benchmark_GoCache_EveryTimeWithoutTimeout(b *testing.B) {
	GetSet(goCache, b, 0)
}

func Benchmark_GoCache_CacheEveryTimeWitTimeout(b *testing.B) {
	GetSet(goCache, b, 1*time.Second)
}

func Benchmark_GoCache_CacheForLargeStringsEveryTime(b *testing.B) {
	GetSetLarge(goCache, b)
}
