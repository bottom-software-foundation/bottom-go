package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/nihaals/bottom-go/bottom"
)

func main() {
	decode := flag.Bool("d", false, "Whether to decode")
	flag.Parse()
	input := flag.Arg(0)
	if input == "" {
		fmt.Println("No input given")
		os.Exit(1)
	}
	if *decode {
		out, err := bottom.Decode(input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(out)
	} else {
		fmt.Println(bottom.Encode(input))
	}
}
