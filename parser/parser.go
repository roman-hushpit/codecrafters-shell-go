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
		case '\\':
			index++
			current += string(input[index])
			index++
		case ' ':
			if current != "" {
				p.args = append(p.args, current)
				current = ""
			}
			for index < len(input) && input[index] == ' ' {
				index++
			}
		case '1', '2':
			if index+1 < len(input) && input[index+1] == '>' {
				if index+2 < len(input) && input[index+2] == '>' {
					current = input[index : index+3]
					p.args = append(p.args, current)
					index += 3
					current = ""
				} else {
					current = input[index : index+2]
					p.args = append(p.args, current)
					index += 2
					current = ""
				}
			} else {
				current += string(input[index])
				index++
			}
		case '>':
			if current != "" {
				p.args = append(p.args, current)
				current = ""
			}
			if index+1 < len(input) && input[index+1] == '>' {
				p.args = append(p.args, ">>")
				index += 2
			} else {
				p.args = append(p.args, ">")
				index++
			}
		default:
		loop:
			for index < len(input) {
				switch input[index] {
				case ' ', '"', '\'', '\\', '>':
					break loop
				default:
					current += string(input[index])
					index++
				}
			}
		}
	}
	if current != "" {
		p.args = append(p.args, current)
	}
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
	sequence := strings.Builder{}
	i := 0
loop:
	for i < len(s) {
		symbol := s[i]
		switch symbol {
		case '\\':
			sequence.WriteString(string(s[i+1]))
			i += 2
		case '"':
			break loop
		default:
			sequence.WriteByte(s[i])
			i++
		}
	}
	return sequence.String(), i + 2
}
