package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const port = ":9000"
const queryParamKey = "key"

type server1 struct {
	data map[string]string
}

type requestCustom struct {
	Key   string
	Value string
}

func main() {
	implemetation := server1{data: make(map[string]string)}
	router := mux.NewRouter()

	router.HandleFunc("/article", func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodPost:
			implemetation.Create(w, req)
		case http.MethodPut:
			implemetation.Update(w, req)
		default:
			fmt.Println("error")
		}
	})

	router.HandleFunc(fmt.Sprintf("/article/{%s:[A-z]+}", queryParamKey), func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			implemetation.Get(w, req)
		case http.MethodDelete:
			implemetation.Delete(w, req)
		default:
			fmt.Println("error")
		}
	})
	http.Handle("/", router)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func (s *server1) Create(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var unm requestCustom
	if err = json.Unmarshal(body, &unm); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if unm.Key == "" || unm.Value == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, ok := s.data[unm.Key]; ok {
		w.WriteHeader(http.StatusConflict)
		return
	}
	s.data[unm.Key] = unm.Value
}

func (s *server1) Update(_ http.ResponseWriter, req *http.Request) {
	fmt.Println("update")
}
func (s *server1) Delete(w http.ResponseWriter, req *http.Request) {
	key, ok := mux.Vars(req)[queryParamKey]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	_, ok = s.data[key]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	delete(s.data, key)

}
func (s *server1) Get(w http.ResponseWriter, req *http.Request) {
	key, ok := mux.Vars(req)[queryParamKey]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	value, ok := s.data[key]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if _, err := w.Write([]byte(value)); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
