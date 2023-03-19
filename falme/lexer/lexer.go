package lexer

import (
	"falme/token"
)

type Lexer struct {
	ch    string
	index int
	input string
}

func Lex() []token.Token {

}

func new(input string) Lexer {
	l := Lexer{
		" ",
		0,
		input,
	}

	return l
}
