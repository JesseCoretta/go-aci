// Code generated from ACIParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package aciparser // ACIParser
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

type ACIParser struct {
	*antlr.BaseParser
}

var ACIParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func aciparserParserInit() {
	staticData := &ACIParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'userdn'", "'groupdn'", "'roledn'", "'userattr'", "'groupattr'",
		"'ssf'", "'ip'", "'dns'", "'authmethod'", "'timeofday'", "'dayofweek'",
		"'target'", "'target_to'", "'target_from'", "'targetscope'", "'targetattr'",
		"'targetfilter'", "'targattrfilters'", "'targetcontrol'", "'extop'",
		"'parent'", "'\"'", "'('", "')'", "','", "';'", "'='", "'!='", "'<'",
		"'<='", "'>'", "'>='", "'&&'", "'||'", "", "'read'", "'write'", "'add'",
		"'compare'", "'search'", "'delete'", "'proxy'", "'export'", "'import'",
		"'selfwrite'", "'all'", "'none'", "'allow'", "'deny'",
	}
	staticData.SymbolicNames = []string{
		"", "BKW_UDN", "BKW_GDN", "BKW_RDN", "BKW_UAT", "BKW_GAT", "BKW_SSF",
		"BKW_IP", "BKW_DNS", "BKW_AM", "BKW_TOD", "BKW_DOW", "TKW_TARGET", "TKW_TO",
		"TKW_FROM", "TKW_SCOPE", "TKW_ATTR", "TKW_FILTER", "TKW_AF", "TKW_CTRL",
		"TKW_EXTOP", "PARENT", "DQUOTE", "LPAREN", "RPAREN", "COMMA", "SEMI",
		"EQ", "NE", "LT", "LE", "GT", "GE", "SYMBOLIC_AND", "SYMBOLIC_OR", "ANCHOR",
		"READ_PRIV", "WRITE_PRIV", "ADD_PRIV", "CMP_PRIV", "SRC_PRIV", "DEL_PRIV",
		"PRX_PRIV", "EXP_PRIV", "IMP_PRIV", "SLF_PRIV", "ALL_PRIV", "NO_PRIV",
		"ALLOW", "DENY", "WORD_AND", "WORD_OR", "WORD_NOT", "NCTF_TO_WHSP",
		"QUOTED_STRING", "WS",
	}
	staticData.RuleNames = []string{
		"parse", "instruction", "permissionBindRule", "permission", "privilege",
		"targetRule", "targetOperator", "targetRules", "bindRule", "wordAnd",
		"wordOr", "wordNot", "bindRules", "quotedValue", "bindOperator", "lessThan",
		"lessThanOrEqual", "greaterThan", "greaterThanOrEqual", "equalTo", "notEqualTo",
		"bindKeyword", "targetKeyword", "openingParenthesis", "closingParenthesis",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 55, 177, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 1, 0, 1, 0, 1, 0,
		1, 1, 3, 1, 55, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 4, 1, 62, 8, 1, 11,
		1, 12, 1, 63, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 5, 3, 77, 8, 3, 10, 3, 12, 3, 80, 9, 3, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 3, 6, 94, 8, 6, 1, 7, 4,
		7, 97, 8, 7, 11, 7, 12, 7, 98, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 3, 8, 111, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1,
		11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 127,
		8, 12, 1, 12, 1, 12, 4, 12, 131, 8, 12, 11, 12, 12, 12, 132, 1, 12, 3,
		12, 136, 8, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 142, 8, 12, 10, 12,
		12, 12, 145, 9, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 3, 14, 155, 8, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1,
		23, 1, 24, 1, 24, 1, 24, 0, 1, 24, 25, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 0, 4, 1, 0,
		48, 49, 1, 0, 36, 47, 1, 0, 1, 11, 1, 0, 12, 20, 167, 0, 50, 1, 0, 0, 0,
		2, 54, 1, 0, 0, 0, 4, 67, 1, 0, 0, 0, 6, 71, 1, 0, 0, 0, 8, 83, 1, 0, 0,
		0, 10, 85, 1, 0, 0, 0, 12, 93, 1, 0, 0, 0, 14, 96, 1, 0, 0, 0, 16, 110,
		1, 0, 0, 0, 18, 112, 1, 0, 0, 0, 20, 114, 1, 0, 0, 0, 22, 116, 1, 0, 0,
		0, 24, 135, 1, 0, 0, 0, 26, 146, 1, 0, 0, 0, 28, 154, 1, 0, 0, 0, 30, 156,
		1, 0, 0, 0, 32, 158, 1, 0, 0, 0, 34, 160, 1, 0, 0, 0, 36, 162, 1, 0, 0,
		0, 38, 164, 1, 0, 0, 0, 40, 166, 1, 0, 0, 0, 42, 168, 1, 0, 0, 0, 44, 170,
		1, 0, 0, 0, 46, 172, 1, 0, 0, 0, 48, 174, 1, 0, 0, 0, 50, 51, 3, 2, 1,
		0, 51, 52, 5, 0, 0, 1, 52, 1, 1, 0, 0, 0, 53, 55, 3, 14, 7, 0, 54, 53,
		1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0, 56, 57, 3, 46, 23,
		0, 57, 58, 5, 35, 0, 0, 58, 59, 3, 26, 13, 0, 59, 61, 5, 26, 0, 0, 60,
		62, 3, 4, 2, 0, 61, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 61, 1, 0, 0,
		0, 63, 64, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 66, 3, 48, 24, 0, 66, 3,
		1, 0, 0, 0, 67, 68, 3, 6, 3, 0, 68, 69, 3, 24, 12, 0, 69, 70, 5, 26, 0,
		0, 70, 5, 1, 0, 0, 0, 71, 72, 7, 0, 0, 0, 72, 73, 3, 46, 23, 0, 73, 78,
		3, 8, 4, 0, 74, 75, 5, 25, 0, 0, 75, 77, 3, 8, 4, 0, 76, 74, 1, 0, 0, 0,
		77, 80, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 81, 1,
		0, 0, 0, 80, 78, 1, 0, 0, 0, 81, 82, 3, 48, 24, 0, 82, 7, 1, 0, 0, 0, 83,
		84, 7, 1, 0, 0, 84, 9, 1, 0, 0, 0, 85, 86, 3, 46, 23, 0, 86, 87, 3, 44,
		22, 0, 87, 88, 3, 12, 6, 0, 88, 89, 3, 26, 13, 0, 89, 90, 3, 48, 24, 0,
		90, 11, 1, 0, 0, 0, 91, 94, 3, 38, 19, 0, 92, 94, 3, 40, 20, 0, 93, 91,
		1, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 13, 1, 0, 0, 0, 95, 97, 3, 10, 5, 0,
		96, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1,
		0, 0, 0, 99, 15, 1, 0, 0, 0, 100, 101, 3, 42, 21, 0, 101, 102, 3, 28, 14,
		0, 102, 103, 3, 26, 13, 0, 103, 111, 1, 0, 0, 0, 104, 105, 3, 46, 23, 0,
		105, 106, 3, 42, 21, 0, 106, 107, 3, 28, 14, 0, 107, 108, 3, 26, 13, 0,
		108, 109, 3, 48, 24, 0, 109, 111, 1, 0, 0, 0, 110, 100, 1, 0, 0, 0, 110,
		104, 1, 0, 0, 0, 111, 17, 1, 0, 0, 0, 112, 113, 5, 50, 0, 0, 113, 19, 1,
		0, 0, 0, 114, 115, 5, 51, 0, 0, 115, 21, 1, 0, 0, 0, 116, 117, 5, 52, 0,
		0, 117, 23, 1, 0, 0, 0, 118, 119, 6, 12, -1, 0, 119, 120, 3, 46, 23, 0,
		120, 121, 3, 24, 12, 0, 121, 122, 3, 48, 24, 0, 122, 136, 1, 0, 0, 0, 123,
		130, 3, 16, 8, 0, 124, 127, 3, 18, 9, 0, 125, 127, 3, 20, 10, 0, 126, 124,
		1, 0, 0, 0, 126, 125, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128, 129, 3, 16,
		8, 0, 129, 131, 1, 0, 0, 0, 130, 126, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0,
		132, 130, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 136, 1, 0, 0, 0, 134,
		136, 3, 16, 8, 0, 135, 118, 1, 0, 0, 0, 135, 123, 1, 0, 0, 0, 135, 134,
		1, 0, 0, 0, 136, 143, 1, 0, 0, 0, 137, 138, 10, 3, 0, 0, 138, 139, 3, 22,
		11, 0, 139, 140, 3, 24, 12, 3, 140, 142, 1, 0, 0, 0, 141, 137, 1, 0, 0,
		0, 142, 145, 1, 0, 0, 0, 143, 141, 1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144,
		25, 1, 0, 0, 0, 145, 143, 1, 0, 0, 0, 146, 147, 5, 54, 0, 0, 147, 27, 1,
		0, 0, 0, 148, 155, 3, 38, 19, 0, 149, 155, 3, 40, 20, 0, 150, 155, 3, 30,
		15, 0, 151, 155, 3, 32, 16, 0, 152, 155, 3, 34, 17, 0, 153, 155, 3, 36,
		18, 0, 154, 148, 1, 0, 0, 0, 154, 149, 1, 0, 0, 0, 154, 150, 1, 0, 0, 0,
		154, 151, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 153, 1, 0, 0, 0, 155,
		29, 1, 0, 0, 0, 156, 157, 5, 29, 0, 0, 157, 31, 1, 0, 0, 0, 158, 159, 5,
		30, 0, 0, 159, 33, 1, 0, 0, 0, 160, 161, 5, 31, 0, 0, 161, 35, 1, 0, 0,
		0, 162, 163, 5, 32, 0, 0, 163, 37, 1, 0, 0, 0, 164, 165, 5, 27, 0, 0, 165,
		39, 1, 0, 0, 0, 166, 167, 5, 28, 0, 0, 167, 41, 1, 0, 0, 0, 168, 169, 7,
		2, 0, 0, 169, 43, 1, 0, 0, 0, 170, 171, 7, 3, 0, 0, 171, 45, 1, 0, 0, 0,
		172, 173, 5, 23, 0, 0, 173, 47, 1, 0, 0, 0, 174, 175, 5, 24, 0, 0, 175,
		49, 1, 0, 0, 0, 11, 54, 63, 78, 93, 98, 110, 126, 132, 135, 143, 154,
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

// ACIParserInit initializes any static state used to implement ACIParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewACIParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ACIParserInit() {
	staticData := &ACIParserParserStaticData
	staticData.once.Do(aciparserParserInit)
}

