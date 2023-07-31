// Code generated from LDAPParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package ldapparser // LDAPParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type LDAPParser struct {
	*antlr.BaseParser
}

var LDAPParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func ldapparserParserInit() {
	staticData := &LDAPParserParserStaticData
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
		"uniformResourceIdentifier", "uRISearchAttributes", "uRIDelimiter",
		"searchScope", "searchFilter", "searchFilterExpr", "and", "or", "not",
		"attributeValueAssertion", "attributeTypeOrMatchingRule", "extensibleMatch",
		"equalityMatch", "greaterThanOrEqualMatch", "lessThanOrEqualMatch",
		"approximateMatch", "objectIdentifier", "openingParenthesis", "closingParenthesis",
		"distinguishedName", "attributeValue",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 22, 208, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 54,
		8, 1, 10, 1, 12, 1, 57, 9, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 3, 3, 64, 8,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 71, 8, 4, 1, 5, 1, 5, 1, 5, 4, 5,
		76, 8, 5, 11, 5, 12, 5, 77, 4, 5, 80, 8, 5, 11, 5, 12, 5, 81, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 4, 5, 89, 8, 5, 11, 5, 12, 5, 90, 4, 5, 93, 8, 5,
		11, 5, 12, 5, 94, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 4, 5,
		105, 8, 5, 11, 5, 12, 5, 106, 3, 5, 109, 8, 5, 1, 6, 1, 6, 1, 7, 1, 7,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 3, 9, 122, 8, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 3, 11, 128, 8, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 134,
		8, 11, 5, 11, 136, 8, 11, 10, 11, 12, 11, 139, 9, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 3, 11, 145, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 3, 12, 155, 8, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		19, 3, 19, 184, 8, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19, 1, 19,
		1, 19, 1, 19, 5, 19, 195, 8, 19, 10, 19, 12, 19, 198, 9, 19, 1, 20, 4,
		20, 201, 8, 20, 11, 20, 12, 20, 202, 1, 20, 3, 20, 206, 8, 20, 1, 20, 3,
		81, 94, 202, 0, 21, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26,
		28, 30, 32, 34, 36, 38, 40, 0, 0, 213, 0, 42, 1, 0, 0, 0, 2, 50, 1, 0,
		0, 0, 4, 58, 1, 0, 0, 0, 6, 63, 1, 0, 0, 0, 8, 70, 1, 0, 0, 0, 10, 108,
		1, 0, 0, 0, 12, 110, 1, 0, 0, 0, 14, 112, 1, 0, 0, 0, 16, 114, 1, 0, 0,
		0, 18, 121, 1, 0, 0, 0, 20, 123, 1, 0, 0, 0, 22, 125, 1, 0, 0, 0, 24, 148,
		1, 0, 0, 0, 26, 158, 1, 0, 0, 0, 28, 164, 1, 0, 0, 0, 30, 170, 1, 0, 0,
		0, 32, 176, 1, 0, 0, 0, 34, 178, 1, 0, 0, 0, 36, 180, 1, 0, 0, 0, 38, 183,
		1, 0, 0, 0, 40, 205, 1, 0, 0, 0, 42, 43, 3, 38, 19, 0, 43, 44, 3, 4, 2,
		0, 44, 45, 3, 2, 1, 0, 45, 46, 3, 4, 2, 0, 46, 47, 3, 6, 3, 0, 47, 48,
		3, 4, 2, 0, 48, 49, 3, 8, 4, 0, 49, 1, 1, 0, 0, 0, 50, 55, 3, 20, 10, 0,
		51, 52, 5, 3, 0, 0, 52, 54, 3, 20, 10, 0, 53, 51, 1, 0, 0, 0, 54, 57, 1,
		0, 0, 0, 55, 53, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 3, 1, 0, 0, 0, 57,
		55, 1, 0, 0, 0, 58, 59, 5, 4, 0, 0, 59, 5, 1, 0, 0, 0, 60, 64, 5, 8, 0,
		0, 61, 64, 5, 9, 0, 0, 62, 64, 5, 10, 0, 0, 63, 60, 1, 0, 0, 0, 63, 61,
		1, 0, 0, 0, 63, 62, 1, 0, 0, 0, 64, 7, 1, 0, 0, 0, 65, 66, 3, 34, 17, 0,
		66, 67, 3, 8, 4, 0, 67, 68, 3, 36, 18, 0, 68, 71, 1, 0, 0, 0, 69, 71, 3,
		10, 5, 0, 70, 65, 1, 0, 0, 0, 70, 69, 1, 0, 0, 0, 71, 9, 1, 0, 0, 0, 72,
		79, 3, 34, 17, 0, 73, 75, 3, 12, 6, 0, 74, 76, 3, 10, 5, 0, 75, 74, 1,
		0, 0, 0, 76, 77, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78,
		80, 1, 0, 0, 0, 79, 73, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 81, 82, 1, 0, 0,
		0, 81, 79, 1, 0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 84, 3, 36, 18, 0, 84, 109,
		1, 0, 0, 0, 85, 92, 3, 34, 17, 0, 86, 88, 3, 14, 7, 0, 87, 89, 3, 10, 5,
		0, 88, 87, 1, 0, 0, 0, 89, 90, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0, 90, 91,
		1, 0, 0, 0, 91, 93, 1, 0, 0, 0, 92, 86, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0,
		94, 95, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 97, 3,
		36, 18, 0, 97, 109, 1, 0, 0, 0, 98, 99, 3, 34, 17, 0, 99, 100, 3, 16, 8,
		0, 100, 101, 3, 10, 5, 0, 101, 102, 3, 36, 18, 0, 102, 109, 1, 0, 0, 0,
		103, 105, 3, 18, 9, 0, 104, 103, 1, 0, 0, 0, 105, 106, 1, 0, 0, 0, 106,
		104, 1, 0, 0, 0, 106, 107, 1, 0, 0, 0, 107, 109, 1, 0, 0, 0, 108, 72, 1,
		0, 0, 0, 108, 85, 1, 0, 0, 0, 108, 98, 1, 0, 0, 0, 108, 104, 1, 0, 0, 0,
		109, 11, 1, 0, 0, 0, 110, 111, 5, 13, 0, 0, 111, 13, 1, 0, 0, 0, 112, 113,
		5, 14, 0, 0, 113, 15, 1, 0, 0, 0, 114, 115, 5, 12, 0, 0, 115, 17, 1, 0,
		0, 0, 116, 122, 3, 24, 12, 0, 117, 122, 3, 26, 13, 0, 118, 122, 3, 28,
		14, 0, 119, 122, 3, 30, 15, 0, 120, 122, 3, 22, 11, 0, 121, 116, 1, 0,
		0, 0, 121, 117, 1, 0, 0, 0, 121, 118, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0,
		121, 120, 1, 0, 0, 0, 122, 19, 1, 0, 0, 0, 123, 124, 5, 16, 0, 0, 124,
		21, 1, 0, 0, 0, 125, 127, 3, 34, 17, 0, 126, 128, 3, 20, 10, 0, 127, 126,
		1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 137, 1, 0, 0, 0, 129, 133, 5, 21,
		0, 0, 130, 134, 5, 6, 0, 0, 131, 134, 3, 20, 10, 0, 132, 134, 3, 32, 16,
		0, 133, 130, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 132, 1, 0, 0, 0, 134,
		136, 1, 0, 0, 0, 135, 129, 1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135,
		1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 140, 1, 0, 0, 0, 139, 137, 1, 0,
		0, 0, 140, 141, 5, 21, 0, 0, 141, 144, 5, 17, 0, 0, 142, 145, 3, 38, 19,
		0, 143, 145, 3, 40, 20, 0, 144, 142, 1, 0, 0, 0, 144, 143, 1, 0, 0, 0,
		145, 146, 1, 0, 0, 0, 146, 147, 3, 36, 18, 0, 147, 23, 1, 0, 0, 0, 148,
		149, 3, 34, 17, 0, 149, 150, 3, 20, 10, 0, 150, 154, 5, 17, 0, 0, 151,
		155, 3, 38, 19, 0, 152, 155, 3, 40, 20, 0, 153, 155, 5, 5, 0, 0, 154, 151,
		1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 153, 1, 0, 0, 0, 155, 156, 1, 0,
		0, 0, 156, 157, 3, 36, 18, 0, 157, 25, 1, 0, 0, 0, 158, 159, 3, 34, 17,
		0, 159, 160, 3, 20, 10, 0, 160, 161, 5, 20, 0, 0, 161, 162, 3, 40, 20,
		0, 162, 163, 3, 36, 18, 0, 163, 27, 1, 0, 0, 0, 164, 165, 3, 34, 17, 0,
		165, 166, 3, 20, 10, 0, 166, 167, 5, 19, 0, 0, 167, 168, 3, 40, 20, 0,
		168, 169, 3, 36, 18, 0, 169, 29, 1, 0, 0, 0, 170, 171, 3, 34, 17, 0, 171,
		172, 3, 20, 10, 0, 172, 173, 5, 18, 0, 0, 173, 174, 3, 40, 20, 0, 174,
		175, 3, 36, 18, 0, 175, 31, 1, 0, 0, 0, 176, 177, 5, 15, 0, 0, 177, 33,
		1, 0, 0, 0, 178, 179, 5, 1, 0, 0, 179, 35, 1, 0, 0, 0, 180, 181, 5, 2,
		0, 0, 181, 37, 1, 0, 0, 0, 182, 184, 5, 11, 0, 0, 183, 182, 1, 0, 0, 0,
		183, 184, 1, 0, 0, 0, 184, 185, 1, 0, 0, 0, 185, 186, 3, 20, 10, 0, 186,
		187, 5, 17, 0, 0, 187, 188, 3, 40, 20, 0, 188, 196, 1, 0, 0, 0, 189, 190,
		5, 3, 0, 0, 190, 191, 3, 20, 10, 0, 191, 192, 5, 17, 0, 0, 192, 193, 3,
		40, 20, 0, 193, 195, 1, 0, 0, 0, 194, 189, 1, 0, 0, 0, 195, 198, 1, 0,
		0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 39, 1, 0, 0, 0,
		198, 196, 1, 0, 0, 0, 199, 201, 9, 0, 0, 0, 200, 199, 1, 0, 0, 0, 201,
		202, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 202, 200, 1, 0, 0, 0, 203, 206,
		1, 0, 0, 0, 204, 206, 9, 0, 0, 0, 205, 200, 1, 0, 0, 0, 205, 204, 1, 0,
		0, 0, 206, 41, 1, 0, 0, 0, 19, 55, 63, 70, 77, 81, 90, 94, 106, 108, 121,
		127, 133, 137, 144, 154, 183, 196, 202, 205,
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

// LDAPParserInit initializes any static state used to implement LDAPParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLDAPParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LDAPParserInit() {
	staticData := &LDAPParserParserStaticData
	staticData.once.Do(ldapparserParserInit)
}

// NewLDAPParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLDAPParser(input antlr.TokenStream) *LDAPParser {
	LDAPParserInit()
	this := new(LDAPParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &LDAPParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "LDAPParser.g4"

	return this
}

// LDAPParser tokens.
const (
	LDAPParserEOF                   = antlr.TokenEOF
	LDAPParserLPAREN                = 1
	LDAPParserRPAREN                = 2
	LDAPParserCOMMA                 = 3
	LDAPParserQMARK                 = 4
	LDAPParserASTERISK              = 5
	LDAPParserDN                    = 6
	LDAPParserWHITESPACE            = 7
	LDAPParserBASEOBJECT_SCOPE      = 8
	LDAPParserSINGLELEVEL_SCOPE     = 9
	LDAPParserWHOLESUBTREE_SCOPE    = 10
	LDAPParserLOCAL_LDAP_SCHEME     = 11
	LDAPParserNOT                   = 12
	LDAPParserAND                   = 13
	LDAPParserOR                    = 14
	LDAPParserOBJECT_IDENTIFIER     = 15
	LDAPParserIDENTIFIER            = 16
	LDAPParserEQUALITY              = 17
	LDAPParserAPPROXIMATE           = 18
	LDAPParserLESS_THAN_OR_EQUAL    = 19
	LDAPParserGREATER_THAN_OR_EQUAL = 20
	LDAPParserEXTENSIBLE_RULE       = 21
	LDAPParserEXCLUSIONS            = 22
)

// LDAPParser rules.
const (
	LDAPParserRULE_uniformResourceIdentifier   = 0
	LDAPParserRULE_uRISearchAttributes         = 1
	LDAPParserRULE_uRIDelimiter                = 2
	LDAPParserRULE_searchScope                 = 3
	LDAPParserRULE_searchFilter                = 4
	LDAPParserRULE_searchFilterExpr            = 5
	LDAPParserRULE_and                         = 6
	LDAPParserRULE_or                          = 7
	LDAPParserRULE_not                         = 8
	LDAPParserRULE_attributeValueAssertion     = 9
	LDAPParserRULE_attributeTypeOrMatchingRule = 10
	LDAPParserRULE_extensibleMatch             = 11
	LDAPParserRULE_equalityMatch               = 12
	LDAPParserRULE_greaterThanOrEqualMatch     = 13
	LDAPParserRULE_lessThanOrEqualMatch        = 14
	LDAPParserRULE_approximateMatch            = 15
	LDAPParserRULE_objectIdentifier            = 16
	LDAPParserRULE_openingParenthesis          = 17
	LDAPParserRULE_closingParenthesis          = 18
	LDAPParserRULE_distinguishedName           = 19
	LDAPParserRULE_attributeValue              = 20
)

// IUniformResourceIdentifierContext is an interface to support dynamic dispatch.
type IUniformResourceIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DistinguishedName() IDistinguishedNameContext
	AllURIDelimiter() []IURIDelimiterContext
	URIDelimiter(i int) IURIDelimiterContext
	URISearchAttributes() IURISearchAttributesContext
	SearchScope() ISearchScopeContext
	SearchFilter() ISearchFilterContext

	// IsUniformResourceIdentifierContext differentiates from other interfaces.
	IsUniformResourceIdentifierContext()
}

type UniformResourceIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUniformResourceIdentifierContext() *UniformResourceIdentifierContext {
	var p = new(UniformResourceIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_uniformResourceIdentifier
	return p
}

func InitEmptyUniformResourceIdentifierContext(p *UniformResourceIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_uniformResourceIdentifier
}

func (*UniformResourceIdentifierContext) IsUniformResourceIdentifierContext() {}

func NewUniformResourceIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UniformResourceIdentifierContext {
	var p = new(UniformResourceIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_uniformResourceIdentifier

	return p
}

func (s *UniformResourceIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *UniformResourceIdentifierContext) DistinguishedName() IDistinguishedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDistinguishedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDistinguishedNameContext)
}

func (s *UniformResourceIdentifierContext) AllURIDelimiter() []IURIDelimiterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IURIDelimiterContext); ok {
			len++
		}
	}

	tst := make([]IURIDelimiterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IURIDelimiterContext); ok {
			tst[i] = t.(IURIDelimiterContext)
			i++
		}
	}

	return tst
}

func (s *UniformResourceIdentifierContext) URIDelimiter(i int) IURIDelimiterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IURIDelimiterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IURIDelimiterContext)
}

func (s *UniformResourceIdentifierContext) URISearchAttributes() IURISearchAttributesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IURISearchAttributesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IURISearchAttributesContext)
}

func (s *UniformResourceIdentifierContext) SearchScope() ISearchScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchScopeContext)
}

func (s *UniformResourceIdentifierContext) SearchFilter() ISearchFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchFilterContext)
}

func (s *UniformResourceIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UniformResourceIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UniformResourceIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterUniformResourceIdentifier(s)
	}
}

func (s *UniformResourceIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitUniformResourceIdentifier(s)
	}
}

