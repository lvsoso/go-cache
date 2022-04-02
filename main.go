package main

import (
	"go-cache/cache"
	"go-cache/http"
)

func main() {
	c := cache.New("inmemory")
	http.New(c).Listen()
}
