package handlers

import (
	"net/http/httptest"
	"testing"
)
func BenchmarkGetPosts(b *testing.B) {
	
	req := httptest.NewRequest("GET", "/posts", nil)
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
			GetPosts(w, req)
	}

}
