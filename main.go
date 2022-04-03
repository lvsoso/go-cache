package main

import (
	"flag"
	"go-cache/cache"
	"go-cache/http"
	"go-cache/tcp"
	"log"

	"go-cache/cluster"
)

func main() {
	typ := flag.String("type", "inmemory", "cache type")
	node := flag.String("node", "127.0.0.1", "node address")
	nodePort := flag.Int("nodePort", 7946, "node port")
	clus := flag.String("cluster", "", "cluster address")
	flag.Parse()
	log.Println("type is", *typ)
	log.Println("node is", *node)
	log.Println("cluster is", *clus)
	c := cache.New(*typ)
	n, e := cluster.New(*node, *nodePort, *clus)
	if e != nil {
		panic(e)
	}
	go tcp.New(c, n).Listen()
	http.New(c, n).Listen()
}
