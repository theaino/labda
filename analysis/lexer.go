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
	currentNumber := ""
	handleNumberEnd := func() {
		if currentNumber != "" {
			tokens = append(tokens, Number(currentNumber))
			currentNumber = ""
		}
	}
	BodyLoop: for idx := 0; idx < len(body); idx++ {
		char := []rune(body)[idx]
		if slices.Contains(Seperators, char) {
			handleWordEnd()
			handleNumberEnd()
			continue BodyLoop
		}
		if slices.Contains(Quotes, char) {
			handleWordEnd()
			handleNumberEnd()
			str, offset := LexString(body[idx:])
			idx += offset
			tokens = append(tokens, String(str))
			continue BodyLoop
		}
		if slices.Contains(Digits, char) && currentWord == "" {
			currentNumber += string(char)
			continue BodyLoop
		}
		for _, single := range Singles {
			if char == rune(single) {
				handleWordEnd()
				handleNumberEnd()
				tokens = append(tokens, single)
				continue BodyLoop
			}
		}
		currentWord += string(char)
		handleNumberEnd()
	}
	handleWordEnd()
	handleNumberEnd()
	return tokens
}
