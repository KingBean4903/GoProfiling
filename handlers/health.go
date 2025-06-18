package handlers

import (
	"net/http"
	"sync/atomic"
)

var ready atomic.Bool

func MarkReady() {
	ready.Store(true)
}

func MarkNotReady() {
	ready.Store(false)
}

func Healthz(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
}

func Readyz(w http.ResponseWriter, r *http.Request) {
			if ready.Load() {
					w.WriteHeader(http.StatusOK)
					w.Write([]byte("ready"))
			} else {
					http.Error(w, "not ready", http.StatusServiceUnavailable)
			}
}
