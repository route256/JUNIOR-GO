package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	defaultPort = ":9000"
	customPort  = ":9001"
)

func main() {
	go func() {
		defaultMux()
	}()
	customMux()
}

func defaultMux() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("default mux")
	})

	if err := http.ListenAndServe(defaultPort, nil); err != nil {
		log.Fatal(err)
	}

}

func customMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("custom mux")
	})
	if err := http.ListenAndServe(customPort, mux); err != nil {
		log.Fatal()
	}
}
