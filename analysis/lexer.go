package analysis

import (
	"slices"
)

func Lex(body string) []Token {
	tokens := make([]Token, 0)
	currentWord := ""
	handleWordEnd := func() {
		if currentWord != "" {
			tokens = append(tokens, Word(currentWord))
			currentWord = ""
		}
	}
	BodyLoop: for idx := 0; idx < len(body); idx++ {
		char := []rune(body)[idx]
		if slices.Contains(Seperators, char) {
			handleWordEnd()
			continue BodyLoop
		}
		if slices.Contains(Quotes, char) {
			handleWordEnd()
			str, offset := LexString(body[idx:])
			idx += offset
			tokens = append(tokens, String(str))
			continue BodyLoop
		}
		for _, single := range Singles {
			if char == rune(single) {
				handleWordEnd()
				tokens = append(tokens, single)
				continue BodyLoop
			}
		}
		currentWord += string(char)
	}
	handleWordEnd()
	return tokens
}