// NewACIParser produces a new parser instance for the optional input antlr.TokenStream.
func NewACIParser(input antlr.TokenStream) *ACIParser {
	ACIParserInit()
	this := new(ACIParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ACIParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ACIParser.g4"

	return this
}

// ACIParser tokens.
const (
	ACIParserEOF           = antlr.TokenEOF
	ACIParserBKW_UDN       = 1
	ACIParserBKW_GDN       = 2
	ACIParserBKW_RDN       = 3
	ACIParserBKW_UAT       = 4
	ACIParserBKW_GAT       = 5
	ACIParserBKW_SSF       = 6
	ACIParserBKW_IP        = 7
	ACIParserBKW_DNS       = 8
	ACIParserBKW_AM        = 9
	ACIParserBKW_TOD       = 10
	ACIParserBKW_DOW       = 11
	ACIParserTKW_TARGET    = 12
	ACIParserTKW_TO        = 13
	ACIParserTKW_FROM      = 14
	ACIParserTKW_SCOPE     = 15
	ACIParserTKW_ATTR      = 16
	ACIParserTKW_FILTER    = 17
	ACIParserTKW_AF        = 18
	ACIParserTKW_CTRL      = 19
	ACIParserTKW_EXTOP     = 20
	ACIParserPARENT        = 21
	ACIParserDQUOTE        = 22
	ACIParserLPAREN        = 23
	ACIParserRPAREN        = 24
	ACIParserCOMMA         = 25
	ACIParserSEMI          = 26
	ACIParserEQ            = 27
	ACIParserNE            = 28
	ACIParserLT            = 29
	ACIParserLE            = 30
	ACIParserGT            = 31
	ACIParserGE            = 32
	ACIParserSYMBOLIC_AND  = 33
	ACIParserSYMBOLIC_OR   = 34
	ACIParserANCHOR        = 35
	ACIParserREAD_PRIV     = 36
	ACIParserWRITE_PRIV    = 37
	ACIParserADD_PRIV      = 38
	ACIParserCMP_PRIV      = 39
	ACIParserSRC_PRIV      = 40
	ACIParserDEL_PRIV      = 41
	ACIParserPRX_PRIV      = 42
	ACIParserEXP_PRIV      = 43
	ACIParserIMP_PRIV      = 44
	ACIParserSLF_PRIV      = 45
	ACIParserALL_PRIV      = 46
	ACIParserNO_PRIV       = 47
	ACIParserALLOW         = 48
	ACIParserDENY          = 49
	ACIParserWORD_AND      = 50
	ACIParserWORD_OR       = 51
	ACIParserWORD_NOT      = 52
	ACIParserNCTF_TO_WHSP  = 53
	ACIParserQUOTED_STRING = 54
	ACIParserWS            = 55
)

// ACIParser rules.
const (
	ACIParserRULE_parse              = 0
	ACIParserRULE_instruction        = 1
	ACIParserRULE_permissionBindRule = 2
	ACIParserRULE_permission         = 3
	ACIParserRULE_privilege          = 4
	ACIParserRULE_targetRule         = 5
	ACIParserRULE_targetOperator     = 6
	ACIParserRULE_targetRules        = 7
	ACIParserRULE_bindRule           = 8
	ACIParserRULE_wordAnd            = 9
	ACIParserRULE_wordOr             = 10
	ACIParserRULE_wordNot            = 11
	ACIParserRULE_bindRules          = 12
	ACIParserRULE_quotedValue        = 13
	ACIParserRULE_bindOperator       = 14
	ACIParserRULE_lessThan           = 15
	ACIParserRULE_lessThanOrEqual    = 16
	ACIParserRULE_greaterThan        = 17
	ACIParserRULE_greaterThanOrEqual = 18
	ACIParserRULE_equalTo            = 19
	ACIParserRULE_notEqualTo         = 20
	ACIParserRULE_bindKeyword        = 21
	ACIParserRULE_targetKeyword      = 22
	ACIParserRULE_openingParenthesis = 23
	ACIParserRULE_closingParenthesis = 24
)

// IParseContext is an interface to support dynamic dispatch.
type IParseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Instruction() IInstructionContext
	EOF() antlr.TerminalNode

	// IsParseContext differentiates from other interfaces.
	IsParseContext()
}

type ParseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParseContext() *ParseContext {
	var p = new(ParseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_parse
	return p
}

func InitEmptyParseContext(p *ParseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_parse
}

func (*ParseContext) IsParseContext() {}

func NewParseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParseContext {
	var p = new(ParseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_parse

	return p
}

func (s *ParseContext) GetParser() antlr.Parser { return s.parser }

func (s *ParseContext) Instruction() IInstructionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *ParseContext) EOF() antlr.TerminalNode {
	return s.GetToken(ACIParserEOF, 0)
}

func (s *ParseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParse(s)
	}
}

func (s *ParseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParse(s)
	}
}

func (p *ACIParser) Parse() (localctx IParseContext) {
	localctx = NewParseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ACIParserRULE_parse)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Instruction()
	}
	{
		p.SetState(51)
		p.Match(ACIParserEOF)
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

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpeningParenthesis() IOpeningParenthesisContext
	ANCHOR() antlr.TerminalNode
	QuotedValue() IQuotedValueContext
	SEMI() antlr.TerminalNode
	ClosingParenthesis() IClosingParenthesisContext
	TargetRules() ITargetRulesContext
	AllPermissionBindRule() []IPermissionBindRuleContext
	PermissionBindRule(i int) IPermissionBindRuleContext

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_instruction
	return p
}

func InitEmptyInstructionContext(p *InstructionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_instruction
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) OpeningParenthesis() IOpeningParenthesisContext {
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

func (s *InstructionContext) ANCHOR() antlr.TerminalNode {
	return s.GetToken(ACIParserANCHOR, 0)
}

func (s *InstructionContext) QuotedValue() IQuotedValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuotedValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuotedValueContext)
}

func (s *InstructionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ACIParserSEMI, 0)
}

func (s *InstructionContext) ClosingParenthesis() IClosingParenthesisContext {
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

func (s *InstructionContext) TargetRules() ITargetRulesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetRulesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetRulesContext)
}

func (s *InstructionContext) AllPermissionBindRule() []IPermissionBindRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPermissionBindRuleContext); ok {
			len++
		}
	}

	tst := make([]IPermissionBindRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPermissionBindRuleContext); ok {
			tst[i] = t.(IPermissionBindRuleContext)
			i++
		}
	}

	return tst
}

func (s *InstructionContext) PermissionBindRule(i int) IPermissionBindRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPermissionBindRuleContext); ok {
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

	return t.(IPermissionBindRuleContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *ACIParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ACIParserRULE_instruction)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(54)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(53)
			p.TargetRules()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(56)
		p.OpeningParenthesis()
	}
	{
		p.SetState(57)
		p.Match(ACIParserANCHOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.QuotedValue()
	}
	{
		p.SetState(59)
		p.Match(ACIParserSEMI)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(61)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ACIParserALLOW || _la == ACIParserDENY {
		{
			p.SetState(60)
			p.PermissionBindRule()
		}

		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(65)
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

// IPermissionBindRuleContext is an interface to support dynamic dispatch.
type IPermissionBindRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Permission() IPermissionContext
	BindRules() IBindRulesContext
	SEMI() antlr.TerminalNode

	// IsPermissionBindRuleContext differentiates from other interfaces.
	IsPermissionBindRuleContext()
}

type PermissionBindRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPermissionBindRuleContext() *PermissionBindRuleContext {
	var p = new(PermissionBindRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_permissionBindRule
	return p
}

func InitEmptyPermissionBindRuleContext(p *PermissionBindRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_permissionBindRule
}

func (*PermissionBindRuleContext) IsPermissionBindRuleContext() {}

func NewPermissionBindRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PermissionBindRuleContext {
	var p = new(PermissionBindRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_permissionBindRule

	return p
}

func (s *PermissionBindRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *PermissionBindRuleContext) Permission() IPermissionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPermissionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPermissionContext)
}

func (s *PermissionBindRuleContext) BindRules() IBindRulesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRulesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRulesContext)
}

func (s *PermissionBindRuleContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ACIParserSEMI, 0)
}

func (s *PermissionBindRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PermissionBindRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PermissionBindRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterPermissionBindRule(s)
	}
}

func (s *PermissionBindRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitPermissionBindRule(s)
	}
}

func (p *ACIParser) PermissionBindRule() (localctx IPermissionBindRuleContext) {
	localctx = NewPermissionBindRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ACIParserRULE_permissionBindRule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Permission()
	}
	{
		p.SetState(68)
		p.bindRules(0)
	}
	{
		p.SetState(69)
		p.Match(ACIParserSEMI)
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

// IPermissionContext is an interface to support dynamic dispatch.
type IPermissionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpeningParenthesis() IOpeningParenthesisContext
	AllPrivilege() []IPrivilegeContext
	Privilege(i int) IPrivilegeContext
	ClosingParenthesis() IClosingParenthesisContext
	ALLOW() antlr.TerminalNode
	DENY() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsPermissionContext differentiates from other interfaces.
	IsPermissionContext()
}

type PermissionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPermissionContext() *PermissionContext {
	var p = new(PermissionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_permission
	return p
}

func InitEmptyPermissionContext(p *PermissionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_permission
}

func (*PermissionContext) IsPermissionContext() {}

func NewPermissionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PermissionContext {
	var p = new(PermissionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_permission

	return p
}

func (s *PermissionContext) GetParser() antlr.Parser { return s.parser }

func (s *PermissionContext) OpeningParenthesis() IOpeningParenthesisContext {
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

func (s *PermissionContext) AllPrivilege() []IPrivilegeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrivilegeContext); ok {
			len++
		}
	}

	tst := make([]IPrivilegeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrivilegeContext); ok {
			tst[i] = t.(IPrivilegeContext)
			i++
		}
	}

	return tst
}

func (s *PermissionContext) Privilege(i int) IPrivilegeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrivilegeContext); ok {
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

	return t.(IPrivilegeContext)
}

func (s *PermissionContext) ClosingParenthesis() IClosingParenthesisContext {
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

func (s *PermissionContext) ALLOW() antlr.TerminalNode {
	return s.GetToken(ACIParserALLOW, 0)
}

func (s *PermissionContext) DENY() antlr.TerminalNode {
	return s.GetToken(ACIParserDENY, 0)
}

func (s *PermissionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ACIParserCOMMA)
}

func (s *PermissionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserCOMMA, i)
}

func (s *PermissionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PermissionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PermissionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterPermission(s)
	}
}

