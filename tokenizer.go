package main

import "strings"

type TokenKind int

const (
	TOKEN_KEYWORD TokenKind = iota
	TOKEN_IDENT
	TOKEN_LPAREN
	TOKEN_RPAREN
	TOKEN_COMMA
	TOKEN_SEMICOLON
	TOKEN_NUMBER
	TOKEN_STRING
)

type Token struct {
	Kind TokenKind
	Value string
}

func tokenize(input string) []Token {
	runes := []rune(input)

	pos := 0
	var tokens []Token

	for pos < len(runes) {
		ch := runes[pos]
		
		if ch == ' ' || ch == '\t' {
			pos++
			continue
		}

		if ch == '(' {
			tokens = append(tokens, Token{Kind: TOKEN_LPAREN, Value: "("})
			pos++
			continue
		}

		if ch == ')' {
			tokens = append(tokens, Token{Kind: TOKEN_RPAREN, Value: ")"})
			pos++
			continue
		}

		if ch == ',' {
			tokens = append(tokens, Token{Kind: TOKEN_COMMA, Value: ","})
			pos++
			continue
		}

		if ch == ';' {
			tokens = append(tokens, Token{Kind: TOKEN_SEMICOLON, Value: ";"})
			pos++
			continue
		}

		if isLetter(ch){
			start := pos
			for pos < len(runes) && isLetter(runes[pos]) {
				pos++
			}
			word := string(runes[start:pos])

			upperWord := strings.ToUpper(word)
			if upperWord == "CREATE" || upperWord == "TABLE" || upperWord == "INT" || upperWord == "TEXT" || upperWord == "INSERT" || upperWord == "INTO" || upperWord == "VALUES" {
				tokens = append(tokens, Token{Kind: TOKEN_KEYWORD, Value: upperWord})
			} else {
				tokens = append(tokens, Token{Kind: TOKEN_IDENT, Value: word})
			}

			continue
		}

		if isDigit(ch) {
			start := pos
			for pos < len(runes) && isDigit(runes[pos]) {
				pos++
			}
			tokens = append(tokens, Token{Kind: TOKEN_NUMBER, Value: string(runes[start:pos])})
			continue
		}

		if ch == '\'' {
			pos++
			start := pos
			for pos < len(runes) && runes[pos] != '\'' {
				pos++
			}
			tokens = append(tokens, Token{Kind: TOKEN_STRING, Value: string(runes[start:pos])})
			pos++
			continue
		}

		pos++
	}

	return tokens
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z' || (ch >= 'A' && ch <= 'Z') || ch == '_')
}

func isDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}
