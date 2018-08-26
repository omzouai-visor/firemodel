package ast

import (
	"io"
	"github.com/alecthomas/participle"
	"github.com/pkg/errors"
	"github.com/alecthomas/participle/lexer"
	"text/scanner"
	"regexp"
	"sort"
	"strings"
)

func ParseSchema(r io.Reader) (*AST, error) {
	parser := participle.MustBuild(
		&AST{},
		participle.Lexer(&lexerDefinition{}),
		participle.Map(func(token lexer.Token) lexer.Token {
			if token.Type == scanner.Comment {
				p := regexp.MustCompile(`//\s*(?P<Comment>.*)`)
				token.Value = p.FindStringSubmatch(token.Value)[1]
			}
			return token
		}),
	)

	s := &AST{}
	if err := parser.Parse(r, s); err != nil {
		return nil, errors.Wrap(err, "firemodel:")
	}
	return s, nil
}

type lexerDefinition struct{}

func (d *lexerDefinition) Lex(r io.Reader) lexer.Lexer {
	s := &scanner.Scanner{}
	l := lexer.LexWithScanner(r, s)
	s.Mode = s.Mode &^ scanner.SkipComments
	return l
}

func (d *lexerDefinition) Symbols() map[string]rune {
	return map[string]rune{
		"EOF":     scanner.EOF,
		"Ident":   scanner.Ident,
		"String":  scanner.String,
		"Int":     scanner.Int,
		"Comment": scanner.Comment,
	}
}

// Read about the magical annotations here: https://github.com/alecthomas/participle/.
type AST struct {
	Types []*ASTElement `parser:"{ @@ }"`
}

type ASTElement struct {
	Comment string     `parser:"{ @Comment }"`
	Model   *ASTModel  `parser:"  'model' @@"`
	Enum    *ASTEnum   `parser:"| 'enum' @@"`
	Option  *ASTOption `parser:"| 'option' @@"`
}

type ASTModel struct {
	Identifier ASTIdentifier      `parser:"@Ident"`
	Elements   []*ASTModelElement `parser:"'{' { @@ } '}'"`
}

type ASTIdentifier string

var (
	reservedIdentifiers = []string{
		// Primitive types.
		"boolean", "integer", "double", "timestamp", "string",
		"bytes", "reference", "geopoint", "array", "map", "url",
		"file", "collection", "file",
		// Keywords.
		"model", "option", "enum",
	}
)

func init() {
	sort.Strings(reservedIdentifiers)
}

func (id ASTIdentifier) IsReserved() bool {
	needle := strings.ToLower(string(id))
	idx := sort.SearchStrings(reservedIdentifiers, needle)
	return reservedIdentifiers[idx] == needle
}

type ASTModelElement struct {
	Option *ASTOption `parser:"  'option' @@"`
	Field  *ASTField  `parser:"| @@"`
}

type ASTEnum struct {
	Identifier ASTIdentifier   `parser:"@Ident '{'"`
	Values     []*ASTEnumValue `parser:"{ @@ } '}'"`
}

type ASTOption struct {
	Language string        `parser:"@Ident '.'"`
	Key      ASTIdentifier `parser:"@Ident '='"`
	Value    string        `parser:"@('true' | 'false' | 'null' | String | Int) ';'"`
}

type ASTEnumValue struct {
	Comment string `parser:"{ @Comment }"`
	Name    string `parser:"@Ident ','"`
}

type ASTField struct {
	Comment string        `parser:"{ @Comment }"`
	Type    *ASTFieldType `parser:"@@"`
	Name    string        `parser:"@Ident ';'"`
}

type ASTFieldType struct {
	Base    ASTType `parser:"@Ident"`
	Generic ASTType `parser:"[ '<' @Ident '>' ]"`
}

type ASTType string

const (
	Boolean   ASTType = "boolean"
	Integer   ASTType = "integer"
	Double    ASTType = "double"
	Timestamp ASTType = "timestamp"
	String    ASTType = "string"
	Bytes     ASTType = "bytes"
	Reference ASTType = "reference"
	GeoPoint  ASTType = "geopoint"
	Array     ASTType = "array"
	Map       ASTType = "map"
	// Fake types.
	URL  ASTType = "URL"
	File ASTType = "File"
	// Non-types.
	collection = "collection"
)

func (s ASTType) IsCollection() bool {
	return s == collection
}

func (s ASTType) IsPrimitive() bool {
	switch s {
	case String,
		Integer,
		Bytes,
		Double,
		Timestamp,
		Boolean,
		Reference,
		GeoPoint,
		Array,
		Map:
		return true
	case collection:
		panic("firemodel/schema: bug. collection should never be treated as primitive type.")
	default:
		return false
	}
}
