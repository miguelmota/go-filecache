# fcache

> Arbitrary data caching to tmp files in Go

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelmota/fcache/master/LICENSE.md) [![Build Status](https://travis-ci.org/miguelmota/fcache.svg?branch=master)](https://travis-ci.org/miguelmota/fcache) [![Go Report Card](https://goreportcard.com/badge/github.com/miguelmota/fcache?)](https://goreportcard.com/report/github.com/miguelmota/fcache) [![GoDoc](https://godoc.org/github.com/miguelmota/fcache?status.svg)](https://godoc.org/github.com/miguelmota/fcache)

## Install

```bash
go get -u github.com/miguelmota/fcache
```

## Documentation

[https://godoc.org/github.com/miguelmota/fcache](https://godoc.org/github.com/miguelmota/fcache)

## Usage

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/miguelmota/fcache"
)

func main() {
	key := "foo"
	data := []byte("bar") // can be of any type
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
```

## License

MIT