func (p *LDAPParser) UniformResourceIdentifier() (localctx IUniformResourceIdentifierContext) {
	localctx = NewUniformResourceIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LDAPParserRULE_uniformResourceIdentifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(42)
		p.DistinguishedName()
	}
	{
		p.SetState(43)
		p.URIDelimiter()
	}
	{
		p.SetState(44)
		p.URISearchAttributes()
	}
	{
		p.SetState(45)
		p.URIDelimiter()
	}
	{
		p.SetState(46)
		p.SearchScope()
	}
	{
		p.SetState(47)
		p.URIDelimiter()
	}
	{
		p.SetState(48)
		p.SearchFilter()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IURISearchAttributesContext is an interface to support dynamic dispatch.
type IURISearchAttributesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAttributeTypeOrMatchingRule() []IAttributeTypeOrMatchingRuleContext
	AttributeTypeOrMatchingRule(i int) IAttributeTypeOrMatchingRuleContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsURISearchAttributesContext differentiates from other interfaces.
	IsURISearchAttributesContext()
}

type URISearchAttributesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyURISearchAttributesContext() *URISearchAttributesContext {
	var p = new(URISearchAttributesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_uRISearchAttributes
	return p
}

func InitEmptyURISearchAttributesContext(p *URISearchAttributesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_uRISearchAttributes
}

func (*URISearchAttributesContext) IsURISearchAttributesContext() {}

func NewURISearchAttributesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *URISearchAttributesContext {
	var p = new(URISearchAttributesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_uRISearchAttributes

	return p
}

func (s *URISearchAttributesContext) GetParser() antlr.Parser { return s.parser }

func (s *URISearchAttributesContext) AllAttributeTypeOrMatchingRule() []IAttributeTypeOrMatchingRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeTypeOrMatchingRuleContext); ok {
			len++
		}
	}

	tst := make([]IAttributeTypeOrMatchingRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeTypeOrMatchingRuleContext); ok {
			tst[i] = t.(IAttributeTypeOrMatchingRuleContext)
			i++
		}
	}

	return tst
}

func (s *URISearchAttributesContext) AttributeTypeOrMatchingRule(i int) IAttributeTypeOrMatchingRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeOrMatchingRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeTypeOrMatchingRuleContext)
}

func (s *URISearchAttributesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LDAPParserCOMMA)
}

func (s *URISearchAttributesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LDAPParserCOMMA, i)
}

func (s *URISearchAttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *URISearchAttributesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *URISearchAttributesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterURISearchAttributes(s)
	}
}

func (s *URISearchAttributesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitURISearchAttributes(s)
	}
}

func (p *LDAPParser) URISearchAttributes() (localctx IURISearchAttributesContext) {
	localctx = NewURISearchAttributesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LDAPParserRULE_uRISearchAttributes)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.AttributeTypeOrMatchingRule()
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LDAPParserCOMMA {
		{
			p.SetState(51)
			p.Match(LDAPParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(52)
			p.AttributeTypeOrMatchingRule()
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IURIDelimiterContext is an interface to support dynamic dispatch.
type IURIDelimiterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QMARK() antlr.TerminalNode

	// IsURIDelimiterContext differentiates from other interfaces.
	IsURIDelimiterContext()
}

type URIDelimiterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyURIDelimiterContext() *URIDelimiterContext {
	var p = new(URIDelimiterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_uRIDelimiter
	return p
}

func InitEmptyURIDelimiterContext(p *URIDelimiterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_uRIDelimiter
}

func (*URIDelimiterContext) IsURIDelimiterContext() {}

func NewURIDelimiterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *URIDelimiterContext {
	var p = new(URIDelimiterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_uRIDelimiter

	return p
}

func (s *URIDelimiterContext) GetParser() antlr.Parser { return s.parser }

func (s *URIDelimiterContext) QMARK() antlr.TerminalNode {
	return s.GetToken(LDAPParserQMARK, 0)
}

func (s *URIDelimiterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *URIDelimiterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *URIDelimiterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterURIDelimiter(s)
	}
}

func (s *URIDelimiterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitURIDelimiter(s)
	}
}

func (p *LDAPParser) URIDelimiter() (localctx IURIDelimiterContext) {
	localctx = NewURIDelimiterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LDAPParserRULE_uRIDelimiter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(58)
		p.Match(LDAPParserQMARK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISearchScopeContext is an interface to support dynamic dispatch.
type ISearchScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSearchScopeContext differentiates from other interfaces.
	IsSearchScopeContext()
}

type SearchScopeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchScopeContext() *SearchScopeContext {
	var p = new(SearchScopeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_searchScope
	return p
}

func InitEmptySearchScopeContext(p *SearchScopeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_searchScope
}

func (*SearchScopeContext) IsSearchScopeContext() {}

func NewSearchScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchScopeContext {
	var p = new(SearchScopeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_searchScope

	return p
}

func (s *SearchScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchScopeContext) CopyAll(ctx *SearchScopeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SearchScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Subtree_scopeContext struct {
	SearchScopeContext
}

func NewSubtree_scopeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Subtree_scopeContext {
	var p = new(Subtree_scopeContext)

	InitEmptySearchScopeContext(&p.SearchScopeContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchScopeContext))

	return p
}

func (s *Subtree_scopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Subtree_scopeContext) WHOLESUBTREE_SCOPE() antlr.TerminalNode {
	return s.GetToken(LDAPParserWHOLESUBTREE_SCOPE, 0)
}

func (s *Subtree_scopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterSubtree_scope(s)
	}
}

func (s *Subtree_scopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitSubtree_scope(s)
	}
}

type Onelevel_scopeContext struct {
	SearchScopeContext
}

func NewOnelevel_scopeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Onelevel_scopeContext {
	var p = new(Onelevel_scopeContext)

	InitEmptySearchScopeContext(&p.SearchScopeContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchScopeContext))

	return p
}

func (s *Onelevel_scopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Onelevel_scopeContext) SINGLELEVEL_SCOPE() antlr.TerminalNode {
	return s.GetToken(LDAPParserSINGLELEVEL_SCOPE, 0)
}

func (s *Onelevel_scopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterOnelevel_scope(s)
	}
}

func (s *Onelevel_scopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitOnelevel_scope(s)
	}
}

type BaseObject_scopeContext struct {
	SearchScopeContext
}

func NewBaseObject_scopeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BaseObject_scopeContext {
	var p = new(BaseObject_scopeContext)

	InitEmptySearchScopeContext(&p.SearchScopeContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchScopeContext))

	return p
}

func (s *BaseObject_scopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseObject_scopeContext) BASEOBJECT_SCOPE() antlr.TerminalNode {
	return s.GetToken(LDAPParserBASEOBJECT_SCOPE, 0)
}

func (s *BaseObject_scopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterBaseObject_scope(s)
	}
}

func (s *BaseObject_scopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitBaseObject_scope(s)
	}
}

func (p *LDAPParser) SearchScope() (localctx ISearchScopeContext) {
	localctx = NewSearchScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LDAPParserRULE_searchScope)
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LDAPParserBASEOBJECT_SCOPE:
		localctx = NewBaseObject_scopeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Match(LDAPParserBASEOBJECT_SCOPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LDAPParserSINGLELEVEL_SCOPE:
		localctx = NewOnelevel_scopeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Match(LDAPParserSINGLELEVEL_SCOPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case LDAPParserWHOLESUBTREE_SCOPE:
		localctx = NewSubtree_scopeContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(62)
			p.Match(LDAPParserWHOLESUBTREE_SCOPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISearchFilterContext is an interface to support dynamic dispatch.
type ISearchFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSearchFilterContext differentiates from other interfaces.
	IsSearchFilterContext()
}

type SearchFilterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchFilterContext() *SearchFilterContext {
	var p = new(SearchFilterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_searchFilter
	return p
}

func InitEmptySearchFilterContext(p *SearchFilterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_searchFilter
}

func (*SearchFilterContext) IsSearchFilterContext() {}

func NewSearchFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchFilterContext {
	var p = new(SearchFilterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_searchFilter

	return p
}

func (s *SearchFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchFilterContext) CopyAll(ctx *SearchFilterContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SearchFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Parenthetical_filterContext struct {
	SearchFilterContext
}

func NewParenthetical_filterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Parenthetical_filterContext {
	var p = new(Parenthetical_filterContext)

	InitEmptySearchFilterContext(&p.SearchFilterContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchFilterContext))

	return p
}

func (s *Parenthetical_filterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Parenthetical_filterContext) OpeningParenthesis() IOpeningParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpeningParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpeningParenthesisContext)
}

func (s *Parenthetical_filterContext) SearchFilter() ISearchFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchFilterContext)
}

