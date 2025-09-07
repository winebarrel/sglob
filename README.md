# sglob

[![CI](https://github.com/winebarrel/sglob/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/sglob/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/winebarrel/sglob.svg)](https://pkg.go.dev/github.com/winebarrel/sglob)
[![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/winebarrel/sglob)](https://pkg.go.dev/github.com/winebarrel/sglob?tab=versions)
[![Go Report Card](https://goreportcard.com/badge/github.com/winebarrel/sglob)](https://goreportcard.com/report/github.com/winebarrel/sglob)

sglob is a small globbing library for string.

## Usage

```go
package main

import "github.com/winebarrel/sglob"

func main() {
	sglob.Match("*@example.com", "scott@example.com")   //=> true
	sglob.Match("scott@example.j?", "scott@example.jp") //=> true
	sglob.Match("scott@example.j?", "scott@example.jo") //=> true
}
```

* `?`: any single character
* `*`: any number of characters

## Benchmark

```
goos: darwin
goarch: arm64
pkg: github.com/winebarrel/sglob
cpu: Apple M4 Pro
BenchmarkGlob/'*@example.com'_match_'scott@example.com'-14         	20158608	        54.63 ns/op	       0 B/op	       0 allocs/op
BenchmarkGlob/'scott@exa?ple.com'_match_'scott@example.com'-14     	19859356	        60.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkGlob/'*@*.com'_match_'scott@example.com'-14               	26311317	        41.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkGlob/'scott@?*.com'_match_'scott@example.com'-14          	23470753	        52.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkRegexp/'.*@example\.com'_match_'scott@example.com'-14     	 6336601	       189.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkRegexp/'scott@exa.ple\.com'_match_'scott@example.com'-14  	13965226	        86.57 ns/op	       0 B/op	       0 allocs/op
BenchmarkRegexp/'.*@.*\.com'_match_'scott@example.com'-14          	 3012747	       397.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkRegexp/'scott@.+\.com'_match_'scott@example.com'-14       	10102154	       119.2 ns/op	       0 B/op	       0 allocs/op
```
