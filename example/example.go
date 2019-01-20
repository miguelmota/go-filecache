package main

import (
	"fmt"
	"log"
	"time"

	filecache "github.com/miguelmota/go-filecache"
)

func main() {
	key := "foo"
	data := []byte("bar") // can be of any time
	expire := 1 * time.Hour

	// caching data
	err := filecache.Set(key, data, expire)
	if err != nil {
		log.Fatal(err)
	}

	// reading cached data
	var dst []byte
	found, err := filecache.Get(key, &dst)
	if err != nil {
		log.Fatal(err)
	}
	if found {
		fmt.Println(string(dst)) // "bar"
	}
}
