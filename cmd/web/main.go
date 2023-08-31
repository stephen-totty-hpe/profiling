package main

import (
	"log"
	"net/http"

	// this import is the key to automatic metrics being added
	_ "net/http/pprof"

	"os"
	"os/signal"
	"syscall"

	profiling "github.com/stephen-totty-hpe/profiling/internal"
)

// see https://pkg.go.dev/runtime/pprof
// ./web
// http://localhost:6060/debug/pprof/
func main() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	runMain := profiling.NewRunMain()
	go func() {
		runMain.Run()
	}()

	go func() {
		sig := <-sigs
		log.Println()
		log.Println(sig)
		done <- true
	}()

	log.Println("awaiting signal")
	<-done
	log.Println("exiting")
}
