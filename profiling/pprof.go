package profiling

import (
	_ "net/http/pprof"
	"log"
	"net/http"
)

func Start() {
	log.Println("Starting pprof server on :6060")
	log.Println(http.ListenAndServe(":6060", nil))
}
