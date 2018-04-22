package main

import (
	"fmt"
	"log"
	"time"

	"github.com/miguelmota/fcache"
)

func main() {
	key := "foo"
	data := []byte("bar") // can be of any time
	expire := 1 * time.Hour

	// caching data
	err := fcache.Set(key, data, expire)
	if err != nil {
		log.Fatal(err)
	}

	// reading cached data
	var dst []byte
	found, err := fcache.Get(key, &dst)
	if err != nil {
		log.Fatal(err)
	}
	if found {
		fmt.Println(string(dst)) // "bar"
	}
}
