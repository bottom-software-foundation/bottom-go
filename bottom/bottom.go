package bottom

import (
	"errors"
	"strings"
)

type character struct {
	value byte
	chars []string
}

var zeroCharacter = character{value: 0, chars: []string{"\u2764", "\uFE0F"}}
var characterValues = [5]character{
	{value: 200, chars: []string{"\U0001FAC2"}},
	{value: 50, chars: []string{"\U0001F496"}},
	{value: 10, chars: []string{"\U00002728"}},
	{value: 5, chars: []string{"\U0001F97A"}},
	{value: 1, chars: []string{"\u002C"}},
}

const sectionSeparator = "\U0001F449\U0001F448"

// Encode encodes a string in bottom
func Encode(s string) (out string) {
	for _, sChar := range []byte(s) {
		if sChar == 0 {
			out += zeroCharacter.chars[0]
			continue
		}
		for sChar != 0 {
			for _, char := range characterValues {
				if sChar >= char.value {
					sChar -= char.value
					out += char.chars[0]
					break
				}
			}
		}
		out += sectionSeparator
	}
	return
}

func validateChar(c string, char character) bool {
	for _, charChar := range char.chars {
		if c == charChar {
			return true
		}
	}
	return false
}

// Validate validates a bottom string
func Validate(bottom string) bool {
	if !strings.HasSuffix(bottom, sectionSeparator) {
		return false
	}
	bottom = strings.Replace(bottom, sectionSeparator, "", -1)
	for _, inputCharRune := range bottom {
		inputChar := string(inputCharRune)
		if validateChar(inputChar, zeroCharacter) {
			continue
		}
		valid := false
		for _, char := range characterValues {
			if validateChar(inputChar, char) {
				valid = true
				break
			}
		}
		if !valid {
			return false
		}
	}
	return true
}

// Decode decodes a bottom string
func Decode(b string) (out string, err error) {
	if !Validate(b) {
		return "", errors.New("Invalid bottom text")
	}
	b = b[:len(b) - 2]
	for _, outCharBlock := range strings.Split(b, sectionSeparator) {
		var sum byte = 0
		for _, bottomCharRune := range outCharBlock {
			bottomChar := string(bottomCharRune)
			for _, char := range characterValues {
				for _, charChar := range char.chars {
					if charChar == bottomChar {
						sum += char.value
					}
				}
			}
		}
		out += string(sum)
	}
	return
}
