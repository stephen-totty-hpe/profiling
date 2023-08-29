# An example on how to profile in Go. 


https://go.dev/doc/diagnostics
https://pkg.go.dev/cmd/go#hdr-Testing_flags
https://stackoverflow.com/questions/23048455/how-to-profile-benchmarks-using-the-pprof-tool


//goland
https://blog.jetbrains.com/go/2019/04/03/profiling-go-applications-and-tests/
https://blog.jetbrains.com/go/2023/02/02/profiling-go-code-with-goland/



https://medium.com/@openmohan/profiling-in-golang-3e51c68eb6a8


go help test
go help testflag

go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
