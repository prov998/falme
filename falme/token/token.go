package token

type Token struct {
	tokenType string
	value     string
}

func (t Token) TokenType() string { return t.tokenType }
func (t Token) Value() string     { return t.value }
