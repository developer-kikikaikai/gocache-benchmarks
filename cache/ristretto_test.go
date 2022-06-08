package cache

import (
	"testing"
	"time"
)

func Benchmark_ristretto_EveryTimeWithoutTimeout(b *testing.B) {
	GetSet(NewRistretto(), b, 0)
}

func Benchmark_ristretto_CacheEveryTimeWitTimeout(b *testing.B) {
	GetSet(NewRistretto(), b, 1*time.Second)
}

func Benchmark_ristretto_CacheForLargeStringsEveryTime(b *testing.B) {
	GetSetLarge(NewRistretto(), b)
}
