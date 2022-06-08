package cache

import (
	"testing"
	"time"
)

func Benchmark_GCache_EveryTimeWithoutTimeout(b *testing.B) {
	GetSet(NewGCache(), b, 0)
}

func Benchmark_GCache_CacheEveryTimeWitTimeout(b *testing.B) {
	GetSet(NewGCache(), b, 1*time.Second)
}

func Benchmark_GCache_CacheForLargeStringsEveryTime(b *testing.B) {
	GetSetLarge(NewGCache(), b)
}

func Benchmark_GCacheLRU_EveryTimeWithoutTimeout(b *testing.B) {
	GetSet(NewGCacheLRU(), b, 0)
}

func Benchmark_GCacheLRU_CacheEveryTimeWitTimeout(b *testing.B) {
	GetSet(NewGCacheLRU(), b, 1*time.Second)
}

func Benchmark_GCacheLRU_CacheForLargeStringsEveryTime(b *testing.B) {
	GetSetLarge(NewGCacheLRU(), b)
}

func Benchmark_GCacheLFU_EveryTimeWithoutTimeout(b *testing.B) {
	GetSet(NewGCacheLFU(), b, 0)
}

func Benchmark_GCacheLFU_CacheEveryTimeWitTimeout(b *testing.B) {
	GetSet(NewGCacheLFU(), b, 1*time.Second)
}

func Benchmark_GCacheLFU_CacheForLargeStringsEveryTime(b *testing.B) {
	GetSetLarge(NewGCacheLFU(), b)
}

func Benchmark_GCacheARC_EveryTimeWithoutTimeout(b *testing.B) {
	GetSet(NewGCacheARC(), b, 0)
}

func Benchmark_GCacheARC_CacheEveryTimeWitTimeout(b *testing.B) {
	GetSet(NewGCacheARC(), b, 1*time.Second)
}

func Benchmark_GCacheARC_CacheForLargeStringsEveryTime(b *testing.B) {
	GetSetLarge(NewGCacheARC(), b)
}
