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
			if upperWord == "CREATE" || upperWord == "TABLE" || upperWord == "INT" || upperWord == "TEXT" {
				tokens = append(tokens, Token{Kind: TOKEN_KEYWORD, Value: upperWord})
			} else {
				tokens = append(tokens, Token{Kind: TOKEN_IDENT, Value: word})
			}

			continue
		}

		pos++
	}

	return tokens
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z' || (ch >= 'A' && ch <= 'Z') || ch == '_')
}
