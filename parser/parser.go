package parser

import (
	"strings"
)

type Parser struct {
	args []string
}

func NewParser() *Parser {
	return &Parser{args: []string{}}
}

func (p *Parser) Args() []string {
	return p.args
}

func (p *Parser) Parse(input string) {
	if input == "" {
		return
	}
	index := 0
	current := ""
	for index < len(input) {
		switch input[index] {
		case '\'':
			token, nextIndex := p.parseSingleQuoteArg(input[index+1:])
			current += token
			index += nextIndex
		case '"':
			token, nextIndex := p.parseDoubleQuoteArg(input[index+1:])
			current += token
			index += nextIndex
		case ' ':
			p.args = append(p.args, current)
			current = ""
			for index < len(input) && input[index] == ' ' {
				index++
			}
		default:
			for index < len(input) && input[index] != ' ' && input[index] != '\'' {
				current += string(input[index])
				index++
			}
		}
	}
	p.args = append(p.args, current)
}

func (p *Parser) parseSingleQuoteArg(s string) (string, int) {
	index := strings.Index(s, "'")
	if index == -1 {
		spaceIndex := strings.Index(s, " ")
		if spaceIndex == -1 {
			return s, len(s)
		}
		return s[:spaceIndex], spaceIndex + 1
	}
	return s[:index], index + 2
}

func (p *Parser) parseDoubleQuoteArg(s string) (string, int) {
	index := strings.Index(s, "\"")
	if index == -1 {
		spaceIndex := strings.Index(s, " ")
		if spaceIndex == -1 {
			return s, len(s)
		}
		return s[:spaceIndex], spaceIndex + 1
	}
	return s[:index], index + 2
}

func (p *Parser) FormatArgs() string {
	return strings.Join(p.args, " ")
}