func (s *Parenthetical_filterContext) ClosingParenthesis() IClosingParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosingParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosingParenthesisContext)
}

func (s *Parenthetical_filterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterParenthetical_filter(s)
	}
}

func (s *Parenthetical_filterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitParenthetical_filter(s)
	}
}

type FilterContext struct {
	SearchFilterContext
}

func NewFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FilterContext {
	var p = new(FilterContext)

	InitEmptySearchFilterContext(&p.SearchFilterContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchFilterContext))

	return p
}

func (s *FilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterContext) SearchFilterExpr() ISearchFilterExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchFilterExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchFilterExprContext)
}

func (s *FilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterFilter(s)
	}
}

func (s *FilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitFilter(s)
	}
}

func (p *LDAPParser) SearchFilter() (localctx ISearchFilterContext) {
	localctx = NewSearchFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LDAPParserRULE_searchFilter)
	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenthetical_filterContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(65)
			p.OpeningParenthesis()
		}
		{
			p.SetState(66)
			p.SearchFilter()
		}
		{
			p.SetState(67)
			p.ClosingParenthesis()
		}

	case 2:
		localctx = NewFilterContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.SearchFilterExpr()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISearchFilterExprContext is an interface to support dynamic dispatch.
type ISearchFilterExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSearchFilterExprContext differentiates from other interfaces.
	IsSearchFilterExprContext()
}

type SearchFilterExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchFilterExprContext() *SearchFilterExprContext {
	var p = new(SearchFilterExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_searchFilterExpr
	return p
}

func InitEmptySearchFilterExprContext(p *SearchFilterExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_searchFilterExpr
}

func (*SearchFilterExprContext) IsSearchFilterExprContext() {}

func NewSearchFilterExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchFilterExprContext {
	var p = new(SearchFilterExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_searchFilterExpr

	return p
}

func (s *SearchFilterExprContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchFilterExprContext) CopyAll(ctx *SearchFilterExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SearchFilterExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchFilterExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type Not_filter_expressionContext struct {
	SearchFilterExprContext
}

func NewNot_filter_expressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Not_filter_expressionContext {
	var p = new(Not_filter_expressionContext)

	InitEmptySearchFilterExprContext(&p.SearchFilterExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchFilterExprContext))

	return p
}

func (s *Not_filter_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Not_filter_expressionContext) OpeningParenthesis() IOpeningParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpeningParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpeningParenthesisContext)
}

func (s *Not_filter_expressionContext) Not() INotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotContext)
}

func (s *Not_filter_expressionContext) SearchFilterExpr() ISearchFilterExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchFilterExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchFilterExprContext)
}

func (s *Not_filter_expressionContext) ClosingParenthesis() IClosingParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosingParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosingParenthesisContext)
}

func (s *Not_filter_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterNot_filter_expression(s)
	}
}

func (s *Not_filter_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitNot_filter_expression(s)
	}
}

type Or_filter_expressionContext struct {
	SearchFilterExprContext
}

func NewOr_filter_expressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Or_filter_expressionContext {
	var p = new(Or_filter_expressionContext)

	InitEmptySearchFilterExprContext(&p.SearchFilterExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchFilterExprContext))

	return p
}

func (s *Or_filter_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Or_filter_expressionContext) OpeningParenthesis() IOpeningParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpeningParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpeningParenthesisContext)
}

func (s *Or_filter_expressionContext) ClosingParenthesis() IClosingParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosingParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosingParenthesisContext)
}

func (s *Or_filter_expressionContext) AllOr() []IOrContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrContext); ok {
			len++
		}
	}

	tst := make([]IOrContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrContext); ok {
			tst[i] = t.(IOrContext)
			i++
		}
	}

	return tst
}

func (s *Or_filter_expressionContext) Or(i int) IOrContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrContext)
}

func (s *Or_filter_expressionContext) AllSearchFilterExpr() []ISearchFilterExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISearchFilterExprContext); ok {
			len++
		}
	}

	tst := make([]ISearchFilterExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISearchFilterExprContext); ok {
			tst[i] = t.(ISearchFilterExprContext)
			i++
		}
	}

	return tst
}

func (s *Or_filter_expressionContext) SearchFilterExpr(i int) ISearchFilterExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchFilterExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchFilterExprContext)
}

func (s *Or_filter_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterOr_filter_expression(s)
	}
}

func (s *Or_filter_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitOr_filter_expression(s)
	}
}

type Ava_expressionContext struct {
	SearchFilterExprContext
}

func NewAva_expressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Ava_expressionContext {
	var p = new(Ava_expressionContext)

	InitEmptySearchFilterExprContext(&p.SearchFilterExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchFilterExprContext))

	return p
}

func (s *Ava_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ava_expressionContext) AllAttributeValueAssertion() []IAttributeValueAssertionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeValueAssertionContext); ok {
			len++
		}
	}

	tst := make([]IAttributeValueAssertionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeValueAssertionContext); ok {
			tst[i] = t.(IAttributeValueAssertionContext)
			i++
		}
	}

	return tst
}

func (s *Ava_expressionContext) AttributeValueAssertion(i int) IAttributeValueAssertionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeValueAssertionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeValueAssertionContext)
}

func (s *Ava_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterAva_expression(s)
	}
}

func (s *Ava_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitAva_expression(s)
	}
}

type And_filter_expressionContext struct {
	SearchFilterExprContext
}

func NewAnd_filter_expressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *And_filter_expressionContext {
	var p = new(And_filter_expressionContext)

	InitEmptySearchFilterExprContext(&p.SearchFilterExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*SearchFilterExprContext))

	return p
}

func (s *And_filter_expressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *And_filter_expressionContext) OpeningParenthesis() IOpeningParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpeningParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpeningParenthesisContext)
}

func (s *And_filter_expressionContext) ClosingParenthesis() IClosingParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosingParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosingParenthesisContext)
}

func (s *And_filter_expressionContext) AllAnd() []IAndContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAndContext); ok {
			len++
		}
	}

	tst := make([]IAndContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAndContext); ok {
			tst[i] = t.(IAndContext)
			i++
		}
	}

	return tst
}

func (s *And_filter_expressionContext) And(i int) IAndContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAndContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAndContext)
}

func (s *And_filter_expressionContext) AllSearchFilterExpr() []ISearchFilterExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISearchFilterExprContext); ok {
			len++
		}
	}

	tst := make([]ISearchFilterExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISearchFilterExprContext); ok {
			tst[i] = t.(ISearchFilterExprContext)
			i++
		}
	}

	return tst
}

func (s *And_filter_expressionContext) SearchFilterExpr(i int) ISearchFilterExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchFilterExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchFilterExprContext)
}

func (s *And_filter_expressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterAnd_filter_expression(s)
	}
}

func (s *And_filter_expressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitAnd_filter_expression(s)
	}
}

func (p *LDAPParser) SearchFilterExpr() (localctx ISearchFilterExprContext) {
	localctx = NewSearchFilterExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LDAPParserRULE_searchFilterExpr)
	var _la int

	var _alt int

	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAnd_filter_expressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.OpeningParenthesis()
		}
		p.SetState(79)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1 + 1
		for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1 + 1:
				{
					p.SetState(73)
					p.And()
				}
				p.SetState(75)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for ok := true; ok; ok = _la == LDAPParserLPAREN {
					{
						p.SetState(74)
						p.SearchFilterExpr()
					}

					p.SetState(77)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(81)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.ClosingParenthesis()
		}

	case 2:
		localctx = NewOr_filter_expressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.OpeningParenthesis()
		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1 + 1
		for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1 + 1:
				{
					p.SetState(86)
					p.Or()
				}
				p.SetState(88)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				for ok := true; ok; ok = _la == LDAPParserLPAREN {
					{
						p.SetState(87)
						p.SearchFilterExpr()
					}

					p.SetState(90)
					p.GetErrorHandler().Sync(p)
					if p.HasError() {
						goto errorExit
					}
					_la = p.GetTokenStream().LA(1)
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(94)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		{
			p.SetState(96)
			p.ClosingParenthesis()
		}

	case 3:
		localctx = NewNot_filter_expressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)
			p.OpeningParenthesis()
		}
		{
			p.SetState(99)
			p.Not()
		}
		{
			p.SetState(100)
			p.SearchFilterExpr()
		}
		{
			p.SetState(101)
			p.ClosingParenthesis()
		}

	case 4:
		localctx = NewAva_expressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(103)
					p.AttributeValueAssertion()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(106)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAndContext is an interface to support dynamic dispatch.
type IAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AND() antlr.TerminalNode

	// IsAndContext differentiates from other interfaces.
	IsAndContext()
}

type AndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAndContext() *AndContext {
	var p = new(AndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_and
	return p
}

func InitEmptyAndContext(p *AndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_and
}

func (*AndContext) IsAndContext() {}

func NewAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AndContext {
	var p = new(AndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_and

	return p
}

func (s *AndContext) GetParser() antlr.Parser { return s.parser }

func (s *AndContext) AND() antlr.TerminalNode {
	return s.GetToken(LDAPParserAND, 0)
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitAnd(s)
	}
}

func (p *LDAPParser) And() (localctx IAndContext) {
	localctx = NewAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LDAPParserRULE_and)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(110)
		p.Match(LDAPParserAND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrContext is an interface to support dynamic dispatch.
type IOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OR() antlr.TerminalNode

	// IsOrContext differentiates from other interfaces.
	IsOrContext()
}

type OrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrContext() *OrContext {
	var p = new(OrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_or
	return p
}

func InitEmptyOrContext(p *OrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_or
}

func (*OrContext) IsOrContext() {}

func NewOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrContext {
	var p = new(OrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_or

	return p
}

func (s *OrContext) GetParser() antlr.Parser { return s.parser }

func (s *OrContext) OR() antlr.TerminalNode {
	return s.GetToken(LDAPParserOR, 0)
}

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitOr(s)
	}
}

func (p *LDAPParser) Or() (localctx IOrContext) {
	localctx = NewOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LDAPParserRULE_or)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(LDAPParserOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INotContext is an interface to support dynamic dispatch.
type INotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NOT() antlr.TerminalNode

	// IsNotContext differentiates from other interfaces.
	IsNotContext()
}

type NotContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotContext() *NotContext {
	var p = new(NotContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_not
	return p
}

func InitEmptyNotContext(p *NotContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_not
}

func (*NotContext) IsNotContext() {}

func NewNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotContext {
	var p = new(NotContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_not

	return p
}

func (s *NotContext) GetParser() antlr.Parser { return s.parser }

func (s *NotContext) NOT() antlr.TerminalNode {
	return s.GetToken(LDAPParserNOT, 0)
}

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterNot(s)
	}
}

func (s *NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitNot(s)
	}
}

func (p *LDAPParser) Not() (localctx INotContext) {
	localctx = NewNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LDAPParserRULE_not)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(LDAPParserNOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttributeValueAssertionContext is an interface to support dynamic dispatch.
type IAttributeValueAssertionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EqualityMatch() IEqualityMatchContext
	GreaterThanOrEqualMatch() IGreaterThanOrEqualMatchContext
	LessThanOrEqualMatch() ILessThanOrEqualMatchContext
	ApproximateMatch() IApproximateMatchContext
	ExtensibleMatch() IExtensibleMatchContext

	// IsAttributeValueAssertionContext differentiates from other interfaces.
	IsAttributeValueAssertionContext()
}

type AttributeValueAssertionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeValueAssertionContext() *AttributeValueAssertionContext {
	var p = new(AttributeValueAssertionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_attributeValueAssertion
	return p
}

func InitEmptyAttributeValueAssertionContext(p *AttributeValueAssertionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_attributeValueAssertion
}

func (*AttributeValueAssertionContext) IsAttributeValueAssertionContext() {}

func NewAttributeValueAssertionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeValueAssertionContext {
	var p = new(AttributeValueAssertionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_attributeValueAssertion

	return p
}

func (s *AttributeValueAssertionContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeValueAssertionContext) EqualityMatch() IEqualityMatchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualityMatchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualityMatchContext)
}

func (s *AttributeValueAssertionContext) GreaterThanOrEqualMatch() IGreaterThanOrEqualMatchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGreaterThanOrEqualMatchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGreaterThanOrEqualMatchContext)
}

func (s *AttributeValueAssertionContext) LessThanOrEqualMatch() ILessThanOrEqualMatchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILessThanOrEqualMatchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILessThanOrEqualMatchContext)
}

func (s *AttributeValueAssertionContext) ApproximateMatch() IApproximateMatchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IApproximateMatchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IApproximateMatchContext)
}

func (s *AttributeValueAssertionContext) ExtensibleMatch() IExtensibleMatchContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensibleMatchContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtensibleMatchContext)
}

func (s *AttributeValueAssertionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeValueAssertionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeValueAssertionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterAttributeValueAssertion(s)
	}
}

func (s *AttributeValueAssertionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitAttributeValueAssertion(s)
	}
}

func (p *LDAPParser) AttributeValueAssertion() (localctx IAttributeValueAssertionContext) {
	localctx = NewAttributeValueAssertionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LDAPParserRULE_attributeValueAssertion)
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(116)
			p.EqualityMatch()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.GreaterThanOrEqualMatch()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(118)
			p.LessThanOrEqualMatch()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(119)
			p.ApproximateMatch()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(120)
			p.ExtensibleMatch()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttributeTypeOrMatchingRuleContext is an interface to support dynamic dispatch.
type IAttributeTypeOrMatchingRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode

	// IsAttributeTypeOrMatchingRuleContext differentiates from other interfaces.
	IsAttributeTypeOrMatchingRuleContext()
}

type AttributeTypeOrMatchingRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeTypeOrMatchingRuleContext() *AttributeTypeOrMatchingRuleContext {
	var p = new(AttributeTypeOrMatchingRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_attributeTypeOrMatchingRule
	return p
}

func InitEmptyAttributeTypeOrMatchingRuleContext(p *AttributeTypeOrMatchingRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_attributeTypeOrMatchingRule
}

func (*AttributeTypeOrMatchingRuleContext) IsAttributeTypeOrMatchingRuleContext() {}

func NewAttributeTypeOrMatchingRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeTypeOrMatchingRuleContext {
	var p = new(AttributeTypeOrMatchingRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_attributeTypeOrMatchingRule

	return p
}

func (s *AttributeTypeOrMatchingRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeTypeOrMatchingRuleContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LDAPParserIDENTIFIER, 0)
}

func (s *AttributeTypeOrMatchingRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeTypeOrMatchingRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeTypeOrMatchingRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterAttributeTypeOrMatchingRule(s)
	}
}

func (s *AttributeTypeOrMatchingRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitAttributeTypeOrMatchingRule(s)
	}
}

