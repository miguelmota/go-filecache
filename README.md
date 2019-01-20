<h3 align="center">
  <br />
  <img src="https://user-images.githubusercontent.com/168240/51435448-5fa07200-1c2c-11e9-99c3-21d4ee2b7a4a.png" alt="logo" width="500" />
  <br />
  <br />
  <br />
</h3>

# go-filecache

> Fast arbitrary data caching to tmp files in Go

[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelmota/go-filecache/master/LICENSE) [![Build Status](https://travis-ci.org/miguelmota/go-filecache.svg?branch=master)](https://travis-ci.org/miguelmota/go-filecache) [![Go Report Card](https://goreportcard.com/badge/github.com/miguelmota/go-filecache?)](https://goreportcard.com/report/github.com/miguelmota/go-filecache) [![GoDoc](https://godoc.org/github.com/miguelmota/go-filecache?status.svg)](https://godoc.org/github.com/miguelmota/go-filecache)

## Install

```bash
go get -u github.com/miguelmota/go-filecache
```

## Documentation

[https://godoc.org/github.com/miguelmota/go-filecache](https://godoc.org/github.com/miguelmota/go-filecache)

## Getting started

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/miguelmota/go-filecache"
)

func main() {
	key := "foo"
	data := []byte("bar") // can be of any type
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
```

## License

[MIT](LICENSE)
