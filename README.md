# Bottom.go

[![Docker Hub: orangutan/bottom-go](https://img.shields.io/badge/Docker%20Hub-orangutan/bottom--go-blue?logo=docker)](https://hub.docker.com/r/orangutan/bottom-go)
[![codecov](https://codecov.io/gh/bottom-software-foundation/bottom-go/branch/main/graph/badge.svg)](https://codecov.io/gh/bottom-software-foundation/bottom-go)
[![Test](https://github.com/bottom-software-foundation/bottom-go/workflows/Test/badge.svg)](https://github.com/bottom-software-foundation/bottom-go/actions?query=workflow%3ATest)
[![Version](https://img.shields.io/github/tag/bottom-software-foundation/bottom-go.svg)](https://github.com/bottom-software-foundation/bottom-go/releases)

## Installation

```
$ go get github.com/bottom-software-foundation/bottom-go/cmd/bottom
```

## Usage

```bash
$ echo "Hello, World!" | bottom
💖✨✨,,👉👈💖💖,👉👈💖💖🥺,,,👉👈💖💖🥺,,,👉👈💖💖✨,👉👈✨✨✨✨,,,,👉👈✨✨✨,,👉👈💖✨✨✨🥺,,👉👈💖💖✨,👉👈💖💖✨,,,,👉👈💖💖🥺,,,👉👈💖💖👉👈✨✨✨,,,👉👈✨👉👈

$ echo "💖✨✨,,👉👈💖💖,👉👈💖💖🥺,,,👉👈💖💖🥺,,,👉👈💖💖✨,👉👈✨✨✨✨,,,,👉👈✨✨✨,,👉👈💖✨✨✨🥺,,👉👈💖💖✨,👉👈💖💖✨,,,,👉👈💖💖🥺,,,👉👈💖💖👉👈✨✨✨,,,👉👈✨👉👈" | bottom -d
Hello, World!

$ echo "Hello, World!" > input.txt
$ bottom input.txt | bottom -d
Hello, World!
```

## Benchmarks

You can see current benchmarks in the [Test workflow](https://github.com/bottom-software-foundation/bottom-go/actions?query=workflow%3ATest) output. Benchmarks for v0.2.0:
```
BenchmarkEncode-2       	 5605858	       230 ns/op	  52.25 MB/s
BenchmarkEncodeTo-2     	  779584	      1567 ns/op	 208.63 MB/s
BenchmarkDecode-2       	  466670	      2585 ns/op	 126.51 MB/s
BenchmarkDecodeFrom-2   	  460870	      2683 ns/op	 121.88 MB/s
```
