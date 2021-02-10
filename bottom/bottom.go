package bottom

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
	"unicode/utf8"

	"github.com/nihaals/bottom-go/bottom/internal/unsafebytes"
)

const (
	char0   = "\u2764\uFE0F"
	char1   = '\u002C'
	char5   = '\U0001F97A'
	char10  = '\U00002728'
	char50  = '\U0001F496'
	char200 = '\U0001FAC2'

	sectionSeparator = "\U0001F449\U0001F448"
)

// characterValues looks up known runes; it returns 0 for unknown runes.
func characterValues(s rune) byte {
	switch s {
	case char1:
		return 1
	case char5:
		return 5
	case char10:
		return 10
	case char50:
		return 50
	case char200:
		return 200
	default:
		return 0 // unknown
	}
}

// valueCharacterBases looks up known values for the corresponding runes; it
// gives 0 for unknown values.
var valueCharacterBases = [255]string{
	1:   string(char1),
	5:   string(char5),
	10:  string(char10),
	50:  string(char50),
	200: string(char200),
}

// valueCharacters looks up all possible bytes for the corresponding bottom
// runes.
var valueCharacters = calculateValueCharacters()

func calculateValueCharacters() [255]string {
	var values = []byte{200, 50, 10, 5, 1}
	var valueCharacters = [255]string{0: char0}
	var buf bytes.Buffer

	for i := byte(1); i < 255; i++ {
		char := i
		for char > 0 {
			for _, v := range values {
				if char >= v {
					char -= v
					buf.WriteString(valueCharacterBases[v])
					break
				}
			}
		}

		buf.WriteString(sectionSeparator)
		valueCharacters[i] = buf.String()
		buf.Reset()
	}

	return valueCharacters
}

// Encode encodes a string in bottom
func Encode(s string) string {
	out := make([]byte, EncodedLen(s))
	sum := 0

	for _, sChar := range []byte(s) {
		sum += copy(out[sum:], valueCharacters[sChar])
	}

	return unsafebytes.String(out)
}

// EncodedLen returns the length of the encoded string in exact.
func EncodedLen(s string) int {
	var l int
	for _, sChar := range []byte(s) {
		l += len(valueCharacters[sChar])
	}

	return l
}

// EncodeTo encodes the given string into the writer.
func EncodeTo(out io.StringWriter, s string) (int, error) {
	var sum int

	for _, sChar := range []byte(s) {
		n, err := out.WriteString(valueCharacters[sChar])
		if err != nil {
			return sum, err
		}

		sum += n
	}

	return sum, nil
}

// Validate validates a bottom string.
func Validate(bottom string) bool {
	return DecodedLen(bottom) > -1
}

// DecodedLen validates the given bottom string and returns the calculated
// length. It returns -1 if the given bottom string is invalid.
func DecodedLen(bottom string) int {
	if !strings.HasSuffix(bottom, sectionSeparator) {
		return -1
	}

	// We used to trim the sectionSeparator suffix here, but since our current
	// method of Index seeking does not account for the last section, we don't
	// need to trim it.
	//
	// This assumption is validated by the above HasSuffix check.

	var length, sum int

	for {
		m := strings.Index(bottom, sectionSeparator)
		if m < 0 {
			break
		}

		sum = 0

		for _, r := range bottom[:m] {
			v := characterValues(r)
			if v == 0 {
				return -1
			}

			// overflow check
			if sum += int(v); sum > 0xFF {
				return -1
			}
		}

		length++
		bottom = bottom[m+len(sectionSeparator):]
	}

	return length
}

// Decode verifies and decodes a bottom string. An error is returned if the
// verification fails.
func Decode(bottom string) (string, error) {
	l := DecodedLen(bottom)
	if l == -1 {
		return "", errors.New("invalid bottom text")
	}

	buf := make([]byte, l)

	var i int
	for i < l {
		m := strings.Index(bottom, sectionSeparator)
		if m < 0 {
			break
		}

		buf[i] = sumByte(bottom[:m])
		bottom = bottom[m+len(sectionSeparator):]
		i++
	}

	return unsafebytes.String(buf), nil
}

// DecodeTo decodes the given bottom string into the given byte writer.
func DecodeTo(w io.ByteWriter, bottom string) error {
	for {
		m := strings.Index(bottom, sectionSeparator)
		if m < 0 {
			break
		}

		if err := w.WriteByte(sumByte(bottom[:m])); err != nil {
			return err
		}

		bottom = bottom[m+len(sectionSeparator):]
	}

	return nil
}

func sumByte(part string) (sum byte) {
	for _, r := range part {
		sum += characterValues(r)
	}
	return
}

// DecodeFrom decodes from a src reader.
func DecodeFrom(w io.ByteWriter, src io.Reader) error {
	scanner := bufio.NewScanner(src)
	scanner.Split(scanUntilSeparator)

	var sum byte
	for scanner.Scan() {
		sum = 0
		bytes := scanner.Bytes()

		for len(bytes) > 0 {
			r, sz := utf8.DecodeRune(bytes)
			if sz == -1 {
				return fmt.Errorf("invalid bytes %q", bytes)
			}

			sum += characterValues(r)
			bytes = bytes[sz:]
		}

		if err := w.WriteByte(sum); err != nil {
			return err
		}
	}

	return scanner.Err()
}

func scanUntilSeparator(data []byte, eof bool) (int, []byte, error) {
	if eof && len(data) == 0 {
		return 0, nil, nil
	}

	if i := bytes.Index(data, []byte(sectionSeparator)); i >= 0 {
		return i + len(sectionSeparator), data[:i], nil
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if eof {
		return len(data), data, nil
	}

	// Request more data.
	return 0, nil, nil
}
