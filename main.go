package main

import (
	"log"
	"net/http"
	"github.com/KingBean4903/router"
	"github.com/KingBean4903/profiling"
)

func main() {
	
	go profiling.Start()

	r := router.Setup()
	log.Println("Start server on port: 8070")
	log.Fatal(http.ListenAndServe(":8070", r))

}
