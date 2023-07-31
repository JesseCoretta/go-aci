// Code generated from LDAPLexer.g4 by ANTLR 4.13.0. DO NOT EDIT.

package ldapparser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type LDAPLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var LDAPLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ldaplexerLexerInit() {
	staticData := &LDAPLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'('", "')'", "','", "'?'", "'*'", "'dn'", "' '",
	}
	staticData.SymbolicNames = []string{
		"", "LPAREN", "RPAREN", "COMMA", "QMARK", "ASTERISK", "DN", "WHITESPACE",
		"BASEOBJECT_SCOPE", "SINGLELEVEL_SCOPE", "WHOLESUBTREE_SCOPE", "LOCAL_LDAP_SCHEME",
		"NOT", "AND", "OR", "OBJECT_IDENTIFIER", "IDENTIFIER", "EQUALITY", "APPROXIMATE",
		"LESS_THAN_OR_EQUAL", "GREATER_THAN_OR_EQUAL", "EXTENSIBLE_RULE", "EXCLUSIONS",
	}
	staticData.RuleNames = []string{
		"LPAREN", "RPAREN", "COMMA", "QMARK", "ASTERISK", "DN", "WHITESPACE",
		"LBRAK", "RBRAK", "DOLLAR", "SOLIDUS", "DOT", "TILDE", "EQ", "EX", "APX",
		"GE", "LE", "ATTR", "LDAP", "BASEOBJECT_SCOPE", "SINGLELEVEL_SCOPE",
		"WHOLESUBTREE_SCOPE", "LOCAL_LDAP_SCHEME", "EXCLAMATION", "NOT", "AMPERSAND",
		"AND", "VERTICAL", "OR", "OBJECT_IDENTIFIER", "IDENTIFIER", "EQUALITY",
		"APPROXIMATE", "LESS_THAN_OR_EQUAL", "GREATER_THAN_OR_EQUAL", "EXTENSIBLE_RULE",
		"EXCLUSIONS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 22, 225, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36,
		7, 36, 2, 37, 7, 37, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20,
		1, 20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1,
		22, 1, 23, 1, 23, 3, 23, 143, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 3, 23, 152, 8, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1,
		23, 3, 23, 160, 8, 23, 1, 23, 5, 23, 163, 8, 23, 10, 23, 12, 23, 166, 9,
		23, 1, 23, 3, 23, 169, 8, 23, 1, 23, 1, 23, 1, 23, 3, 23, 174, 8, 23, 3,
		23, 176, 8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26,
		1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 4, 30, 193, 8, 30, 11,
		30, 12, 30, 194, 1, 30, 1, 30, 4, 30, 199, 8, 30, 11, 30, 12, 30, 200,
		4, 30, 203, 8, 30, 11, 30, 12, 30, 204, 1, 31, 4, 31, 208, 8, 31, 11, 31,
		12, 31, 209, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1,
		36, 1, 36, 1, 37, 1, 37, 1, 37, 1, 37, 0, 0, 38, 1, 1, 3, 2, 5, 3, 7, 4,
		9, 5, 11, 6, 13, 7, 15, 0, 17, 0, 19, 0, 21, 0, 23, 0, 25, 0, 27, 0, 29,
		0, 31, 0, 33, 0, 35, 0, 37, 0, 39, 0, 41, 8, 43, 9, 45, 10, 47, 11, 49,
		0, 51, 12, 53, 0, 55, 13, 57, 0, 59, 14, 61, 15, 63, 16, 65, 17, 67, 18,
		69, 19, 71, 20, 73, 21, 75, 22, 1, 0, 11, 2, 0, 66, 66, 98, 98, 2, 0, 65,
		65, 97, 97, 2, 0, 83, 83, 115, 115, 2, 0, 69, 69, 101, 101, 2, 0, 79, 79,
		111, 111, 2, 0, 78, 78, 110, 110, 2, 0, 85, 85, 117, 117, 2, 0, 105, 105,
		115, 115, 1, 0, 48, 57, 5, 0, 45, 45, 48, 57, 59, 59, 65, 90, 97, 122,
		3, 0, 9, 10, 13, 13, 40, 41, 221, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0,
		0, 47, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 59, 1, 0, 0,
		0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0,
		0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0, 0, 75, 1,
		0, 0, 0, 1, 77, 1, 0, 0, 0, 3, 79, 1, 0, 0, 0, 5, 81, 1, 0, 0, 0, 7, 83,
		1, 0, 0, 0, 9, 85, 1, 0, 0, 0, 11, 87, 1, 0, 0, 0, 13, 90, 1, 0, 0, 0,
		15, 92, 1, 0, 0, 0, 17, 94, 1, 0, 0, 0, 19, 96, 1, 0, 0, 0, 21, 98, 1,
		0, 0, 0, 23, 100, 1, 0, 0, 0, 25, 102, 1, 0, 0, 0, 27, 104, 1, 0, 0, 0,
		29, 106, 1, 0, 0, 0, 31, 108, 1, 0, 0, 0, 33, 111, 1, 0, 0, 0, 35, 114,
		1, 0, 0, 0, 37, 117, 1, 0, 0, 0, 39, 122, 1, 0, 0, 0, 41, 127, 1, 0, 0,
		0, 43, 132, 1, 0, 0, 0, 45, 136, 1, 0, 0, 0, 47, 140, 1, 0, 0, 0, 49, 179,
		1, 0, 0, 0, 51, 181, 1, 0, 0, 0, 53, 183, 1, 0, 0, 0, 55, 185, 1, 0, 0,
		0, 57, 187, 1, 0, 0, 0, 59, 189, 1, 0, 0, 0, 61, 192, 1, 0, 0, 0, 63, 207,
		1, 0, 0, 0, 65, 211, 1, 0, 0, 0, 67, 213, 1, 0, 0, 0, 69, 215, 1, 0, 0,
		0, 71, 217, 1, 0, 0, 0, 73, 219, 1, 0, 0, 0, 75, 221, 1, 0, 0, 0, 77, 78,
		5, 40, 0, 0, 78, 2, 1, 0, 0, 0, 79, 80, 5, 41, 0, 0, 80, 4, 1, 0, 0, 0,
		81, 82, 5, 44, 0, 0, 82, 6, 1, 0, 0, 0, 83, 84, 5, 63, 0, 0, 84, 8, 1,
		0, 0, 0, 85, 86, 5, 42, 0, 0, 86, 10, 1, 0, 0, 0, 87, 88, 5, 100, 0, 0,
		88, 89, 5, 110, 0, 0, 89, 12, 1, 0, 0, 0, 90, 91, 5, 32, 0, 0, 91, 14,
		1, 0, 0, 0, 92, 93, 5, 91, 0, 0, 93, 16, 1, 0, 0, 0, 94, 95, 5, 93, 0,
		0, 95, 18, 1, 0, 0, 0, 96, 97, 5, 36, 0, 0, 97, 20, 1, 0, 0, 0, 98, 99,
		5, 47, 0, 0, 99, 22, 1, 0, 0, 0, 100, 101, 5, 46, 0, 0, 101, 24, 1, 0,
		0, 0, 102, 103, 5, 126, 0, 0, 103, 26, 1, 0, 0, 0, 104, 105, 5, 61, 0,
		0, 105, 28, 1, 0, 0, 0, 106, 107, 5, 58, 0, 0, 107, 30, 1, 0, 0, 0, 108,
		109, 3, 25, 12, 0, 109, 110, 3, 27, 13, 0, 110, 32, 1, 0, 0, 0, 111, 112,
		5, 62, 0, 0, 112, 113, 3, 27, 13, 0, 113, 34, 1, 0, 0, 0, 114, 115, 5,
		60, 0, 0, 115, 116, 3, 27, 13, 0, 116, 36, 1, 0, 0, 0, 117, 118, 5, 97,
		0, 0, 118, 119, 5, 116, 0, 0, 119, 120, 5, 116, 0, 0, 120, 121, 5, 114,
		0, 0, 121, 38, 1, 0, 0, 0, 122, 123, 5, 108, 0, 0, 123, 124, 5, 100, 0,
		0, 124, 125, 5, 97, 0, 0, 125, 126, 5, 112, 0, 0, 126, 40, 1, 0, 0, 0,
		127, 128, 7, 0, 0, 0, 128, 129, 7, 1, 0, 0, 129, 130, 7, 2, 0, 0, 130,
		131, 7, 3, 0, 0, 131, 42, 1, 0, 0, 0, 132, 133, 7, 4, 0, 0, 133, 134, 7,
		5, 0, 0, 134, 135, 7, 3, 0, 0, 135, 44, 1, 0, 0, 0, 136, 137, 7, 2, 0,
		0, 137, 138, 7, 6, 0, 0, 138, 139, 7, 0, 0, 0, 139, 46, 1, 0, 0, 0, 140,
		142, 3, 39, 19, 0, 141, 143, 7, 7, 0, 0, 142, 141, 1, 0, 0, 0, 142, 143,
		1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 145, 3, 29, 14, 0, 145, 146, 3,
		21, 10, 0, 146, 175, 3, 21, 10, 0, 147, 152, 3, 15, 7, 0, 148, 149, 5,
		37, 0, 0, 149, 150, 5, 50, 0, 0, 150, 152, 5, 102, 0, 0, 151, 147, 1, 0,
		0, 0, 151, 148, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0,
		153, 164, 3, 63, 31, 0, 154, 160, 3, 23, 11, 0, 155, 160, 3, 29, 14, 0,
		156, 157, 5, 37, 0, 0, 157, 158, 5, 50, 0, 0, 158, 160, 5, 102, 0, 0, 159,
		154, 1, 0, 0, 0, 159, 155, 1, 0, 0, 0, 159, 156, 1, 0, 0, 0, 160, 161,
		1, 0, 0, 0, 161, 163, 3, 63, 31, 0, 162, 159, 1, 0, 0, 0, 163, 166, 1,
		0, 0, 0, 164, 162, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 168, 1, 0, 0,
		0, 166, 164, 1, 0, 0, 0, 167, 169, 3, 17, 8, 0, 168, 167, 1, 0, 0, 0, 168,
		169, 1, 0, 0, 0, 169, 173, 1, 0, 0, 0, 170, 171, 3, 29, 14, 0, 171, 172,
		3, 63, 31, 0, 172, 174, 1, 0, 0, 0, 173, 170, 1, 0, 0, 0, 173, 174, 1,
		0, 0, 0, 174, 176, 1, 0, 0, 0, 175, 151, 1, 0, 0, 0, 175, 176, 1, 0, 0,
		0, 176, 177, 1, 0, 0, 0, 177, 178, 3, 21, 10, 0, 178, 48, 1, 0, 0, 0, 179,
		180, 5, 33, 0, 0, 180, 50, 1, 0, 0, 0, 181, 182, 3, 49, 24, 0, 182, 52,
		1, 0, 0, 0, 183, 184, 5, 38, 0, 0, 184, 54, 1, 0, 0, 0, 185, 186, 3, 53,
		26, 0, 186, 56, 1, 0, 0, 0, 187, 188, 5, 124, 0, 0, 188, 58, 1, 0, 0, 0,
		189, 190, 3, 57, 28, 0, 190, 60, 1, 0, 0, 0, 191, 193, 7, 8, 0, 0, 192,
		191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 192, 1, 0, 0, 0, 194, 195,
		1, 0, 0, 0, 195, 202, 1, 0, 0, 0, 196, 198, 3, 23, 11, 0, 197, 199, 7,
		8, 0, 0, 198, 197, 1, 0, 0, 0, 199, 200, 1, 0, 0, 0, 200, 198, 1, 0, 0,
		0, 200, 201, 1, 0, 0, 0, 201, 203, 1, 0, 0, 0, 202, 196, 1, 0, 0, 0, 203,
		204, 1, 0, 0, 0, 204, 202, 1, 0, 0, 0, 204, 205, 1, 0, 0, 0, 205, 62, 1,
		0, 0, 0, 206, 208, 7, 9, 0, 0, 207, 206, 1, 0, 0, 0, 208, 209, 1, 0, 0,
		0, 209, 207, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 64, 1, 0, 0, 0, 211,
		212, 3, 27, 13, 0, 212, 66, 1, 0, 0, 0, 213, 214, 3, 31, 15, 0, 214, 68,
		1, 0, 0, 0, 215, 216, 3, 35, 17, 0, 216, 70, 1, 0, 0, 0, 217, 218, 3, 33,
		16, 0, 218, 72, 1, 0, 0, 0, 219, 220, 3, 29, 14, 0, 220, 74, 1, 0, 0, 0,
		221, 222, 7, 10, 0, 0, 222, 223, 1, 0, 0, 0, 223, 224, 6, 37, 0, 0, 224,
		76, 1, 0, 0, 0, 12, 0, 142, 151, 159, 164, 168, 173, 175, 194, 200, 204,
		209, 1, 6, 0, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// LDAPLexerInit initializes any static state used to implement LDAPLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewLDAPLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func LDAPLexerInit() {
	staticData := &LDAPLexerLexerStaticData
	staticData.once.Do(ldaplexerLexerInit)
}

// NewLDAPLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewLDAPLexer(input antlr.CharStream) *LDAPLexer {
	LDAPLexerInit()
	l := new(LDAPLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &LDAPLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "LDAPLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LDAPLexer tokens.
const (
	LDAPLexerLPAREN                = 1
	LDAPLexerRPAREN                = 2
	LDAPLexerCOMMA                 = 3
	LDAPLexerQMARK                 = 4
	LDAPLexerASTERISK              = 5
	LDAPLexerDN                    = 6
	LDAPLexerWHITESPACE            = 7
	LDAPLexerBASEOBJECT_SCOPE      = 8
	LDAPLexerSINGLELEVEL_SCOPE     = 9
	LDAPLexerWHOLESUBTREE_SCOPE    = 10
	LDAPLexerLOCAL_LDAP_SCHEME     = 11
	LDAPLexerNOT                   = 12
	LDAPLexerAND                   = 13
	LDAPLexerOR                    = 14
	LDAPLexerOBJECT_IDENTIFIER     = 15
	LDAPLexerIDENTIFIER            = 16
	LDAPLexerEQUALITY              = 17
	LDAPLexerAPPROXIMATE           = 18
	LDAPLexerLESS_THAN_OR_EQUAL    = 19
	LDAPLexerGREATER_THAN_OR_EQUAL = 20
	LDAPLexerEXTENSIBLE_RULE       = 21
	LDAPLexerEXCLUSIONS            = 22
)
