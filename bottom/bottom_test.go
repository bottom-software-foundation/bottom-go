package bottom

import (
	"strings"
	"testing"
)

type testCase struct {
	in  string
	out string
}

var testCases = []testCase{
	{"test", "💖💖✨🥺,👉👈💖💖,👉👈💖💖✨🥺👉👈💖💖✨🥺,👉👈"},
	{"Hello World!", "💖✨✨,,👉👈💖💖,👉👈💖💖🥺,,,👉👈💖💖🥺,,,👉👈💖💖✨,👉👈✨✨✨,,👉👈💖✨✨✨🥺,,👉👈💖💖✨,👉👈💖💖✨,,,,👉👈💖💖🥺,,,👉👈💖💖👉👈✨✨✨,,,👉👈"},
	{"がんばれ", "🫂✨✨🥺,,👉👈💖💖✨✨🥺,,,,👉👈💖💖✨✨✨✨👉👈🫂✨✨🥺,,👉👈💖💖✨✨✨👉👈💖💖✨✨✨✨🥺,,👉👈🫂✨✨🥺,,👉👈💖💖✨✨🥺,,,,👉👈💖💖💖✨✨🥺,👉👈" +
		"🫂✨✨🥺,,👉👈💖💖✨✨✨👉👈💖💖✨✨✨✨👉👈",
	},
	{"Te\x00st", "💖✨✨✨,,,,👉👈💖💖,👉👈❤️👉👈💖💖✨🥺👉👈💖💖✨🥺,👉👈"},
}

func TestEncode(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.in, func(t *testing.T) {
			if out := Encode(c.in); out != c.out {
				t.Fatalf("expected %q, got %q", c.out, out)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.in, func(t *testing.T) {
			o, err := Decode(c.out)
			if err != nil {
				t.Fatal(err)
			}

			if o != c.in {
				t.Fatalf("expected %q, got %q", c.in, o)
			}
		})
	}
}

func TestDecodeTo(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.in, func(t *testing.T) {
			var buf strings.Builder

			if err := DecodeTo(&buf, c.out); err != nil {
				t.Fatal(err)
			}

			if buf.String() != c.in {
				t.Fatalf("expected %q, got %q", c.in, buf.String())
			}
		})
	}
}

func TestDecodeFrom(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.in, func(t *testing.T) {
			var dst strings.Builder

			if err := DecodeFrom(&dst, strings.NewReader(c.out)); err != nil {
				t.Fatal(err)
			}

			if dst.String() != c.in {
				t.Fatalf("expected %q, got %q", c.in, dst.String())
			}
		})
	}
}

func TestDecodedLen(t *testing.T) {
	for _, c := range testCases {
		t.Run(c.in, func(t *testing.T) {
			var dst strings.Builder

			// use DecodeTo to skip the length check
			_ = DecodeTo(&dst, c.out)
			out, _ := Decode(c.out)

			if dst.Len() != len(out) {
				t.Errorf("expected len %d, got %d", len(out), dst.Len())
			}
		})
	}
}

func BenchmarkEncode(b *testing.B) {
	testCase := testCases[2]
	b.SetBytes(int64(len(testCase.in)))

	for i := 0; i < b.N; i++ {
		_ = Encode(testCase.in)
	}
}

func BenchmarkEncodeTo(b *testing.B) {
	testCase := testCases[2]
	b.SetBytes(int64(len(testCase.out)))

	var w noopWriter
	for i := 0; i < b.N; i++ {
		_ = EncodeTo(w, testCase.out)
	}
}

func BenchmarkDecode(b *testing.B) {
	testCase := testCases[2]
	b.SetBytes(int64(len(testCase.out)))

	for i := 0; i < b.N; i++ {
		_, _ = Decode(testCase.out)
	}
}

func BenchmarkDecodeFrom(b *testing.B) {
	testCase := testCases[2]
	b.SetBytes(int64(len(testCase.out)))

	var w noopWriter
	for i := 0; i < b.N; i++ {
		_ = DecodeFrom(w, strings.NewReader(testCase.out))
	}
}

type noopWriter struct{}

func (noopWriter) WriteByte(c byte) error { return nil }

func (noopWriter) WriteString(s string) (int, error) { return len(s), nil }
