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
type String string

func (s Single) isToken() {}
func (w Word) isToken() {}
func (s String) isToken() {}

var Seperators = []rune(" \n\r\t")
var Quotes = []rune("'\"")

func (s Single) String() string {
	return string(rune(s))
}

func (s String) String() string {
	return fmt.Sprintf("String(%#v)", string(s))
}

func (w Word) String() string {
	return fmt.Sprintf("Word(%v)", string(w))
}
