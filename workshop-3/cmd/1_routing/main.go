package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":9000"

func main() {
	//многоуровевый базовый путь(роут)
	http.HandleFunc("/", rootHandler)

	// в случае повторной регистрации - паника
	//http.HandleFunc("/", rootHandler)

	// фиксированный путь
	http.HandleFunc("/home", homeHandler)

	//многоуровневый путь
	http.HandleFunc("/article/", articleHandler)

	// фиксированный путь
	http.HandleFunc("/article/hello", helloHandler)

	// регулярки не работают
	//http.HandleFunc("article/*/hello/world", nil)

	//редирект
	http.HandleFunc("/redirect", http.RedirectHandler("http://localhost:9000/home", http.StatusPermanentRedirect).ServeHTTP)
	http.HandleFunc("/redirect/2", redirectHandler)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func rootHandler(_ http.ResponseWriter, req *http.Request) {
	fmt.Println("root")
}

func homeHandler(_ http.ResponseWriter, req *http.Request) {
	fmt.Println("home")
}

func articleHandler(_ http.ResponseWriter, req *http.Request) {
	fmt.Println("article")
}

func helloHandler(_ http.ResponseWriter, req *http.Request) {
	fmt.Println("hello")
}

func redirectHandler(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "http://localhost:9000/article", http.StatusPermanentRedirect)
}
