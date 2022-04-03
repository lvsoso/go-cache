package main

import (
	"flag"
	"go-cache/cache"
	"go-cache/http"
	"go-cache/tcp"
	"log"
)

func main() {
	typ := flag.String("type", "inmemory", "cache type")
	flag.Parse()
	log.Println("type is", *typ)
	c := cache.New(*typ)
	go tcp.New(c).Listen()
	http.New(c).Listen()
}
