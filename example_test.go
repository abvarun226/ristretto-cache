package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/abvarun226/ristretto-cache/cache"
)

func BenchmarkSetByTags(b *testing.B) {
	c := cache.New()
	for n := 0; n < b.N; n++ {
		key := fmt.Sprintf("key%d", n)
		c.SetByTags(key, value, expiry, tags)
	}
}

func BenchmarkSetWithoutTags(b *testing.B) {
	c := cache.New()
	for n := 0; n < b.N; n++ {
		key := fmt.Sprintf("key%d", n)
		c.SetByTags(key, value, expiry, nil)
	}
}

func BenchmarkInvalidate(b *testing.B) {
	c := cache.New()
	c.SetByTags("key1", value, expiry, tags)
	for n := 0; n < b.N; n++ {
		c.Invalidate(tags)
	}
}

func BenchmarkGet(b *testing.B) {
	c := cache.New()
	key := "key1"
	c.SetByTags(key, value, expiry, tags)
	for n := 0; n < b.N; n++ {
		c.Get(key)
	}
}

const (
	value  = "setting a value in cache"
	expiry = 15 * time.Minute
)

var (
	tags = []string{"tag1", "tag2"}
)
