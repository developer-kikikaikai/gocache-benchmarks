package cache

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

const (
	keySize = 2048
	//valueSize   = 1 * 1024 * 1024 * 1024
	valueSize   = 1 * 1024 * 1024
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var largeKey string
var largeValue string

func init() {
	largeKey = NewLargeKey()
	largeValue = NewLargeValue()
}

func GetSet(c Cache, b *testing.B, t time.Duration) {
	getSet(c, b, t, "a", "b")
}

func GetSetLarge(c Cache, b *testing.B) {
	getSet(c, b, 0, largeKey, largeValue)
}

func getSet(c Cache, b *testing.B, t time.Duration, k, v string) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Set(k, v, t)
		res, ok := c.Get(k)
		if !ok {
			b.Errorf("Get vaule failed")
		}
		if res != v {
			b.Errorf("Get vaule is different from set value")
		}
	}
}

func NewLargeKey() string {
	return RandStringBytes(keySize)
}

func NewLargeValue() string {
	return RandStringBytes(valueSize)
}

//get code from https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return fmt.Sprintf("%s", b)
}
