package bottom

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"
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

// sumByte sums up the given part string. If the sum overflows a byte or an
// invalid input is encountered, then an error wrapped with wrapChunkError is
// returned.
func sumByte(part string) (byte, error) {
	partRune := []rune(part)

	var sum int

	for i := 0; i < len(partRune); i++ {
		switch r := partRune[i]; r {
		case []rune(char0)[0]:
			// Edge case: null-byte emoji must have a valid byte after it.
			// Ensure that we can peak to the next byte for this.
			if i >= len(partRune)-1 {
				// The string stopped short when we're expecting another rune,
				// so an UnexpectedEOF is valid.
				return 0, wrapChunkError(part, io.ErrUnexpectedEOF)
			}
			if []rune(char0)[1] != partRune[i+1] {
				return 0, wrapChunkError(part, InvalidRuneError(r))
			}
			i++ // skip peeked rune
			sum += 0
		case char1:
			sum += 1
		case char5:
			sum += 5
		case char10:
			sum += 10
		case char50:
			sum += 50
		case char200:
			sum += 200
		default:
			return 0, wrapChunkError(part, InvalidRuneError(r))
		}
	}

	if sum > 0xFF {
		return 0, wrapChunkError(part, ErrByteOverflow)
	}

	return byte(sum), nil
}

// wrapChunkError wraps the given error with a "failed to decode chunk" error.
func wrapChunkError(part string, err error) error {
	return fmt.Errorf("failed to decode chunk %q: %w", part, err)
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
	var buf bytes.Buffer

	values := []byte{200, 50, 10, 5, 1}
	valueCharacters := [255]string{
		0: char0 + sectionSeparator,
	}

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

// ErrByteOverflow is returned when the given input string overflows a byte when
// decoded.
var ErrByteOverflow = errors.New("sum overflows byte")

// InvalidRuneError is returned when an invalid rune is encountered when
// decoding.
type InvalidRuneError rune

// Error formats InvalidRuneError to show the quoted rune and the Unicode
// codepoint notation.
func (r InvalidRuneError) Error() string {
	return fmt.Sprintf("unexpected rune %q (%U)", rune(r), rune(r))
}

// Encode encodes a string in bottom
func Encode(s string) string {
	builder := strings.Builder{}
	builder.Grow(EncodedLen(s))

	for _, sChar := range []byte(s) {
		builder.WriteString(valueCharacters[sChar])
	}

	return builder.String()
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
func EncodeTo(out io.StringWriter, s string) error {
	for _, sChar := range []byte(s) {
		_, err := out.WriteString(valueCharacters[sChar])
		if err != nil {
			return err
		}
	}

	return nil
}

// EncodeFrom encodes from the given src reader to out.
func EncodeFrom(out io.StringWriter, src io.ByteReader) error {
	for {
		b, err := src.ReadByte()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		if _, err = out.WriteString(valueCharacters[b]); err != nil {
			return err
		}
	}
}

// Validate validates a bottom string. False is returned if the validation
// fails.
func Validate(bottom string) bool {
	return DecodedLen(bottom) > -1
}

// DecodedLen validates the given bottom string and returns the calculated
// length. It returns -1 if the given bottom string is invalid.
func DecodedLen(bottom string) int {
	l, _ := decodedLen(bottom, true)
	return l
}

// decodedLen is the implementation of DecodedLen that returns an error if the
// input bottom string is invalid.
func decodedLen(bottom string, verify bool) (int, error) {
	if !strings.HasSuffix(bottom, sectionSeparator) {
		return -1, errors.New("missing trailing separator")
	}

	// We used to trim the sectionSeparator suffix here, but since our current
	// method of Index seeking does not account for the last section, we don't
	// need to trim it.
	//
	// This assumption is validated by the above HasSuffix check.

	var length int

	for {
		m := strings.Index(bottom, sectionSeparator)
		if m < 0 {
			break
		}

		if verify {
			_, err := sumByte(string(bottom[:m]))
			if err != nil {
				return -1, err
			}
		}

		length++
		bottom = bottom[m+len(sectionSeparator):]
	}

	return length, nil
}

// Decode verifies and decodes a bottom string. An error is returned if the
// verification fails.
func Decode(bottom string) (string, error) {
	// Skip verification, since we're doing it in the loop.
	l, err := decodedLen(bottom, false)
	if err != nil {
		return "", err
	}

	builder := strings.Builder{}
	builder.Grow(l)

	var i int
	for i < l {
		m := strings.Index(bottom, sectionSeparator)
		if m < 0 {
			break
		}

		sum, err := sumByte(bottom[:m])
		if err != nil {
			return "", err
		}

		builder.WriteByte(sum)
		bottom = bottom[m+len(sectionSeparator):]
		i++
	}

	return builder.String(), nil
}

// DecodeTo decodes the given bottom string into the given byte writer.
func DecodeTo(w io.ByteWriter, bottom string) error {
	for {
		m := strings.Index(bottom, sectionSeparator)
		if m < 0 {
			break
		}

		sum, err := sumByte(bottom[:m])
		if err != nil {
			return err
		}

		if err := w.WriteByte(sum); err != nil {
			return err
		}

		bottom = bottom[m+len(sectionSeparator):]
	}

	return nil
}

// DecodeFrom decodes from a src reader.
func DecodeFrom(w io.ByteWriter, src io.Reader) error {
	scanner := bufio.NewScanner(src)
	scanner.Split(scanUntilSeparator)

	for scanner.Scan() {
		sum, err := sumByte(scanner.Text())
		if err != nil {
			return err
		}

		if err := w.WriteByte(byte(sum)); err != nil {
			return err
		}
	}

	return scanner.Err()
}

// scanUntilSeparator is used with bufio.Scanner to scan chunks separated by
// sectionSeparator.
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
