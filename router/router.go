package router

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/KingBean4903/GoProfiling/handlers"
)

func Setup() mux.Router {

		r := mux.NewRouter()

		r.HandleFunc("/posts", handler.GetPosts).Methods("GET")
		r.HandleFunc("/posts/{id}", handler.GetPost).Methods("GET")
		r.HandleFunc("/posts", handler.CreatePost).Methods("POST")
		r.HandleFunc("/posts/{id}", handler.UpdatePost).Methods("PUT")
		r.HandleFunc("/posts/{id}", handler.DeletePost).Methods("DELETE")

		return r
}
