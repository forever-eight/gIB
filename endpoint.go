package main

import (
	"encoding/json"
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
	defer r.Body.Close()
	switch r.Method {
	case get:
		value := e.Service.Get(r.URL.Path[1:])
		if value == nil {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.Header().Set("Content-Type", "application/json")
			resp := make(map[string]string)
			resp["result"] = *value
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("error in JSON marshal: %s", err)
			}
			w.Write(jsonResp)
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
