package bottom_test

import (
	"testing"

	"github.com/nihaals/bottom-go/bottom"
)

var testCases = [][]string{
	{"test", "💖💖✨🥺,👉👈💖💖,👉👈💖💖✨🥺👉👈💖💖✨🥺,👉👈"},
	{"Hello World!", "💖✨✨,,👉👈💖💖,👉👈💖💖🥺,,,👉👈💖💖🥺,,,👉👈💖💖✨,👉👈✨✨✨,,👉👈💖✨✨✨🥺,,👉👈💖💖✨,👉👈💖💖✨,,,,👉👈💖💖🥺,,,👉👈💖💖👉👈✨✨✨,,,👉👈"},
	{"がんばれ", "🫂✨✨🥺,,👉👈💖💖✨✨🥺,,,,👉👈💖💖✨✨✨✨👉👈🫂✨✨🥺,,👉👈💖💖✨✨✨👉👈💖💖✨✨✨✨🥺,,👉👈🫂✨✨🥺,,👉👈💖💖✨✨🥺,,,,👉👈💖💖💖✨✨🥺,👉👈🫂✨✨🥺,,👉👈💖💖✨✨✨👉👈💖💖✨✨✨✨👉👈"},
}

func TestEncode(t *testing.T) {
	for _, c := range testCases {
		if out := bottom.Encode(c[0]); out != c[1] {
			t.Errorf("expected: \"%v\", got \"%v\"", c[1], out)
		}
	}
}

func TestDecode(t *testing.T) {
	for _, c := range testCases {
		out, err := bottom.Decode(c[1])
		if err != nil {
			t.Error(err)
		}

		if out != c[0] {
			t.Errorf("expected: \"%v\", got \"%v\"", c[0], out)
		}
	}
}
