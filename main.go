package main

import (
	"context"
	"os"
	"runtime"
	"runtime/pprof"
	"os/signal"
	"syscall"
	"time"

	"log"
	"net/http"
	"github.com/KingBean4903/GoProfiling/router"
	"github.com/KingBean4903/GoProfiling/profiling"
	"github.com/KingBean4903/GoProfiling/handlers"
)

func main() {
	
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	
	defer stop()

	handlers.StartPostWithWorkers(ctx)

	go profiling.Start()


	srv := &http.Server{
			Addr: ":8070",
			Handler: router.Router(),
	}

	go func() {
			log.Println("HTTP server running on :8070")
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
						log.Fatalf("Listen error: %s", err)
			}
	}()

	handlers.MarkReady()

	<-ctx.Done()

	handlers.MarkNotReady()

	log.Println("Shutting down gracefully....")
	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Server shutdown Failed: %+v", err)
	}

	log.Println("Server exited properly")

 //Capture CPU profile before shutdown
 cpuProfile, err := os.Create("cpu_shutdown.pprof")
 if err != nil {
			log.Printf("Could not create CPU profile: %v", err)
 } else {
		
	 pprof.StartCPUProfile(cpuProfile)
	 defer func() {
				pprof.StopCPUProfile()
				cpuProfile.Close()
				log.Println("CPU profile captured to cpu_shutdown.pprof")
	 }()
 }

 memProfile, err := os.Create("mem_shutdown.pprof")
 if err == nil {
		runtime.GC()
		pprof.WriteHeapProfile(memProfile)
		memProfile.Close()
		log.Println("Memory profile written to mem_shutdown.pprof")
 }


 goroutineProfile, err := os.Create("goroutines_shutdown.pprof")
 if err == nil {
			pprof.Lookup("goroutine").WriteTo(goroutineProfile, 0)
			goroutineProfile.Close()
			log.Println("Goroutine profile written to gorotines_shutdown.pprof")
 }


 time.Sleep(5 * time.Second)


}
