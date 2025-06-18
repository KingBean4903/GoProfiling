package handlers

import (
	"encoding/json"
	"sync"
	"context"
	"sync"
	"log"
	"time"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
	"github.com/KingBean4903/GoProfiling/models"
)



var postsSlicePool = sync.Pool {
		New: func() any {
			slice := make([]models.Post, 0, 100)
			return &slice
		},
}

var postPool = sync.Pool {
		New: func() any {
				return new(models.Post)
		},
}

var (
	postQueue  = make(chan models.Post, 100)
	numWorkers = 4
	wg			 		 sync.WaitGroup
)

func StartPostWithWorkers(ctx context.Context) {
	for i:= 0; i < numWorkers; i++ {
			go func(workerID int) {
				for {
						select {					
						case <- ctx.Done():
								log.Printf("[Worker %d] context cancelled, draining queue... ", workerID)
								for post := range postQueue {
										models.AddPost(post)
										wg.Done()
								}
								log.Printf("[Worker %d] drained and exciting:", workerID)
								return
						case post := <- postQueue:
									start := time.Now()
									models.AddPost(post)
									elapsed := time.Since(start)
									fmt.Printf("[Worker %d] processed post in %s\n", workerID, elapsed)
									wg.Done()
					}
				}
			}(i)
	}

}

/*func init() {	
	StartPostWithWorkers()
	go func() {
			for post := range postQueue {
					models.AddPost(post)
			}
	}()
}*/

func GetPosts(w http.ResponseWriter, r *http.Request) {
	pooled := postsSlicePool.Get().(*[]models.Post)
	*pooled = (*pooled)[:0]

	all := models.GetAllPosts()
	*pooled = append(*pooled, all...)
	json.NewEncoder(w).Encode(*pooled)

	postsSlicePool.Put(pooled)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, _ := strconv.Atoi(idStr)
		post := models.GetPost(id)
		json.NewEncoder(w).Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
		post := postPool.Get().(*models.Post)
		defer postPool.Put(post)
		*post = models.Post{}

		if err := json.NewDecoder(r.Body).Decode(post); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}

		wg.Add(1)
		postQueue <- *post


}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, _ := strconv.Atoi(idStr)
		post := postPool.Get().(*models.Post)
		defer postPool.Put(post)
		*post = models.Post{}

		if err := json.NewDecoder(r.Body).Decode(post); err != nil {
				http.Error(w, "Invalid input", http.StatusBadRequest)
		}
		models.UpdatePost(id, *post)
		w.WriteHeader(http.StatusNoContent)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, _ := strconv.Atoi(idStr)
		models.DeletePost(id)
		w.WriteHeader(http.StatusNoContent)
}
