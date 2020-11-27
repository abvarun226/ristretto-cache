package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/abvarun226/ristretto-cache/cache"
)

func main() {
	// create a new cache.
	c := cache.New()

	// test the cache with multiple concurrent workers.
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i+1, c, wg)
	}
	wg.Wait()

	// done with the test, now close the cache.
	c.Close()
}

func worker(w int, c *cache.Store, wg *sync.WaitGroup) {
	defer wg.Done()

	key1 := "key1"
	val1 := "value1"
	tags := []string{"cachetag:k1:t1", "cachetag:k2:t1"}
	invalidateTags := []string{tags[0]}
	expiry := 15 * time.Minute

	fmt.Printf("[worker %2d] set key `%s` by tags: `%s`\n", w, key1, val1)
	c.SetByTags(key1, val1, expiry, tags)

	// wait for value to pass through buffers
	time.Sleep(5 * time.Millisecond)

	if res, found := c.Get(key1); found {
		fmt.Printf("[worker %2d] (1) get key `%s`: `%s`\n", w, key1, res)
	} else {
		fmt.Printf("[worker %2d] (1) key `%s` not found\n", w, key1)
	}

	time.Sleep(5 * time.Millisecond)

	// fmt.Printf("invalidate cache by tag `%s`", tags[0])
	c.Invalidate(invalidateTags)

	if res, found := c.Get(key1); found {
		fmt.Printf("[worker %2d] (2) get key `%s`: `%s`\n", w, key1, res)
	} else {
		fmt.Printf("[worker %2d] (2) key `%s` not found\n", w, key1)
	}
}
