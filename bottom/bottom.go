package bottom

import "fmt"

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
				fmt.Println(char)
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