func (p *LDAPParser) AttributeTypeOrMatchingRule() (localctx IAttributeTypeOrMatchingRuleContext) {
	localctx = NewAttributeTypeOrMatchingRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LDAPParserRULE_attributeTypeOrMatchingRule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(123)
		p.Match(LDAPParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExtensibleMatchContext is an interface to support dynamic dispatch.
type IExtensibleMatchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpeningParenthesis() IOpeningParenthesisContext
	AllEXTENSIBLE_RULE() []antlr.TerminalNode
	EXTENSIBLE_RULE(i int) antlr.TerminalNode
	EQUALITY() antlr.TerminalNode
	ClosingParenthesis() IClosingParenthesisContext
	DistinguishedName() IDistinguishedNameContext
	AttributeValue() IAttributeValueContext
	AllAttributeTypeOrMatchingRule() []IAttributeTypeOrMatchingRuleContext
	AttributeTypeOrMatchingRule(i int) IAttributeTypeOrMatchingRuleContext
	AllDN() []antlr.TerminalNode
	DN(i int) antlr.TerminalNode
	AllObjectIdentifier() []IObjectIdentifierContext
	ObjectIdentifier(i int) IObjectIdentifierContext

	// IsExtensibleMatchContext differentiates from other interfaces.
	IsExtensibleMatchContext()
}

type ExtensibleMatchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensibleMatchContext() *ExtensibleMatchContext {
	var p = new(ExtensibleMatchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_extensibleMatch
	return p
}

func InitEmptyExtensibleMatchContext(p *ExtensibleMatchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_extensibleMatch
}

func (*ExtensibleMatchContext) IsExtensibleMatchContext() {}

func NewExtensibleMatchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensibleMatchContext {
	var p = new(ExtensibleMatchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_extensibleMatch

	return p
}

func (s *ExtensibleMatchContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensibleMatchContext) OpeningParenthesis() IOpeningParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpeningParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpeningParenthesisContext)
}

func (s *ExtensibleMatchContext) AllEXTENSIBLE_RULE() []antlr.TerminalNode {
	return s.GetTokens(LDAPParserEXTENSIBLE_RULE)
}

func (s *ExtensibleMatchContext) EXTENSIBLE_RULE(i int) antlr.TerminalNode {
	return s.GetToken(LDAPParserEXTENSIBLE_RULE, i)
}

func (s *ExtensibleMatchContext) EQUALITY() antlr.TerminalNode {
	return s.GetToken(LDAPParserEQUALITY, 0)
}

func (s *ExtensibleMatchContext) ClosingParenthesis() IClosingParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosingParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosingParenthesisContext)
}

func (s *ExtensibleMatchContext) DistinguishedName() IDistinguishedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDistinguishedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDistinguishedNameContext)
}

func (s *ExtensibleMatchContext) AttributeValue() IAttributeValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeValueContext)
}

func (s *ExtensibleMatchContext) AllAttributeTypeOrMatchingRule() []IAttributeTypeOrMatchingRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeTypeOrMatchingRuleContext); ok {
			len++
		}
	}

	tst := make([]IAttributeTypeOrMatchingRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeTypeOrMatchingRuleContext); ok {
			tst[i] = t.(IAttributeTypeOrMatchingRuleContext)
			i++
		}
	}

	return tst
}

func (s *ExtensibleMatchContext) AttributeTypeOrMatchingRule(i int) IAttributeTypeOrMatchingRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeOrMatchingRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeTypeOrMatchingRuleContext)
}

func (s *ExtensibleMatchContext) AllDN() []antlr.TerminalNode {
	return s.GetTokens(LDAPParserDN)
}

func (s *ExtensibleMatchContext) DN(i int) antlr.TerminalNode {
	return s.GetToken(LDAPParserDN, i)
}

func (s *ExtensibleMatchContext) AllObjectIdentifier() []IObjectIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObjectIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IObjectIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObjectIdentifierContext); ok {
			tst[i] = t.(IObjectIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *ExtensibleMatchContext) ObjectIdentifier(i int) IObjectIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectIdentifierContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectIdentifierContext)
}

func (s *ExtensibleMatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensibleMatchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensibleMatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterExtensibleMatch(s)
	}
}

func (s *ExtensibleMatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitExtensibleMatch(s)
	}
}

func (p *LDAPParser) ExtensibleMatch() (localctx IExtensibleMatchContext) {
	localctx = NewExtensibleMatchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LDAPParserRULE_extensibleMatch)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(125)
		p.OpeningParenthesis()
	}
	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LDAPParserIDENTIFIER {
		{
			p.SetState(126)
			p.AttributeTypeOrMatchingRule()
		}

	}
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(129)
				p.Match(LDAPParserEXTENSIBLE_RULE)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			p.SetState(133)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case LDAPParserDN:
				{
					p.SetState(130)
					p.Match(LDAPParserDN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			case LDAPParserIDENTIFIER:
				{
					p.SetState(131)
					p.AttributeTypeOrMatchingRule()
				}

			case LDAPParserOBJECT_IDENTIFIER:
				{
					p.SetState(132)
					p.ObjectIdentifier()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(140)
		p.Match(LDAPParserEXTENSIBLE_RULE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(LDAPParserEQUALITY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(142)
			p.DistinguishedName()
		}

	case 2:
		{
			p.SetState(143)
			p.AttributeValue()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(146)
		p.ClosingParenthesis()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEqualityMatchContext is an interface to support dynamic dispatch.
type IEqualityMatchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpeningParenthesis() IOpeningParenthesisContext
	AttributeTypeOrMatchingRule() IAttributeTypeOrMatchingRuleContext
	EQUALITY() antlr.TerminalNode
	ClosingParenthesis() IClosingParenthesisContext
	DistinguishedName() IDistinguishedNameContext
	AttributeValue() IAttributeValueContext
	ASTERISK() antlr.TerminalNode

	// IsEqualityMatchContext differentiates from other interfaces.
	IsEqualityMatchContext()
}

type EqualityMatchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualityMatchContext() *EqualityMatchContext {
	var p = new(EqualityMatchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_equalityMatch
	return p
}

func InitEmptyEqualityMatchContext(p *EqualityMatchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_equalityMatch
}

func (*EqualityMatchContext) IsEqualityMatchContext() {}

func NewEqualityMatchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualityMatchContext {
	var p = new(EqualityMatchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_equalityMatch

	return p
}

func (s *EqualityMatchContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualityMatchContext) OpeningParenthesis() IOpeningParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpeningParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpeningParenthesisContext)
}

func (s *EqualityMatchContext) AttributeTypeOrMatchingRule() IAttributeTypeOrMatchingRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeOrMatchingRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeTypeOrMatchingRuleContext)
}

func (s *EqualityMatchContext) EQUALITY() antlr.TerminalNode {
	return s.GetToken(LDAPParserEQUALITY, 0)
}

func (s *EqualityMatchContext) ClosingParenthesis() IClosingParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosingParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosingParenthesisContext)
}

func (s *EqualityMatchContext) DistinguishedName() IDistinguishedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDistinguishedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDistinguishedNameContext)
}

func (s *EqualityMatchContext) AttributeValue() IAttributeValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeValueContext)
}

func (s *EqualityMatchContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(LDAPParserASTERISK, 0)
}

func (s *EqualityMatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualityMatchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualityMatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterEqualityMatch(s)
	}
}

func (s *EqualityMatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitEqualityMatch(s)
	}
}

func (p *LDAPParser) EqualityMatch() (localctx IEqualityMatchContext) {
	localctx = NewEqualityMatchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LDAPParserRULE_equalityMatch)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(148)
		p.OpeningParenthesis()
	}
	{
		p.SetState(149)
		p.AttributeTypeOrMatchingRule()
	}
	{
		p.SetState(150)
		p.Match(LDAPParserEQUALITY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 14, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(151)
			p.DistinguishedName()
		}

	case 2:
		{
			p.SetState(152)
			p.AttributeValue()
		}

	case 3:
		{
			p.SetState(153)
			p.Match(LDAPParserASTERISK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	{
		p.SetState(156)
		p.ClosingParenthesis()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGreaterThanOrEqualMatchContext is an interface to support dynamic dispatch.
type IGreaterThanOrEqualMatchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpeningParenthesis() IOpeningParenthesisContext
	AttributeTypeOrMatchingRule() IAttributeTypeOrMatchingRuleContext
	GREATER_THAN_OR_EQUAL() antlr.TerminalNode
	AttributeValue() IAttributeValueContext
	ClosingParenthesis() IClosingParenthesisContext

	// IsGreaterThanOrEqualMatchContext differentiates from other interfaces.
	IsGreaterThanOrEqualMatchContext()
}

type GreaterThanOrEqualMatchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGreaterThanOrEqualMatchContext() *GreaterThanOrEqualMatchContext {
	var p = new(GreaterThanOrEqualMatchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_greaterThanOrEqualMatch
	return p
}

func InitEmptyGreaterThanOrEqualMatchContext(p *GreaterThanOrEqualMatchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_greaterThanOrEqualMatch
}

func (*GreaterThanOrEqualMatchContext) IsGreaterThanOrEqualMatchContext() {}

func NewGreaterThanOrEqualMatchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GreaterThanOrEqualMatchContext {
	var p = new(GreaterThanOrEqualMatchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_greaterThanOrEqualMatch

	return p
}

func (s *GreaterThanOrEqualMatchContext) GetParser() antlr.Parser { return s.parser }

func (s *GreaterThanOrEqualMatchContext) OpeningParenthesis() IOpeningParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpeningParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpeningParenthesisContext)
}

func (s *GreaterThanOrEqualMatchContext) AttributeTypeOrMatchingRule() IAttributeTypeOrMatchingRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeOrMatchingRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeTypeOrMatchingRuleContext)
}

func (s *GreaterThanOrEqualMatchContext) GREATER_THAN_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(LDAPParserGREATER_THAN_OR_EQUAL, 0)
}

