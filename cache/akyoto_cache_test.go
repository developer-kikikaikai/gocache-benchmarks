package cache

import (
	"testing"
	"time"
)

func Benchmark_AkyotoCache_EveryTimeWithoutTimeout(b *testing.B) {
	GetSet(NewAkyotoCache(), b, 0)
}

func Benchmark_AkyotoCache_CacheEveryTimeWitTimeout(b *testing.B) {
	GetSet(NewAkyotoCache(), b, 1*time.Second)
}

func Benchmark_AkyotoCache_CacheForLargeStringsEveryTime(b *testing.B) {
	GetSetLarge(NewAkyotoCache(), b)
}
