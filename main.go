package main

import (
	"log"
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
	s := &Service{queue: queue}
	e := &Endpoint{Service: s}

	err := e.Start()
	if err != nil {
		log.Println("start server error:", err)
	}
	log.Println("http server started at http://127.0.0.1:3333")

}
