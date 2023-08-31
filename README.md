# How to profile in Go. 

This discusses a few go profiling mechanisms and when and where to use each. Which tool you use depends heavily on if the code is local or on a deployed system. Also, it depends on when you need access to the profiles that are generated.


# Go profiling basics

Go provides some very useful diagnostic [documentation](https://go.dev/doc/diagnostics) for profiling, tracing, and debugging go code. There are many links off of this documentation that describes how to use the [pprof tool](https://go.dev/blog/pprof). 

Here is a great [gophercon](https://www.youtube.com/watch?v=nok0aYiGiYA) video from Dave Cheney on profiling that I highly recommend watching.  Also, here is a [workshop](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html) from Dave Cheney that is also excellent.

Another useful link is from [DataDog](https://github.com/DataDog/go-profiler-notes/blob/main/guide/README.md), just know they point out limitations is the hope of selling something.

Lastly, another good link [here](https://www.infoq.com/articles/debugging-go-programs-pprof-trace/)

# Types of profiles

* CPU Profiler

* Memory Profiler

* Block Profiler

* Mutex Profiler

* Goroutine Profiler

* Trace Profiler

# Go profiling from unit tests

It is sometimes useful to profile using unit tests. The easiest way to do this is to create a benchmark test. The function should have the signature ```func BenchmarkXXXXX(b *testing.B)```. Here are some examples of [usage](https://dave.cheney.net/high-performance-go-workshop/dotgo-paris.html#profiling_benchmarks).

There are several [testing flags](https://pkg.go.dev/cmd/go#hdr-Testing_flags) that can help generate the correct profiles needed to pass to the [pprof tool](https://go.dev/blog/pprof).

You can also run like so:
* ```go test -cpuprofile cpu.pprof```
* ```go test -memprofile mem.pprof```
* ```go test -blockprofile block.pprof```
* ```go test -mutexprofile mutex.pprof```
* ```go test -trace trace.out```
* ```go test -race```

# Profiling in a standalone program

As mentioned in the [pprof documentation](https://pkg.go.dev/runtime/pprof#hdr-Profiling_a_Go_program), you can embed the creating of profiles into your code.  An example can be seen [here](cmd/standalone/main.go).

## Exposed profiles
Here is a list of profiles in pprof
```
	profiles.m = map[string]*Profile{
		"goroutine":    goroutineProfile,
		"threadcreate": threadcreateProfile,
		"heap":         heapProfile,
		"allocs":       allocsProfile,
		"block":        blockProfile,
		"mutex":        mutexProfile,
	}
```

## pprof.StartCPUProfile
```
    import "runtime/pprof"

	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()
```
Then run: ```go tool pprof -http=:8080 cpu.prof```

## pprof.WriteHeapProfile
```
    import "runtime/pprof"

	f, err := os.Create(*memprofile)
	if err != nil {
		log.Fatal("could not create memory profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	runtime.GC()    // get up-to-date statistics
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("could not write memory profile: ", err)
	}
```
Then run: ```go tool pprof -http=:8080 mem.prof```

## pprof.Lookup("block")
```
	runtime.SetBlockProfileRate(100000000) // WARNING: Can cause some CPU overhead
	f, err := os.Create(*blockprofile)
	if err != nil {
		log.Fatal("could not create block profile: ", err)
	}
	defer f.Close() // error handling omitted for example
	defer pprof.Lookup("block").WriteTo(f, 0)
```
Then run: ```go tool pprof -http=:8080 block.prof```

## pprof.Lookup("mutex")
```
    import "runtime/pprof"

	runtime.SetMutexProfileFraction(100)
	f, err := os.Create(*mutexprofile)
	if err != nil {
		log.Fatal("could not create mutex profile: ", err)
	}
	defer pprof.Lookup("mutex").WriteTo(f, 0)
```
Then run: ```go tool pprof -http=:8080 mutex.prof```

## pprof.Lookup("goroutine")
```
	pprof.Lookup("goroutine")
	f, err := os.Create(*goroutineprofile)
	if err != nil {
		log.Fatal("could not create goroutine profile: ", err)
	}
	defer pprof.Lookup("goroutine").WriteTo(f, 0)
```
Then run: ```go tool pprof -http=:8080 goroutine.prof```

## trace.Start
```
    import "runtime/trace"

	traceFile, err := os.Create(*traceprofile)
	if err != nil {
		panic(err)
	}
	defer traceFile.Close()

	if err := trace.Start(traceFile); err != nil {
		panic(err)
	}
	defer trace.Stop()
```
Then run: ```go tool trace trace.prof```

# Profiling with a REST handler

You can also enable a REST handler using [http pprof](https://pkg.go.dev/net/http/pprof) to make calls into a running program. There is an example [here](cmd/web/main.go) 

# Profiling with GoLand

If you have access to GoLand and can simulate the issues locally, 
you can check out [article](https://blog.jetbrains.com/go/2023/02/02/profiling-go-code-with-goland/).  There is also a slightly older [article](https://blog.jetbrains.com/go/2019/04/03/profiling-go-applications-and-tests/) that might provide some more information.