func (s *PermissionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitPermission(s)
	}
}

func (p *ACIParser) Permission() (localctx IPermissionContext) {
	localctx = NewPermissionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ACIParserRULE_permission)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ACIParserALLOW || _la == ACIParserDENY) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(72)
		p.OpeningParenthesis()
	}
	{
		p.SetState(73)
		p.Privilege()
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ACIParserCOMMA {
		{
			p.SetState(74)
			p.Match(ACIParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(75)
			p.Privilege()
		}

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(81)
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

// IPrivilegeContext is an interface to support dynamic dispatch.
type IPrivilegeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	READ_PRIV() antlr.TerminalNode
	WRITE_PRIV() antlr.TerminalNode
	SLF_PRIV() antlr.TerminalNode
	CMP_PRIV() antlr.TerminalNode
	SRC_PRIV() antlr.TerminalNode
	PRX_PRIV() antlr.TerminalNode
	ADD_PRIV() antlr.TerminalNode
	DEL_PRIV() antlr.TerminalNode
	IMP_PRIV() antlr.TerminalNode
	EXP_PRIV() antlr.TerminalNode
	ALL_PRIV() antlr.TerminalNode
	NO_PRIV() antlr.TerminalNode

	// IsPrivilegeContext differentiates from other interfaces.
	IsPrivilegeContext()
}

type PrivilegeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrivilegeContext() *PrivilegeContext {
	var p = new(PrivilegeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_privilege
	return p
}

func InitEmptyPrivilegeContext(p *PrivilegeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_privilege
}

func (*PrivilegeContext) IsPrivilegeContext() {}

func NewPrivilegeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrivilegeContext {
	var p = new(PrivilegeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_privilege

	return p
}

func (s *PrivilegeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrivilegeContext) READ_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserREAD_PRIV, 0)
}

func (s *PrivilegeContext) WRITE_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserWRITE_PRIV, 0)
}

func (s *PrivilegeContext) SLF_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserSLF_PRIV, 0)
}

func (s *PrivilegeContext) CMP_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserCMP_PRIV, 0)
}

func (s *PrivilegeContext) SRC_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserSRC_PRIV, 0)
}

func (s *PrivilegeContext) PRX_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserPRX_PRIV, 0)
}

func (s *PrivilegeContext) ADD_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserADD_PRIV, 0)
}

func (s *PrivilegeContext) DEL_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserDEL_PRIV, 0)
}

func (s *PrivilegeContext) IMP_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserIMP_PRIV, 0)
}

func (s *PrivilegeContext) EXP_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserEXP_PRIV, 0)
}

func (s *PrivilegeContext) ALL_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserALL_PRIV, 0)
}

func (s *PrivilegeContext) NO_PRIV() antlr.TerminalNode {
	return s.GetToken(ACIParserNO_PRIV, 0)
}

func (s *PrivilegeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrivilegeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrivilegeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterPrivilege(s)
	}
}

func (s *PrivilegeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitPrivilege(s)
	}
}

func (p *ACIParser) Privilege() (localctx IPrivilegeContext) {
	localctx = NewPrivilegeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ACIParserRULE_privilege)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&281406257233920) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// ITargetRuleContext is an interface to support dynamic dispatch.
type ITargetRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpeningParenthesis() IOpeningParenthesisContext
	TargetKeyword() ITargetKeywordContext
	TargetOperator() ITargetOperatorContext
	QuotedValue() IQuotedValueContext
	ClosingParenthesis() IClosingParenthesisContext

	// IsTargetRuleContext differentiates from other interfaces.
	IsTargetRuleContext()
}

type TargetRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetRuleContext() *TargetRuleContext {
	var p = new(TargetRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetRule
	return p
}

func InitEmptyTargetRuleContext(p *TargetRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetRule
}

func (*TargetRuleContext) IsTargetRuleContext() {}

func NewTargetRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetRuleContext {
	var p = new(TargetRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetRule

	return p
}

func (s *TargetRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetRuleContext) OpeningParenthesis() IOpeningParenthesisContext {
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

func (s *TargetRuleContext) TargetKeyword() ITargetKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetKeywordContext)
}

func (s *TargetRuleContext) TargetOperator() ITargetOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetOperatorContext)
}

func (s *TargetRuleContext) QuotedValue() IQuotedValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuotedValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuotedValueContext)
}

func (s *TargetRuleContext) ClosingParenthesis() IClosingParenthesisContext {
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

func (s *TargetRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetRule(s)
	}
}

func (s *TargetRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetRule(s)
	}
}

func (p *ACIParser) TargetRule() (localctx ITargetRuleContext) {
	localctx = NewTargetRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ACIParserRULE_targetRule)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.OpeningParenthesis()
	}
	{
		p.SetState(86)
		p.TargetKeyword()
	}
	{
		p.SetState(87)
		p.TargetOperator()
	}
	{
		p.SetState(88)
		p.QuotedValue()
	}
	{
		p.SetState(89)
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

// ITargetOperatorContext is an interface to support dynamic dispatch.
type ITargetOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EqualTo() IEqualToContext
	NotEqualTo() INotEqualToContext

	// IsTargetOperatorContext differentiates from other interfaces.
	IsTargetOperatorContext()
}

type TargetOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetOperatorContext() *TargetOperatorContext {
	var p = new(TargetOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetOperator
	return p
}

func InitEmptyTargetOperatorContext(p *TargetOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetOperator
}

func (*TargetOperatorContext) IsTargetOperatorContext() {}

func NewTargetOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetOperatorContext {
	var p = new(TargetOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetOperator

	return p
}

func (s *TargetOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetOperatorContext) EqualTo() IEqualToContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualToContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualToContext)
}

func (s *TargetOperatorContext) NotEqualTo() INotEqualToContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotEqualToContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotEqualToContext)
}

func (s *TargetOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetOperator(s)
	}
}

func (s *TargetOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetOperator(s)
	}
}

func (p *ACIParser) TargetOperator() (localctx ITargetOperatorContext) {
	localctx = NewTargetOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ACIParserRULE_targetOperator)
	p.SetState(93)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserEQ:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(91)
			p.EqualTo()
		}

	case ACIParserNE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(92)
			p.NotEqualTo()
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

// ITargetRulesContext is an interface to support dynamic dispatch.
type ITargetRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllTargetRule() []ITargetRuleContext
	TargetRule(i int) ITargetRuleContext

	// IsTargetRulesContext differentiates from other interfaces.
	IsTargetRulesContext()
}

type TargetRulesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetRulesContext() *TargetRulesContext {
	var p = new(TargetRulesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetRules
	return p
}

func InitEmptyTargetRulesContext(p *TargetRulesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetRules
}

func (*TargetRulesContext) IsTargetRulesContext() {}

func NewTargetRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetRulesContext {
	var p = new(TargetRulesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetRules

	return p
}

func (s *TargetRulesContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetRulesContext) AllTargetRule() []ITargetRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITargetRuleContext); ok {
			len++
		}
	}

	tst := make([]ITargetRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITargetRuleContext); ok {
			tst[i] = t.(ITargetRuleContext)
			i++
		}
	}

	return tst
}

