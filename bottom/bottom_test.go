package bottom_test

import (
	"testing"

	"github.com/nihaals/bottom-go/bottom"
)

func TestEncode(t *testing.T) {
	if bottom.Encode("Test") != "💖✨✨✨,,,,👉👈💖💖,👉👈💖💖✨🥺👉👈💖💖✨🥺,👉👈" {
		t.Error()
	}
}
