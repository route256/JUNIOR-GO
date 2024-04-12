package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":9000"

func main() {
	var implemetation server1
	http.HandleFunc("/article", authMiddleware(articleHandler(implemetation)))
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func articleHandler(implemetation server1) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			implemetation.Get(w, req)
		case http.MethodPost:
			implemetation.Create(w, req)
		case http.MethodPut:
			implemetation.Update(w, req)
		case http.MethodDelete:
			implemetation.Delete(w, req)
		default:
			log.Fatal("unsupported method")
		}
	}
}

type server1 struct {
	a string
}

func (s *server1) Create(_ http.ResponseWriter, req *http.Request) {
	fmt.Println("create")
}
func (s *server1) Update(_ http.ResponseWriter, req *http.Request) {
	fmt.Println("update")
}
func (s *server1) Delete(_ http.ResponseWriter, req *http.Request) {
	fmt.Println("delete")

}
func (s *server1) Get(_ http.ResponseWriter, req *http.Request) {
	fmt.Println("get")

}

// Authorization: Basic (<user:password> в base64)
func authMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		username, password, ok := req.BasicAuth()
		if ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println(username, password)
		handler.ServeHTTP(w, req)
	}
}
