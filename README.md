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

## pprof.StartCPUProfile
```
    import "runtime/pprof"

    // Create a CPU profile file
    f, err := os.Create("cpu.prof")
    if err != nil {
        panic(err)
    }
    defer f.Close()

    // Start CPU profiling
    if err := pprof.StartCPUProfile(f); err != nil {
        panic(err)
    }
    defer pprof.StopCPUProfile()
```
Then run: ```go tool pprof -http=:8080 cpu.prof```

## pprof.WriteHeapProfile
```
    import "runtime/pprof"

    // Create a memory profile file
    memProfileFile, err := os.Create("mem.prof")
    if err != nil {
        panic(err)
    }
    defer memProfileFile.Close()

    // Write memory profile to file
    if err := pprof.WriteHeapProfile(memProfileFile); err != nil {
        panic(err)
    }
```
Then run: ```go tool pprof -http=:8080 mem.prof```

## runtime.SetBlockProfileRate
```
runtime.SetBlockProfileRate(100_000_000) // WARNING: Can cause some CPU overhead
file, _ := os.Create("./block.prof")
defer pprof.Lookup("block").WriteTo(file, 0)
```
Then run: ```go tool pprof -http=:8080 block.prof```

## runtime.SetMutexProfileRate(rate)
```
runtime.SetMutexProfileFraction(100)
file, _ := os.Create("./mutex.prof")
defer pprof.Lookup("mutex").WriteTo(file, 0)
```
Then run: ```go tool pprof -http=:8080 mutex.prof```

## runtime.GoroutineProfile()

## profile := pprof.Lookup("goroutine")
```
Each Profile has a unique name. A few profiles are predefined:

goroutine - stack traces of all current goroutines
heap - a sampling of all heap allocations
threadcreate - stack traces that led to the creation of new OS threads
block - stack traces that led to blocking on synchronization primitives

file, _ := os.Create("./goroutine.prof")
pprof.Lookup("goroutine").(file, 0)
```

## trace.Start
```
    import "runtime/trace"

    // Start tracing
    traceFile, err := os.Create("trace.prof")
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

# 

https://www.youtube.com/watch?v=xxDZuPEgbBU

https://github.com/bradfitz/talk-yapc-asia-2015/blob/master/talk.md

https://www.youtube.com/watch?v=HEwSkhr_8_M

https://github.com/DataDog/go-profiler-notes/blob/main/guide/README.md