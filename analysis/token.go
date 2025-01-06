package analysis

import (
	"fmt"
)

type Token interface {
	isToken()
}

type Single rune
const (
	Lambda Single =  '$'
	Dot Single = '.'
	LParen Single = '('
	RParen Single = ')'
	Bar Single = '|'
	Equal Single = '='
)

var Singles = []Single{Lambda, Dot, LParen, RParen, Bar, Equal}

type Word string
type Number string
type String string

func (s Single) isToken() {}
func (w Word) isToken() {}
func (n Number) isToken() {}
func (s String) isToken() {}

var Seperators = []rune(" \n\r\t")
var Quotes = []rune("'\"")
var Digits = []rune("1234567890")

func (s Single) String() string {
	return string(rune(s))
}

func (w Word) String() string {
	return fmt.Sprintf("Word(%v)", string(w))
}

func (n Number) String() string {
	return fmt.Sprintf("Number(%v)", string(n))
}
func (s String) String() string {
	return fmt.Sprintf("String(%#v)", string(s))
}

