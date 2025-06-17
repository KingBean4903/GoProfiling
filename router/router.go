package router

import (
	"github.com/gorilla/mux"
	"github.com/KingBean4903/GoProfiling/handlers"
)

func Setup() *mux.Router {

		r := mux.NewRouter()

		r.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
		r.HandleFunc("/posts/{id}", handlers.GetPost).Methods("GET")
		r.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
		r.HandleFunc("/posts/{id}", handlers.UpdatePost).Methods("PUT")
		r.HandleFunc("/posts/{id}", handlers.DeletePost).Methods("DELETE")

		return r
}
