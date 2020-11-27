# Ristretto Cache

An experiment designing a cache library in Go using [Ristretto](https://github.com/dgraph-io/ristretto) as the store.

## How to execute?

Execute an example program that uses this cache library.
```sh
$ go run example.go
```

Execute the benchmark.
```sh
$ go test -benchmem -bench .
goos: darwin
goarch: amd64
pkg: github.com/abvarun226/ristretto-cache
BenchmarkSetByTags-8              317388              3917 ns/op            1213 B/op         16 allocs/op
BenchmarkSetWithoutTags-8        1268894              1037 ns/op             528 B/op          5 allocs/op
BenchmarkInvalidate-8             719690              1760 ns/op             507 B/op         12 allocs/op
BenchmarkGet-8                   3980583               270 ns/op              29 B/op          1 allocs/op
PASS
ok      github.com/abvarun226/ristretto-cache   9.720s
```