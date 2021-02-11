package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/nihaals/bottom-go/bottom"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"Usage: %s [flags...] [input]\n\n", filepath.Base(os.Args[0]))
		fmt.Fprintf(flag.CommandLine.Output(),
			"Flags:\n")
		flag.PrintDefaults()
	}
}

func main() {
	var (
		decode bool
		output string = "-"
	)

	flag.BoolVar(&decode, "d", decode, "Decode instead of encode")
	flag.StringVar(&output, "o", output, "Output path, - for stdout")
	flag.Parse()

	var err error

	var src = os.Stdin
	if input := flag.Arg(0); input != "" && input != "-" {
		src, err = os.Open(input)
		if err != nil {
			log.Fatalln("failed to open input:", err)
		}
		defer src.Close()
	}

	var dst = os.Stdout
	if output != "" && output != "-" {
		dst, err = os.Create(output)
		if err != nil {
			log.Fatalln("failed to create output:", err)
		}
		defer dst.Close()
	}

	// Make a new buffered writer of size 1024.
	bufdst := bufio.NewWriterSize(dst, 1024)
	defer bufdst.Flush()

	// Make a new buffered reader also of size 1024.
	bufsrc := bufio.NewReaderSize(src, 1024)

	if decode {
		if err := bottom.DecodeFrom(bufdst, bufsrc); err != nil {
			log.Fatalln("failed to decode:", err)
		}
	}

	if err = bottom.EncodeFrom(bufdst, bufsrc); err != nil {
		log.Fatalln("failed to encode:", err)
	}

	if err != nil {
		log.Fatalln(err)
	}
}
