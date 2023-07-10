package main

import (
	"fmt"
	_ "time"
	//"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/antlr4-go/antlr/v4"
	"github.com/JesseCoretta/parser"
)

const (
	definition = `key = "value" AND ( someother = "value" OR something = "value2" )`
)

type listener struct {
	*aciparser.BaseACIExprListener
}

func (r *listener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("%v\n", node.GetText())
}

func Parse(program string) {
        is := antlr.NewInputStream(definition)
        lexer := aciparser.NewACIExprLexer(is)
        tstream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
        parser := aciparser.NewACIExprParser(tstream)
	exprtree := parser.Expression()
        antlr.NewParseTreeWalker().Walk(&listener{}, exprtree)

	fmt.Printf("%v\n", exprtree.GetChildren())
}

func main() {
	Parse(definition)

	/*

	fmt.Printf("RULENAMES: %T\n", lexer.GetATN())

	for {
		t := lexer.NextToken()
		if t.GetTokenType() == antlr.TokenEOF {
			break
		}

		if t.GetTokenType() == int(aciparser.ACIExprLexerKEYWORD) {
			fmt.Printf("TType is a KEYWORD: '%s'\n", t.GetText())
		} else if t.GetTokenType() == int(aciparser.ACIExprLexerASSERTION) {
			fmt.Printf("TType is an ASSERTION: '%s'\n", t.GetText())
		} else if t.GetTokenType() == int(aciparser.ACIExprLexerWOP) {
			fmt.Printf("TType is a BOOLEAN OPERATOR: '%s'\n", t.GetText())
		} else if t.GetTokenType() == int(aciparser.ACIExprLexerCOP) {
			fmt.Printf("TType is a COMPARISON OPERATOR: '%s'\n", t.GetText())
		}
		time.Sleep(50*time.Millisecond)
	}
	*/
}