func (s *TargetRulesContext) TargetRule(i int) ITargetRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetRuleContext); ok {
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

	return t.(ITargetRuleContext)
}

func (s *TargetRulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetRulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetRulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetRules(s)
	}
}

func (s *TargetRulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetRules(s)
	}
}

func (p *ACIParser) TargetRules() (localctx ITargetRulesContext) {
	localctx = NewTargetRulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ACIParserRULE_targetRules)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(95)
				p.TargetRule()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
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

// IBindRuleContext is an interface to support dynamic dispatch.
type IBindRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BindKeyword() IBindKeywordContext
	BindOperator() IBindOperatorContext
	QuotedValue() IQuotedValueContext
	OpeningParenthesis() IOpeningParenthesisContext
	ClosingParenthesis() IClosingParenthesisContext

	// IsBindRuleContext differentiates from other interfaces.
	IsBindRuleContext()
}

type BindRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleContext() *BindRuleContext {
	var p = new(BindRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRule
	return p
}

func InitEmptyBindRuleContext(p *BindRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRule
}

func (*BindRuleContext) IsBindRuleContext() {}

func NewBindRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleContext {
	var p = new(BindRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRule

	return p
}

func (s *BindRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleContext) BindKeyword() IBindKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindKeywordContext)
}

func (s *BindRuleContext) BindOperator() IBindOperatorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindOperatorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindOperatorContext)
}

func (s *BindRuleContext) QuotedValue() IQuotedValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IQuotedValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IQuotedValueContext)
}

func (s *BindRuleContext) OpeningParenthesis() IOpeningParenthesisContext {
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

func (s *BindRuleContext) ClosingParenthesis() IClosingParenthesisContext {
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

func (s *BindRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindRule(s)
	}
}

func (s *BindRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindRule(s)
	}
}

func (p *ACIParser) BindRule() (localctx IBindRuleContext) {
	localctx = NewBindRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ACIParserRULE_bindRule)
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserBKW_UDN, ACIParserBKW_GDN, ACIParserBKW_RDN, ACIParserBKW_UAT, ACIParserBKW_GAT, ACIParserBKW_SSF, ACIParserBKW_IP, ACIParserBKW_DNS, ACIParserBKW_AM, ACIParserBKW_TOD, ACIParserBKW_DOW:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(100)
			p.BindKeyword()
		}
		{
			p.SetState(101)
			p.BindOperator()
		}
		{
			p.SetState(102)
			p.QuotedValue()
		}

	case ACIParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.OpeningParenthesis()
		}
		{
			p.SetState(105)
			p.BindKeyword()
		}
		{
			p.SetState(106)
			p.BindOperator()
		}
		{
			p.SetState(107)
			p.QuotedValue()
		}
		{
			p.SetState(108)
			p.ClosingParenthesis()
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

// IWordAndContext is an interface to support dynamic dispatch.
type IWordAndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD_AND() antlr.TerminalNode

	// IsWordAndContext differentiates from other interfaces.
	IsWordAndContext()
}

type WordAndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordAndContext() *WordAndContext {
	var p = new(WordAndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_wordAnd
	return p
}

func InitEmptyWordAndContext(p *WordAndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_wordAnd
}

func (*WordAndContext) IsWordAndContext() {}

func NewWordAndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordAndContext {
	var p = new(WordAndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_wordAnd

	return p
}

func (s *WordAndContext) GetParser() antlr.Parser { return s.parser }

func (s *WordAndContext) WORD_AND() antlr.TerminalNode {
	return s.GetToken(ACIParserWORD_AND, 0)
}

func (s *WordAndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordAndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordAndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterWordAnd(s)
	}
}

func (s *WordAndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitWordAnd(s)
	}
}

func (p *ACIParser) WordAnd() (localctx IWordAndContext) {
	localctx = NewWordAndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ACIParserRULE_wordAnd)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(112)
		p.Match(ACIParserWORD_AND)
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

// IWordOrContext is an interface to support dynamic dispatch.
type IWordOrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD_OR() antlr.TerminalNode

	// IsWordOrContext differentiates from other interfaces.
	IsWordOrContext()
}

type WordOrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordOrContext() *WordOrContext {
	var p = new(WordOrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_wordOr
	return p
}

func InitEmptyWordOrContext(p *WordOrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_wordOr
}

func (*WordOrContext) IsWordOrContext() {}

func NewWordOrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordOrContext {
	var p = new(WordOrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_wordOr

	return p
}

func (s *WordOrContext) GetParser() antlr.Parser { return s.parser }

func (s *WordOrContext) WORD_OR() antlr.TerminalNode {
	return s.GetToken(ACIParserWORD_OR, 0)
}

func (s *WordOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordOrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterWordOr(s)
	}
}

func (s *WordOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitWordOr(s)
	}
}

func (p *ACIParser) WordOr() (localctx IWordOrContext) {
	localctx = NewWordOrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ACIParserRULE_wordOr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(ACIParserWORD_OR)
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

// IWordNotContext is an interface to support dynamic dispatch.
type IWordNotContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	WORD_NOT() antlr.TerminalNode

	// IsWordNotContext differentiates from other interfaces.
	IsWordNotContext()
}

type WordNotContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWordNotContext() *WordNotContext {
	var p = new(WordNotContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_wordNot
	return p
}

func InitEmptyWordNotContext(p *WordNotContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_wordNot
}

func (*WordNotContext) IsWordNotContext() {}

func NewWordNotContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WordNotContext {
	var p = new(WordNotContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_wordNot

	return p
}

func (s *WordNotContext) GetParser() antlr.Parser { return s.parser }

func (s *WordNotContext) WORD_NOT() antlr.TerminalNode {
	return s.GetToken(ACIParserWORD_NOT, 0)
}

func (s *WordNotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WordNotContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WordNotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterWordNot(s)
	}
}

func (s *WordNotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitWordNot(s)
	}
}

func (p *ACIParser) WordNot() (localctx IWordNotContext) {
	localctx = NewWordNotContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ACIParserRULE_wordNot)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(ACIParserWORD_NOT)
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

// IBindRulesContext is an interface to support dynamic dispatch.
type IBindRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OpeningParenthesis() IOpeningParenthesisContext
	AllBindRules() []IBindRulesContext
	BindRules(i int) IBindRulesContext
	ClosingParenthesis() IClosingParenthesisContext
	AllBindRule() []IBindRuleContext
	BindRule(i int) IBindRuleContext
	AllWordAnd() []IWordAndContext
	WordAnd(i int) IWordAndContext
	AllWordOr() []IWordOrContext
	WordOr(i int) IWordOrContext
	WordNot() IWordNotContext

	// IsBindRulesContext differentiates from other interfaces.
	IsBindRulesContext()
}

type BindRulesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRulesContext() *BindRulesContext {
	var p = new(BindRulesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRules
	return p
}

func InitEmptyBindRulesContext(p *BindRulesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRules
}

func (*BindRulesContext) IsBindRulesContext() {}

func NewBindRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRulesContext {
	var p = new(BindRulesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRules

	return p
}

func (s *BindRulesContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRulesContext) OpeningParenthesis() IOpeningParenthesisContext {
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

func (s *BindRulesContext) AllBindRules() []IBindRulesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBindRulesContext); ok {
			len++
		}
	}

	tst := make([]IBindRulesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBindRulesContext); ok {
			tst[i] = t.(IBindRulesContext)
			i++
		}
	}

	return tst
}

func (s *BindRulesContext) BindRules(i int) IBindRulesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRulesContext); ok {
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

	return t.(IBindRulesContext)
}

func (s *BindRulesContext) ClosingParenthesis() IClosingParenthesisContext {
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

func (s *BindRulesContext) AllBindRule() []IBindRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBindRuleContext); ok {
			len++
		}
	}

	tst := make([]IBindRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBindRuleContext); ok {
			tst[i] = t.(IBindRuleContext)
			i++
		}
	}

	return tst
}

func (s *BindRulesContext) BindRule(i int) IBindRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleContext); ok {
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

	return t.(IBindRuleContext)
}

func (s *BindRulesContext) AllWordAnd() []IWordAndContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordAndContext); ok {
			len++
		}
	}

	tst := make([]IWordAndContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordAndContext); ok {
			tst[i] = t.(IWordAndContext)
			i++
		}
	}

	return tst
}

func (s *BindRulesContext) WordAnd(i int) IWordAndContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordAndContext); ok {
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

	return t.(IWordAndContext)
}

func (s *BindRulesContext) AllWordOr() []IWordOrContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IWordOrContext); ok {
			len++
		}
	}

	tst := make([]IWordOrContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IWordOrContext); ok {
			tst[i] = t.(IWordOrContext)
			i++
		}
	}

	return tst
}

func (s *BindRulesContext) WordOr(i int) IWordOrContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordOrContext); ok {
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

	return t.(IWordOrContext)
}

func (s *BindRulesContext) WordNot() IWordNotContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWordNotContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWordNotContext)
}

func (s *BindRulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindRulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindRules(s)
	}
}

func (s *BindRulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindRules(s)
	}
}

func (p *ACIParser) BindRules() (localctx IBindRulesContext) {
	return p.bindRules(0)
}

func (p *ACIParser) bindRules(_p int) (localctx IBindRulesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewBindRulesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBindRulesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 24
	p.EnterRecursionRule(localctx, 24, ACIParserRULE_bindRules, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(119)
			p.OpeningParenthesis()
		}
		{
			p.SetState(120)
			p.bindRules(0)
		}
		{
			p.SetState(121)
			p.ClosingParenthesis()
		}

	case 2:
		{
			p.SetState(123)
			p.BindRule()
		}
		p.SetState(130)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				p.SetState(126)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}

				switch p.GetTokenStream().LA(1) {
				case ACIParserWORD_AND:
					{
						p.SetState(124)
						p.WordAnd()
					}

				case ACIParserWORD_OR:
					{
						p.SetState(125)
						p.WordOr()
					}

				default:
					p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
					goto errorExit
				}
				{
					p.SetState(128)
					p.BindRule()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(132)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(134)
			p.BindRule()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewBindRulesContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, ACIParserRULE_bindRules)
			p.SetState(137)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(138)
				p.WordNot()
			}
			{
				p.SetState(139)
				p.bindRules(3)
			}

		}
		p.SetState(145)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
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
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IQuotedValueContext is an interface to support dynamic dispatch.
type IQuotedValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	QUOTED_STRING() antlr.TerminalNode

	// IsQuotedValueContext differentiates from other interfaces.
	IsQuotedValueContext()
}

type QuotedValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQuotedValueContext() *QuotedValueContext {
	var p = new(QuotedValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_quotedValue
	return p
}

func InitEmptyQuotedValueContext(p *QuotedValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_quotedValue
}

func (*QuotedValueContext) IsQuotedValueContext() {}

func NewQuotedValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QuotedValueContext {
	var p = new(QuotedValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_quotedValue

	return p
}

func (s *QuotedValueContext) GetParser() antlr.Parser { return s.parser }

func (s *QuotedValueContext) QUOTED_STRING() antlr.TerminalNode {
	return s.GetToken(ACIParserQUOTED_STRING, 0)
}

func (s *QuotedValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QuotedValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterQuotedValue(s)
	}
}

func (s *QuotedValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitQuotedValue(s)
	}
}

func (p *ACIParser) QuotedValue() (localctx IQuotedValueContext) {
	localctx = NewQuotedValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ACIParserRULE_quotedValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(ACIParserQUOTED_STRING)
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

// IBindOperatorContext is an interface to support dynamic dispatch.
type IBindOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EqualTo() IEqualToContext
	NotEqualTo() INotEqualToContext
	LessThan() ILessThanContext
	LessThanOrEqual() ILessThanOrEqualContext
	GreaterThan() IGreaterThanContext
	GreaterThanOrEqual() IGreaterThanOrEqualContext

	// IsBindOperatorContext differentiates from other interfaces.
	IsBindOperatorContext()
}

type BindOperatorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindOperatorContext() *BindOperatorContext {
	var p = new(BindOperatorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindOperator
	return p
}

func InitEmptyBindOperatorContext(p *BindOperatorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindOperator
}

func (*BindOperatorContext) IsBindOperatorContext() {}

func NewBindOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindOperatorContext {
	var p = new(BindOperatorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindOperator

	return p
}

func (s *BindOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *BindOperatorContext) EqualTo() IEqualToContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEqualToContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEqualToContext)
}

func (s *BindOperatorContext) NotEqualTo() INotEqualToContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INotEqualToContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INotEqualToContext)
}

func (s *BindOperatorContext) LessThan() ILessThanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILessThanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILessThanContext)
}

func (s *BindOperatorContext) LessThanOrEqual() ILessThanOrEqualContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILessThanOrEqualContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILessThanOrEqualContext)
}

func (s *BindOperatorContext) GreaterThan() IGreaterThanContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGreaterThanContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGreaterThanContext)
}

func (s *BindOperatorContext) GreaterThanOrEqual() IGreaterThanOrEqualContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGreaterThanOrEqualContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGreaterThanOrEqualContext)
}

func (s *BindOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindOperator(s)
	}
}

func (s *BindOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindOperator(s)
	}
}

func (p *ACIParser) BindOperator() (localctx IBindOperatorContext) {
	localctx = NewBindOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ACIParserRULE_bindOperator)
	p.SetState(154)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserEQ:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(148)
			p.EqualTo()
		}

	case ACIParserNE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(149)
			p.NotEqualTo()
		}

	case ACIParserLT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(150)
			p.LessThan()
		}

	case ACIParserLE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(151)
			p.LessThanOrEqual()
		}

	case ACIParserGT:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(152)
			p.GreaterThan()
		}

	case ACIParserGE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(153)
			p.GreaterThanOrEqual()
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

// ILessThanContext is an interface to support dynamic dispatch.
type ILessThanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LT() antlr.TerminalNode

	// IsLessThanContext differentiates from other interfaces.
	IsLessThanContext()
}

type LessThanContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLessThanContext() *LessThanContext {
	var p = new(LessThanContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lessThan
	return p
}

func InitEmptyLessThanContext(p *LessThanContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lessThan
}

func (*LessThanContext) IsLessThanContext() {}

func NewLessThanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LessThanContext {
	var p = new(LessThanContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_lessThan

	return p
}

func (s *LessThanContext) GetParser() antlr.Parser { return s.parser }

func (s *LessThanContext) LT() antlr.TerminalNode {
	return s.GetToken(ACIParserLT, 0)
}

func (s *LessThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LessThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterLessThan(s)
	}
}

func (s *LessThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitLessThan(s)
	}
}

func (p *ACIParser) LessThan() (localctx ILessThanContext) {
	localctx = NewLessThanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ACIParserRULE_lessThan)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(156)
		p.Match(ACIParserLT)
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

// ILessThanOrEqualContext is an interface to support dynamic dispatch.
type ILessThanOrEqualContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LE() antlr.TerminalNode

	// IsLessThanOrEqualContext differentiates from other interfaces.
	IsLessThanOrEqualContext()
}

type LessThanOrEqualContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLessThanOrEqualContext() *LessThanOrEqualContext {
	var p = new(LessThanOrEqualContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lessThanOrEqual
	return p
}

func InitEmptyLessThanOrEqualContext(p *LessThanOrEqualContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lessThanOrEqual
}

func (*LessThanOrEqualContext) IsLessThanOrEqualContext() {}

func NewLessThanOrEqualContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LessThanOrEqualContext {
	var p = new(LessThanOrEqualContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_lessThanOrEqual

	return p
}

func (s *LessThanOrEqualContext) GetParser() antlr.Parser { return s.parser }

func (s *LessThanOrEqualContext) LE() antlr.TerminalNode {
	return s.GetToken(ACIParserLE, 0)
}

func (s *LessThanOrEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanOrEqualContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LessThanOrEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterLessThanOrEqual(s)
	}
}

func (s *LessThanOrEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitLessThanOrEqual(s)
	}
}

func (p *ACIParser) LessThanOrEqual() (localctx ILessThanOrEqualContext) {
	localctx = NewLessThanOrEqualContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ACIParserRULE_lessThanOrEqual)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(158)
		p.Match(ACIParserLE)
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

// IGreaterThanContext is an interface to support dynamic dispatch.
type IGreaterThanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GT() antlr.TerminalNode

	// IsGreaterThanContext differentiates from other interfaces.
	IsGreaterThanContext()
}

type GreaterThanContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGreaterThanContext() *GreaterThanContext {
	var p = new(GreaterThanContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_greaterThan
	return p
}

func InitEmptyGreaterThanContext(p *GreaterThanContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_greaterThan
}

func (*GreaterThanContext) IsGreaterThanContext() {}

func NewGreaterThanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GreaterThanContext {
	var p = new(GreaterThanContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_greaterThan

	return p
}

func (s *GreaterThanContext) GetParser() antlr.Parser { return s.parser }

func (s *GreaterThanContext) GT() antlr.TerminalNode {
	return s.GetToken(ACIParserGT, 0)
}

func (s *GreaterThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GreaterThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterGreaterThan(s)
	}
}

func (s *GreaterThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitGreaterThan(s)
	}
}

func (p *ACIParser) GreaterThan() (localctx IGreaterThanContext) {
	localctx = NewGreaterThanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ACIParserRULE_greaterThan)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(160)
		p.Match(ACIParserGT)
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

// IGreaterThanOrEqualContext is an interface to support dynamic dispatch.
type IGreaterThanOrEqualContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GE() antlr.TerminalNode

	// IsGreaterThanOrEqualContext differentiates from other interfaces.
	IsGreaterThanOrEqualContext()
}

type GreaterThanOrEqualContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGreaterThanOrEqualContext() *GreaterThanOrEqualContext {
	var p = new(GreaterThanOrEqualContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_greaterThanOrEqual
	return p
}

func InitEmptyGreaterThanOrEqualContext(p *GreaterThanOrEqualContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_greaterThanOrEqual
}

func (*GreaterThanOrEqualContext) IsGreaterThanOrEqualContext() {}

func NewGreaterThanOrEqualContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GreaterThanOrEqualContext {
	var p = new(GreaterThanOrEqualContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_greaterThanOrEqual

	return p
}

func (s *GreaterThanOrEqualContext) GetParser() antlr.Parser { return s.parser }

func (s *GreaterThanOrEqualContext) GE() antlr.TerminalNode {
	return s.GetToken(ACIParserGE, 0)
}

func (s *GreaterThanOrEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanOrEqualContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GreaterThanOrEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterGreaterThanOrEqual(s)
	}
}

func (s *GreaterThanOrEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitGreaterThanOrEqual(s)
	}
}

func (p *ACIParser) GreaterThanOrEqual() (localctx IGreaterThanOrEqualContext) {
	localctx = NewGreaterThanOrEqualContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ACIParserRULE_greaterThanOrEqual)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(162)
		p.Match(ACIParserGE)
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

// IEqualToContext is an interface to support dynamic dispatch.
type IEqualToContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EQ() antlr.TerminalNode

	// IsEqualToContext differentiates from other interfaces.
	IsEqualToContext()
}

type EqualToContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEqualToContext() *EqualToContext {
	var p = new(EqualToContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_equalTo
	return p
}

func InitEmptyEqualToContext(p *EqualToContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_equalTo
}

func (*EqualToContext) IsEqualToContext() {}

func NewEqualToContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EqualToContext {
	var p = new(EqualToContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_equalTo

	return p
}

func (s *EqualToContext) GetParser() antlr.Parser { return s.parser }

func (s *EqualToContext) EQ() antlr.TerminalNode {
	return s.GetToken(ACIParserEQ, 0)
}

func (s *EqualToContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualToContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EqualToContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterEqualTo(s)
	}
}

func (s *EqualToContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitEqualTo(s)
	}
}

func (p *ACIParser) EqualTo() (localctx IEqualToContext) {
	localctx = NewEqualToContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ACIParserRULE_equalTo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(164)
		p.Match(ACIParserEQ)
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

// INotEqualToContext is an interface to support dynamic dispatch.
type INotEqualToContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NE() antlr.TerminalNode

	// IsNotEqualToContext differentiates from other interfaces.
	IsNotEqualToContext()
}

type NotEqualToContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNotEqualToContext() *NotEqualToContext {
	var p = new(NotEqualToContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_notEqualTo
	return p
}

func InitEmptyNotEqualToContext(p *NotEqualToContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_notEqualTo
}

func (*NotEqualToContext) IsNotEqualToContext() {}

func NewNotEqualToContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NotEqualToContext {
	var p = new(NotEqualToContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_notEqualTo

	return p
}

func (s *NotEqualToContext) GetParser() antlr.Parser { return s.parser }

func (s *NotEqualToContext) NE() antlr.TerminalNode {
	return s.GetToken(ACIParserNE, 0)
}

func (s *NotEqualToContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqualToContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NotEqualToContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterNotEqualTo(s)
	}
}

func (s *NotEqualToContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitNotEqualTo(s)
	}
}

