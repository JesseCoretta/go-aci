package ldapparser

import (
	"github.com/antlr4-go/antlr/v4"
)

type ldapFilterListener struct {
    tempList []string
    *BaseLDAPParserListener
}

/*
VisitTerminal is the listener component that bridges the internal
ANTLR4 framework with the go-aci package. It is not intended for
use by users, and exists only to satisfy the particular ANTLR4
interface signature requirements needed for a Listener.
*/
func (r *ldapFilterListener) VisitTerminal(node antlr.TerminalNode) {
        r.tempList = append(r.tempList, node.GetText())
}

/*
FilterToTokens takes a non-zero string value representing an LDAP
Search Filter and tokenizes it using the ANTLR framework for Go.
An instance of []string -- hopefully containing all the tokenized
components of the input -- is returned.
*/
func FilterToTokens(filter string) (tokens []string) {
        // Create input stream using filter
        // input string.
        is := antlr.NewInputStream(filter)

        // prepare lexer and a token streamer
        lx := NewLDAPLexer(is)
        ts := antlr.NewCommonTokenStream(lx, 0)

        // prepare the parser and declare our
        // intent build a parse tree
        p := NewLDAPParser(ts)
        p.BuildParseTrees = true

        // Prepare parser tree using the SearchFilter
        // parser rule as input.
        tree := p.SearchFilter()

        // Initialize our filter expr listener type and
        // initialize a temporary []string instance for
        // storage of said tokens.
        listen := new(ldapFilterListener)
        listen.tempList = make([]string,0)

        // Trigger the antlr parse tree walker
        // which will build said temporary list
        // for the user to traverse.
        walker := antlr.NewParseTreeWalker()
        walker.Walk(listen, tree)

	return listen.tempList
}
