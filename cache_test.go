package cachebench_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/allegro/bigcache"
	"github.com/bluele/gcache"
	"github.com/coocood/freecache"
	hashicorp "github.com/hashicorp/golang-lru"
	"github.com/patrickmn/go-cache"
)

const (
	// Standard cache size for all implementations
	cacheSize = 10000
	// Standard TTL for all implementations
	cacheTTL = 10 * time.Minute
)

// runBenchmark provides a consistent benchmark structure for all cache implementations
func runBenchmark(b *testing.B, name string, set func(key, val string), get func(key string) bool) {
	b.Run(name+"/Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			set(fmt.Sprintf("key%d", i), "value")
		}
	})

	b.Run(name+"/Get", func(b *testing.B) {
		// Pre-populate cache
		for i := 0; i < cacheSize; i++ {
			set(fmt.Sprintf("key%d", i), "value")
		}
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			get(fmt.Sprintf("key%d", i%cacheSize))
		}
	})

	b.Run(name+"/Mixed", func(b *testing.B) {
		// Pre-populate cache to 50%
		for i := 0; i < cacheSize/2; i++ {
			set(fmt.Sprintf("key%d", i), "value")
		}
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			if i%5 == 0 { // 20% writes, 80% reads
				set(fmt.Sprintf("key%d", i), "value")
			} else {
				get(fmt.Sprintf("key%d", i%cacheSize))
			}
		}
	})
}

func BenchmarkHashicorpLRU(b *testing.B) {
	c, _ := hashicorp.New(cacheSize)

	runBenchmark(b, "HashicorpLRU",
		func(key, val string) {
			c.Add(key, val)
		},
		func(key string) bool {
			_, ok := c.Get(key)
			return ok
		},
	)
}

func BenchmarkGoCache(b *testing.B) {
	c := cache.New(cacheTTL, cacheTTL*2)

	runBenchmark(b, "GoCache",
		func(key, val string) {
			c.Set(key, val, cache.DefaultExpiration)
		},
		func(key string) bool {
			_, found := c.Get(key)
			return found
		},
	)
}

func BenchmarkFreecache(b *testing.B) {
	// 100MB cache size - typical production setting
	c := freecache.NewCache(100 * 1024 * 1024)

	runBenchmark(b, "Freecache",
		func(key, val string) {
			c.Set([]byte(key), []byte(val), int(cacheTTL.Seconds()))
		},
		func(key string) bool {
			_, err := c.Get([]byte(key))
			return err == nil
		},
	)
}

func BenchmarkBigCache(b *testing.B) {
	config := bigcache.Config{
		Shards:             1024,
		LifeWindow:         cacheTTL,
		MaxEntriesInWindow: cacheSize * 10,
		MaxEntrySize:       500,
		Verbose:            false,
		HardMaxCacheSize:   100, // 100MB
		CleanWindow:        1 * time.Minute,
	}
	c, _ := bigcache.NewBigCache(config)

	runBenchmark(b, "BigCache",
		func(key, val string) {
			c.Set(key, []byte(val))
		},
		func(key string) bool {
			_, err := c.Get(key)
			return err == nil
		},
	)
}

func BenchmarkGCache(b *testing.B) {
	c := gcache.New(cacheSize).
		LRU().
		Expiration(cacheTTL).
		Build()

	runBenchmark(b, "GCache",
		func(key, val string) {
			c.Set(key, val)
		},
		func(key string) bool {
			_, err := c.Get(key)
			return err == nil
		},
	)
}

// BenchmarkRealWorld simulates common caching scenarios
func BenchmarkRealWorld(b *testing.B) {
	scenarios := []struct {
		name        string
		keyPattern  func(i int) string
		valueSize   int
		readPercent int
	}{
		{
			name:        "UserSession",
			keyPattern:  func(i int) string { return fmt.Sprintf("user%d", i%1000) }, // 1000 unique users
			valueSize:   200,                                                          // Small session data
			readPercent: 95,                                                           // Mostly reads
		},
		{
			name:        "APICache",
			keyPattern:  func(i int) string { return fmt.Sprintf("api%d", i%100) }, // 100 endpoints
			valueSize:   2048,                                                       // Medium JSON responses
			readPercent: 80,                                                         // Typical read/write ratio
		},
	}

	for _, scenario := range scenarios {
		b.Run(scenario.name, func(b *testing.B) {
			// Test with go-cache as a representative implementation
			c := cache.New(cacheTTL, cacheTTL*2)
			value := make([]byte, scenario.valueSize)

			for i := 0; i < b.N; i++ {
				key := scenario.keyPattern(i)
				if i%100 < scenario.readPercent {
					c.Get(key)
				} else {
					c.Set(key, value, cache.DefaultExpiration)
				}
			}
		})
	}
}