func (s *GreaterThanOrEqualMatchContext) AttributeValue() IAttributeValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeValueContext)
}

func (s *GreaterThanOrEqualMatchContext) ClosingParenthesis() IClosingParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosingParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosingParenthesisContext)
}

func (s *GreaterThanOrEqualMatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanOrEqualMatchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GreaterThanOrEqualMatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterGreaterThanOrEqualMatch(s)
	}
}

func (s *GreaterThanOrEqualMatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitGreaterThanOrEqualMatch(s)
	}
}

func (p *LDAPParser) GreaterThanOrEqualMatch() (localctx IGreaterThanOrEqualMatchContext) {
	localctx = NewGreaterThanOrEqualMatchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LDAPParserRULE_greaterThanOrEqualMatch)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.OpeningParenthesis()
	}
	{
		p.SetState(159)
		p.AttributeTypeOrMatchingRule()
	}
	{
		p.SetState(160)
		p.Match(LDAPParserGREATER_THAN_OR_EQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(161)
		p.AttributeValue()
	}
	{
		p.SetState(162)
		p.ClosingParenthesis()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILessThanOrEqualMatchContext is an interface to support dynamic dispatch.
type ILessThanOrEqualMatchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpeningParenthesis() IOpeningParenthesisContext
	AttributeTypeOrMatchingRule() IAttributeTypeOrMatchingRuleContext
	LESS_THAN_OR_EQUAL() antlr.TerminalNode
	AttributeValue() IAttributeValueContext
	ClosingParenthesis() IClosingParenthesisContext

	// IsLessThanOrEqualMatchContext differentiates from other interfaces.
	IsLessThanOrEqualMatchContext()
}

type LessThanOrEqualMatchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLessThanOrEqualMatchContext() *LessThanOrEqualMatchContext {
	var p = new(LessThanOrEqualMatchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_lessThanOrEqualMatch
	return p
}

func InitEmptyLessThanOrEqualMatchContext(p *LessThanOrEqualMatchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_lessThanOrEqualMatch
}

func (*LessThanOrEqualMatchContext) IsLessThanOrEqualMatchContext() {}

func NewLessThanOrEqualMatchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LessThanOrEqualMatchContext {
	var p = new(LessThanOrEqualMatchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_lessThanOrEqualMatch

	return p
}

func (s *LessThanOrEqualMatchContext) GetParser() antlr.Parser { return s.parser }

func (s *LessThanOrEqualMatchContext) OpeningParenthesis() IOpeningParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpeningParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpeningParenthesisContext)
}

func (s *LessThanOrEqualMatchContext) AttributeTypeOrMatchingRule() IAttributeTypeOrMatchingRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeOrMatchingRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeTypeOrMatchingRuleContext)
}

func (s *LessThanOrEqualMatchContext) LESS_THAN_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(LDAPParserLESS_THAN_OR_EQUAL, 0)
}

func (s *LessThanOrEqualMatchContext) AttributeValue() IAttributeValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeValueContext)
}

func (s *LessThanOrEqualMatchContext) ClosingParenthesis() IClosingParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosingParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosingParenthesisContext)
}

func (s *LessThanOrEqualMatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanOrEqualMatchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LessThanOrEqualMatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterLessThanOrEqualMatch(s)
	}
}

func (s *LessThanOrEqualMatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitLessThanOrEqualMatch(s)
	}
}

func (p *LDAPParser) LessThanOrEqualMatch() (localctx ILessThanOrEqualMatchContext) {
	localctx = NewLessThanOrEqualMatchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LDAPParserRULE_lessThanOrEqualMatch)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.OpeningParenthesis()
	}
	{
		p.SetState(165)
		p.AttributeTypeOrMatchingRule()
	}
	{
		p.SetState(166)
		p.Match(LDAPParserLESS_THAN_OR_EQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.AttributeValue()
	}
	{
		p.SetState(168)
		p.ClosingParenthesis()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IApproximateMatchContext is an interface to support dynamic dispatch.
type IApproximateMatchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpeningParenthesis() IOpeningParenthesisContext
	AttributeTypeOrMatchingRule() IAttributeTypeOrMatchingRuleContext
	APPROXIMATE() antlr.TerminalNode
	AttributeValue() IAttributeValueContext
	ClosingParenthesis() IClosingParenthesisContext

	// IsApproximateMatchContext differentiates from other interfaces.
	IsApproximateMatchContext()
}

type ApproximateMatchContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyApproximateMatchContext() *ApproximateMatchContext {
	var p = new(ApproximateMatchContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_approximateMatch
	return p
}

func InitEmptyApproximateMatchContext(p *ApproximateMatchContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_approximateMatch
}

func (*ApproximateMatchContext) IsApproximateMatchContext() {}

func NewApproximateMatchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ApproximateMatchContext {
	var p = new(ApproximateMatchContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_approximateMatch

	return p
}

func (s *ApproximateMatchContext) GetParser() antlr.Parser { return s.parser }

func (s *ApproximateMatchContext) OpeningParenthesis() IOpeningParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOpeningParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOpeningParenthesisContext)
}

func (s *ApproximateMatchContext) AttributeTypeOrMatchingRule() IAttributeTypeOrMatchingRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeOrMatchingRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeTypeOrMatchingRuleContext)
}

func (s *ApproximateMatchContext) APPROXIMATE() antlr.TerminalNode {
	return s.GetToken(LDAPParserAPPROXIMATE, 0)
}

func (s *ApproximateMatchContext) AttributeValue() IAttributeValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeValueContext)
}

func (s *ApproximateMatchContext) ClosingParenthesis() IClosingParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClosingParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClosingParenthesisContext)
}

func (s *ApproximateMatchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ApproximateMatchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ApproximateMatchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterApproximateMatch(s)
	}
}

func (s *ApproximateMatchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitApproximateMatch(s)
	}
}

func (p *LDAPParser) ApproximateMatch() (localctx IApproximateMatchContext) {
	localctx = NewApproximateMatchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LDAPParserRULE_approximateMatch)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.OpeningParenthesis()
	}
	{
		p.SetState(171)
		p.AttributeTypeOrMatchingRule()
	}
	{
		p.SetState(172)
		p.Match(LDAPParserAPPROXIMATE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.AttributeValue()
	}
	{
		p.SetState(174)
		p.ClosingParenthesis()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectIdentifierContext is an interface to support dynamic dispatch.
type IObjectIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OBJECT_IDENTIFIER() antlr.TerminalNode

	// IsObjectIdentifierContext differentiates from other interfaces.
	IsObjectIdentifierContext()
}

type ObjectIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectIdentifierContext() *ObjectIdentifierContext {
	var p = new(ObjectIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_objectIdentifier
	return p
}

func InitEmptyObjectIdentifierContext(p *ObjectIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_objectIdentifier
}

func (*ObjectIdentifierContext) IsObjectIdentifierContext() {}

func NewObjectIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectIdentifierContext {
	var p = new(ObjectIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_objectIdentifier

	return p
}

func (s *ObjectIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectIdentifierContext) OBJECT_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(LDAPParserOBJECT_IDENTIFIER, 0)
}

func (s *ObjectIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterObjectIdentifier(s)
	}
}

func (s *ObjectIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitObjectIdentifier(s)
	}
}

