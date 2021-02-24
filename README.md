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

You can see current benchmarks in the [Test workflow](https://github.com/bottom-software-foundation/bottom-go/actions?query=workflow%3ATest) output. Benchmarks for v0.2.1:
```
BenchmarkEncode-2       	 6773714	       190 ns/op	  63.27 MB/s
BenchmarkEncodeTo-2     	  971070	      1335 ns/op	 244.88 MB/s
BenchmarkDecode-2       	  579741	      2276 ns/op	 143.70 MB/s
BenchmarkDecodeFrom-2   	  477393	      2331 ns/op	 140.28 MB/s
```
