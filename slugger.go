// Package slugger provides a parser extension that generates GitHub-style slugs for headings.
package slugger

import (
	"strings"
	"unicode"

	"github.com/yuin/goldmark/v2/ast"
	"github.com/yuin/goldmark/v2/parser"
)

var _ parser.IDGenerator = (*IDGenerator)(nil)

// IDGenerator generates GitHub-style slugs.
type IDGenerator struct {
}

// Generate generates a GitHub-style slug.
func (g *IDGenerator) Generate(value []byte, _ ast.NodeKind) []byte {
	s := strings.ToLower(string(value))

	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		switch {
		case r == ' ':
			fallthrough
		case r == '-':
			// ASCII space -> '-'
			b.WriteByte('-')

		case isGitHubWordChar(r):
			b.WriteRune(r)

		default:
			// punctuation, symbols, emoji, other whitespace, etc.
			continue
		}
	}

	return []byte(b.String())
}

func isGitHubWordChar(r rune) bool {
	if r == '_' || r == '\u200C' || r == '\u200D' {
		return true
	}

	if unicode.IsLetter(r) {
		return true
	}

	if unicode.IsNumber(r) {
		return true
	}

	if unicode.IsMark(r) {
		return true
	}

	return false
}

type parserExtension struct {
}

// Parser is a parser.Extension that generates GitHub-style slugs.
var Parser = &parserExtension{}

func (e *parserExtension) ParserOptions(_ *parser.Config) []parser.Option {
	return []parser.Option{
		parser.WithIDGenerator(&IDGenerator{}),
	}
}