func (p *LDAPParser) ObjectIdentifier() (localctx IObjectIdentifierContext) {
	localctx = NewObjectIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LDAPParserRULE_objectIdentifier)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(176)
		p.Match(LDAPParserOBJECT_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOpeningParenthesisContext is an interface to support dynamic dispatch.
type IOpeningParenthesisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LPAREN() antlr.TerminalNode

	// IsOpeningParenthesisContext differentiates from other interfaces.
	IsOpeningParenthesisContext()
}

type OpeningParenthesisContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOpeningParenthesisContext() *OpeningParenthesisContext {
	var p = new(OpeningParenthesisContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_openingParenthesis
	return p
}

func InitEmptyOpeningParenthesisContext(p *OpeningParenthesisContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_openingParenthesis
}

func (*OpeningParenthesisContext) IsOpeningParenthesisContext() {}

func NewOpeningParenthesisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpeningParenthesisContext {
	var p = new(OpeningParenthesisContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_openingParenthesis

	return p
}

func (s *OpeningParenthesisContext) GetParser() antlr.Parser { return s.parser }

func (s *OpeningParenthesisContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LDAPParserLPAREN, 0)
}

func (s *OpeningParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpeningParenthesisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpeningParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterOpeningParenthesis(s)
	}
}

func (s *OpeningParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitOpeningParenthesis(s)
	}
}

func (p *LDAPParser) OpeningParenthesis() (localctx IOpeningParenthesisContext) {
	localctx = NewOpeningParenthesisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LDAPParserRULE_openingParenthesis)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
		p.Match(LDAPParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClosingParenthesisContext is an interface to support dynamic dispatch.
type IClosingParenthesisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RPAREN() antlr.TerminalNode

	// IsClosingParenthesisContext differentiates from other interfaces.
	IsClosingParenthesisContext()
}

type ClosingParenthesisContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClosingParenthesisContext() *ClosingParenthesisContext {
	var p = new(ClosingParenthesisContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_closingParenthesis
	return p
}

func InitEmptyClosingParenthesisContext(p *ClosingParenthesisContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_closingParenthesis
}

func (*ClosingParenthesisContext) IsClosingParenthesisContext() {}

func NewClosingParenthesisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClosingParenthesisContext {
	var p = new(ClosingParenthesisContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_closingParenthesis

	return p
}

func (s *ClosingParenthesisContext) GetParser() antlr.Parser { return s.parser }

func (s *ClosingParenthesisContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LDAPParserRPAREN, 0)
}

func (s *ClosingParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClosingParenthesisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClosingParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterClosingParenthesis(s)
	}
}

func (s *ClosingParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitClosingParenthesis(s)
	}
}

func (p *LDAPParser) ClosingParenthesis() (localctx IClosingParenthesisContext) {
	localctx = NewClosingParenthesisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LDAPParserRULE_closingParenthesis)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(LDAPParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDistinguishedNameContext is an interface to support dynamic dispatch.
type IDistinguishedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAttributeTypeOrMatchingRule() []IAttributeTypeOrMatchingRuleContext
	AttributeTypeOrMatchingRule(i int) IAttributeTypeOrMatchingRuleContext
	AllEQUALITY() []antlr.TerminalNode
	EQUALITY(i int) antlr.TerminalNode
	AllAttributeValue() []IAttributeValueContext
	AttributeValue(i int) IAttributeValueContext
	LOCAL_LDAP_SCHEME() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsDistinguishedNameContext differentiates from other interfaces.
	IsDistinguishedNameContext()
}

type DistinguishedNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDistinguishedNameContext() *DistinguishedNameContext {
	var p = new(DistinguishedNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_distinguishedName
	return p
}

func InitEmptyDistinguishedNameContext(p *DistinguishedNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_distinguishedName
}

func (*DistinguishedNameContext) IsDistinguishedNameContext() {}

func NewDistinguishedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistinguishedNameContext {
	var p = new(DistinguishedNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_distinguishedName

	return p
}

func (s *DistinguishedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DistinguishedNameContext) AllAttributeTypeOrMatchingRule() []IAttributeTypeOrMatchingRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeTypeOrMatchingRuleContext); ok {
			len++
		}
	}

	tst := make([]IAttributeTypeOrMatchingRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeTypeOrMatchingRuleContext); ok {
			tst[i] = t.(IAttributeTypeOrMatchingRuleContext)
			i++
		}
	}

	return tst
}

func (s *DistinguishedNameContext) AttributeTypeOrMatchingRule(i int) IAttributeTypeOrMatchingRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeOrMatchingRuleContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeTypeOrMatchingRuleContext)
}

func (s *DistinguishedNameContext) AllEQUALITY() []antlr.TerminalNode {
	return s.GetTokens(LDAPParserEQUALITY)
}

func (s *DistinguishedNameContext) EQUALITY(i int) antlr.TerminalNode {
	return s.GetToken(LDAPParserEQUALITY, i)
}

func (s *DistinguishedNameContext) AllAttributeValue() []IAttributeValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeValueContext); ok {
			len++
		}
	}

	tst := make([]IAttributeValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeValueContext); ok {
			tst[i] = t.(IAttributeValueContext)
			i++
		}
	}

	return tst
}

func (s *DistinguishedNameContext) AttributeValue(i int) IAttributeValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeValueContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeValueContext)
}

func (s *DistinguishedNameContext) LOCAL_LDAP_SCHEME() antlr.TerminalNode {
	return s.GetToken(LDAPParserLOCAL_LDAP_SCHEME, 0)
}

func (s *DistinguishedNameContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(LDAPParserCOMMA)
}

func (s *DistinguishedNameContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(LDAPParserCOMMA, i)
}

func (s *DistinguishedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistinguishedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DistinguishedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterDistinguishedName(s)
	}
}

func (s *DistinguishedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitDistinguishedName(s)
	}
}

func (p *LDAPParser) DistinguishedName() (localctx IDistinguishedNameContext) {
	localctx = NewDistinguishedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LDAPParserRULE_distinguishedName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(183)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == LDAPParserLOCAL_LDAP_SCHEME {
		{
			p.SetState(182)
			p.Match(LDAPParserLOCAL_LDAP_SCHEME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

	{
		p.SetState(185)
		p.AttributeTypeOrMatchingRule()
	}
	{
		p.SetState(186)
		p.Match(LDAPParserEQUALITY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.AttributeValue()
	}

	p.SetState(196)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LDAPParserCOMMA {
		{
			p.SetState(189)
			p.Match(LDAPParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

		{
			p.SetState(190)
			p.AttributeTypeOrMatchingRule()
		}
		{
			p.SetState(191)
			p.Match(LDAPParserEQUALITY)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(192)
			p.AttributeValue()
		}

		p.SetState(198)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAttributeValueContext is an interface to support dynamic dispatch.
type IAttributeValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAttributeValueContext differentiates from other interfaces.
	IsAttributeValueContext()
}

type AttributeValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeValueContext() *AttributeValueContext {
	var p = new(AttributeValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_attributeValue
	return p
}

func InitEmptyAttributeValueContext(p *AttributeValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LDAPParserRULE_attributeValue
}

func (*AttributeValueContext) IsAttributeValueContext() {}

func NewAttributeValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeValueContext {
	var p = new(AttributeValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDAPParserRULE_attributeValue

	return p
}

func (s *AttributeValueContext) GetParser() antlr.Parser { return s.parser }
func (s *AttributeValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.EnterAttributeValue(s)
	}
}

func (s *AttributeValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDAPParserListener); ok {
		listenerT.ExitAttributeValue(s)
	}
}

func (p *LDAPParser) AttributeValue() (localctx IAttributeValueContext) {
	localctx = NewAttributeValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LDAPParserRULE_attributeValue)
	var _alt int

	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1 + 1
		for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1 + 1:
				p.SetState(199)
				p.MatchWildcard()

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(202)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(204)
		p.MatchWildcard()

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
