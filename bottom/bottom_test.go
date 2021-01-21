package bottom_test

import (
	"testing"

	"github.com/nihaals/bottom-go/bottom"
)

// "Test" encoded as bottom
const testEncoded =
"\U0001F496" +
"\u2728" +
"\u2728" +
"\u2728" +
"\u002C" +
"\u002C" +
"\u002C" +
"\u002C" +
"\U0001F449" +
"\U0001F448" +
"\U0001F496" +
"\U0001F496" +
"\u002C" +
"\U0001F449" +
"\U0001F448" +
"\U0001F496" +
"\U0001F496" +
"\u2728" +
"\U0001F97A" +
"\U0001F449" +
"\U0001F448" +
"\U0001F496" +
"\U0001F496" +
"\u2728" +
"\U0001F97A" +
"\u002C" +
"\U0001F449" +
"\U0001F448"

func TestEncode(t *testing.T) {
	if bottom.Encode("Test") != testEncoded {
		t.Error()
	}
}

func TestDecode(t *testing.T) {
	out, err := bottom.Decode(testEncoded)
	if err != nil {
		t.Error()
	}
	if out != "Test" {
		t.Error()
	}
}
