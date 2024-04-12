package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	securePort   = ":9000"
	insecurePort = ":9001"
)

func main() {
	go func() {
		secure()
	}()

	insecure()
}

func secure() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("secure")
	})

	if err := http.ListenAndServeTLS(securePort, "./server.crt", "./server.key", nil); err != nil {
		log.Fatal(err)
	}
}

func insecure() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println("insecure")
	})

	if err := http.ListenAndServe(insecurePort, mux); err != nil {
		log.Fatal(err)
	}
}
