package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	get = "GET"
	put = "PUT"
)

type Endpoint struct {
	Service *Service
}

func (e *Endpoint) handler(w http.ResponseWriter, r *http.Request) {
	t := r.Method
	switch t {
	case get:
		log.Println(get)
		// передаем название очереди и параметр для того, чтобы их добавить
	case put:
		e.Service.Add(r.URL.Path[1:], r.URL.Query().Get("v"))
	}

}

func (e *Endpoint) Start() error {
	http.HandleFunc("/", e.handler)
	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
