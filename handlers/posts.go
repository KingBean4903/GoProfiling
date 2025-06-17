package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/KingBean4903/GoProfiling/models"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	posts := models.GetAllPosts()
	json.NewEncoder(w).Encode(posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, _ := strconv.Atoi(idStr)
		post := models.GetPost(id)
		json.NewEncoder(w).Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	
		var post models.Post
		_ = json.NewDecoder(r.Body).Decode(&post)
		models.AddPost(post)
		w.WriteHeader(http.StatusCreated)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, _ := strconv.Atoi(idStr)
		var post models.Post
		_ = json.NewDecoder(r.Body).Decode(&post)
		models.UpdatePost(id, post)
		w.WriteHeader(http.StatusNoContent)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, _ := strconv.Atoi(idStr)
		models.DeletePost(id)
		w.WriteHeader(http.StatusNoContent)
}
