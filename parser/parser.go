package parser

import (
	"fmt"
	"strings"
)

type Token interface {
	Value() string
}

type StringToken struct {
	value string
}

type SpaceToken struct {
}

func (s SpaceToken) Value() string {
	return " "
}

func (s StringToken) Value() string {
	return s.value
}

type SingleQuoteToken struct {
	value string
}

func (s SingleQuoteToken) Value() string {
	return s.value
}

type Parser struct {
	tokens []Token
}

func NewParser() *Parser {
	return &Parser{tokens: []Token{}}
}

func (p *Parser) Parse(input string) {
	if input == "" {
		return
	}
	index := 0
	for index < len(input) {
		switch input[index] {
		case '\'':
			token, nextIndex := p.parseSingleQuoteToken(input[index+1:])
			p.tokens = append(p.tokens, token)
			index += nextIndex
		case ' ':
			p.tokens = append(p.tokens, SpaceToken{})
			index++
			for index < len(input) && input[index] == ' ' {
				index++
			}
		default:
			word := ""
			for index < len(input) && input[index] != ' ' && input[index] != '\'' {
				word += string(input[index])
				index++
			}
			p.tokens = append(p.tokens, StringToken{value: word})
		}
	}
}

func (p *Parser) parseSingleQuoteToken(s string) (Token, int) {
	index := strings.Index(s, "'")
	if index == -1 {
		spaceIndex := strings.Index(s, " ")
		if spaceIndex == -1 {
			return StringToken{value: s}, len(s)
		}
		return StringToken{value: s[:spaceIndex]}, spaceIndex + 1
	}
	return SingleQuoteToken{value: s[:index]}, index + 2
}

func (p *Parser) EchoTokens() {
	for _, token := range p.tokens {
		fmt.Print(token.Value())
	}
	fmt.Print("\n")
}
