package main

import (
	"go-cache/cache"
	"go-cache/http"
	"go-cache/tcp"
)

func main() {
	c := cache.New("inmemory")
	go tcp.New(c).Listen()
	http.New(c).Listen()
}
