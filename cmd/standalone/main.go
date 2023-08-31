package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"runtime/trace"
	"syscall"

	profiling "github.com/stephen-totty-hpe/profiling/internal"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")
var blockprofile = flag.String("blockprofile", "", "write block profile to `file`")
var mutexprofile = flag.String("mutexprofile", "", "write mutex profile to `file`")
var goroutineprofile = flag.String("goroutineprofile", "", "write goroutine profile to `file`")
var traceprofile = flag.String("traceprofile", "", "write trace profile to `file`")

// see https://pkg.go.dev/runtime/pprof
// ./standalone -cpuprofile cpuprofile.out -memprofile memprofile.out -blockprofile blockprofile.out -mutexprofile mutexprofile.out -goroutineprofile goroutineprofile.out -traceprofile traceprofile.out
// go tool pprof -http=:8080 cpuprofile.out
// go tool pprof -http=:8080 memprofile.out
// go tool pprof -http=:8080 blockprofile.out
// go tool pprof -http=:8080 mutexprofile.out
// go tool pprof -http=:8080 goroutineprofile.out
// go tool trace traceprofile.out
func main() {

	// put this block at the beginning
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if *blockprofile != "" {
		runtime.SetBlockProfileRate(100000000) // WARNING: Can cause some CPU overhead
		f, err := os.Create(*blockprofile)
		if err != nil {
			log.Fatal("could not create block profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		defer pprof.Lookup("block").WriteTo(f, 0)
	}

	if *mutexprofile != "" {
		runtime.SetMutexProfileFraction(100)
		f, err := os.Create(*mutexprofile)
		if err != nil {
			log.Fatal("could not create mutex profile: ", err)
		}
		defer pprof.Lookup("mutex").WriteTo(f, 0)
	}

	if *goroutineprofile != "" {
		pprof.Lookup("goroutine")
		f, err := os.Create(*goroutineprofile)
		if err != nil {
			log.Fatal("could not create goroutine profile: ", err)
		}
		defer pprof.Lookup("goroutine").WriteTo(f, 0)
	}

	if *traceprofile != "" {
		// Start tracing
		traceFile, err := os.Create(*traceprofile)
		if err != nil {
			panic(err)
		}
		defer traceFile.Close()

		if err := trace.Start(traceFile); err != nil {
			panic(err)
		}
		defer trace.Stop()
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)

	runMain := profiling.NewRunMain()
	go func() {
		runMain.Run()
	}()

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")

	// put this block at the end
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}
