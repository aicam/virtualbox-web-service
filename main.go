package main

import (
	"log"
	"net/http"
)

func main() {

	s := NewServer()
	s.Route()

	err := http.ListenAndServe("0.0.0.0:2020", s.Router)
	if err != nil {
		log.Print(err)
	}
}
