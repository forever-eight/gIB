package main

import (
	"log"
	"sync"
)

type Node struct {
	Current string
	Next    *Node
}

// Queue может быть не нужно
type Queue struct {
	// first value
	First *Node
	// last value
	Last *Node
}

func main() {
	queue := make(map[string]*Queue)
	waiter := make(map[string]int)
	mu := &sync.Mutex{}
	c := make(chan string)

	s := &Service{
		queue:  queue,
		mu:     mu,
		waiter: waiter,
		c:      c,
	}
	e := &Endpoint{Service: s}

	err := e.Start()
	if err != nil {
		log.Println("start server error:", err)
	}
	log.Println("http server started at http://127.0.0.1:3333")

}