func (p *ACIParser) NotEqualTo() (localctx INotEqualToContext) {
	localctx = NewNotEqualToContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ACIParserRULE_notEqualTo)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(ACIParserNE)
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

// IBindKeywordContext is an interface to support dynamic dispatch.
type IBindKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BKW_UDN() antlr.TerminalNode
	BKW_GDN() antlr.TerminalNode
	BKW_RDN() antlr.TerminalNode
	BKW_UAT() antlr.TerminalNode
	BKW_GAT() antlr.TerminalNode
	BKW_SSF() antlr.TerminalNode
	BKW_IP() antlr.TerminalNode
	BKW_DNS() antlr.TerminalNode
	BKW_DOW() antlr.TerminalNode
	BKW_TOD() antlr.TerminalNode
	BKW_AM() antlr.TerminalNode

	// IsBindKeywordContext differentiates from other interfaces.
	IsBindKeywordContext()
}

type BindKeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindKeywordContext() *BindKeywordContext {
	var p = new(BindKeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindKeyword
	return p
}

func InitEmptyBindKeywordContext(p *BindKeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindKeyword
}

func (*BindKeywordContext) IsBindKeywordContext() {}

func NewBindKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindKeywordContext {
	var p = new(BindKeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindKeyword

	return p
}

func (s *BindKeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *BindKeywordContext) BKW_UDN() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_UDN, 0)
}

func (s *BindKeywordContext) BKW_GDN() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_GDN, 0)
}

func (s *BindKeywordContext) BKW_RDN() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_RDN, 0)
}

func (s *BindKeywordContext) BKW_UAT() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_UAT, 0)
}

func (s *BindKeywordContext) BKW_GAT() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_GAT, 0)
}

func (s *BindKeywordContext) BKW_SSF() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_SSF, 0)
}

func (s *BindKeywordContext) BKW_IP() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_IP, 0)
}

func (s *BindKeywordContext) BKW_DNS() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_DNS, 0)
}

func (s *BindKeywordContext) BKW_DOW() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_DOW, 0)
}

func (s *BindKeywordContext) BKW_TOD() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_TOD, 0)
}

func (s *BindKeywordContext) BKW_AM() antlr.TerminalNode {
	return s.GetToken(ACIParserBKW_AM, 0)
}

func (s *BindKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindKeyword(s)
	}
}

func (s *BindKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindKeyword(s)
	}
}

func (p *ACIParser) BindKeyword() (localctx IBindKeywordContext) {
	localctx = NewBindKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ACIParserRULE_bindKeyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4094) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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

// ITargetKeywordContext is an interface to support dynamic dispatch.
type ITargetKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TKW_TARGET() antlr.TerminalNode
	TKW_TO() antlr.TerminalNode
	TKW_FROM() antlr.TerminalNode
	TKW_SCOPE() antlr.TerminalNode
	TKW_ATTR() antlr.TerminalNode
	TKW_FILTER() antlr.TerminalNode
	TKW_AF() antlr.TerminalNode
	TKW_CTRL() antlr.TerminalNode
	TKW_EXTOP() antlr.TerminalNode

	// IsTargetKeywordContext differentiates from other interfaces.
	IsTargetKeywordContext()
}

type TargetKeywordContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetKeywordContext() *TargetKeywordContext {
	var p = new(TargetKeywordContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetKeyword
	return p
}

func InitEmptyTargetKeywordContext(p *TargetKeywordContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetKeyword
}

func (*TargetKeywordContext) IsTargetKeywordContext() {}

func NewTargetKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetKeywordContext {
	var p = new(TargetKeywordContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetKeyword

	return p
}

func (s *TargetKeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetKeywordContext) TKW_TARGET() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_TARGET, 0)
}

func (s *TargetKeywordContext) TKW_TO() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_TO, 0)
}

func (s *TargetKeywordContext) TKW_FROM() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_FROM, 0)
}

func (s *TargetKeywordContext) TKW_SCOPE() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_SCOPE, 0)
}

func (s *TargetKeywordContext) TKW_ATTR() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_ATTR, 0)
}

func (s *TargetKeywordContext) TKW_FILTER() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_FILTER, 0)
}

func (s *TargetKeywordContext) TKW_AF() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_AF, 0)
}

func (s *TargetKeywordContext) TKW_CTRL() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_CTRL, 0)
}

func (s *TargetKeywordContext) TKW_EXTOP() antlr.TerminalNode {
	return s.GetToken(ACIParserTKW_EXTOP, 0)
}

func (s *TargetKeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetKeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetKeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetKeyword(s)
	}
}

func (s *TargetKeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetKeyword(s)
	}
}

func (p *ACIParser) TargetKeyword() (localctx ITargetKeywordContext) {
	localctx = NewTargetKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ACIParserRULE_targetKeyword)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2093056) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	p.RuleIndex = ACIParserRULE_openingParenthesis
	return p
}

func InitEmptyOpeningParenthesisContext(p *OpeningParenthesisContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_openingParenthesis
}

func (*OpeningParenthesisContext) IsOpeningParenthesisContext() {}

func NewOpeningParenthesisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OpeningParenthesisContext {
	var p = new(OpeningParenthesisContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_openingParenthesis

	return p
}

func (s *OpeningParenthesisContext) GetParser() antlr.Parser { return s.parser }

func (s *OpeningParenthesisContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *OpeningParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpeningParenthesisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OpeningParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterOpeningParenthesis(s)
	}
}

func (s *OpeningParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitOpeningParenthesis(s)
	}
}

func (p *ACIParser) OpeningParenthesis() (localctx IOpeningParenthesisContext) {
	localctx = NewOpeningParenthesisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ACIParserRULE_openingParenthesis)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(ACIParserLPAREN)
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
	p.RuleIndex = ACIParserRULE_closingParenthesis
	return p
}

func InitEmptyClosingParenthesisContext(p *ClosingParenthesisContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_closingParenthesis
}

func (*ClosingParenthesisContext) IsClosingParenthesisContext() {}

func NewClosingParenthesisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClosingParenthesisContext {
	var p = new(ClosingParenthesisContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_closingParenthesis

	return p
}

func (s *ClosingParenthesisContext) GetParser() antlr.Parser { return s.parser }

func (s *ClosingParenthesisContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ClosingParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClosingParenthesisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClosingParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterClosingParenthesis(s)
	}
}

func (s *ClosingParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitClosingParenthesis(s)
	}
}

func (p *ACIParser) ClosingParenthesis() (localctx IClosingParenthesisContext) {
	localctx = NewClosingParenthesisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ACIParserRULE_closingParenthesis)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(ACIParserRPAREN)
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

func (p *ACIParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 12:
		var t *BindRulesContext = nil
		if localctx != nil {
			t = localctx.(*BindRulesContext)
		}
		return p.BindRules_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ACIParser) BindRules_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
