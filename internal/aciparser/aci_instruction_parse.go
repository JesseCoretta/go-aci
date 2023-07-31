package aciparser

import (
        "github.com/antlr4-go/antlr/v4"
)

/*
instructionListener is a private type used to facilitate the
parsing of an ACIv3 instruction expression through ANTLR4.
*/
type instructionListener struct {
	tempList []string
	*BaseACIParserListener
}

/*
VisitTerminal is the listener component that bridges the internal
ANTLR4 framework with the go-aci package. It is not intended for
use by users, and exists only to satisfy the particular ANTLR4
interface signature requirements needed for a Listener.

The main purpose of this particular function, for us, is to build
the list of tokens (tempList) which we can read outside the ANTLR
framework as needed.
*/
func (r *instructionListener) VisitTerminal(node antlr.TerminalNode) {
	value := node.GetSymbol()
	if value.GetTokenType() == antlr.TokenEOF {
		return
	}

        r.tempList = append(r.tempList, value.GetText())
}

/*
InstructionToTokens parses the input value (instruct) into ANTLR
tokens ([]string). An instance of []string -- hopefully containing
all the tokenized components of the input -- is returned.
*/
func InstructionToTokens(instruct string) (tokens []string) {
        // Create input stream using instruct
        // input string.
        is := antlr.NewInputStream(instruct)

        // prepare lexer and a token streamer
        lx := NewACILexer(is)
        ts := antlr.NewCommonTokenStream(lx, 0)

        // prepare the parser and declare our
        // intent build a parse tree
        p := NewACIParser(ts)
        p.BuildParseTrees = true
	tree := p.Parse()

        // Initialize our instruction listener type and
        // initialize a temporary []string instance for
	// storage of said tokens.
        listen := new(instructionListener)
        listen.tempList = make([]string,0)

        // Trigger the antlr parse tree walker
        // which will build said temporary list
        // for the user to traverse.
        walker := antlr.NewParseTreeWalker()
        walker.Walk(listen, tree)

        return listen.tempList
}
