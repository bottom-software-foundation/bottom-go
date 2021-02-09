package bottom

import (
	"errors"
	"strings"
)

type character struct {
	value byte
	char  string
}

var zeroCharacter = character{value: 0, char: "\u2764\uFE0F"}
var characterValues = [5]character{
	{value: 200, char: "\U0001FAC2"},
	{value: 50, char: "\U0001F496"},
	{value: 10, char: "\U00002728"},
	{value: 5, char: "\U0001F97A"},
	{value: 1, char: "\u002C"},
}

const sectionSeparator = "\U0001F449\U0001F448"

// Encode encodes a string in bottom
func Encode(s string) (out string) {
	for _, sChar := range []byte(s) {
		if sChar == 0 {
			out += zeroCharacter.char
			continue
		}
		for sChar != 0 {
			for _, char := range characterValues {
				if sChar >= char.value {
					sChar -= char.value
					out += char.char
					break
				}
			}
		}
		out += sectionSeparator
	}
	return
}

// Validate validates a bottom string
func Validate(bottom string) bool {
	if !strings.HasSuffix(bottom, sectionSeparator) {
		return false
	}
	bottom = strings.Replace(bottom, sectionSeparator, "", -1)
	for _, inputCharRune := range bottom {
		inputChar := string(inputCharRune)
		if inputChar == zeroCharacter.char {
			continue
		}
		valid := false
		for _, char := range characterValues {
			if char.char == inputChar {
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
	b = b[:len(b)-2]
	o := []byte{}
	for _, outCharBlock := range strings.Split(b, sectionSeparator) {
		var sum byte = 0
		for _, bottomCharRune := range outCharBlock {
			bottomChar := string(bottomCharRune)
			for _, char := range characterValues {
				if char.char == bottomChar {
					sum += char.value
				}
			}
		}
		o = append(o, sum)
	}
	out = string(o)
	return
}
