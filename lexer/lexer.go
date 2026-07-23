package lexer

type TokenType int

const (
	TokenError TokenType = iota
	TokenEOF
	TokenEOFLine
	TokenLabel
	TokenInstruction
	TokenOperateCommand
	TokenIdentifier
	TokenOctal
	TokenDecimal
	TokenComma
	TokenColon
	TokenIndirectFlag
)
