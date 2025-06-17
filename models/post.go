package models

import "sync"

type Post struct {
	ID      int `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var (
	posts = make(map[int]Post)
	nextID = 1
	mu sync.RWMutex
)

func GetAllPosts() []Post {
	mu.RLock()
	defer mu.RUnlock()
	result := make([]Post, 0, len(posts))
	for _, post := range posts {
			result = append(result, post)
	}
	 return result
}

func GetPost(id int) Post {
	mu.RLock()
	defer mu.RUnlock()
	return posts[id]
}

func AddPost(post Post) {
		mu.Lock()
		defer mu.Unlock()
		post.ID = nextID
		nextID++
		posts[post.ID] = post
}

func UpdatePost(id int, post Post) {
		mu.Lock()
		defer mu.Unlock()
		post.ID = id
		posts[id] = post
}

func DeletePost(id int) {
		mu.Lock()
		defer mu.Unlock()
		delete(posts, id)
}
