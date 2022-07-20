package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const (
	get = "GET"
	put = "PUT"
)

type Endpoint struct {
	Service *Service
}

func (e *Endpoint) handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	switch r.Method {
	case get:
		timeout := r.URL.Query().Get("timeout")
		var value *string
		if timeout != "" {
			t, err := strconv.Atoi(timeout)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}

			value = e.Service.Wait(t, r.URL.Path[1:])
		} else {
			value = e.Service.Get(r.URL.Path[1:])
		}
		if value == nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			io.WriteString(w, *value)
		}
	case put:
		if r.URL.Query().Get("v") == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
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
