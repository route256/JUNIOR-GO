package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gitlab.ozon.dev/workshop3/workshop-3/internal/pkg/db"
	"gitlab.ozon.dev/workshop3/workshop-3/internal/pkg/repository"
	"gitlab.ozon.dev/workshop3/workshop-3/internal/pkg/repository/postgresql"
)

const port = ":9000"
const queryParamKey = "key"

type server1 struct {
	repo repository.ArticlesRepo
}

type addArticleRequest struct {
	Name   string `json:"name"`
	Rating int64  `json:"rating"`
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	database, err := db.NewDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer database.GetPool(ctx).Close()

	articleRepo := postgresql.NewArticles(database)

	implemetation := server1{repo: articleRepo}
	http.Handle("/", createRouter(implemetation))
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func createRouter(implemetation server1) *mux.Router {
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

	router.HandleFunc(fmt.Sprintf("/article/{%s:[0-9]+}", queryParamKey), func(w http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			key, status := parseGetID(req)
			if status != http.StatusOK {
				w.WriteHeader(status)
			}
			data, status := implemetation.Get(req.Context(), key)
			w.WriteHeader(status)
			w.Write(data)
		case http.MethodDelete:
			implemetation.Delete(w, req)
		default:
			fmt.Println("error")
		}
	})
	return router
}

func (s *server1) Get(ctx context.Context, key int64) ([]byte, int) {
	article, err := s.repo.GetByID(ctx, key)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return nil, http.StatusNotFound

		}
		return nil, http.StatusInternalServerError

	}
	articleJson, _ := json.Marshal(article)
	return articleJson, http.StatusOK
}

func parseGetID(req *http.Request) (int64, int) {
	key, ok := mux.Vars(req)[queryParamKey]
	if !ok {
		return 0, http.StatusBadRequest
	}
	keyInt, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return 0, http.StatusBadRequest
	}
	return keyInt, http.StatusOK
}

func (s *server1) Create(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var unm addArticleRequest
	if err = json.Unmarshal(body, &unm); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	articleRepo := &repository.Article{
		Name:   unm.Name,
		Rating: unm.Rating,
	}
	id, err := s.repo.Add(req.Context(), articleRepo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	articleRepo.ID = id
	articleJson, _ := json.Marshal(articleRepo)
	w.Write(articleJson)
}

func (s *server1) Update(_ http.ResponseWriter, req *http.Request) {
	fmt.Println("update")
}
func (s *server1) Delete(w http.ResponseWriter, req *http.Request) {
	fmt.Println("delete")

	//key, ok := mux.Vars(req)[queryParamKey]
	//if !ok {
	//	w.WriteHeader(http.StatusBadRequest)
	//	return
	//}
	//_, ok = s.data[key]
	//if !ok {
	//	w.WriteHeader(http.StatusNotFound)
	//	return
	//}
	//
	//delete(s.data, key)

}
