package cache

import (
	"testing"
	"time"
)

func Benchmark_go_cache_EveryTimeWithoutTimeout(b *testing.B) {
	GetSet(NewPatrickmnGoCache(), b, 0)
}

func Benchmark_go_cache_CacheEveryTimeWitTimeout(b *testing.B) {
	GetSet(NewPatrickmnGoCache(), b, 1*time.Second)
}

func Benchmark_go_cache_CacheForLargeStringsEveryTime(b *testing.B) {
	GetSetLarge(NewPatrickmnGoCache(), b)
}
