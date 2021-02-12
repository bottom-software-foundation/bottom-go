# Bottom.go

## Installation

`$ go get github.com/bottom-software-foundation/bottom-go/cmd/bottom`

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
