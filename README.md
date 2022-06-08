- github.com/patrickmn/go-cache is one of the simplest and fastest tool.
- If you'd like to configure some parameters, Serialize logic, etc, it's better to use github.com/bluele/gcache.
- github.com/dgraph-io/ristretto is more customizable package I guess. 


```
d0303tujihi1noMacBook-Pro:gocache-benchmarks tsuji$ make bench
go1.18 test ./... -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: gocache-benchmarks/cache
cpu: Intel(R) Core(TM) i5-7360U CPU @ 2.30GHz
Benchmark_AkyotoCache_EveryTimeWithoutTimeout-4         	 5225078	       220.1 ns/op	      72 B/op	       4 allocs/op
Benchmark_AkyotoCache_CacheEveryTimeWitTimeout-4        	 3179415	       385.0 ns/op	      72 B/op	       4 allocs/op
Benchmark_AkyotoCache_CacheForLargeStringsEveryTime-4   	 3604564	       343.1 ns/op	      72 B/op	       4 allocs/op
Benchmark_GCache_EveryTimeWithoutTimeout-4              	 2931758	       411.5 ns/op	      72 B/op	       4 allocs/op
Benchmark_GCache_CacheEveryTimeWitTimeout-4             	 2311746	       520.1 ns/op	      96 B/op	       5 allocs/op
Benchmark_GCache_CacheForLargeStringsEveryTime-4        	 2237434	       539.5 ns/op	      72 B/op	       4 allocs/op
Benchmark_GCacheLRU_EveryTimeWithoutTimeout-4           	 2987014	       404.0 ns/op	      72 B/op	       4 allocs/op
Benchmark_GCacheLRU_CacheEveryTimeWitTimeout-4          	 2307759	       524.6 ns/op	      96 B/op	       5 allocs/op
Benchmark_GCacheLRU_CacheForLargeStringsEveryTime-4     	 2216824	       541.6 ns/op	      72 B/op	       4 allocs/op
Benchmark_GCacheLFU_EveryTimeWithoutTimeout-4           	 2730613	       447.9 ns/op	      72 B/op	       4 allocs/op
Benchmark_GCacheLFU_CacheEveryTimeWitTimeout-4          	 2142141	       558.7 ns/op	      96 B/op	       5 allocs/op
Benchmark_GCacheLFU_CacheForLargeStringsEveryTime-4     	 2109086	       575.2 ns/op	      72 B/op	       4 allocs/op
Benchmark_GCacheARC_EveryTimeWithoutTimeout-4           	 2529687	       479.7 ns/op	      72 B/op	       4 allocs/op
Benchmark_GCacheARC_CacheEveryTimeWitTimeout-4          	 2010477	       616.8 ns/op	      96 B/op	       5 allocs/op
Benchmark_GCacheARC_CacheForLargeStringsEveryTime-4     	 1333800	       876.2 ns/op	      72 B/op	       4 allocs/op
Benchmark_go_cache_EveryTimeWithoutTimeout-4            	 4762456	       254.6 ns/op	      16 B/op	       1 allocs/op
Benchmark_go_cache_CacheEveryTimeWitTimeout-4           	 4695368	       252.3 ns/op	      16 B/op	       1 allocs/op
Benchmark_go_cache_CacheForLargeStringsEveryTime-4      	 3645369	       321.3 ns/op	      16 B/op	       1 allocs/op
--- FAIL: Benchmark_GoCache_EveryTimeWithoutTimeout
    common_test.go:42: Get vaule is different from set value
--- FAIL: Benchmark_GoCache_CacheEveryTimeWitTimeout
    common_test.go:42: Get vaule is different from set value
2022/06/09 04:14:13 Allocated new queue in 214.449µs; Capacity: 2686292 
--- FAIL: Benchmark_GoCache_CacheForLargeStringsEveryTime
    common_test.go:42: Get vaule is different from set value
Benchmark_ristretto_EveryTimeWithoutTimeout-4           	  661338	      1800 ns/op	     231 B/op	       6 allocs/op
Benchmark_ristretto_CacheEveryTimeWitTimeout-4          	  639409	      1764 ns/op	     231 B/op	       6 allocs/op
Benchmark_ristretto_CacheForLargeStringsEveryTime-4     	  511674	      2349 ns/op	     231 B/op	       6 allocs/op
FAIL
exit status 1
FAIL	gocache-benchmarks/cache	35.552s
FAIL
make: *** [bench] Error 1
```
