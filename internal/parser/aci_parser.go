// Code generated from ACIParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // ACIParser
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
		"", "", "", "", "':'", "", "", "", "", "'~='", "'!='", "'>='", "'<='",
		"", "", "", "", "", "", "", "'target'", "'targetscope'", "'target_to'",
		"'target_from'", "'targetattr'", "'targetfilter'", "'targattrfilters'",
		"'targetcontrol'", "'extop'", "'none'", "'simple'", "'SASL'", "'SSL'",
		"'roledn'", "'userdn'", "'groupdn'", "'userattr'", "'groupattr'", "'ssf'",
		"'dns'", "'ip'", "'authmethod'", "'timeofday'", "'dayofweek'", "", "",
		"", "", "", "", "", "", "'parent['", "", "';'", "'='", "'>'", "'<'",
		"'ldap'", "'version 3.0; acl '", "'dn'", "'&'", "'|'", "'!'", "'/'",
		"'('", "'{'", "'['", "')'", "']'", "'}'", "'\"'", "','", "'~'", "'#'",
		"'$'", "'@'", "'.'", "'-'", "'*'", "'?'",
	}
	staticData.SymbolicNames = []string{
		"", "LEVELS", "INT", "ID", "COLON", "ExtensibleRuleDNMatch", "ExtensibleRuleDNOIDMatch",
		"ExtensibleRuleAttrMatch", "ExtensibleRuleMatch", "ApproximateMatch",
		"NotEqualTo", "GreaterThanOrEqual", "LessThanOrEqual", "DAMP", "DPIPE",
		"BooleanWordOperators", "BooleanAndWord", "BooleanOrWord", "BooleanNotWord",
		"DAYS", "TargetKeyword", "TargetScopeKeyword", "TargetToKeyword", "TargetFromKeyword",
		"TargetAttrKeyword", "TargetFilterKeyword", "TargetAttrFiltersKeyword",
		"TargetControlKeyword", "ExtOpKeyword", "ANONYMOUS", "SIMPLE", "SASL",
		"SSL", "RoleDNKeyword", "UserDNKeyword", "GroupDNKeyword", "UserAttrKeyword",
		"GroupAttrKeyword", "SSFKeyword", "DNSKeyword", "IPKeyword", "AuthMethodKeyword",
		"TimeOfDayKeyword", "DayOfWeekKeyword", "DISPOSITION", "RIGHTS", "AttributeFilterOperation",
		"AddOperation", "DeleteOperation", "TargetRuleSearchScopes", "LDAPSearchScopes",
		"BINDTYPES", "INHERITANCEPREFIX", "LocalLDAPScheme", "SEMI", "EqualTo",
		"GreaterThan", "LessThan", "LDAP", "ANCHOR", "DN", "AMP", "PIPE", "BANG",
		"SOLIDUS", "LPAREN", "LBRAC", "LBRAK", "RPAREN", "RBRAK", "RBRAC", "DQUOTE",
		"COMMA", "TILDE", "HASH", "DOLLAR", "ATSIGN", "DOT", "DASH", "STAR",
		"QMARK", "LineTerminator", "WhiteSpaces", "NumericLiteral", "Literal",
		"DelimitedAddress", "DelimitedNumbers", "StringLiteral", "WildcardString",
		"MacroValue", "ANY",
	}
	staticData.RuleNames = []string{
		"parse", "instruction", "version", "permBindRules", "permBindRule",
		"permission", "targetRules", "targetRule", "targetControl", "targetExtOp",
		"targetFilter", "targetFilterValue", "target", "targetTo", "targetFrom",
		"targetAttrFilters", "targetAttrFiltersValue", "targetScope", "targetAttr",
		"attributeTypes", "bindRule", "bindRuleExprParen", "bindRuleExpr", "bindRuleUserDN",
		"bindRuleRoleDN", "bindRuleGroupDN", "bindRuleUserAttr", "bindRuleGroupAttr",
		"bindRuleAuthMethod", "bindRuleDNS", "bindRuleTimeOfDay", "bindRuleDayOfWeek",
		"bindRuleIP", "bindRuleSecurityStrengthFactor", "dayOfWeek", "fQDN",
		"objectIdentifiers", "objectIdentifier", "iPV6Address", "iPV4Address",
		"securityStrengthFactor", "timeOfDay", "objectIdentifierArc", "inheritance",
		"inheritanceLevels", "attributeBindTypeOrValue", "attributeFilters",
		"attributeFilterSet", "doubleAmpersand", "attributeFilter", "distinguishedNames",
		"doublePipe", "lDAPURIAndBindType", "lDAPURI", "uRISearchFilter", "uRISearchScopes",
		"uRIAttributeList", "distinguishedName", "relativeDistinguishedName",
		"lDAPFilter", "lDAPFilterExpr", "attributeValueAssertion", "attributeType",
		"attributeValue", "attributeOperators",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 90, 594, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 5, 3, 147, 8, 3, 10, 3,
		12, 3, 150, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		5, 5, 161, 8, 5, 10, 5, 12, 5, 164, 9, 5, 3, 5, 166, 8, 5, 1, 5, 1, 5,
		1, 6, 5, 6, 171, 8, 6, 10, 6, 12, 6, 174, 9, 6, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 185, 8, 7, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 3, 16, 241, 8, 16, 1,
		17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 19, 1, 19, 1, 19, 1, 19, 5, 19, 259, 8, 19, 10, 19, 12, 19, 262,
		9, 19, 1, 20, 1, 20, 1, 20, 1, 20, 5, 20, 268, 8, 20, 10, 20, 12, 20, 271,
		9, 20, 3, 20, 273, 8, 20, 1, 21, 1, 21, 1, 21, 1, 21, 5, 21, 279, 8, 21,
		10, 21, 12, 21, 282, 9, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1,
		21, 1, 21, 1, 21, 3, 21, 293, 8, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 3,
		22, 310, 8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 3, 23,
		319, 8, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 3, 24, 328,
		8, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 3, 25, 337, 8,
		25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 347,
		8, 26, 3, 26, 349, 8, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 3, 27, 359, 8, 27, 3, 27, 361, 8, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 3, 28, 370, 8, 28, 1, 29, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 3, 29, 379, 8, 29, 1, 30, 1, 30, 1, 30, 1, 30, 1,
		30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 391, 8, 30, 1, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 403, 8,
		31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 413,
		8, 32, 3, 32, 415, 8, 32, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1, 33, 1,
		33, 3, 33, 424, 8, 33, 1, 34, 1, 34, 1, 34, 5, 34, 429, 8, 34, 10, 34,
		12, 34, 432, 9, 34, 3, 34, 434, 8, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36,
		1, 36, 4, 36, 442, 8, 36, 11, 36, 12, 36, 443, 1, 37, 1, 37, 1, 38, 1,
		38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43,
		1, 43, 1, 43, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 45, 1, 45, 3,
		45, 470, 8, 45, 1, 46, 1, 46, 1, 46, 4, 46, 475, 8, 46, 11, 46, 12, 46,
		476, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 5, 47, 485, 8, 47, 10, 47,
		12, 47, 488, 9, 47, 3, 47, 490, 8, 47, 1, 48, 1, 48, 1, 49, 1, 49, 1, 49,
		1, 49, 1, 50, 1, 50, 1, 50, 1, 50, 5, 50, 502, 8, 50, 10, 50, 12, 50, 505,
		9, 50, 4, 50, 507, 8, 50, 11, 50, 12, 50, 508, 1, 51, 1, 51, 1, 52, 1,
		52, 1, 52, 1, 52, 1, 53, 1, 53, 1, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 54,
		1, 55, 1, 55, 3, 55, 527, 8, 55, 1, 56, 1, 56, 1, 56, 1, 56, 5, 56, 533,
		8, 56, 10, 56, 12, 56, 536, 9, 56, 3, 56, 538, 8, 56, 1, 57, 1, 57, 1,
		57, 1, 57, 5, 57, 544, 8, 57, 10, 57, 12, 57, 547, 9, 57, 1, 58, 1, 58,
		1, 58, 1, 58, 1, 58, 3, 58, 554, 8, 58, 1, 59, 1, 59, 1, 59, 1, 59, 1,
		59, 5, 59, 561, 8, 59, 10, 59, 12, 59, 564, 9, 59, 3, 59, 566, 8, 59, 1,
		60, 1, 60, 3, 60, 570, 8, 60, 1, 60, 1, 60, 1, 60, 4, 60, 575, 8, 60, 11,
		60, 12, 60, 576, 1, 60, 1, 60, 1, 60, 3, 60, 582, 8, 60, 1, 61, 1, 61,
		1, 61, 1, 61, 1, 62, 1, 62, 1, 63, 1, 63, 1, 64, 1, 64, 1, 64, 2, 476,
		576, 0, 65, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
		32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66,
		68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102,
		104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126, 128, 0, 8,
		2, 0, 10, 10, 55, 55, 1, 0, 29, 32, 2, 0, 10, 12, 55, 57, 2, 0, 54, 54,
		72, 72, 1, 0, 47, 48, 1, 0, 61, 63, 3, 0, 2, 3, 88, 88, 90, 90, 3, 0, 5,
		9, 11, 12, 55, 55, 594, 0, 130, 1, 0, 0, 0, 2, 133, 1, 0, 0, 0, 4, 140,
		1, 0, 0, 0, 6, 148, 1, 0, 0, 0, 8, 151, 1, 0, 0, 0, 10, 155, 1, 0, 0, 0,
		12, 172, 1, 0, 0, 0, 14, 184, 1, 0, 0, 0, 16, 186, 1, 0, 0, 0, 18, 192,
		1, 0, 0, 0, 20, 198, 1, 0, 0, 0, 22, 203, 1, 0, 0, 0, 24, 207, 1, 0, 0,
		0, 26, 212, 1, 0, 0, 0, 28, 217, 1, 0, 0, 0, 30, 222, 1, 0, 0, 0, 32, 240,
		1, 0, 0, 0, 34, 242, 1, 0, 0, 0, 36, 248, 1, 0, 0, 0, 38, 254, 1, 0, 0,
		0, 40, 272, 1, 0, 0, 0, 42, 292, 1, 0, 0, 0, 44, 309, 1, 0, 0, 0, 46, 318,
		1, 0, 0, 0, 48, 327, 1, 0, 0, 0, 50, 336, 1, 0, 0, 0, 52, 348, 1, 0, 0,
		0, 54, 360, 1, 0, 0, 0, 56, 369, 1, 0, 0, 0, 58, 378, 1, 0, 0, 0, 60, 390,
		1, 0, 0, 0, 62, 402, 1, 0, 0, 0, 64, 414, 1, 0, 0, 0, 66, 423, 1, 0, 0,
		0, 68, 433, 1, 0, 0, 0, 70, 435, 1, 0, 0, 0, 72, 437, 1, 0, 0, 0, 74, 445,
		1, 0, 0, 0, 76, 447, 1, 0, 0, 0, 78, 449, 1, 0, 0, 0, 80, 451, 1, 0, 0,
		0, 82, 453, 1, 0, 0, 0, 84, 455, 1, 0, 0, 0, 86, 457, 1, 0, 0, 0, 88, 463,
		1, 0, 0, 0, 90, 465, 1, 0, 0, 0, 92, 471, 1, 0, 0, 0, 94, 478, 1, 0, 0,
		0, 96, 491, 1, 0, 0, 0, 98, 493, 1, 0, 0, 0, 100, 506, 1, 0, 0, 0, 102,
		510, 1, 0, 0, 0, 104, 512, 1, 0, 0, 0, 106, 516, 1, 0, 0, 0, 108, 521,
		1, 0, 0, 0, 110, 524, 1, 0, 0, 0, 112, 528, 1, 0, 0, 0, 114, 539, 1, 0,
		0, 0, 116, 553, 1, 0, 0, 0, 118, 565, 1, 0, 0, 0, 120, 581, 1, 0, 0, 0,
		122, 583, 1, 0, 0, 0, 124, 587, 1, 0, 0, 0, 126, 589, 1, 0, 0, 0, 128,
		591, 1, 0, 0, 0, 130, 131, 3, 2, 1, 0, 131, 132, 5, 0, 0, 1, 132, 1, 1,
		0, 0, 0, 133, 134, 5, 65, 0, 0, 134, 135, 3, 14, 7, 0, 135, 136, 3, 4,
		2, 0, 136, 137, 3, 6, 3, 0, 137, 138, 5, 68, 0, 0, 138, 139, 5, 54, 0,
		0, 139, 3, 1, 0, 0, 0, 140, 141, 5, 65, 0, 0, 141, 142, 5, 59, 0, 0, 142,
		143, 5, 84, 0, 0, 143, 144, 5, 54, 0, 0, 144, 5, 1, 0, 0, 0, 145, 147,
		3, 8, 4, 0, 146, 145, 1, 0, 0, 0, 147, 150, 1, 0, 0, 0, 148, 146, 1, 0,
		0, 0, 148, 149, 1, 0, 0, 0, 149, 7, 1, 0, 0, 0, 150, 148, 1, 0, 0, 0, 151,
		152, 3, 10, 5, 0, 152, 153, 3, 40, 20, 0, 153, 154, 5, 54, 0, 0, 154, 9,
		1, 0, 0, 0, 155, 156, 5, 44, 0, 0, 156, 165, 5, 65, 0, 0, 157, 162, 5,
		45, 0, 0, 158, 159, 5, 72, 0, 0, 159, 161, 5, 45, 0, 0, 160, 158, 1, 0,
		0, 0, 161, 164, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 162, 163, 1, 0, 0, 0,
		163, 166, 1, 0, 0, 0, 164, 162, 1, 0, 0, 0, 165, 157, 1, 0, 0, 0, 165,
		166, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 168, 5, 68, 0, 0, 168, 11,
		1, 0, 0, 0, 169, 171, 3, 14, 7, 0, 170, 169, 1, 0, 0, 0, 171, 174, 1, 0,
		0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 13, 1, 0, 0, 0,
		174, 172, 1, 0, 0, 0, 175, 185, 3, 16, 8, 0, 176, 185, 3, 18, 9, 0, 177,
		185, 3, 20, 10, 0, 178, 185, 3, 30, 15, 0, 179, 185, 3, 34, 17, 0, 180,
		185, 3, 36, 18, 0, 181, 185, 3, 24, 12, 0, 182, 185, 3, 26, 13, 0, 183,
		185, 3, 28, 14, 0, 184, 175, 1, 0, 0, 0, 184, 176, 1, 0, 0, 0, 184, 177,
		1, 0, 0, 0, 184, 178, 1, 0, 0, 0, 184, 179, 1, 0, 0, 0, 184, 180, 1, 0,
		0, 0, 184, 181, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 184, 183, 1, 0, 0, 0,
		185, 15, 1, 0, 0, 0, 186, 187, 5, 65, 0, 0, 187, 188, 5, 27, 0, 0, 188,
		189, 7, 0, 0, 0, 189, 190, 3, 72, 36, 0, 190, 191, 5, 68, 0, 0, 191, 17,
		1, 0, 0, 0, 192, 193, 5, 65, 0, 0, 193, 194, 5, 28, 0, 0, 194, 195, 7,
		0, 0, 0, 195, 196, 3, 72, 36, 0, 196, 197, 5, 68, 0, 0, 197, 19, 1, 0,
		0, 0, 198, 199, 5, 65, 0, 0, 199, 200, 5, 25, 0, 0, 200, 201, 7, 0, 0,
		0, 201, 202, 3, 22, 11, 0, 202, 21, 1, 0, 0, 0, 203, 204, 5, 71, 0, 0,
		204, 205, 3, 118, 59, 0, 205, 206, 5, 71, 0, 0, 206, 23, 1, 0, 0, 0, 207,
		208, 5, 65, 0, 0, 208, 209, 5, 20, 0, 0, 209, 210, 7, 0, 0, 0, 210, 211,
		3, 100, 50, 0, 211, 25, 1, 0, 0, 0, 212, 213, 5, 65, 0, 0, 213, 214, 5,
		22, 0, 0, 214, 215, 7, 0, 0, 0, 215, 216, 3, 114, 57, 0, 216, 27, 1, 0,
		0, 0, 217, 218, 5, 65, 0, 0, 218, 219, 5, 23, 0, 0, 219, 220, 7, 0, 0,
		0, 220, 221, 3, 114, 57, 0, 221, 29, 1, 0, 0, 0, 222, 223, 5, 65, 0, 0,
		223, 224, 5, 26, 0, 0, 224, 225, 5, 55, 0, 0, 225, 226, 3, 32, 16, 0, 226,
		227, 5, 68, 0, 0, 227, 31, 1, 0, 0, 0, 228, 229, 5, 71, 0, 0, 229, 230,
		3, 92, 46, 0, 230, 231, 5, 71, 0, 0, 231, 241, 1, 0, 0, 0, 232, 233, 5,
		71, 0, 0, 233, 234, 3, 94, 47, 0, 234, 235, 5, 71, 0, 0, 235, 241, 1, 0,
		0, 0, 236, 237, 5, 71, 0, 0, 237, 238, 3, 98, 49, 0, 238, 239, 5, 71, 0,
		0, 239, 241, 1, 0, 0, 0, 240, 228, 1, 0, 0, 0, 240, 232, 1, 0, 0, 0, 240,
		236, 1, 0, 0, 0, 241, 33, 1, 0, 0, 0, 242, 243, 5, 65, 0, 0, 243, 244,
		5, 21, 0, 0, 244, 245, 5, 55, 0, 0, 245, 246, 5, 49, 0, 0, 246, 247, 5,
		68, 0, 0, 247, 35, 1, 0, 0, 0, 248, 249, 5, 65, 0, 0, 249, 250, 5, 24,
		0, 0, 250, 251, 7, 0, 0, 0, 251, 252, 3, 124, 62, 0, 252, 253, 5, 68, 0,
		0, 253, 37, 1, 0, 0, 0, 254, 260, 3, 124, 62, 0, 255, 256, 3, 102, 51,
		0, 256, 257, 3, 124, 62, 0, 257, 259, 1, 0, 0, 0, 258, 255, 1, 0, 0, 0,
		259, 262, 1, 0, 0, 0, 260, 258, 1, 0, 0, 0, 260, 261, 1, 0, 0, 0, 261,
		39, 1, 0, 0, 0, 262, 260, 1, 0, 0, 0, 263, 273, 3, 44, 22, 0, 264, 269,
		3, 42, 21, 0, 265, 266, 5, 15, 0, 0, 266, 268, 3, 42, 21, 0, 267, 265,
		1, 0, 0, 0, 268, 271, 1, 0, 0, 0, 269, 267, 1, 0, 0, 0, 269, 270, 1, 0,
		0, 0, 270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0, 272, 263, 1, 0, 0, 0,
		272, 264, 1, 0, 0, 0, 273, 41, 1, 0, 0, 0, 274, 275, 5, 65, 0, 0, 275,
		280, 3, 44, 22, 0, 276, 277, 5, 15, 0, 0, 277, 279, 3, 44, 22, 0, 278,
		276, 1, 0, 0, 0, 279, 282, 1, 0, 0, 0, 280, 278, 1, 0, 0, 0, 280, 281,
		1, 0, 0, 0, 281, 283, 1, 0, 0, 0, 282, 280, 1, 0, 0, 0, 283, 284, 5, 68,
		0, 0, 284, 293, 1, 0, 0, 0, 285, 286, 5, 18, 0, 0, 286, 293, 3, 44, 22,
		0, 287, 288, 5, 65, 0, 0, 288, 289, 3, 44, 22, 0, 289, 290, 5, 68, 0, 0,
		290, 293, 1, 0, 0, 0, 291, 293, 3, 44, 22, 0, 292, 274, 1, 0, 0, 0, 292,
		285, 1, 0, 0, 0, 292, 287, 1, 0, 0, 0, 292, 291, 1, 0, 0, 0, 293, 43, 1,
		0, 0, 0, 294, 295, 5, 65, 0, 0, 295, 296, 3, 44, 22, 0, 296, 297, 5, 68,
		0, 0, 297, 310, 1, 0, 0, 0, 298, 310, 3, 46, 23, 0, 299, 310, 3, 52, 26,
		0, 300, 310, 3, 50, 25, 0, 301, 310, 3, 54, 27, 0, 302, 310, 3, 48, 24,
		0, 303, 310, 3, 58, 29, 0, 304, 310, 3, 64, 32, 0, 305, 310, 3, 60, 30,
		0, 306, 310, 3, 62, 31, 0, 307, 310, 3, 66, 33, 0, 308, 310, 3, 56, 28,
		0, 309, 294, 1, 0, 0, 0, 309, 298, 1, 0, 0, 0, 309, 299, 1, 0, 0, 0, 309,
		300, 1, 0, 0, 0, 309, 301, 1, 0, 0, 0, 309, 302, 1, 0, 0, 0, 309, 303,
		1, 0, 0, 0, 309, 304, 1, 0, 0, 0, 309, 305, 1, 0, 0, 0, 309, 306, 1, 0,
		0, 0, 309, 307, 1, 0, 0, 0, 309, 308, 1, 0, 0, 0, 310, 45, 1, 0, 0, 0,
		311, 312, 5, 65, 0, 0, 312, 313, 3, 46, 23, 0, 313, 314, 5, 68, 0, 0, 314,
		319, 1, 0, 0, 0, 315, 316, 5, 34, 0, 0, 316, 317, 7, 0, 0, 0, 317, 319,
		3, 114, 57, 0, 318, 311, 1, 0, 0, 0, 318, 315, 1, 0, 0, 0, 319, 47, 1,
		0, 0, 0, 320, 321, 5, 65, 0, 0, 321, 322, 3, 48, 24, 0, 322, 323, 5, 68,
		0, 0, 323, 328, 1, 0, 0, 0, 324, 325, 5, 33, 0, 0, 325, 326, 7, 0, 0, 0,
		326, 328, 3, 114, 57, 0, 327, 320, 1, 0, 0, 0, 327, 324, 1, 0, 0, 0, 328,
		49, 1, 0, 0, 0, 329, 330, 5, 65, 0, 0, 330, 331, 3, 50, 25, 0, 331, 332,
		5, 68, 0, 0, 332, 337, 1, 0, 0, 0, 333, 334, 5, 35, 0, 0, 334, 335, 7,
		0, 0, 0, 335, 337, 3, 114, 57, 0, 336, 329, 1, 0, 0, 0, 336, 333, 1, 0,
		0, 0, 337, 51, 1, 0, 0, 0, 338, 339, 5, 65, 0, 0, 339, 340, 3, 52, 26,
		0, 340, 341, 5, 68, 0, 0, 341, 349, 1, 0, 0, 0, 342, 343, 5, 36, 0, 0,
		343, 346, 7, 0, 0, 0, 344, 347, 3, 90, 45, 0, 345, 347, 3, 86, 43, 0, 346,
		344, 1, 0, 0, 0, 346, 345, 1, 0, 0, 0, 347, 349, 1, 0, 0, 0, 348, 338,
		1, 0, 0, 0, 348, 342, 1, 0, 0, 0, 349, 53, 1, 0, 0, 0, 350, 351, 5, 65,
		0, 0, 351, 352, 3, 54, 27, 0, 352, 353, 5, 68, 0, 0, 353, 361, 1, 0, 0,
		0, 354, 355, 5, 37, 0, 0, 355, 358, 7, 0, 0, 0, 356, 359, 3, 90, 45, 0,
		357, 359, 3, 86, 43, 0, 358, 356, 1, 0, 0, 0, 358, 357, 1, 0, 0, 0, 359,
		361, 1, 0, 0, 0, 360, 350, 1, 0, 0, 0, 360, 354, 1, 0, 0, 0, 361, 55, 1,
		0, 0, 0, 362, 363, 5, 65, 0, 0, 363, 364, 3, 56, 28, 0, 364, 365, 5, 68,
		0, 0, 365, 370, 1, 0, 0, 0, 366, 367, 5, 41, 0, 0, 367, 368, 7, 0, 0, 0,
		368, 370, 7, 1, 0, 0, 369, 362, 1, 0, 0, 0, 369, 366, 1, 0, 0, 0, 370,
		57, 1, 0, 0, 0, 371, 372, 5, 65, 0, 0, 372, 373, 3, 58, 29, 0, 373, 374,
		5, 68, 0, 0, 374, 379, 1, 0, 0, 0, 375, 376, 5, 39, 0, 0, 376, 377, 7,
		0, 0, 0, 377, 379, 3, 70, 35, 0, 378, 371, 1, 0, 0, 0, 378, 375, 1, 0,
		0, 0, 379, 59, 1, 0, 0, 0, 380, 381, 5, 65, 0, 0, 381, 382, 3, 60, 30,
		0, 382, 383, 5, 68, 0, 0, 383, 391, 1, 0, 0, 0, 384, 385, 5, 42, 0, 0,
		385, 386, 7, 2, 0, 0, 386, 387, 5, 71, 0, 0, 387, 388, 3, 82, 41, 0, 388,
		389, 5, 71, 0, 0, 389, 391, 1, 0, 0, 0, 390, 380, 1, 0, 0, 0, 390, 384,
		1, 0, 0, 0, 391, 61, 1, 0, 0, 0, 392, 393, 5, 65, 0, 0, 393, 394, 3, 62,
		31, 0, 394, 395, 5, 68, 0, 0, 395, 403, 1, 0, 0, 0, 396, 397, 5, 43, 0,
		0, 397, 398, 7, 0, 0, 0, 398, 399, 5, 71, 0, 0, 399, 400, 3, 68, 34, 0,
		400, 401, 5, 71, 0, 0, 401, 403, 1, 0, 0, 0, 402, 392, 1, 0, 0, 0, 402,
		396, 1, 0, 0, 0, 403, 63, 1, 0, 0, 0, 404, 405, 5, 65, 0, 0, 405, 406,
		3, 64, 32, 0, 406, 407, 5, 68, 0, 0, 407, 415, 1, 0, 0, 0, 408, 409, 5,
		40, 0, 0, 409, 412, 7, 0, 0, 0, 410, 413, 3, 78, 39, 0, 411, 413, 3, 76,
		38, 0, 412, 410, 1, 0, 0, 0, 412, 411, 1, 0, 0, 0, 413, 415, 1, 0, 0, 0,
		414, 404, 1, 0, 0, 0, 414, 408, 1, 0, 0, 0, 415, 65, 1, 0, 0, 0, 416, 417,
		5, 65, 0, 0, 417, 418, 3, 66, 33, 0, 418, 419, 5, 68, 0, 0, 419, 424, 1,
		0, 0, 0, 420, 421, 5, 38, 0, 0, 421, 422, 7, 2, 0, 0, 422, 424, 3, 80,
		40, 0, 423, 416, 1, 0, 0, 0, 423, 420, 1, 0, 0, 0, 424, 67, 1, 0, 0, 0,
		425, 430, 5, 19, 0, 0, 426, 427, 5, 72, 0, 0, 427, 429, 5, 19, 0, 0, 428,
		426, 1, 0, 0, 0, 429, 432, 1, 0, 0, 0, 430, 428, 1, 0, 0, 0, 430, 431,
		1, 0, 0, 0, 431, 434, 1, 0, 0, 0, 432, 430, 1, 0, 0, 0, 433, 425, 1, 0,
		0, 0, 433, 434, 1, 0, 0, 0, 434, 69, 1, 0, 0, 0, 435, 436, 5, 85, 0, 0,
		436, 71, 1, 0, 0, 0, 437, 441, 3, 74, 37, 0, 438, 439, 3, 102, 51, 0, 439,
		440, 3, 74, 37, 0, 440, 442, 1, 0, 0, 0, 441, 438, 1, 0, 0, 0, 442, 443,
		1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 443, 444, 1, 0, 0, 0, 444, 73, 1, 0,
		0, 0, 445, 446, 5, 85, 0, 0, 446, 75, 1, 0, 0, 0, 447, 448, 5, 85, 0, 0,
		448, 77, 1, 0, 0, 0, 449, 450, 5, 85, 0, 0, 450, 79, 1, 0, 0, 0, 451, 452,
		5, 2, 0, 0, 452, 81, 1, 0, 0, 0, 453, 454, 5, 2, 0, 0, 454, 83, 1, 0, 0,
		0, 455, 456, 5, 2, 0, 0, 456, 85, 1, 0, 0, 0, 457, 458, 5, 52, 0, 0, 458,
		459, 3, 88, 44, 0, 459, 460, 5, 69, 0, 0, 460, 461, 5, 77, 0, 0, 461, 462,
		3, 90, 45, 0, 462, 87, 1, 0, 0, 0, 463, 464, 5, 86, 0, 0, 464, 89, 1, 0,
		0, 0, 465, 466, 3, 124, 62, 0, 466, 469, 5, 74, 0, 0, 467, 470, 5, 51,
		0, 0, 468, 470, 3, 126, 63, 0, 469, 467, 1, 0, 0, 0, 469, 468, 1, 0, 0,
		0, 470, 91, 1, 0, 0, 0, 471, 474, 3, 94, 47, 0, 472, 473, 7, 3, 0, 0, 473,
		475, 3, 94, 47, 0, 474, 472, 1, 0, 0, 0, 475, 476, 1, 0, 0, 0, 476, 477,
		1, 0, 0, 0, 476, 474, 1, 0, 0, 0, 477, 93, 1, 0, 0, 0, 478, 479, 7, 4,
		0, 0, 479, 489, 5, 55, 0, 0, 480, 486, 3, 98, 49, 0, 481, 482, 3, 96, 48,
		0, 482, 483, 3, 98, 49, 0, 483, 485, 1, 0, 0, 0, 484, 481, 1, 0, 0, 0,
		485, 488, 1, 0, 0, 0, 486, 484, 1, 0, 0, 0, 486, 487, 1, 0, 0, 0, 487,
		490, 1, 0, 0, 0, 488, 486, 1, 0, 0, 0, 489, 480, 1, 0, 0, 0, 489, 490,
		1, 0, 0, 0, 490, 95, 1, 0, 0, 0, 491, 492, 5, 13, 0, 0, 492, 97, 1, 0,
		0, 0, 493, 494, 3, 124, 62, 0, 494, 495, 5, 4, 0, 0, 495, 496, 3, 118,
		59, 0, 496, 99, 1, 0, 0, 0, 497, 503, 3, 114, 57, 0, 498, 499, 3, 102,
		51, 0, 499, 500, 3, 114, 57, 0, 500, 502, 1, 0, 0, 0, 501, 498, 1, 0, 0,
		0, 502, 505, 1, 0, 0, 0, 503, 501, 1, 0, 0, 0, 503, 504, 1, 0, 0, 0, 504,
		507, 1, 0, 0, 0, 505, 503, 1, 0, 0, 0, 506, 497, 1, 0, 0, 0, 507, 508,
		1, 0, 0, 0, 508, 506, 1, 0, 0, 0, 508, 509, 1, 0, 0, 0, 509, 101, 1, 0,
		0, 0, 510, 511, 5, 14, 0, 0, 511, 103, 1, 0, 0, 0, 512, 513, 3, 114, 57,
		0, 513, 514, 5, 80, 0, 0, 514, 515, 3, 90, 45, 0, 515, 105, 1, 0, 0, 0,
		516, 517, 3, 114, 57, 0, 517, 518, 3, 112, 56, 0, 518, 519, 3, 110, 55,
		0, 519, 520, 3, 108, 54, 0, 520, 107, 1, 0, 0, 0, 521, 522, 5, 80, 0, 0,
		522, 523, 3, 118, 59, 0, 523, 109, 1, 0, 0, 0, 524, 526, 5, 80, 0, 0, 525,
		527, 5, 50, 0, 0, 526, 525, 1, 0, 0, 0, 526, 527, 1, 0, 0, 0, 527, 111,
		1, 0, 0, 0, 528, 537, 5, 80, 0, 0, 529, 534, 3, 124, 62, 0, 530, 531, 5,
		72, 0, 0, 531, 533, 3, 124, 62, 0, 532, 530, 1, 0, 0, 0, 533, 536, 1, 0,
		0, 0, 534, 532, 1, 0, 0, 0, 534, 535, 1, 0, 0, 0, 535, 538, 1, 0, 0, 0,
		536, 534, 1, 0, 0, 0, 537, 529, 1, 0, 0, 0, 537, 538, 1, 0, 0, 0, 538,
		113, 1, 0, 0, 0, 539, 540, 5, 53, 0, 0, 540, 545, 3, 116, 58, 0, 541, 542,
		5, 72, 0, 0, 542, 544, 3, 116, 58, 0, 543, 541, 1, 0, 0, 0, 544, 547, 1,
		0, 0, 0, 545, 543, 1, 0, 0, 0, 545, 546, 1, 0, 0, 0, 546, 115, 1, 0, 0,
		0, 547, 545, 1, 0, 0, 0, 548, 549, 3, 124, 62, 0, 549, 550, 5, 55, 0, 0,
		550, 551, 3, 126, 63, 0, 551, 554, 1, 0, 0, 0, 552, 554, 5, 89, 0, 0, 553,
		548, 1, 0, 0, 0, 553, 552, 1, 0, 0, 0, 554, 117, 1, 0, 0, 0, 555, 556,
		5, 65, 0, 0, 556, 557, 3, 120, 60, 0, 557, 558, 5, 68, 0, 0, 558, 566,
		1, 0, 0, 0, 559, 561, 3, 120, 60, 0, 560, 559, 1, 0, 0, 0, 561, 564, 1,
		0, 0, 0, 562, 560, 1, 0, 0, 0, 562, 563, 1, 0, 0, 0, 563, 566, 1, 0, 0,
		0, 564, 562, 1, 0, 0, 0, 565, 555, 1, 0, 0, 0, 565, 562, 1, 0, 0, 0, 566,
		119, 1, 0, 0, 0, 567, 569, 5, 65, 0, 0, 568, 570, 7, 5, 0, 0, 569, 568,
		1, 0, 0, 0, 569, 570, 1, 0, 0, 0, 570, 571, 1, 0, 0, 0, 571, 572, 3, 120,
		60, 0, 572, 573, 5, 68, 0, 0, 573, 575, 1, 0, 0, 0, 574, 567, 1, 0, 0,
		0, 575, 576, 1, 0, 0, 0, 576, 577, 1, 0, 0, 0, 576, 574, 1, 0, 0, 0, 577,
		582, 1, 0, 0, 0, 578, 579, 5, 63, 0, 0, 579, 582, 3, 120, 60, 0, 580, 582,
		3, 122, 61, 0, 581, 574, 1, 0, 0, 0, 581, 578, 1, 0, 0, 0, 581, 580, 1,
		0, 0, 0, 582, 121, 1, 0, 0, 0, 583, 584, 3, 124, 62, 0, 584, 585, 3, 128,
		64, 0, 585, 586, 3, 126, 63, 0, 586, 123, 1, 0, 0, 0, 587, 588, 5, 3, 0,
		0, 588, 125, 1, 0, 0, 0, 589, 590, 7, 6, 0, 0, 590, 127, 1, 0, 0, 0, 591,
		592, 7, 7, 0, 0, 592, 129, 1, 0, 0, 0, 45, 148, 162, 165, 172, 184, 240,
		260, 269, 272, 280, 292, 309, 318, 327, 336, 346, 348, 358, 360, 369, 378,
		390, 402, 412, 414, 423, 430, 433, 443, 469, 476, 486, 489, 503, 508, 526,
		534, 537, 545, 553, 562, 565, 569, 576, 581,
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
	ACIParserEOF                      = antlr.TokenEOF
	ACIParserLEVELS                   = 1
	ACIParserINT                      = 2
	ACIParserID                       = 3
	ACIParserCOLON                    = 4
	ACIParserExtensibleRuleDNMatch    = 5
	ACIParserExtensibleRuleDNOIDMatch = 6
	ACIParserExtensibleRuleAttrMatch  = 7
	ACIParserExtensibleRuleMatch      = 8
	ACIParserApproximateMatch         = 9
	ACIParserNotEqualTo               = 10
	ACIParserGreaterThanOrEqual       = 11
	ACIParserLessThanOrEqual          = 12
	ACIParserDAMP                     = 13
	ACIParserDPIPE                    = 14
	ACIParserBooleanWordOperators     = 15
	ACIParserBooleanAndWord           = 16
	ACIParserBooleanOrWord            = 17
	ACIParserBooleanNotWord           = 18
	ACIParserDAYS                     = 19
	ACIParserTargetKeyword            = 20
	ACIParserTargetScopeKeyword       = 21
	ACIParserTargetToKeyword          = 22
	ACIParserTargetFromKeyword        = 23
	ACIParserTargetAttrKeyword        = 24
	ACIParserTargetFilterKeyword      = 25
	ACIParserTargetAttrFiltersKeyword = 26
	ACIParserTargetControlKeyword     = 27
	ACIParserExtOpKeyword             = 28
	ACIParserANONYMOUS                = 29
	ACIParserSIMPLE                   = 30
	ACIParserSASL                     = 31
	ACIParserSSL                      = 32
	ACIParserRoleDNKeyword            = 33
	ACIParserUserDNKeyword            = 34
	ACIParserGroupDNKeyword           = 35
	ACIParserUserAttrKeyword          = 36
	ACIParserGroupAttrKeyword         = 37
	ACIParserSSFKeyword               = 38
	ACIParserDNSKeyword               = 39
	ACIParserIPKeyword                = 40
	ACIParserAuthMethodKeyword        = 41
	ACIParserTimeOfDayKeyword         = 42
	ACIParserDayOfWeekKeyword         = 43
	ACIParserDISPOSITION              = 44
	ACIParserRIGHTS                   = 45
	ACIParserAttributeFilterOperation = 46
	ACIParserAddOperation             = 47
	ACIParserDeleteOperation          = 48
	ACIParserTargetRuleSearchScopes   = 49
	ACIParserLDAPSearchScopes         = 50
	ACIParserBINDTYPES                = 51
	ACIParserINHERITANCEPREFIX        = 52
	ACIParserLocalLDAPScheme          = 53
	ACIParserSEMI                     = 54
	ACIParserEqualTo                  = 55
	ACIParserGreaterThan              = 56
	ACIParserLessThan                 = 57
	ACIParserLDAP                     = 58
	ACIParserANCHOR                   = 59
	ACIParserDN                       = 60
	ACIParserAMP                      = 61
	ACIParserPIPE                     = 62
	ACIParserBANG                     = 63
	ACIParserSOLIDUS                  = 64
	ACIParserLPAREN                   = 65
	ACIParserLBRAC                    = 66
	ACIParserLBRAK                    = 67
	ACIParserRPAREN                   = 68
	ACIParserRBRAK                    = 69
	ACIParserRBRAC                    = 70
	ACIParserDQUOTE                   = 71
	ACIParserCOMMA                    = 72
	ACIParserTILDE                    = 73
	ACIParserHASH                     = 74
	ACIParserDOLLAR                   = 75
	ACIParserATSIGN                   = 76
	ACIParserDOT                      = 77
	ACIParserDASH                     = 78
	ACIParserSTAR                     = 79
	ACIParserQMARK                    = 80
	ACIParserLineTerminator           = 81
	ACIParserWhiteSpaces              = 82
	ACIParserNumericLiteral           = 83
	ACIParserLiteral                  = 84
	ACIParserDelimitedAddress         = 85
	ACIParserDelimitedNumbers         = 86
	ACIParserStringLiteral            = 87
	ACIParserWildcardString           = 88
	ACIParserMacroValue               = 89
	ACIParserANY                      = 90
)

// ACIParser rules.
const (
	ACIParserRULE_parse                          = 0
	ACIParserRULE_instruction                    = 1
	ACIParserRULE_version                        = 2
	ACIParserRULE_permBindRules                  = 3
	ACIParserRULE_permBindRule                   = 4
	ACIParserRULE_permission                     = 5
	ACIParserRULE_targetRules                    = 6
	ACIParserRULE_targetRule                     = 7
	ACIParserRULE_targetControl                  = 8
	ACIParserRULE_targetExtOp                    = 9
	ACIParserRULE_targetFilter                   = 10
	ACIParserRULE_targetFilterValue              = 11
	ACIParserRULE_target                         = 12
	ACIParserRULE_targetTo                       = 13
	ACIParserRULE_targetFrom                     = 14
	ACIParserRULE_targetAttrFilters              = 15
	ACIParserRULE_targetAttrFiltersValue         = 16
	ACIParserRULE_targetScope                    = 17
	ACIParserRULE_targetAttr                     = 18
	ACIParserRULE_attributeTypes                 = 19
	ACIParserRULE_bindRule                       = 20
	ACIParserRULE_bindRuleExprParen              = 21
	ACIParserRULE_bindRuleExpr                   = 22
	ACIParserRULE_bindRuleUserDN                 = 23
	ACIParserRULE_bindRuleRoleDN                 = 24
	ACIParserRULE_bindRuleGroupDN                = 25
	ACIParserRULE_bindRuleUserAttr               = 26
	ACIParserRULE_bindRuleGroupAttr              = 27
	ACIParserRULE_bindRuleAuthMethod             = 28
	ACIParserRULE_bindRuleDNS                    = 29
	ACIParserRULE_bindRuleTimeOfDay              = 30
	ACIParserRULE_bindRuleDayOfWeek              = 31
	ACIParserRULE_bindRuleIP                     = 32
	ACIParserRULE_bindRuleSecurityStrengthFactor = 33
	ACIParserRULE_dayOfWeek                      = 34
	ACIParserRULE_fQDN                           = 35
	ACIParserRULE_objectIdentifiers              = 36
	ACIParserRULE_objectIdentifier               = 37
	ACIParserRULE_iPV6Address                    = 38
	ACIParserRULE_iPV4Address                    = 39
	ACIParserRULE_securityStrengthFactor         = 40
	ACIParserRULE_timeOfDay                      = 41
	ACIParserRULE_objectIdentifierArc            = 42
	ACIParserRULE_inheritance                    = 43
	ACIParserRULE_inheritanceLevels              = 44
	ACIParserRULE_attributeBindTypeOrValue       = 45
	ACIParserRULE_attributeFilters               = 46
	ACIParserRULE_attributeFilterSet             = 47
	ACIParserRULE_doubleAmpersand                = 48
	ACIParserRULE_attributeFilter                = 49
	ACIParserRULE_distinguishedNames             = 50
	ACIParserRULE_doublePipe                     = 51
	ACIParserRULE_lDAPURIAndBindType             = 52
	ACIParserRULE_lDAPURI                        = 53
	ACIParserRULE_uRISearchFilter                = 54
	ACIParserRULE_uRISearchScopes                = 55
	ACIParserRULE_uRIAttributeList               = 56
	ACIParserRULE_distinguishedName              = 57
	ACIParserRULE_relativeDistinguishedName      = 58
	ACIParserRULE_lDAPFilter                     = 59
	ACIParserRULE_lDAPFilterExpr                 = 60
	ACIParserRULE_attributeValueAssertion        = 61
	ACIParserRULE_attributeType                  = 62
	ACIParserRULE_attributeValue                 = 63
	ACIParserRULE_attributeOperators             = 64
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
		p.SetState(130)
		p.Instruction()
	}
	{
		p.SetState(131)
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

func (s *InstructionContext) CopyAll(ctx *InstructionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AccessControlInstructionContext struct {
	InstructionContext
}

func NewAccessControlInstructionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccessControlInstructionContext {
	var p = new(AccessControlInstructionContext)

	InitEmptyInstructionContext(&p.InstructionContext)
	p.parser = parser
	p.CopyAll(ctx.(*InstructionContext))

	return p
}

func (s *AccessControlInstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessControlInstructionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *AccessControlInstructionContext) TargetRule() ITargetRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetRuleContext)
}

func (s *AccessControlInstructionContext) Version() IVersionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVersionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVersionContext)
}

func (s *AccessControlInstructionContext) PermBindRules() IPermBindRulesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPermBindRulesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPermBindRulesContext)
}

func (s *AccessControlInstructionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *AccessControlInstructionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ACIParserSEMI, 0)
}

func (s *AccessControlInstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAccessControlInstruction(s)
	}
}

func (s *AccessControlInstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAccessControlInstruction(s)
	}
}

func (p *ACIParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ACIParserRULE_instruction)
	localctx = NewAccessControlInstructionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(133)
		p.Match(ACIParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(134)
		p.TargetRule()
	}
	{
		p.SetState(135)
		p.Version()
	}
	{
		p.SetState(136)
		p.PermBindRules()
	}
	{
		p.SetState(137)
		p.Match(ACIParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
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

// IVersionContext is an interface to support dynamic dispatch.
type IVersionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVersionContext differentiates from other interfaces.
	IsVersionContext()
}

type VersionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVersionContext() *VersionContext {
	var p = new(VersionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_version
	return p
}

func InitEmptyVersionContext(p *VersionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_version
}

func (*VersionContext) IsVersionContext() {}

func NewVersionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VersionContext {
	var p = new(VersionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_version

	return p
}

func (s *VersionContext) GetParser() antlr.Parser { return s.parser }

func (s *VersionContext) CopyAll(ctx *VersionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VersionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AccessControlInstructionNameAndVersionContext struct {
	VersionContext
}

func NewAccessControlInstructionNameAndVersionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AccessControlInstructionNameAndVersionContext {
	var p = new(AccessControlInstructionNameAndVersionContext)

	InitEmptyVersionContext(&p.VersionContext)
	p.parser = parser
	p.CopyAll(ctx.(*VersionContext))

	return p
}

func (s *AccessControlInstructionNameAndVersionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AccessControlInstructionNameAndVersionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *AccessControlInstructionNameAndVersionContext) ANCHOR() antlr.TerminalNode {
	return s.GetToken(ACIParserANCHOR, 0)
}

func (s *AccessControlInstructionNameAndVersionContext) Literal() antlr.TerminalNode {
	return s.GetToken(ACIParserLiteral, 0)
}

func (s *AccessControlInstructionNameAndVersionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ACIParserSEMI, 0)
}

func (s *AccessControlInstructionNameAndVersionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAccessControlInstructionNameAndVersion(s)
	}
}

func (s *AccessControlInstructionNameAndVersionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAccessControlInstructionNameAndVersion(s)
	}
}

func (p *ACIParser) Version() (localctx IVersionContext) {
	localctx = NewVersionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ACIParserRULE_version)
	localctx = NewAccessControlInstructionNameAndVersionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(140)
		p.Match(ACIParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(141)
		p.Match(ACIParserANCHOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(142)
		p.Match(ACIParserLiteral)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(143)
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

// IPermBindRulesContext is an interface to support dynamic dispatch.
type IPermBindRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPermBindRulesContext differentiates from other interfaces.
	IsPermBindRulesContext()
}

type PermBindRulesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPermBindRulesContext() *PermBindRulesContext {
	var p = new(PermBindRulesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_permBindRules
	return p
}

func InitEmptyPermBindRulesContext(p *PermBindRulesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_permBindRules
}

func (*PermBindRulesContext) IsPermBindRulesContext() {}

func NewPermBindRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PermBindRulesContext {
	var p = new(PermBindRulesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_permBindRules

	return p
}

func (s *PermBindRulesContext) GetParser() antlr.Parser { return s.parser }

func (s *PermBindRulesContext) CopyAll(ctx *PermBindRulesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PermBindRulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PermBindRulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PermissionBindRulesContext struct {
	PermBindRulesContext
}

func NewPermissionBindRulesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PermissionBindRulesContext {
	var p = new(PermissionBindRulesContext)

	InitEmptyPermBindRulesContext(&p.PermBindRulesContext)
	p.parser = parser
	p.CopyAll(ctx.(*PermBindRulesContext))

	return p
}

func (s *PermissionBindRulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PermissionBindRulesContext) AllPermBindRule() []IPermBindRuleContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPermBindRuleContext); ok {
			len++
		}
	}

	tst := make([]IPermBindRuleContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPermBindRuleContext); ok {
			tst[i] = t.(IPermBindRuleContext)
			i++
		}
	}

	return tst
}

func (s *PermissionBindRulesContext) PermBindRule(i int) IPermBindRuleContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPermBindRuleContext); ok {
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

	return t.(IPermBindRuleContext)
}

func (s *PermissionBindRulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterPermissionBindRules(s)
	}
}

func (s *PermissionBindRulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitPermissionBindRules(s)
	}
}

func (p *ACIParser) PermBindRules() (localctx IPermBindRulesContext) {
	localctx = NewPermBindRulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ACIParserRULE_permBindRules)
	var _la int

	localctx = NewPermissionBindRulesContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(148)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ACIParserDISPOSITION {
		{
			p.SetState(145)
			p.PermBindRule()
		}

		p.SetState(150)
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

// IPermBindRuleContext is an interface to support dynamic dispatch.
type IPermBindRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPermBindRuleContext differentiates from other interfaces.
	IsPermBindRuleContext()
}

type PermBindRuleContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPermBindRuleContext() *PermBindRuleContext {
	var p = new(PermBindRuleContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_permBindRule
	return p
}

func InitEmptyPermBindRuleContext(p *PermBindRuleContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_permBindRule
}

func (*PermBindRuleContext) IsPermBindRuleContext() {}

func NewPermBindRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PermBindRuleContext {
	var p = new(PermBindRuleContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_permBindRule

	return p
}

func (s *PermBindRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *PermBindRuleContext) CopyAll(ctx *PermBindRuleContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PermBindRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PermBindRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PermissionBindRuleContext struct {
	PermBindRuleContext
}

func NewPermissionBindRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PermissionBindRuleContext {
	var p = new(PermissionBindRuleContext)

	InitEmptyPermBindRuleContext(&p.PermBindRuleContext)
	p.parser = parser
	p.CopyAll(ctx.(*PermBindRuleContext))

	return p
}

func (s *PermissionBindRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

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

func (s *PermissionBindRuleContext) BindRule() IBindRuleContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleContext)
}

func (s *PermissionBindRuleContext) SEMI() antlr.TerminalNode {
	return s.GetToken(ACIParserSEMI, 0)
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

func (p *ACIParser) PermBindRule() (localctx IPermBindRuleContext) {
	localctx = NewPermBindRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ACIParserRULE_permBindRule)
	localctx = NewPermissionBindRuleContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Permission()
	}
	{
		p.SetState(152)
		p.BindRule()
	}
	{
		p.SetState(153)
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

func (s *PermissionContext) CopyAll(ctx *PermissionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PermissionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PermissionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PermissionExpressionContext struct {
	PermissionContext
}

func NewPermissionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PermissionExpressionContext {
	var p = new(PermissionExpressionContext)

	InitEmptyPermissionContext(&p.PermissionContext)
	p.parser = parser
	p.CopyAll(ctx.(*PermissionContext))

	return p
}

func (s *PermissionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PermissionExpressionContext) DISPOSITION() antlr.TerminalNode {
	return s.GetToken(ACIParserDISPOSITION, 0)
}

func (s *PermissionExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *PermissionExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *PermissionExpressionContext) AllRIGHTS() []antlr.TerminalNode {
	return s.GetTokens(ACIParserRIGHTS)
}

func (s *PermissionExpressionContext) RIGHTS(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserRIGHTS, i)
}

func (s *PermissionExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ACIParserCOMMA)
}

func (s *PermissionExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserCOMMA, i)
}

func (s *PermissionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterPermissionExpression(s)
	}
}

func (s *PermissionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitPermissionExpression(s)
	}
}

func (p *ACIParser) Permission() (localctx IPermissionContext) {
	localctx = NewPermissionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ACIParserRULE_permission)
	var _la int

	localctx = NewPermissionExpressionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(ACIParserDISPOSITION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(156)
		p.Match(ACIParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ACIParserRIGHTS {
		{
			p.SetState(157)
			p.Match(ACIParserRIGHTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(162)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ACIParserCOMMA {
			{
				p.SetState(158)
				p.Match(ACIParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(159)
				p.Match(ACIParserRIGHTS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(164)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(167)
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

// ITargetRulesContext is an interface to support dynamic dispatch.
type ITargetRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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

func (s *TargetRulesContext) CopyAll(ctx *TargetRulesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TargetRulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetRulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TargetRuleExpressionsContext struct {
	TargetRulesContext
}

func NewTargetRuleExpressionsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetRuleExpressionsContext {
	var p = new(TargetRuleExpressionsContext)

	InitEmptyTargetRulesContext(&p.TargetRulesContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetRulesContext))

	return p
}

func (s *TargetRuleExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetRuleExpressionsContext) AllTargetRule() []ITargetRuleContext {
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

func (s *TargetRuleExpressionsContext) TargetRule(i int) ITargetRuleContext {
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

func (s *TargetRuleExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetRuleExpressions(s)
	}
}

func (s *TargetRuleExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetRuleExpressions(s)
	}
}

func (p *ACIParser) TargetRules() (localctx ITargetRulesContext) {
	localctx = NewTargetRulesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ACIParserRULE_targetRules)
	var _la int

	localctx = NewTargetRuleExpressionsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ACIParserLPAREN {
		{
			p.SetState(169)
			p.TargetRule()
		}

		p.SetState(174)
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

// ITargetRuleContext is an interface to support dynamic dispatch.
type ITargetRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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

func (s *TargetRuleContext) CopyAll(ctx *TargetRuleContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TargetRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TargetscopeContext struct {
	TargetRuleContext
}

func NewTargetscopeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetscopeContext {
	var p = new(TargetscopeContext)

	InitEmptyTargetRuleContext(&p.TargetRuleContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetRuleContext))

	return p
}

func (s *TargetscopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetscopeContext) TargetScope() ITargetScopeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetScopeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetScopeContext)
}

func (s *TargetscopeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetscope(s)
	}
}

func (s *TargetscopeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetscope(s)
	}
}

type TargetdnContext struct {
	TargetRuleContext
}

func NewTargetdnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetdnContext {
	var p = new(TargetdnContext)

	InitEmptyTargetRuleContext(&p.TargetRuleContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetRuleContext))

	return p
}

func (s *TargetdnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetdnContext) Target() ITargetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *TargetdnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetdn(s)
	}
}

func (s *TargetdnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetdn(s)
	}
}

type TargetfromDNContext struct {
	TargetRuleContext
}

func NewTargetfromDNContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetfromDNContext {
	var p = new(TargetfromDNContext)

	InitEmptyTargetRuleContext(&p.TargetRuleContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetRuleContext))

	return p
}

func (s *TargetfromDNContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetfromDNContext) TargetFrom() ITargetFromContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetFromContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetFromContext)
}

func (s *TargetfromDNContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetfromDN(s)
	}
}

func (s *TargetfromDNContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetfromDN(s)
	}
}

type ExtopContext struct {
	TargetRuleContext
}

func NewExtopContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExtopContext {
	var p = new(ExtopContext)

	InitEmptyTargetRuleContext(&p.TargetRuleContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetRuleContext))

	return p
}

func (s *ExtopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtopContext) TargetExtOp() ITargetExtOpContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetExtOpContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetExtOpContext)
}

func (s *ExtopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterExtop(s)
	}
}

func (s *ExtopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitExtop(s)
	}
}

type TargetattrContext struct {
	TargetRuleContext
}

func NewTargetattrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetattrContext {
	var p = new(TargetattrContext)

	InitEmptyTargetRuleContext(&p.TargetRuleContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetRuleContext))

	return p
}

func (s *TargetattrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetattrContext) TargetAttr() ITargetAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetAttrContext)
}

func (s *TargetattrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetattr(s)
	}
}

func (s *TargetattrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetattr(s)
	}
}

type TargattrfiltersContext struct {
	TargetRuleContext
}

func NewTargattrfiltersContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargattrfiltersContext {
	var p = new(TargattrfiltersContext)

	InitEmptyTargetRuleContext(&p.TargetRuleContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetRuleContext))

	return p
}

func (s *TargattrfiltersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargattrfiltersContext) TargetAttrFilters() ITargetAttrFiltersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetAttrFiltersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetAttrFiltersContext)
}

func (s *TargattrfiltersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargattrfilters(s)
	}
}

func (s *TargattrfiltersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargattrfilters(s)
	}
}

type TargettodnContext struct {
	TargetRuleContext
}

func NewTargettodnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargettodnContext {
	var p = new(TargettodnContext)

	InitEmptyTargetRuleContext(&p.TargetRuleContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetRuleContext))

	return p
}

func (s *TargettodnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargettodnContext) TargetTo() ITargetToContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetToContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetToContext)
}

func (s *TargettodnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargettodn(s)
	}
}

func (s *TargettodnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargettodn(s)
	}
}

type TargetcontrolContext struct {
	TargetRuleContext
}

func NewTargetcontrolContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetcontrolContext {
	var p = new(TargetcontrolContext)

	InitEmptyTargetRuleContext(&p.TargetRuleContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetRuleContext))

	return p
}

func (s *TargetcontrolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetcontrolContext) TargetControl() ITargetControlContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetControlContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetControlContext)
}

func (s *TargetcontrolContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetcontrol(s)
	}
}

func (s *TargetcontrolContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetcontrol(s)
	}
}

type TargetfilterContext struct {
	TargetRuleContext
}

func NewTargetfilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetfilterContext {
	var p = new(TargetfilterContext)

	InitEmptyTargetRuleContext(&p.TargetRuleContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetRuleContext))

	return p
}

func (s *TargetfilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetfilterContext) TargetFilter() ITargetFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetFilterContext)
}

func (s *TargetfilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetfilter(s)
	}
}

func (s *TargetfilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetfilter(s)
	}
}

func (p *ACIParser) TargetRule() (localctx ITargetRuleContext) {
	localctx = NewTargetRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ACIParserRULE_targetRule)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		localctx = NewTargetcontrolContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(175)
			p.TargetControl()
		}

	case 2:
		localctx = NewExtopContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(176)
			p.TargetExtOp()
		}

	case 3:
		localctx = NewTargetfilterContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(177)
			p.TargetFilter()
		}

	case 4:
		localctx = NewTargattrfiltersContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(178)
			p.TargetAttrFilters()
		}

	case 5:
		localctx = NewTargetscopeContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(179)
			p.TargetScope()
		}

	case 6:
		localctx = NewTargetattrContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(180)
			p.TargetAttr()
		}

	case 7:
		localctx = NewTargetdnContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(181)
			p.Target()
		}

	case 8:
		localctx = NewTargettodnContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(182)
			p.TargetTo()
		}

	case 9:
		localctx = NewTargetfromDNContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(183)
			p.TargetFrom()
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

// ITargetControlContext is an interface to support dynamic dispatch.
type ITargetControlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTargetControlContext differentiates from other interfaces.
	IsTargetControlContext()
}

type TargetControlContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetControlContext() *TargetControlContext {
	var p = new(TargetControlContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetControl
	return p
}

func InitEmptyTargetControlContext(p *TargetControlContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetControl
}

func (*TargetControlContext) IsTargetControlContext() {}

func NewTargetControlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetControlContext {
	var p = new(TargetControlContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetControl

	return p
}

func (s *TargetControlContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetControlContext) CopyAll(ctx *TargetControlContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TargetControlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetControlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentheticalControlsContext struct {
	TargetControlContext
}

func NewParentheticalControlsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalControlsContext {
	var p = new(ParentheticalControlsContext)

	InitEmptyTargetControlContext(&p.TargetControlContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetControlContext))

	return p
}

func (s *ParentheticalControlsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalControlsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalControlsContext) TargetControlKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserTargetControlKeyword, 0)
}

func (s *ParentheticalControlsContext) ObjectIdentifiers() IObjectIdentifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectIdentifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectIdentifiersContext)
}

func (s *ParentheticalControlsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalControlsContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *ParentheticalControlsContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *ParentheticalControlsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalControls(s)
	}
}

func (s *ParentheticalControlsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalControls(s)
	}
}

func (p *ACIParser) TargetControl() (localctx ITargetControlContext) {
	localctx = NewTargetControlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ACIParserRULE_targetControl)
	var _la int

	localctx = NewParentheticalControlsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)
		p.Match(ACIParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(187)
		p.Match(ACIParserTargetControlKeyword)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(188)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(189)
		p.ObjectIdentifiers()
	}
	{
		p.SetState(190)
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

// ITargetExtOpContext is an interface to support dynamic dispatch.
type ITargetExtOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTargetExtOpContext differentiates from other interfaces.
	IsTargetExtOpContext()
}

type TargetExtOpContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetExtOpContext() *TargetExtOpContext {
	var p = new(TargetExtOpContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetExtOp
	return p
}

func InitEmptyTargetExtOpContext(p *TargetExtOpContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetExtOp
}

func (*TargetExtOpContext) IsTargetExtOpContext() {}

func NewTargetExtOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetExtOpContext {
	var p = new(TargetExtOpContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetExtOp

	return p
}

func (s *TargetExtOpContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetExtOpContext) CopyAll(ctx *TargetExtOpContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TargetExtOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetExtOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentheticalExtendedOperationsContext struct {
	TargetExtOpContext
}

func NewParentheticalExtendedOperationsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalExtendedOperationsContext {
	var p = new(ParentheticalExtendedOperationsContext)

	InitEmptyTargetExtOpContext(&p.TargetExtOpContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetExtOpContext))

	return p
}

func (s *ParentheticalExtendedOperationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalExtendedOperationsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalExtendedOperationsContext) ExtOpKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserExtOpKeyword, 0)
}

func (s *ParentheticalExtendedOperationsContext) ObjectIdentifiers() IObjectIdentifiersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectIdentifiersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectIdentifiersContext)
}

func (s *ParentheticalExtendedOperationsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalExtendedOperationsContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *ParentheticalExtendedOperationsContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *ParentheticalExtendedOperationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalExtendedOperations(s)
	}
}

func (s *ParentheticalExtendedOperationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalExtendedOperations(s)
	}
}

func (p *ACIParser) TargetExtOp() (localctx ITargetExtOpContext) {
	localctx = NewTargetExtOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ACIParserRULE_targetExtOp)
	var _la int

	localctx = NewParentheticalExtendedOperationsContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(192)
		p.Match(ACIParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(193)
		p.Match(ACIParserExtOpKeyword)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(194)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(195)
		p.ObjectIdentifiers()
	}
	{
		p.SetState(196)
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

// ITargetFilterContext is an interface to support dynamic dispatch.
type ITargetFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTargetFilterContext differentiates from other interfaces.
	IsTargetFilterContext()
}

type TargetFilterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetFilterContext() *TargetFilterContext {
	var p = new(TargetFilterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetFilter
	return p
}

func InitEmptyTargetFilterContext(p *TargetFilterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetFilter
}

func (*TargetFilterContext) IsTargetFilterContext() {}

func NewTargetFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetFilterContext {
	var p = new(TargetFilterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetFilter

	return p
}

func (s *TargetFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetFilterContext) CopyAll(ctx *TargetFilterContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TargetFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentheticalTargetFilterExpressionContext struct {
	TargetFilterContext
}

func NewParentheticalTargetFilterExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalTargetFilterExpressionContext {
	var p = new(ParentheticalTargetFilterExpressionContext)

	InitEmptyTargetFilterContext(&p.TargetFilterContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetFilterContext))

	return p
}

func (s *ParentheticalTargetFilterExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalTargetFilterExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalTargetFilterExpressionContext) TargetFilterKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserTargetFilterKeyword, 0)
}

func (s *ParentheticalTargetFilterExpressionContext) TargetFilterValue() ITargetFilterValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetFilterValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetFilterValueContext)
}

func (s *ParentheticalTargetFilterExpressionContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *ParentheticalTargetFilterExpressionContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *ParentheticalTargetFilterExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalTargetFilterExpression(s)
	}
}

func (s *ParentheticalTargetFilterExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalTargetFilterExpression(s)
	}
}

func (p *ACIParser) TargetFilter() (localctx ITargetFilterContext) {
	localctx = NewTargetFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ACIParserRULE_targetFilter)
	var _la int

	localctx = NewParentheticalTargetFilterExpressionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(198)
		p.Match(ACIParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(199)
		p.Match(ACIParserTargetFilterKeyword)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(201)
		p.TargetFilterValue()
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

// ITargetFilterValueContext is an interface to support dynamic dispatch.
type ITargetFilterValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTargetFilterValueContext differentiates from other interfaces.
	IsTargetFilterValueContext()
}

type TargetFilterValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetFilterValueContext() *TargetFilterValueContext {
	var p = new(TargetFilterValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetFilterValue
	return p
}

func InitEmptyTargetFilterValueContext(p *TargetFilterValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetFilterValue
}

func (*TargetFilterValueContext) IsTargetFilterValueContext() {}

func NewTargetFilterValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetFilterValueContext {
	var p = new(TargetFilterValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetFilterValue

	return p
}

func (s *TargetFilterValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetFilterValueContext) CopyAll(ctx *TargetFilterValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TargetFilterValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetFilterValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type QuotedFilterExpressionContext struct {
	TargetFilterValueContext
}

func NewQuotedFilterExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QuotedFilterExpressionContext {
	var p = new(QuotedFilterExpressionContext)

	InitEmptyTargetFilterValueContext(&p.TargetFilterValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetFilterValueContext))

	return p
}

func (s *QuotedFilterExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedFilterExpressionContext) AllDQUOTE() []antlr.TerminalNode {
	return s.GetTokens(ACIParserDQUOTE)
}

func (s *QuotedFilterExpressionContext) DQUOTE(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserDQUOTE, i)
}

func (s *QuotedFilterExpressionContext) LDAPFilter() ILDAPFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILDAPFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILDAPFilterContext)
}

func (s *QuotedFilterExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterQuotedFilterExpression(s)
	}
}

func (s *QuotedFilterExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitQuotedFilterExpression(s)
	}
}

func (p *ACIParser) TargetFilterValue() (localctx ITargetFilterValueContext) {
	localctx = NewTargetFilterValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ACIParserRULE_targetFilterValue)
	localctx = NewQuotedFilterExpressionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(ACIParserDQUOTE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.LDAPFilter()
	}
	{
		p.SetState(205)
		p.Match(ACIParserDQUOTE)
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

// ITargetContext is an interface to support dynamic dispatch.
type ITargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTargetContext differentiates from other interfaces.
	IsTargetContext()
}

type TargetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetContext() *TargetContext {
	var p = new(TargetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_target
	return p
}

func InitEmptyTargetContext(p *TargetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_target
}

func (*TargetContext) IsTargetContext() {}

func NewTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetContext {
	var p = new(TargetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_target

	return p
}

func (s *TargetContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetContext) CopyAll(ctx *TargetContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TargetDistinguishedNamesContext struct {
	TargetContext
}

func NewTargetDistinguishedNamesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetDistinguishedNamesContext {
	var p = new(TargetDistinguishedNamesContext)

	InitEmptyTargetContext(&p.TargetContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetContext))

	return p
}

func (s *TargetDistinguishedNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetDistinguishedNamesContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *TargetDistinguishedNamesContext) TargetKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserTargetKeyword, 0)
}

func (s *TargetDistinguishedNamesContext) DistinguishedNames() IDistinguishedNamesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDistinguishedNamesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDistinguishedNamesContext)
}

func (s *TargetDistinguishedNamesContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *TargetDistinguishedNamesContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *TargetDistinguishedNamesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetDistinguishedNames(s)
	}
}

func (s *TargetDistinguishedNamesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetDistinguishedNames(s)
	}
}

func (p *ACIParser) Target() (localctx ITargetContext) {
	localctx = NewTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ACIParserRULE_target)
	var _la int

	localctx = NewTargetDistinguishedNamesContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(207)
		p.Match(ACIParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(208)
		p.Match(ACIParserTargetKeyword)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(210)
		p.DistinguishedNames()
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

// ITargetToContext is an interface to support dynamic dispatch.
type ITargetToContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTargetToContext differentiates from other interfaces.
	IsTargetToContext()
}

type TargetToContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetToContext() *TargetToContext {
	var p = new(TargetToContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetTo
	return p
}

func InitEmptyTargetToContext(p *TargetToContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetTo
}

func (*TargetToContext) IsTargetToContext() {}

func NewTargetToContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetToContext {
	var p = new(TargetToContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetTo

	return p
}

func (s *TargetToContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetToContext) CopyAll(ctx *TargetToContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TargetToContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetToContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TargetToDistinguishedNameContext struct {
	TargetToContext
}

func NewTargetToDistinguishedNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetToDistinguishedNameContext {
	var p = new(TargetToDistinguishedNameContext)

	InitEmptyTargetToContext(&p.TargetToContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetToContext))

	return p
}

func (s *TargetToDistinguishedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetToDistinguishedNameContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *TargetToDistinguishedNameContext) TargetToKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserTargetToKeyword, 0)
}

func (s *TargetToDistinguishedNameContext) DistinguishedName() IDistinguishedNameContext {
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

func (s *TargetToDistinguishedNameContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *TargetToDistinguishedNameContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *TargetToDistinguishedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetToDistinguishedName(s)
	}
}

func (s *TargetToDistinguishedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetToDistinguishedName(s)
	}
}

func (p *ACIParser) TargetTo() (localctx ITargetToContext) {
	localctx = NewTargetToContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ACIParserRULE_targetTo)
	var _la int

	localctx = NewTargetToDistinguishedNameContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(ACIParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.Match(ACIParserTargetToKeyword)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(214)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(215)
		p.DistinguishedName()
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

// ITargetFromContext is an interface to support dynamic dispatch.
type ITargetFromContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTargetFromContext differentiates from other interfaces.
	IsTargetFromContext()
}

type TargetFromContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetFromContext() *TargetFromContext {
	var p = new(TargetFromContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetFrom
	return p
}

func InitEmptyTargetFromContext(p *TargetFromContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetFrom
}

func (*TargetFromContext) IsTargetFromContext() {}

func NewTargetFromContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetFromContext {
	var p = new(TargetFromContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetFrom

	return p
}

func (s *TargetFromContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetFromContext) CopyAll(ctx *TargetFromContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TargetFromContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetFromContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TargetFromDistinguishedNameContext struct {
	TargetFromContext
}

func NewTargetFromDistinguishedNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetFromDistinguishedNameContext {
	var p = new(TargetFromDistinguishedNameContext)

	InitEmptyTargetFromContext(&p.TargetFromContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetFromContext))

	return p
}

func (s *TargetFromDistinguishedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetFromDistinguishedNameContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *TargetFromDistinguishedNameContext) TargetFromKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserTargetFromKeyword, 0)
}

func (s *TargetFromDistinguishedNameContext) DistinguishedName() IDistinguishedNameContext {
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

func (s *TargetFromDistinguishedNameContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *TargetFromDistinguishedNameContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *TargetFromDistinguishedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetFromDistinguishedName(s)
	}
}

func (s *TargetFromDistinguishedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetFromDistinguishedName(s)
	}
}

func (p *ACIParser) TargetFrom() (localctx ITargetFromContext) {
	localctx = NewTargetFromContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ACIParserRULE_targetFrom)
	var _la int

	localctx = NewTargetFromDistinguishedNameContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(217)
		p.Match(ACIParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(218)
		p.Match(ACIParserTargetFromKeyword)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(220)
		p.DistinguishedName()
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

// ITargetAttrFiltersContext is an interface to support dynamic dispatch.
type ITargetAttrFiltersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTargetAttrFiltersContext differentiates from other interfaces.
	IsTargetAttrFiltersContext()
}

type TargetAttrFiltersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetAttrFiltersContext() *TargetAttrFiltersContext {
	var p = new(TargetAttrFiltersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetAttrFilters
	return p
}

func InitEmptyTargetAttrFiltersContext(p *TargetAttrFiltersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetAttrFilters
}

func (*TargetAttrFiltersContext) IsTargetAttrFiltersContext() {}

func NewTargetAttrFiltersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetAttrFiltersContext {
	var p = new(TargetAttrFiltersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetAttrFilters

	return p
}

func (s *TargetAttrFiltersContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetAttrFiltersContext) CopyAll(ctx *TargetAttrFiltersContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TargetAttrFiltersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetAttrFiltersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentheticalTargetAttrFiltersContext struct {
	TargetAttrFiltersContext
}

func NewParentheticalTargetAttrFiltersContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalTargetAttrFiltersContext {
	var p = new(ParentheticalTargetAttrFiltersContext)

	InitEmptyTargetAttrFiltersContext(&p.TargetAttrFiltersContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetAttrFiltersContext))

	return p
}

func (s *ParentheticalTargetAttrFiltersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalTargetAttrFiltersContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalTargetAttrFiltersContext) TargetAttrFiltersKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserTargetAttrFiltersKeyword, 0)
}

func (s *ParentheticalTargetAttrFiltersContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *ParentheticalTargetAttrFiltersContext) TargetAttrFiltersValue() ITargetAttrFiltersValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITargetAttrFiltersValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITargetAttrFiltersValueContext)
}

func (s *ParentheticalTargetAttrFiltersContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalTargetAttrFiltersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalTargetAttrFilters(s)
	}
}

func (s *ParentheticalTargetAttrFiltersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalTargetAttrFilters(s)
	}
}

func (p *ACIParser) TargetAttrFilters() (localctx ITargetAttrFiltersContext) {
	localctx = NewTargetAttrFiltersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, ACIParserRULE_targetAttrFilters)
	localctx = NewParentheticalTargetAttrFiltersContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		p.Match(ACIParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(223)
		p.Match(ACIParserTargetAttrFiltersKeyword)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(224)
		p.Match(ACIParserEqualTo)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(225)
		p.TargetAttrFiltersValue()
	}
	{
		p.SetState(226)
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

// ITargetAttrFiltersValueContext is an interface to support dynamic dispatch.
type ITargetAttrFiltersValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTargetAttrFiltersValueContext differentiates from other interfaces.
	IsTargetAttrFiltersValueContext()
}

type TargetAttrFiltersValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetAttrFiltersValueContext() *TargetAttrFiltersValueContext {
	var p = new(TargetAttrFiltersValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetAttrFiltersValue
	return p
}

func InitEmptyTargetAttrFiltersValueContext(p *TargetAttrFiltersValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetAttrFiltersValue
}

func (*TargetAttrFiltersValueContext) IsTargetAttrFiltersValueContext() {}

func NewTargetAttrFiltersValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetAttrFiltersValueContext {
	var p = new(TargetAttrFiltersValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetAttrFiltersValue

	return p
}

func (s *TargetAttrFiltersValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetAttrFiltersValueContext) CopyAll(ctx *TargetAttrFiltersValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TargetAttrFiltersValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetAttrFiltersValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type QuotedAttributeFiltersContext struct {
	TargetAttrFiltersValueContext
}

func NewQuotedAttributeFiltersContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QuotedAttributeFiltersContext {
	var p = new(QuotedAttributeFiltersContext)

	InitEmptyTargetAttrFiltersValueContext(&p.TargetAttrFiltersValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetAttrFiltersValueContext))

	return p
}

func (s *QuotedAttributeFiltersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedAttributeFiltersContext) AllDQUOTE() []antlr.TerminalNode {
	return s.GetTokens(ACIParserDQUOTE)
}

func (s *QuotedAttributeFiltersContext) DQUOTE(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserDQUOTE, i)
}

func (s *QuotedAttributeFiltersContext) AttributeFilters() IAttributeFiltersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeFiltersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeFiltersContext)
}

func (s *QuotedAttributeFiltersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterQuotedAttributeFilters(s)
	}
}

func (s *QuotedAttributeFiltersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitQuotedAttributeFilters(s)
	}
}

type QuotedAttributeFilterSetContext struct {
	TargetAttrFiltersValueContext
}

func NewQuotedAttributeFilterSetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QuotedAttributeFilterSetContext {
	var p = new(QuotedAttributeFilterSetContext)

	InitEmptyTargetAttrFiltersValueContext(&p.TargetAttrFiltersValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetAttrFiltersValueContext))

	return p
}

func (s *QuotedAttributeFilterSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedAttributeFilterSetContext) AllDQUOTE() []antlr.TerminalNode {
	return s.GetTokens(ACIParserDQUOTE)
}

func (s *QuotedAttributeFilterSetContext) DQUOTE(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserDQUOTE, i)
}

func (s *QuotedAttributeFilterSetContext) AttributeFilterSet() IAttributeFilterSetContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeFilterSetContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeFilterSetContext)
}

func (s *QuotedAttributeFilterSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterQuotedAttributeFilterSet(s)
	}
}

func (s *QuotedAttributeFilterSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitQuotedAttributeFilterSet(s)
	}
}

type QuotedAttributeFilterContext struct {
	TargetAttrFiltersValueContext
}

func NewQuotedAttributeFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *QuotedAttributeFilterContext {
	var p = new(QuotedAttributeFilterContext)

	InitEmptyTargetAttrFiltersValueContext(&p.TargetAttrFiltersValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetAttrFiltersValueContext))

	return p
}

func (s *QuotedAttributeFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QuotedAttributeFilterContext) AllDQUOTE() []antlr.TerminalNode {
	return s.GetTokens(ACIParserDQUOTE)
}

func (s *QuotedAttributeFilterContext) DQUOTE(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserDQUOTE, i)
}

func (s *QuotedAttributeFilterContext) AttributeFilter() IAttributeFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeFilterContext)
}

func (s *QuotedAttributeFilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterQuotedAttributeFilter(s)
	}
}

func (s *QuotedAttributeFilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitQuotedAttributeFilter(s)
	}
}

func (p *ACIParser) TargetAttrFiltersValue() (localctx ITargetAttrFiltersValueContext) {
	localctx = NewTargetAttrFiltersValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ACIParserRULE_targetAttrFiltersValue)
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewQuotedAttributeFiltersContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(228)
			p.Match(ACIParserDQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(229)
			p.AttributeFilters()
		}
		{
			p.SetState(230)
			p.Match(ACIParserDQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewQuotedAttributeFilterSetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(232)
			p.Match(ACIParserDQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(233)
			p.AttributeFilterSet()
		}
		{
			p.SetState(234)
			p.Match(ACIParserDQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewQuotedAttributeFilterContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(236)
			p.Match(ACIParserDQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(237)
			p.AttributeFilter()
		}
		{
			p.SetState(238)
			p.Match(ACIParserDQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
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

// ITargetScopeContext is an interface to support dynamic dispatch.
type ITargetScopeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTargetScopeContext differentiates from other interfaces.
	IsTargetScopeContext()
}

type TargetScopeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetScopeContext() *TargetScopeContext {
	var p = new(TargetScopeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetScope
	return p
}

func InitEmptyTargetScopeContext(p *TargetScopeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetScope
}

func (*TargetScopeContext) IsTargetScopeContext() {}

func NewTargetScopeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetScopeContext {
	var p = new(TargetScopeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetScope

	return p
}

func (s *TargetScopeContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetScopeContext) CopyAll(ctx *TargetScopeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TargetScopeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetScopeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TargetScopeBindRuleContext struct {
	TargetScopeContext
}

func NewTargetScopeBindRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetScopeBindRuleContext {
	var p = new(TargetScopeBindRuleContext)

	InitEmptyTargetScopeContext(&p.TargetScopeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetScopeContext))

	return p
}

func (s *TargetScopeBindRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetScopeBindRuleContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *TargetScopeBindRuleContext) TargetScopeKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserTargetScopeKeyword, 0)
}

func (s *TargetScopeBindRuleContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *TargetScopeBindRuleContext) TargetRuleSearchScopes() antlr.TerminalNode {
	return s.GetToken(ACIParserTargetRuleSearchScopes, 0)
}

func (s *TargetScopeBindRuleContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *TargetScopeBindRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetScopeBindRule(s)
	}
}

func (s *TargetScopeBindRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetScopeBindRule(s)
	}
}

func (p *ACIParser) TargetScope() (localctx ITargetScopeContext) {
	localctx = NewTargetScopeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, ACIParserRULE_targetScope)
	localctx = NewTargetScopeBindRuleContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(242)
		p.Match(ACIParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(243)
		p.Match(ACIParserTargetScopeKeyword)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(244)
		p.Match(ACIParserEqualTo)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(245)
		p.Match(ACIParserTargetRuleSearchScopes)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(246)
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

// ITargetAttrContext is an interface to support dynamic dispatch.
type ITargetAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTargetAttrContext differentiates from other interfaces.
	IsTargetAttrContext()
}

type TargetAttrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetAttrContext() *TargetAttrContext {
	var p = new(TargetAttrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetAttr
	return p
}

func InitEmptyTargetAttrContext(p *TargetAttrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_targetAttr
}

func (*TargetAttrContext) IsTargetAttrContext() {}

func NewTargetAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetAttrContext {
	var p = new(TargetAttrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_targetAttr

	return p
}

func (s *TargetAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetAttrContext) CopyAll(ctx *TargetAttrContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TargetAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TargetAttrBindRuleContext struct {
	TargetAttrContext
}

func NewTargetAttrBindRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TargetAttrBindRuleContext {
	var p = new(TargetAttrBindRuleContext)

	InitEmptyTargetAttrContext(&p.TargetAttrContext)
	p.parser = parser
	p.CopyAll(ctx.(*TargetAttrContext))

	return p
}

func (s *TargetAttrBindRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetAttrBindRuleContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *TargetAttrBindRuleContext) TargetAttrKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserTargetAttrKeyword, 0)
}

func (s *TargetAttrBindRuleContext) AttributeType() IAttributeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeTypeContext)
}

func (s *TargetAttrBindRuleContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *TargetAttrBindRuleContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *TargetAttrBindRuleContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *TargetAttrBindRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTargetAttrBindRule(s)
	}
}

func (s *TargetAttrBindRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTargetAttrBindRule(s)
	}
}

func (p *ACIParser) TargetAttr() (localctx ITargetAttrContext) {
	localctx = NewTargetAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, ACIParserRULE_targetAttr)
	var _la int

	localctx = NewTargetAttrBindRuleContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(248)
		p.Match(ACIParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(249)
		p.Match(ACIParserTargetAttrKeyword)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(250)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(251)
		p.AttributeType()
	}
	{
		p.SetState(252)
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

// IAttributeTypesContext is an interface to support dynamic dispatch.
type IAttributeTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAttributeTypesContext differentiates from other interfaces.
	IsAttributeTypesContext()
}

type AttributeTypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeTypesContext() *AttributeTypesContext {
	var p = new(AttributeTypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeTypes
	return p
}

func InitEmptyAttributeTypesContext(p *AttributeTypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeTypes
}

func (*AttributeTypesContext) IsAttributeTypesContext() {}

func NewAttributeTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeTypesContext {
	var p = new(AttributeTypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_attributeTypes

	return p
}

func (s *AttributeTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeTypesContext) CopyAll(ctx *AttributeTypesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AttributeTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AttributeTypesListContext struct {
	AttributeTypesContext
}

func NewAttributeTypesListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttributeTypesListContext {
	var p = new(AttributeTypesListContext)

	InitEmptyAttributeTypesContext(&p.AttributeTypesContext)
	p.parser = parser
	p.CopyAll(ctx.(*AttributeTypesContext))

	return p
}

func (s *AttributeTypesListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeTypesListContext) AllAttributeType() []IAttributeTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeTypeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeTypeContext); ok {
			tst[i] = t.(IAttributeTypeContext)
			i++
		}
	}

	return tst
}

func (s *AttributeTypesListContext) AttributeType(i int) IAttributeTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeContext); ok {
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

	return t.(IAttributeTypeContext)
}

func (s *AttributeTypesListContext) AllDoublePipe() []IDoublePipeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoublePipeContext); ok {
			len++
		}
	}

	tst := make([]IDoublePipeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoublePipeContext); ok {
			tst[i] = t.(IDoublePipeContext)
			i++
		}
	}

	return tst
}

func (s *AttributeTypesListContext) DoublePipe(i int) IDoublePipeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoublePipeContext); ok {
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

	return t.(IDoublePipeContext)
}

func (s *AttributeTypesListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAttributeTypesList(s)
	}
}

func (s *AttributeTypesListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAttributeTypesList(s)
	}
}

func (p *ACIParser) AttributeTypes() (localctx IAttributeTypesContext) {
	localctx = NewAttributeTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ACIParserRULE_attributeTypes)
	var _la int

	localctx = NewAttributeTypesListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.AttributeType()
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ACIParserDPIPE {
		{
			p.SetState(255)
			p.DoublePipe()
		}
		{
			p.SetState(256)
			p.AttributeType()
		}

		p.SetState(262)
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

// IBindRuleContext is an interface to support dynamic dispatch.
type IBindRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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

func (s *BindRuleContext) CopyAll(ctx *BindRuleContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BindRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentheticalBindRuleInstanceWithRequiredBooleanOperatorContext struct {
	BindRuleContext
}

func NewParentheticalBindRuleInstanceWithRequiredBooleanOperatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalBindRuleInstanceWithRequiredBooleanOperatorContext {
	var p = new(ParentheticalBindRuleInstanceWithRequiredBooleanOperatorContext)

	InitEmptyBindRuleContext(&p.BindRuleContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleContext))

	return p
}

func (s *ParentheticalBindRuleInstanceWithRequiredBooleanOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalBindRuleInstanceWithRequiredBooleanOperatorContext) AllBindRuleExprParen() []IBindRuleExprParenContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBindRuleExprParenContext); ok {
			len++
		}
	}

	tst := make([]IBindRuleExprParenContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBindRuleExprParenContext); ok {
			tst[i] = t.(IBindRuleExprParenContext)
			i++
		}
	}

	return tst
}

func (s *ParentheticalBindRuleInstanceWithRequiredBooleanOperatorContext) BindRuleExprParen(i int) IBindRuleExprParenContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleExprParenContext); ok {
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

	return t.(IBindRuleExprParenContext)
}

func (s *ParentheticalBindRuleInstanceWithRequiredBooleanOperatorContext) AllBooleanWordOperators() []antlr.TerminalNode {
	return s.GetTokens(ACIParserBooleanWordOperators)
}

func (s *ParentheticalBindRuleInstanceWithRequiredBooleanOperatorContext) BooleanWordOperators(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserBooleanWordOperators, i)
}

func (s *ParentheticalBindRuleInstanceWithRequiredBooleanOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalBindRuleInstanceWithRequiredBooleanOperator(s)
	}
}

func (s *ParentheticalBindRuleInstanceWithRequiredBooleanOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalBindRuleInstanceWithRequiredBooleanOperator(s)
	}
}

type BindRuleInstanceContext struct {
	BindRuleContext
}

func NewBindRuleInstanceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BindRuleInstanceContext {
	var p = new(BindRuleInstanceContext)

	InitEmptyBindRuleContext(&p.BindRuleContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleContext))

	return p
}

func (s *BindRuleInstanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleInstanceContext) BindRuleExpr() IBindRuleExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleExprContext)
}

func (s *BindRuleInstanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindRuleInstance(s)
	}
}

func (s *BindRuleInstanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindRuleInstance(s)
	}
}

func (p *ACIParser) BindRule() (localctx IBindRuleContext) {
	localctx = NewBindRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ACIParserRULE_bindRule)
	var _la int

	p.SetState(272)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBindRuleInstanceContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(263)
			p.BindRuleExpr()
		}

	case 2:
		localctx = NewParentheticalBindRuleInstanceWithRequiredBooleanOperatorContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(264)
			p.BindRuleExprParen()
		}
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ACIParserBooleanWordOperators {
			{
				p.SetState(265)
				p.Match(ACIParserBooleanWordOperators)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(266)
				p.BindRuleExprParen()
			}

			p.SetState(271)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// IBindRuleExprParenContext is an interface to support dynamic dispatch.
type IBindRuleExprParenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBindRuleExprParenContext differentiates from other interfaces.
	IsBindRuleExprParenContext()
}

type BindRuleExprParenContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleExprParenContext() *BindRuleExprParenContext {
	var p = new(BindRuleExprParenContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleExprParen
	return p
}

func InitEmptyBindRuleExprParenContext(p *BindRuleExprParenContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleExprParen
}

func (*BindRuleExprParenContext) IsBindRuleExprParenContext() {}

func NewBindRuleExprParenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleExprParenContext {
	var p = new(BindRuleExprParenContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRuleExprParen

	return p
}

func (s *BindRuleExprParenContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleExprParenContext) CopyAll(ctx *BindRuleExprParenContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BindRuleExprParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleExprParenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BindRuleExpressionRecursionContext struct {
	BindRuleExprParenContext
}

func NewBindRuleExpressionRecursionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BindRuleExpressionRecursionContext {
	var p = new(BindRuleExpressionRecursionContext)

	InitEmptyBindRuleExprParenContext(&p.BindRuleExprParenContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprParenContext))

	return p
}

func (s *BindRuleExpressionRecursionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleExpressionRecursionContext) BindRuleExpr() IBindRuleExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleExprContext)
}

func (s *BindRuleExpressionRecursionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindRuleExpressionRecursion(s)
	}
}

func (s *BindRuleExpressionRecursionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindRuleExpressionRecursion(s)
	}
}

type ParentheticalBindRuleExpressionRecursionContext struct {
	BindRuleExprParenContext
}

func NewParentheticalBindRuleExpressionRecursionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalBindRuleExpressionRecursionContext {
	var p = new(ParentheticalBindRuleExpressionRecursionContext)

	InitEmptyBindRuleExprParenContext(&p.BindRuleExprParenContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprParenContext))

	return p
}

func (s *ParentheticalBindRuleExpressionRecursionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalBindRuleExpressionRecursionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalBindRuleExpressionRecursionContext) BindRuleExpr() IBindRuleExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleExprContext)
}

func (s *ParentheticalBindRuleExpressionRecursionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalBindRuleExpressionRecursionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalBindRuleExpressionRecursion(s)
	}
}

func (s *ParentheticalBindRuleExpressionRecursionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalBindRuleExpressionRecursion(s)
	}
}

type ParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext struct {
	BindRuleExprParenContext
}

func NewParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext {
	var p = new(ParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext)

	InitEmptyBindRuleExprParenContext(&p.BindRuleExprParenContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprParenContext))

	return p
}

func (s *ParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext) AllBindRuleExpr() []IBindRuleExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBindRuleExprContext); ok {
			len++
		}
	}

	tst := make([]IBindRuleExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBindRuleExprContext); ok {
			tst[i] = t.(IBindRuleExprContext)
			i++
		}
	}

	return tst
}

func (s *ParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext) BindRuleExpr(i int) IBindRuleExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleExprContext); ok {
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

	return t.(IBindRuleExprContext)
}

func (s *ParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext) AllBooleanWordOperators() []antlr.TerminalNode {
	return s.GetTokens(ACIParserBooleanWordOperators)
}

func (s *ParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext) BooleanWordOperators(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserBooleanWordOperators, i)
}

func (s *ParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursion(s)
	}
}

func (s *ParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursion(s)
	}
}

type NegatedBindRuleExpressionRecursionContext struct {
	BindRuleExprParenContext
}

func NewNegatedBindRuleExpressionRecursionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegatedBindRuleExpressionRecursionContext {
	var p = new(NegatedBindRuleExpressionRecursionContext)

	InitEmptyBindRuleExprParenContext(&p.BindRuleExprParenContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprParenContext))

	return p
}

func (s *NegatedBindRuleExpressionRecursionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegatedBindRuleExpressionRecursionContext) BooleanNotWord() antlr.TerminalNode {
	return s.GetToken(ACIParserBooleanNotWord, 0)
}

func (s *NegatedBindRuleExpressionRecursionContext) BindRuleExpr() IBindRuleExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleExprContext)
}

func (s *NegatedBindRuleExpressionRecursionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterNegatedBindRuleExpressionRecursion(s)
	}
}

func (s *NegatedBindRuleExpressionRecursionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitNegatedBindRuleExpressionRecursion(s)
	}
}

func (p *ACIParser) BindRuleExprParen() (localctx IBindRuleExprParenContext) {
	localctx = NewBindRuleExprParenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ACIParserRULE_bindRuleExprParen)
	var _la int

	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(275)
			p.BindRuleExpr()
		}
		p.SetState(280)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ACIParserBooleanWordOperators {
			{
				p.SetState(276)
				p.Match(ACIParserBooleanWordOperators)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(277)
				p.BindRuleExpr()
			}

			p.SetState(282)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(283)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewNegatedBindRuleExpressionRecursionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(285)
			p.Match(ACIParserBooleanNotWord)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(286)
			p.BindRuleExpr()
		}

	case 3:
		localctx = NewParentheticalBindRuleExpressionRecursionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(287)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(288)
			p.BindRuleExpr()
		}
		{
			p.SetState(289)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		localctx = NewBindRuleExpressionRecursionContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(291)
			p.BindRuleExpr()
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

// IBindRuleExprContext is an interface to support dynamic dispatch.
type IBindRuleExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBindRuleExprContext differentiates from other interfaces.
	IsBindRuleExprContext()
}

type BindRuleExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleExprContext() *BindRuleExprContext {
	var p = new(BindRuleExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleExpr
	return p
}

func InitEmptyBindRuleExprContext(p *BindRuleExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleExpr
}

func (*BindRuleExprContext) IsBindRuleExprContext() {}

func NewBindRuleExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleExprContext {
	var p = new(BindRuleExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRuleExpr

	return p
}

func (s *BindRuleExprContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleExprContext) CopyAll(ctx *BindRuleExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BindRuleExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GroupattrContext struct {
	BindRuleExprContext
}

func NewGroupattrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupattrContext {
	var p = new(GroupattrContext)

	InitEmptyBindRuleExprContext(&p.BindRuleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprContext))

	return p
}

func (s *GroupattrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupattrContext) BindRuleGroupAttr() IBindRuleGroupAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleGroupAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleGroupAttrContext)
}

func (s *GroupattrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterGroupattr(s)
	}
}

func (s *GroupattrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitGroupattr(s)
	}
}

type TimeofdayContext struct {
	BindRuleExprContext
}

func NewTimeofdayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimeofdayContext {
	var p = new(TimeofdayContext)

	InitEmptyBindRuleExprContext(&p.BindRuleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprContext))

	return p
}

func (s *TimeofdayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeofdayContext) BindRuleTimeOfDay() IBindRuleTimeOfDayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleTimeOfDayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleTimeOfDayContext)
}

func (s *TimeofdayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTimeofday(s)
	}
}

func (s *TimeofdayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTimeofday(s)
	}
}

type GroupdnContext struct {
	BindRuleExprContext
}

func NewGroupdnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupdnContext {
	var p = new(GroupdnContext)

	InitEmptyBindRuleExprContext(&p.BindRuleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprContext))

	return p
}

func (s *GroupdnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupdnContext) BindRuleGroupDN() IBindRuleGroupDNContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleGroupDNContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleGroupDNContext)
}

func (s *GroupdnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterGroupdn(s)
	}
}

func (s *GroupdnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitGroupdn(s)
	}
}

type RolednContext struct {
	BindRuleExprContext
}

func NewRolednContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RolednContext {
	var p = new(RolednContext)

	InitEmptyBindRuleExprContext(&p.BindRuleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprContext))

	return p
}

func (s *RolednContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RolednContext) BindRuleRoleDN() IBindRuleRoleDNContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleRoleDNContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleRoleDNContext)
}

func (s *RolednContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterRoledn(s)
	}
}

func (s *RolednContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitRoledn(s)
	}
}

type IpContext struct {
	BindRuleExprContext
}

func NewIpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IpContext {
	var p = new(IpContext)

	InitEmptyBindRuleExprContext(&p.BindRuleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprContext))

	return p
}

func (s *IpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpContext) BindRuleIP() IBindRuleIPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleIPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleIPContext)
}

func (s *IpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterIp(s)
	}
}

func (s *IpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitIp(s)
	}
}

type SsfContext struct {
	BindRuleExprContext
}

func NewSsfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SsfContext {
	var p = new(SsfContext)

	InitEmptyBindRuleExprContext(&p.BindRuleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprContext))

	return p
}

func (s *SsfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SsfContext) BindRuleSecurityStrengthFactor() IBindRuleSecurityStrengthFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleSecurityStrengthFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleSecurityStrengthFactorContext)
}

func (s *SsfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterSsf(s)
	}
}

func (s *SsfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitSsf(s)
	}
}

type DnsContext struct {
	BindRuleExprContext
}

func NewDnsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DnsContext {
	var p = new(DnsContext)

	InitEmptyBindRuleExprContext(&p.BindRuleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprContext))

	return p
}

func (s *DnsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DnsContext) BindRuleDNS() IBindRuleDNSContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleDNSContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleDNSContext)
}

func (s *DnsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterDns(s)
	}
}

func (s *DnsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitDns(s)
	}
}

type DayofweekContext struct {
	BindRuleExprContext
}

func NewDayofweekContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DayofweekContext {
	var p = new(DayofweekContext)

	InitEmptyBindRuleExprContext(&p.BindRuleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprContext))

	return p
}

func (s *DayofweekContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayofweekContext) BindRuleDayOfWeek() IBindRuleDayOfWeekContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleDayOfWeekContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleDayOfWeekContext)
}

func (s *DayofweekContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterDayofweek(s)
	}
}

func (s *DayofweekContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitDayofweek(s)
	}
}

type AuthmethodContext struct {
	BindRuleExprContext
}

func NewAuthmethodContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AuthmethodContext {
	var p = new(AuthmethodContext)

	InitEmptyBindRuleExprContext(&p.BindRuleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprContext))

	return p
}

func (s *AuthmethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuthmethodContext) BindRuleAuthMethod() IBindRuleAuthMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleAuthMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleAuthMethodContext)
}

func (s *AuthmethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAuthmethod(s)
	}
}

func (s *AuthmethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAuthmethod(s)
	}
}

type BindRuleExpressionContext struct {
	BindRuleExprContext
}

func NewBindRuleExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BindRuleExpressionContext {
	var p = new(BindRuleExpressionContext)

	InitEmptyBindRuleExprContext(&p.BindRuleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprContext))

	return p
}

func (s *BindRuleExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *BindRuleExpressionContext) BindRuleExpr() IBindRuleExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleExprContext)
}

func (s *BindRuleExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *BindRuleExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterBindRuleExpression(s)
	}
}

func (s *BindRuleExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitBindRuleExpression(s)
	}
}

type UserdnContext struct {
	BindRuleExprContext
}

func NewUserdnContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UserdnContext {
	var p = new(UserdnContext)

	InitEmptyBindRuleExprContext(&p.BindRuleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprContext))

	return p
}

func (s *UserdnContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserdnContext) BindRuleUserDN() IBindRuleUserDNContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleUserDNContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleUserDNContext)
}

func (s *UserdnContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterUserdn(s)
	}
}

func (s *UserdnContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitUserdn(s)
	}
}

type UserattrContext struct {
	BindRuleExprContext
}

func NewUserattrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UserattrContext {
	var p = new(UserattrContext)

	InitEmptyBindRuleExprContext(&p.BindRuleExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleExprContext))

	return p
}

func (s *UserattrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserattrContext) BindRuleUserAttr() IBindRuleUserAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleUserAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleUserAttrContext)
}

func (s *UserattrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterUserattr(s)
	}
}

func (s *UserattrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitUserattr(s)
	}
}

func (p *ACIParser) BindRuleExpr() (localctx IBindRuleExprContext) {
	localctx = NewBindRuleExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, ACIParserRULE_bindRuleExpr)
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBindRuleExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(294)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)
			p.BindRuleExpr()
		}
		{
			p.SetState(296)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewUserdnContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(298)
			p.BindRuleUserDN()
		}

	case 3:
		localctx = NewUserattrContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(299)
			p.BindRuleUserAttr()
		}

	case 4:
		localctx = NewGroupdnContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(300)
			p.BindRuleGroupDN()
		}

	case 5:
		localctx = NewGroupattrContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(301)
			p.BindRuleGroupAttr()
		}

	case 6:
		localctx = NewRolednContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(302)
			p.BindRuleRoleDN()
		}

	case 7:
		localctx = NewDnsContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(303)
			p.BindRuleDNS()
		}

	case 8:
		localctx = NewIpContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(304)
			p.BindRuleIP()
		}

	case 9:
		localctx = NewTimeofdayContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(305)
			p.BindRuleTimeOfDay()
		}

	case 10:
		localctx = NewDayofweekContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(306)
			p.BindRuleDayOfWeek()
		}

	case 11:
		localctx = NewSsfContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(307)
			p.BindRuleSecurityStrengthFactor()
		}

	case 12:
		localctx = NewAuthmethodContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(308)
			p.BindRuleAuthMethod()
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

// IBindRuleUserDNContext is an interface to support dynamic dispatch.
type IBindRuleUserDNContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBindRuleUserDNContext differentiates from other interfaces.
	IsBindRuleUserDNContext()
}

type BindRuleUserDNContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleUserDNContext() *BindRuleUserDNContext {
	var p = new(BindRuleUserDNContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleUserDN
	return p
}

func InitEmptyBindRuleUserDNContext(p *BindRuleUserDNContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleUserDN
}

func (*BindRuleUserDNContext) IsBindRuleUserDNContext() {}

func NewBindRuleUserDNContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleUserDNContext {
	var p = new(BindRuleUserDNContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRuleUserDN

	return p
}

func (s *BindRuleUserDNContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleUserDNContext) CopyAll(ctx *BindRuleUserDNContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BindRuleUserDNContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleUserDNContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentheticalUserDistinguishedNameContext struct {
	BindRuleUserDNContext
}

func NewParentheticalUserDistinguishedNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalUserDistinguishedNameContext {
	var p = new(ParentheticalUserDistinguishedNameContext)

	InitEmptyBindRuleUserDNContext(&p.BindRuleUserDNContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleUserDNContext))

	return p
}

func (s *ParentheticalUserDistinguishedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalUserDistinguishedNameContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalUserDistinguishedNameContext) BindRuleUserDN() IBindRuleUserDNContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleUserDNContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleUserDNContext)
}

func (s *ParentheticalUserDistinguishedNameContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalUserDistinguishedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalUserDistinguishedName(s)
	}
}

func (s *ParentheticalUserDistinguishedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalUserDistinguishedName(s)
	}
}

type UserDistinguishedNameExpressionContext struct {
	BindRuleUserDNContext
}

func NewUserDistinguishedNameExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UserDistinguishedNameExpressionContext {
	var p = new(UserDistinguishedNameExpressionContext)

	InitEmptyBindRuleUserDNContext(&p.BindRuleUserDNContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleUserDNContext))

	return p
}

func (s *UserDistinguishedNameExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserDistinguishedNameExpressionContext) UserDNKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserUserDNKeyword, 0)
}

func (s *UserDistinguishedNameExpressionContext) DistinguishedName() IDistinguishedNameContext {
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

func (s *UserDistinguishedNameExpressionContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *UserDistinguishedNameExpressionContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *UserDistinguishedNameExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterUserDistinguishedNameExpression(s)
	}
}

func (s *UserDistinguishedNameExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitUserDistinguishedNameExpression(s)
	}
}

func (p *ACIParser) BindRuleUserDN() (localctx IBindRuleUserDNContext) {
	localctx = NewBindRuleUserDNContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ACIParserRULE_bindRuleUserDN)
	var _la int

	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserLPAREN:
		localctx = NewParentheticalUserDistinguishedNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(311)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(312)
			p.BindRuleUserDN()
		}
		{
			p.SetState(313)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ACIParserUserDNKeyword:
		localctx = NewUserDistinguishedNameExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(315)
			p.Match(ACIParserUserDNKeyword)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(316)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(317)
			p.DistinguishedName()
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

// IBindRuleRoleDNContext is an interface to support dynamic dispatch.
type IBindRuleRoleDNContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBindRuleRoleDNContext differentiates from other interfaces.
	IsBindRuleRoleDNContext()
}

type BindRuleRoleDNContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleRoleDNContext() *BindRuleRoleDNContext {
	var p = new(BindRuleRoleDNContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleRoleDN
	return p
}

func InitEmptyBindRuleRoleDNContext(p *BindRuleRoleDNContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleRoleDN
}

func (*BindRuleRoleDNContext) IsBindRuleRoleDNContext() {}

func NewBindRuleRoleDNContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleRoleDNContext {
	var p = new(BindRuleRoleDNContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRuleRoleDN

	return p
}

func (s *BindRuleRoleDNContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleRoleDNContext) CopyAll(ctx *BindRuleRoleDNContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BindRuleRoleDNContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleRoleDNContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentheticalRoleDistinguishedNameContext struct {
	BindRuleRoleDNContext
}

func NewParentheticalRoleDistinguishedNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalRoleDistinguishedNameContext {
	var p = new(ParentheticalRoleDistinguishedNameContext)

	InitEmptyBindRuleRoleDNContext(&p.BindRuleRoleDNContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleRoleDNContext))

	return p
}

func (s *ParentheticalRoleDistinguishedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalRoleDistinguishedNameContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalRoleDistinguishedNameContext) BindRuleRoleDN() IBindRuleRoleDNContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleRoleDNContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleRoleDNContext)
}

func (s *ParentheticalRoleDistinguishedNameContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalRoleDistinguishedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalRoleDistinguishedName(s)
	}
}

func (s *ParentheticalRoleDistinguishedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalRoleDistinguishedName(s)
	}
}

type RoleDistinguishedNameExpressionContext struct {
	BindRuleRoleDNContext
}

func NewRoleDistinguishedNameExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoleDistinguishedNameExpressionContext {
	var p = new(RoleDistinguishedNameExpressionContext)

	InitEmptyBindRuleRoleDNContext(&p.BindRuleRoleDNContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleRoleDNContext))

	return p
}

func (s *RoleDistinguishedNameExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoleDistinguishedNameExpressionContext) RoleDNKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserRoleDNKeyword, 0)
}

func (s *RoleDistinguishedNameExpressionContext) DistinguishedName() IDistinguishedNameContext {
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

func (s *RoleDistinguishedNameExpressionContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *RoleDistinguishedNameExpressionContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *RoleDistinguishedNameExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterRoleDistinguishedNameExpression(s)
	}
}

func (s *RoleDistinguishedNameExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitRoleDistinguishedNameExpression(s)
	}
}

func (p *ACIParser) BindRuleRoleDN() (localctx IBindRuleRoleDNContext) {
	localctx = NewBindRuleRoleDNContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, ACIParserRULE_bindRuleRoleDN)
	var _la int

	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserLPAREN:
		localctx = NewParentheticalRoleDistinguishedNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(320)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(321)
			p.BindRuleRoleDN()
		}
		{
			p.SetState(322)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ACIParserRoleDNKeyword:
		localctx = NewRoleDistinguishedNameExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(324)
			p.Match(ACIParserRoleDNKeyword)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(325)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(326)
			p.DistinguishedName()
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

// IBindRuleGroupDNContext is an interface to support dynamic dispatch.
type IBindRuleGroupDNContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBindRuleGroupDNContext differentiates from other interfaces.
	IsBindRuleGroupDNContext()
}

type BindRuleGroupDNContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleGroupDNContext() *BindRuleGroupDNContext {
	var p = new(BindRuleGroupDNContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleGroupDN
	return p
}

func InitEmptyBindRuleGroupDNContext(p *BindRuleGroupDNContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleGroupDN
}

func (*BindRuleGroupDNContext) IsBindRuleGroupDNContext() {}

func NewBindRuleGroupDNContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleGroupDNContext {
	var p = new(BindRuleGroupDNContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRuleGroupDN

	return p
}

func (s *BindRuleGroupDNContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleGroupDNContext) CopyAll(ctx *BindRuleGroupDNContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BindRuleGroupDNContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleGroupDNContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type GroupDistinguishedNameExpressionContext struct {
	BindRuleGroupDNContext
}

func NewGroupDistinguishedNameExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupDistinguishedNameExpressionContext {
	var p = new(GroupDistinguishedNameExpressionContext)

	InitEmptyBindRuleGroupDNContext(&p.BindRuleGroupDNContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleGroupDNContext))

	return p
}

func (s *GroupDistinguishedNameExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupDistinguishedNameExpressionContext) GroupDNKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserGroupDNKeyword, 0)
}

func (s *GroupDistinguishedNameExpressionContext) DistinguishedName() IDistinguishedNameContext {
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

func (s *GroupDistinguishedNameExpressionContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *GroupDistinguishedNameExpressionContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *GroupDistinguishedNameExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterGroupDistinguishedNameExpression(s)
	}
}

func (s *GroupDistinguishedNameExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitGroupDistinguishedNameExpression(s)
	}
}

type ParentheticalGroupDistinguishedNameContext struct {
	BindRuleGroupDNContext
}

func NewParentheticalGroupDistinguishedNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalGroupDistinguishedNameContext {
	var p = new(ParentheticalGroupDistinguishedNameContext)

	InitEmptyBindRuleGroupDNContext(&p.BindRuleGroupDNContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleGroupDNContext))

	return p
}

func (s *ParentheticalGroupDistinguishedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalGroupDistinguishedNameContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalGroupDistinguishedNameContext) BindRuleGroupDN() IBindRuleGroupDNContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleGroupDNContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleGroupDNContext)
}

func (s *ParentheticalGroupDistinguishedNameContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalGroupDistinguishedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalGroupDistinguishedName(s)
	}
}

func (s *ParentheticalGroupDistinguishedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalGroupDistinguishedName(s)
	}
}

func (p *ACIParser) BindRuleGroupDN() (localctx IBindRuleGroupDNContext) {
	localctx = NewBindRuleGroupDNContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ACIParserRULE_bindRuleGroupDN)
	var _la int

	p.SetState(336)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserLPAREN:
		localctx = NewParentheticalGroupDistinguishedNameContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(329)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(330)
			p.BindRuleGroupDN()
		}
		{
			p.SetState(331)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ACIParserGroupDNKeyword:
		localctx = NewGroupDistinguishedNameExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(333)
			p.Match(ACIParserGroupDNKeyword)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(334)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(335)
			p.DistinguishedName()
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

// IBindRuleUserAttrContext is an interface to support dynamic dispatch.
type IBindRuleUserAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBindRuleUserAttrContext differentiates from other interfaces.
	IsBindRuleUserAttrContext()
}

type BindRuleUserAttrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleUserAttrContext() *BindRuleUserAttrContext {
	var p = new(BindRuleUserAttrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleUserAttr
	return p
}

func InitEmptyBindRuleUserAttrContext(p *BindRuleUserAttrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleUserAttr
}

func (*BindRuleUserAttrContext) IsBindRuleUserAttrContext() {}

func NewBindRuleUserAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleUserAttrContext {
	var p = new(BindRuleUserAttrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRuleUserAttr

	return p
}

func (s *BindRuleUserAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleUserAttrContext) CopyAll(ctx *BindRuleUserAttrContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BindRuleUserAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleUserAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentheticalUserAttributesContext struct {
	BindRuleUserAttrContext
}

func NewParentheticalUserAttributesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalUserAttributesContext {
	var p = new(ParentheticalUserAttributesContext)

	InitEmptyBindRuleUserAttrContext(&p.BindRuleUserAttrContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleUserAttrContext))

	return p
}

func (s *ParentheticalUserAttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalUserAttributesContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalUserAttributesContext) BindRuleUserAttr() IBindRuleUserAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleUserAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleUserAttrContext)
}

func (s *ParentheticalUserAttributesContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalUserAttributesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalUserAttributes(s)
	}
}

func (s *ParentheticalUserAttributesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalUserAttributes(s)
	}
}

type UserAttributesExpressionContext struct {
	BindRuleUserAttrContext
}

func NewUserAttributesExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UserAttributesExpressionContext {
	var p = new(UserAttributesExpressionContext)

	InitEmptyBindRuleUserAttrContext(&p.BindRuleUserAttrContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleUserAttrContext))

	return p
}

func (s *UserAttributesExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UserAttributesExpressionContext) UserAttrKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserUserAttrKeyword, 0)
}

func (s *UserAttributesExpressionContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *UserAttributesExpressionContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *UserAttributesExpressionContext) AttributeBindTypeOrValue() IAttributeBindTypeOrValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeBindTypeOrValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeBindTypeOrValueContext)
}

func (s *UserAttributesExpressionContext) Inheritance() IInheritanceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInheritanceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInheritanceContext)
}

func (s *UserAttributesExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterUserAttributesExpression(s)
	}
}

func (s *UserAttributesExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitUserAttributesExpression(s)
	}
}

func (p *ACIParser) BindRuleUserAttr() (localctx IBindRuleUserAttrContext) {
	localctx = NewBindRuleUserAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ACIParserRULE_bindRuleUserAttr)
	var _la int

	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserLPAREN:
		localctx = NewParentheticalUserAttributesContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(338)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.BindRuleUserAttr()
		}
		{
			p.SetState(340)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ACIParserUserAttrKeyword:
		localctx = NewUserAttributesExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(342)
			p.Match(ACIParserUserAttrKeyword)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(343)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ACIParserID:
			{
				p.SetState(344)
				p.AttributeBindTypeOrValue()
			}

		case ACIParserINHERITANCEPREFIX:
			{
				p.SetState(345)
				p.Inheritance()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
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

// IBindRuleGroupAttrContext is an interface to support dynamic dispatch.
type IBindRuleGroupAttrContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBindRuleGroupAttrContext differentiates from other interfaces.
	IsBindRuleGroupAttrContext()
}

type BindRuleGroupAttrContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleGroupAttrContext() *BindRuleGroupAttrContext {
	var p = new(BindRuleGroupAttrContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleGroupAttr
	return p
}

func InitEmptyBindRuleGroupAttrContext(p *BindRuleGroupAttrContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleGroupAttr
}

func (*BindRuleGroupAttrContext) IsBindRuleGroupAttrContext() {}

func NewBindRuleGroupAttrContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleGroupAttrContext {
	var p = new(BindRuleGroupAttrContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRuleGroupAttr

	return p
}

func (s *BindRuleGroupAttrContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleGroupAttrContext) CopyAll(ctx *BindRuleGroupAttrContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BindRuleGroupAttrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleGroupAttrContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentheticalGroupAttributesContext struct {
	BindRuleGroupAttrContext
}

func NewParentheticalGroupAttributesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalGroupAttributesContext {
	var p = new(ParentheticalGroupAttributesContext)

	InitEmptyBindRuleGroupAttrContext(&p.BindRuleGroupAttrContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleGroupAttrContext))

	return p
}

func (s *ParentheticalGroupAttributesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalGroupAttributesContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalGroupAttributesContext) BindRuleGroupAttr() IBindRuleGroupAttrContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleGroupAttrContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleGroupAttrContext)
}

func (s *ParentheticalGroupAttributesContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalGroupAttributesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalGroupAttributes(s)
	}
}

func (s *ParentheticalGroupAttributesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalGroupAttributes(s)
	}
}

type GroupAttributesExpressionContext struct {
	BindRuleGroupAttrContext
}

func NewGroupAttributesExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GroupAttributesExpressionContext {
	var p = new(GroupAttributesExpressionContext)

	InitEmptyBindRuleGroupAttrContext(&p.BindRuleGroupAttrContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleGroupAttrContext))

	return p
}

func (s *GroupAttributesExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupAttributesExpressionContext) GroupAttrKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserGroupAttrKeyword, 0)
}

func (s *GroupAttributesExpressionContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *GroupAttributesExpressionContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *GroupAttributesExpressionContext) AttributeBindTypeOrValue() IAttributeBindTypeOrValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeBindTypeOrValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeBindTypeOrValueContext)
}

func (s *GroupAttributesExpressionContext) Inheritance() IInheritanceContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInheritanceContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInheritanceContext)
}

func (s *GroupAttributesExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterGroupAttributesExpression(s)
	}
}

func (s *GroupAttributesExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitGroupAttributesExpression(s)
	}
}

func (p *ACIParser) BindRuleGroupAttr() (localctx IBindRuleGroupAttrContext) {
	localctx = NewBindRuleGroupAttrContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ACIParserRULE_bindRuleGroupAttr)
	var _la int

	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserLPAREN:
		localctx = NewParentheticalGroupAttributesContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(350)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(351)
			p.BindRuleGroupAttr()
		}
		{
			p.SetState(352)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ACIParserGroupAttrKeyword:
		localctx = NewGroupAttributesExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(354)
			p.Match(ACIParserGroupAttrKeyword)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(355)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case ACIParserID:
			{
				p.SetState(356)
				p.AttributeBindTypeOrValue()
			}

		case ACIParserINHERITANCEPREFIX:
			{
				p.SetState(357)
				p.Inheritance()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
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

// IBindRuleAuthMethodContext is an interface to support dynamic dispatch.
type IBindRuleAuthMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBindRuleAuthMethodContext differentiates from other interfaces.
	IsBindRuleAuthMethodContext()
}

type BindRuleAuthMethodContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleAuthMethodContext() *BindRuleAuthMethodContext {
	var p = new(BindRuleAuthMethodContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleAuthMethod
	return p
}

func InitEmptyBindRuleAuthMethodContext(p *BindRuleAuthMethodContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleAuthMethod
}

func (*BindRuleAuthMethodContext) IsBindRuleAuthMethodContext() {}

func NewBindRuleAuthMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleAuthMethodContext {
	var p = new(BindRuleAuthMethodContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRuleAuthMethod

	return p
}

func (s *BindRuleAuthMethodContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleAuthMethodContext) CopyAll(ctx *BindRuleAuthMethodContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BindRuleAuthMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleAuthMethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AuthenticationMethodExpressionContext struct {
	BindRuleAuthMethodContext
}

func NewAuthenticationMethodExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AuthenticationMethodExpressionContext {
	var p = new(AuthenticationMethodExpressionContext)

	InitEmptyBindRuleAuthMethodContext(&p.BindRuleAuthMethodContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleAuthMethodContext))

	return p
}

func (s *AuthenticationMethodExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AuthenticationMethodExpressionContext) AuthMethodKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserAuthMethodKeyword, 0)
}

func (s *AuthenticationMethodExpressionContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *AuthenticationMethodExpressionContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *AuthenticationMethodExpressionContext) ANONYMOUS() antlr.TerminalNode {
	return s.GetToken(ACIParserANONYMOUS, 0)
}

func (s *AuthenticationMethodExpressionContext) SIMPLE() antlr.TerminalNode {
	return s.GetToken(ACIParserSIMPLE, 0)
}

func (s *AuthenticationMethodExpressionContext) SSL() antlr.TerminalNode {
	return s.GetToken(ACIParserSSL, 0)
}

func (s *AuthenticationMethodExpressionContext) SASL() antlr.TerminalNode {
	return s.GetToken(ACIParserSASL, 0)
}

func (s *AuthenticationMethodExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAuthenticationMethodExpression(s)
	}
}

func (s *AuthenticationMethodExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAuthenticationMethodExpression(s)
	}
}

type ParentheticalAuthenticationMethodContext struct {
	BindRuleAuthMethodContext
}

func NewParentheticalAuthenticationMethodContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalAuthenticationMethodContext {
	var p = new(ParentheticalAuthenticationMethodContext)

	InitEmptyBindRuleAuthMethodContext(&p.BindRuleAuthMethodContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleAuthMethodContext))

	return p
}

func (s *ParentheticalAuthenticationMethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalAuthenticationMethodContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalAuthenticationMethodContext) BindRuleAuthMethod() IBindRuleAuthMethodContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleAuthMethodContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleAuthMethodContext)
}

func (s *ParentheticalAuthenticationMethodContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalAuthenticationMethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalAuthenticationMethod(s)
	}
}

func (s *ParentheticalAuthenticationMethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalAuthenticationMethod(s)
	}
}

func (p *ACIParser) BindRuleAuthMethod() (localctx IBindRuleAuthMethodContext) {
	localctx = NewBindRuleAuthMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ACIParserRULE_bindRuleAuthMethod)
	var _la int

	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserLPAREN:
		localctx = NewParentheticalAuthenticationMethodContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(362)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(363)
			p.BindRuleAuthMethod()
		}
		{
			p.SetState(364)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ACIParserAuthMethodKeyword:
		localctx = NewAuthenticationMethodExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(366)
			p.Match(ACIParserAuthMethodKeyword)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(367)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(368)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&8053063680) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
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

// IBindRuleDNSContext is an interface to support dynamic dispatch.
type IBindRuleDNSContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBindRuleDNSContext differentiates from other interfaces.
	IsBindRuleDNSContext()
}

type BindRuleDNSContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleDNSContext() *BindRuleDNSContext {
	var p = new(BindRuleDNSContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleDNS
	return p
}

func InitEmptyBindRuleDNSContext(p *BindRuleDNSContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleDNS
}

func (*BindRuleDNSContext) IsBindRuleDNSContext() {}

func NewBindRuleDNSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleDNSContext {
	var p = new(BindRuleDNSContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRuleDNS

	return p
}

func (s *BindRuleDNSContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleDNSContext) CopyAll(ctx *BindRuleDNSContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BindRuleDNSContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleDNSContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentheticalDNSContext struct {
	BindRuleDNSContext
}

func NewParentheticalDNSContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalDNSContext {
	var p = new(ParentheticalDNSContext)

	InitEmptyBindRuleDNSContext(&p.BindRuleDNSContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleDNSContext))

	return p
}

func (s *ParentheticalDNSContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalDNSContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalDNSContext) BindRuleDNS() IBindRuleDNSContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleDNSContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleDNSContext)
}

func (s *ParentheticalDNSContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalDNSContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalDNS(s)
	}
}

func (s *ParentheticalDNSContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalDNS(s)
	}
}

type DNSBindRuleContext struct {
	BindRuleDNSContext
}

func NewDNSBindRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DNSBindRuleContext {
	var p = new(DNSBindRuleContext)

	InitEmptyBindRuleDNSContext(&p.BindRuleDNSContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleDNSContext))

	return p
}

func (s *DNSBindRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DNSBindRuleContext) DNSKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserDNSKeyword, 0)
}

func (s *DNSBindRuleContext) FQDN() IFQDNContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFQDNContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFQDNContext)
}

func (s *DNSBindRuleContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *DNSBindRuleContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *DNSBindRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterDNSBindRule(s)
	}
}

func (s *DNSBindRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitDNSBindRule(s)
	}
}

func (p *ACIParser) BindRuleDNS() (localctx IBindRuleDNSContext) {
	localctx = NewBindRuleDNSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ACIParserRULE_bindRuleDNS)
	var _la int

	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserLPAREN:
		localctx = NewParentheticalDNSContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(371)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(372)
			p.BindRuleDNS()
		}
		{
			p.SetState(373)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ACIParserDNSKeyword:
		localctx = NewDNSBindRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(375)
			p.Match(ACIParserDNSKeyword)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(376)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(377)
			p.FQDN()
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

// IBindRuleTimeOfDayContext is an interface to support dynamic dispatch.
type IBindRuleTimeOfDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBindRuleTimeOfDayContext differentiates from other interfaces.
	IsBindRuleTimeOfDayContext()
}

type BindRuleTimeOfDayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleTimeOfDayContext() *BindRuleTimeOfDayContext {
	var p = new(BindRuleTimeOfDayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleTimeOfDay
	return p
}

func InitEmptyBindRuleTimeOfDayContext(p *BindRuleTimeOfDayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleTimeOfDay
}

func (*BindRuleTimeOfDayContext) IsBindRuleTimeOfDayContext() {}

func NewBindRuleTimeOfDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleTimeOfDayContext {
	var p = new(BindRuleTimeOfDayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRuleTimeOfDay

	return p
}

func (s *BindRuleTimeOfDayContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleTimeOfDayContext) CopyAll(ctx *BindRuleTimeOfDayContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BindRuleTimeOfDayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleTimeOfDayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TimeOfDayBindRuleContext struct {
	BindRuleTimeOfDayContext
}

func NewTimeOfDayBindRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimeOfDayBindRuleContext {
	var p = new(TimeOfDayBindRuleContext)

	InitEmptyBindRuleTimeOfDayContext(&p.BindRuleTimeOfDayContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleTimeOfDayContext))

	return p
}

func (s *TimeOfDayBindRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeOfDayBindRuleContext) TimeOfDayKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserTimeOfDayKeyword, 0)
}

func (s *TimeOfDayBindRuleContext) AllDQUOTE() []antlr.TerminalNode {
	return s.GetTokens(ACIParserDQUOTE)
}

func (s *TimeOfDayBindRuleContext) DQUOTE(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserDQUOTE, i)
}

func (s *TimeOfDayBindRuleContext) TimeOfDay() ITimeOfDayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeOfDayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeOfDayContext)
}

func (s *TimeOfDayBindRuleContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *TimeOfDayBindRuleContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *TimeOfDayBindRuleContext) GreaterThan() antlr.TerminalNode {
	return s.GetToken(ACIParserGreaterThan, 0)
}

func (s *TimeOfDayBindRuleContext) GreaterThanOrEqual() antlr.TerminalNode {
	return s.GetToken(ACIParserGreaterThanOrEqual, 0)
}

func (s *TimeOfDayBindRuleContext) LessThan() antlr.TerminalNode {
	return s.GetToken(ACIParserLessThan, 0)
}

func (s *TimeOfDayBindRuleContext) LessThanOrEqual() antlr.TerminalNode {
	return s.GetToken(ACIParserLessThanOrEqual, 0)
}

func (s *TimeOfDayBindRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTimeOfDayBindRule(s)
	}
}

func (s *TimeOfDayBindRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTimeOfDayBindRule(s)
	}
}

type ParentheticalTimeOfDayContext struct {
	BindRuleTimeOfDayContext
}

func NewParentheticalTimeOfDayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalTimeOfDayContext {
	var p = new(ParentheticalTimeOfDayContext)

	InitEmptyBindRuleTimeOfDayContext(&p.BindRuleTimeOfDayContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleTimeOfDayContext))

	return p
}

func (s *ParentheticalTimeOfDayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalTimeOfDayContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalTimeOfDayContext) BindRuleTimeOfDay() IBindRuleTimeOfDayContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleTimeOfDayContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleTimeOfDayContext)
}

func (s *ParentheticalTimeOfDayContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalTimeOfDayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalTimeOfDay(s)
	}
}

func (s *ParentheticalTimeOfDayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalTimeOfDay(s)
	}
}

func (p *ACIParser) BindRuleTimeOfDay() (localctx IBindRuleTimeOfDayContext) {
	localctx = NewBindRuleTimeOfDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, ACIParserRULE_bindRuleTimeOfDay)
	var _la int

	p.SetState(390)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserLPAREN:
		localctx = NewParentheticalTimeOfDayContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(380)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(381)
			p.BindRuleTimeOfDay()
		}
		{
			p.SetState(382)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ACIParserTimeOfDayKeyword:
		localctx = NewTimeOfDayBindRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(384)
			p.Match(ACIParserTimeOfDayKeyword)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(385)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&252201579132754944) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(386)
			p.Match(ACIParserDQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(387)
			p.TimeOfDay()
		}
		{
			p.SetState(388)
			p.Match(ACIParserDQUOTE)
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

// IBindRuleDayOfWeekContext is an interface to support dynamic dispatch.
type IBindRuleDayOfWeekContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBindRuleDayOfWeekContext differentiates from other interfaces.
	IsBindRuleDayOfWeekContext()
}

type BindRuleDayOfWeekContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleDayOfWeekContext() *BindRuleDayOfWeekContext {
	var p = new(BindRuleDayOfWeekContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleDayOfWeek
	return p
}

func InitEmptyBindRuleDayOfWeekContext(p *BindRuleDayOfWeekContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleDayOfWeek
}

func (*BindRuleDayOfWeekContext) IsBindRuleDayOfWeekContext() {}

func NewBindRuleDayOfWeekContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleDayOfWeekContext {
	var p = new(BindRuleDayOfWeekContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRuleDayOfWeek

	return p
}

func (s *BindRuleDayOfWeekContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleDayOfWeekContext) CopyAll(ctx *BindRuleDayOfWeekContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BindRuleDayOfWeekContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleDayOfWeekContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentheticalDayOfWeekContext struct {
	BindRuleDayOfWeekContext
}

func NewParentheticalDayOfWeekContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalDayOfWeekContext {
	var p = new(ParentheticalDayOfWeekContext)

	InitEmptyBindRuleDayOfWeekContext(&p.BindRuleDayOfWeekContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleDayOfWeekContext))

	return p
}

func (s *ParentheticalDayOfWeekContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalDayOfWeekContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalDayOfWeekContext) BindRuleDayOfWeek() IBindRuleDayOfWeekContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleDayOfWeekContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleDayOfWeekContext)
}

func (s *ParentheticalDayOfWeekContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalDayOfWeekContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalDayOfWeek(s)
	}
}

func (s *ParentheticalDayOfWeekContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalDayOfWeek(s)
	}
}

type DayOfWeekExpressionContext struct {
	BindRuleDayOfWeekContext
}

func NewDayOfWeekExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DayOfWeekExpressionContext {
	var p = new(DayOfWeekExpressionContext)

	InitEmptyBindRuleDayOfWeekContext(&p.BindRuleDayOfWeekContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleDayOfWeekContext))

	return p
}

func (s *DayOfWeekExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayOfWeekExpressionContext) DayOfWeekKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserDayOfWeekKeyword, 0)
}

func (s *DayOfWeekExpressionContext) AllDQUOTE() []antlr.TerminalNode {
	return s.GetTokens(ACIParserDQUOTE)
}

func (s *DayOfWeekExpressionContext) DQUOTE(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserDQUOTE, i)
}

func (s *DayOfWeekExpressionContext) DayOfWeek() IDayOfWeekContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDayOfWeekContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDayOfWeekContext)
}

func (s *DayOfWeekExpressionContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *DayOfWeekExpressionContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *DayOfWeekExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterDayOfWeekExpression(s)
	}
}

func (s *DayOfWeekExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitDayOfWeekExpression(s)
	}
}

func (p *ACIParser) BindRuleDayOfWeek() (localctx IBindRuleDayOfWeekContext) {
	localctx = NewBindRuleDayOfWeekContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, ACIParserRULE_bindRuleDayOfWeek)
	var _la int

	p.SetState(402)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserLPAREN:
		localctx = NewParentheticalDayOfWeekContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(392)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(393)
			p.BindRuleDayOfWeek()
		}
		{
			p.SetState(394)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ACIParserDayOfWeekKeyword:
		localctx = NewDayOfWeekExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(396)
			p.Match(ACIParserDayOfWeekKeyword)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(397)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(398)
			p.Match(ACIParserDQUOTE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(399)
			p.DayOfWeek()
		}
		{
			p.SetState(400)
			p.Match(ACIParserDQUOTE)
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

// IBindRuleIPContext is an interface to support dynamic dispatch.
type IBindRuleIPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBindRuleIPContext differentiates from other interfaces.
	IsBindRuleIPContext()
}

type BindRuleIPContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleIPContext() *BindRuleIPContext {
	var p = new(BindRuleIPContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleIP
	return p
}

func InitEmptyBindRuleIPContext(p *BindRuleIPContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleIP
}

func (*BindRuleIPContext) IsBindRuleIPContext() {}

func NewBindRuleIPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleIPContext {
	var p = new(BindRuleIPContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRuleIP

	return p
}

func (s *BindRuleIPContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleIPContext) CopyAll(ctx *BindRuleIPContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BindRuleIPContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleIPContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IpAddressBindRuleContext struct {
	BindRuleIPContext
}

func NewIpAddressBindRuleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IpAddressBindRuleContext {
	var p = new(IpAddressBindRuleContext)

	InitEmptyBindRuleIPContext(&p.BindRuleIPContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleIPContext))

	return p
}

func (s *IpAddressBindRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IpAddressBindRuleContext) IPKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserIPKeyword, 0)
}

func (s *IpAddressBindRuleContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *IpAddressBindRuleContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *IpAddressBindRuleContext) IPV4Address() IIPV4AddressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIPV4AddressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIPV4AddressContext)
}

func (s *IpAddressBindRuleContext) IPV6Address() IIPV6AddressContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIPV6AddressContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIPV6AddressContext)
}

func (s *IpAddressBindRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterIpAddressBindRule(s)
	}
}

func (s *IpAddressBindRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitIpAddressBindRule(s)
	}
}

type ParentheticalIPAddressContext struct {
	BindRuleIPContext
}

func NewParentheticalIPAddressContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalIPAddressContext {
	var p = new(ParentheticalIPAddressContext)

	InitEmptyBindRuleIPContext(&p.BindRuleIPContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleIPContext))

	return p
}

func (s *ParentheticalIPAddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalIPAddressContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalIPAddressContext) BindRuleIP() IBindRuleIPContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleIPContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleIPContext)
}

func (s *ParentheticalIPAddressContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalIPAddressContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalIPAddress(s)
	}
}

func (s *ParentheticalIPAddressContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalIPAddress(s)
	}
}

func (p *ACIParser) BindRuleIP() (localctx IBindRuleIPContext) {
	localctx = NewBindRuleIPContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, ACIParserRULE_bindRuleIP)
	var _la int

	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserLPAREN:
		localctx = NewParentheticalIPAddressContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(404)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(405)
			p.BindRuleIP()
		}
		{
			p.SetState(406)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ACIParserIPKeyword:
		localctx = NewIpAddressBindRuleContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(408)
			p.Match(ACIParserIPKeyword)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(409)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ACIParserNotEqualTo || _la == ACIParserEqualTo) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		p.SetState(412)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(410)
				p.IPV4Address()
			}

		case 2:
			{
				p.SetState(411)
				p.IPV6Address()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
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

// IBindRuleSecurityStrengthFactorContext is an interface to support dynamic dispatch.
type IBindRuleSecurityStrengthFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBindRuleSecurityStrengthFactorContext differentiates from other interfaces.
	IsBindRuleSecurityStrengthFactorContext()
}

type BindRuleSecurityStrengthFactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindRuleSecurityStrengthFactorContext() *BindRuleSecurityStrengthFactorContext {
	var p = new(BindRuleSecurityStrengthFactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleSecurityStrengthFactor
	return p
}

func InitEmptyBindRuleSecurityStrengthFactorContext(p *BindRuleSecurityStrengthFactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_bindRuleSecurityStrengthFactor
}

func (*BindRuleSecurityStrengthFactorContext) IsBindRuleSecurityStrengthFactorContext() {}

func NewBindRuleSecurityStrengthFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindRuleSecurityStrengthFactorContext {
	var p = new(BindRuleSecurityStrengthFactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_bindRuleSecurityStrengthFactor

	return p
}

func (s *BindRuleSecurityStrengthFactorContext) GetParser() antlr.Parser { return s.parser }

func (s *BindRuleSecurityStrengthFactorContext) CopyAll(ctx *BindRuleSecurityStrengthFactorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *BindRuleSecurityStrengthFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindRuleSecurityStrengthFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentheticalSecurityStrengthFactorContext struct {
	BindRuleSecurityStrengthFactorContext
}

func NewParentheticalSecurityStrengthFactorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalSecurityStrengthFactorContext {
	var p = new(ParentheticalSecurityStrengthFactorContext)

	InitEmptyBindRuleSecurityStrengthFactorContext(&p.BindRuleSecurityStrengthFactorContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleSecurityStrengthFactorContext))

	return p
}

func (s *ParentheticalSecurityStrengthFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalSecurityStrengthFactorContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalSecurityStrengthFactorContext) BindRuleSecurityStrengthFactor() IBindRuleSecurityStrengthFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindRuleSecurityStrengthFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindRuleSecurityStrengthFactorContext)
}

func (s *ParentheticalSecurityStrengthFactorContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalSecurityStrengthFactorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalSecurityStrengthFactor(s)
	}
}

func (s *ParentheticalSecurityStrengthFactorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalSecurityStrengthFactor(s)
	}
}

type SecurityStrengthFactorExpressionContext struct {
	BindRuleSecurityStrengthFactorContext
}

func NewSecurityStrengthFactorExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SecurityStrengthFactorExpressionContext {
	var p = new(SecurityStrengthFactorExpressionContext)

	InitEmptyBindRuleSecurityStrengthFactorContext(&p.BindRuleSecurityStrengthFactorContext)
	p.parser = parser
	p.CopyAll(ctx.(*BindRuleSecurityStrengthFactorContext))

	return p
}

func (s *SecurityStrengthFactorExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SecurityStrengthFactorExpressionContext) SSFKeyword() antlr.TerminalNode {
	return s.GetToken(ACIParserSSFKeyword, 0)
}

func (s *SecurityStrengthFactorExpressionContext) SecurityStrengthFactor() ISecurityStrengthFactorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISecurityStrengthFactorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISecurityStrengthFactorContext)
}

func (s *SecurityStrengthFactorExpressionContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *SecurityStrengthFactorExpressionContext) NotEqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserNotEqualTo, 0)
}

func (s *SecurityStrengthFactorExpressionContext) GreaterThan() antlr.TerminalNode {
	return s.GetToken(ACIParserGreaterThan, 0)
}

func (s *SecurityStrengthFactorExpressionContext) GreaterThanOrEqual() antlr.TerminalNode {
	return s.GetToken(ACIParserGreaterThanOrEqual, 0)
}

func (s *SecurityStrengthFactorExpressionContext) LessThan() antlr.TerminalNode {
	return s.GetToken(ACIParserLessThan, 0)
}

func (s *SecurityStrengthFactorExpressionContext) LessThanOrEqual() antlr.TerminalNode {
	return s.GetToken(ACIParserLessThanOrEqual, 0)
}

func (s *SecurityStrengthFactorExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterSecurityStrengthFactorExpression(s)
	}
}

func (s *SecurityStrengthFactorExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitSecurityStrengthFactorExpression(s)
	}
}

func (p *ACIParser) BindRuleSecurityStrengthFactor() (localctx IBindRuleSecurityStrengthFactorContext) {
	localctx = NewBindRuleSecurityStrengthFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, ACIParserRULE_bindRuleSecurityStrengthFactor)
	var _la int

	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserLPAREN:
		localctx = NewParentheticalSecurityStrengthFactorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(416)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(417)
			p.BindRuleSecurityStrengthFactor()
		}
		{
			p.SetState(418)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ACIParserSSFKeyword:
		localctx = NewSecurityStrengthFactorExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(420)
			p.Match(ACIParserSSFKeyword)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(421)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&252201579132754944) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(422)
			p.SecurityStrengthFactor()
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

// IDayOfWeekContext is an interface to support dynamic dispatch.
type IDayOfWeekContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDayOfWeekContext differentiates from other interfaces.
	IsDayOfWeekContext()
}

type DayOfWeekContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDayOfWeekContext() *DayOfWeekContext {
	var p = new(DayOfWeekContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_dayOfWeek
	return p
}

func InitEmptyDayOfWeekContext(p *DayOfWeekContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_dayOfWeek
}

func (*DayOfWeekContext) IsDayOfWeekContext() {}

func NewDayOfWeekContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DayOfWeekContext {
	var p = new(DayOfWeekContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_dayOfWeek

	return p
}

func (s *DayOfWeekContext) GetParser() antlr.Parser { return s.parser }

func (s *DayOfWeekContext) CopyAll(ctx *DayOfWeekContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DayOfWeekContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayOfWeekContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DayOfWeekValueContext struct {
	DayOfWeekContext
}

func NewDayOfWeekValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DayOfWeekValueContext {
	var p = new(DayOfWeekValueContext)

	InitEmptyDayOfWeekContext(&p.DayOfWeekContext)
	p.parser = parser
	p.CopyAll(ctx.(*DayOfWeekContext))

	return p
}

func (s *DayOfWeekValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DayOfWeekValueContext) AllDAYS() []antlr.TerminalNode {
	return s.GetTokens(ACIParserDAYS)
}

func (s *DayOfWeekValueContext) DAYS(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserDAYS, i)
}

func (s *DayOfWeekValueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ACIParserCOMMA)
}

func (s *DayOfWeekValueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserCOMMA, i)
}

func (s *DayOfWeekValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterDayOfWeekValue(s)
	}
}

func (s *DayOfWeekValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitDayOfWeekValue(s)
	}
}

func (p *ACIParser) DayOfWeek() (localctx IDayOfWeekContext) {
	localctx = NewDayOfWeekContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ACIParserRULE_dayOfWeek)
	var _la int

	localctx = NewDayOfWeekValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(433)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ACIParserDAYS {
		{
			p.SetState(425)
			p.Match(ACIParserDAYS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ACIParserCOMMA {
			{
				p.SetState(426)
				p.Match(ACIParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(427)
				p.Match(ACIParserDAYS)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(432)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// IFQDNContext is an interface to support dynamic dispatch.
type IFQDNContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsFQDNContext differentiates from other interfaces.
	IsFQDNContext()
}

type FQDNContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFQDNContext() *FQDNContext {
	var p = new(FQDNContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_fQDN
	return p
}

func InitEmptyFQDNContext(p *FQDNContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_fQDN
}

func (*FQDNContext) IsFQDNContext() {}

func NewFQDNContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FQDNContext {
	var p = new(FQDNContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_fQDN

	return p
}

func (s *FQDNContext) GetParser() antlr.Parser { return s.parser }

func (s *FQDNContext) CopyAll(ctx *FQDNContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FQDNContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FQDNContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FullyQualifiedDomainNameValueContext struct {
	FQDNContext
}

func NewFullyQualifiedDomainNameValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FullyQualifiedDomainNameValueContext {
	var p = new(FullyQualifiedDomainNameValueContext)

	InitEmptyFQDNContext(&p.FQDNContext)
	p.parser = parser
	p.CopyAll(ctx.(*FQDNContext))

	return p
}

func (s *FullyQualifiedDomainNameValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullyQualifiedDomainNameValueContext) DelimitedAddress() antlr.TerminalNode {
	return s.GetToken(ACIParserDelimitedAddress, 0)
}

func (s *FullyQualifiedDomainNameValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterFullyQualifiedDomainNameValue(s)
	}
}

func (s *FullyQualifiedDomainNameValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitFullyQualifiedDomainNameValue(s)
	}
}

func (p *ACIParser) FQDN() (localctx IFQDNContext) {
	localctx = NewFQDNContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, ACIParserRULE_fQDN)
	localctx = NewFullyQualifiedDomainNameValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(435)
		p.Match(ACIParserDelimitedAddress)
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

// IObjectIdentifiersContext is an interface to support dynamic dispatch.
type IObjectIdentifiersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsObjectIdentifiersContext differentiates from other interfaces.
	IsObjectIdentifiersContext()
}

type ObjectIdentifiersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectIdentifiersContext() *ObjectIdentifiersContext {
	var p = new(ObjectIdentifiersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_objectIdentifiers
	return p
}

func InitEmptyObjectIdentifiersContext(p *ObjectIdentifiersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_objectIdentifiers
}

func (*ObjectIdentifiersContext) IsObjectIdentifiersContext() {}

func NewObjectIdentifiersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectIdentifiersContext {
	var p = new(ObjectIdentifiersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_objectIdentifiers

	return p
}

func (s *ObjectIdentifiersContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectIdentifiersContext) CopyAll(ctx *ObjectIdentifiersContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ObjectIdentifiersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectIdentifiersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ObjectIdentifierValuesContext struct {
	ObjectIdentifiersContext
}

func NewObjectIdentifierValuesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectIdentifierValuesContext {
	var p = new(ObjectIdentifierValuesContext)

	InitEmptyObjectIdentifiersContext(&p.ObjectIdentifiersContext)
	p.parser = parser
	p.CopyAll(ctx.(*ObjectIdentifiersContext))

	return p
}

func (s *ObjectIdentifierValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectIdentifierValuesContext) AllObjectIdentifier() []IObjectIdentifierContext {
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

func (s *ObjectIdentifierValuesContext) ObjectIdentifier(i int) IObjectIdentifierContext {
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

func (s *ObjectIdentifierValuesContext) AllDoublePipe() []IDoublePipeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoublePipeContext); ok {
			len++
		}
	}

	tst := make([]IDoublePipeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoublePipeContext); ok {
			tst[i] = t.(IDoublePipeContext)
			i++
		}
	}

	return tst
}

func (s *ObjectIdentifierValuesContext) DoublePipe(i int) IDoublePipeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoublePipeContext); ok {
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

	return t.(IDoublePipeContext)
}

func (s *ObjectIdentifierValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterObjectIdentifierValues(s)
	}
}

func (s *ObjectIdentifierValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitObjectIdentifierValues(s)
	}
}

func (p *ACIParser) ObjectIdentifiers() (localctx IObjectIdentifiersContext) {
	localctx = NewObjectIdentifiersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ACIParserRULE_objectIdentifiers)
	var _la int

	localctx = NewObjectIdentifierValuesContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(437)
		p.ObjectIdentifier()
	}
	p.SetState(441)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ACIParserDPIPE {
		{
			p.SetState(438)
			p.DoublePipe()
		}
		{
			p.SetState(439)
			p.ObjectIdentifier()
		}

		p.SetState(443)
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

// IObjectIdentifierContext is an interface to support dynamic dispatch.
type IObjectIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = ACIParserRULE_objectIdentifier
	return p
}

func InitEmptyObjectIdentifierContext(p *ObjectIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_objectIdentifier
}

func (*ObjectIdentifierContext) IsObjectIdentifierContext() {}

func NewObjectIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectIdentifierContext {
	var p = new(ObjectIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_objectIdentifier

	return p
}

func (s *ObjectIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectIdentifierContext) CopyAll(ctx *ObjectIdentifierContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ObjectIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ObjectIdentifierValueContext struct {
	ObjectIdentifierContext
}

func NewObjectIdentifierValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ObjectIdentifierValueContext {
	var p = new(ObjectIdentifierValueContext)

	InitEmptyObjectIdentifierContext(&p.ObjectIdentifierContext)
	p.parser = parser
	p.CopyAll(ctx.(*ObjectIdentifierContext))

	return p
}

func (s *ObjectIdentifierValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectIdentifierValueContext) DelimitedAddress() antlr.TerminalNode {
	return s.GetToken(ACIParserDelimitedAddress, 0)
}

func (s *ObjectIdentifierValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterObjectIdentifierValue(s)
	}
}

func (s *ObjectIdentifierValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitObjectIdentifierValue(s)
	}
}

func (p *ACIParser) ObjectIdentifier() (localctx IObjectIdentifierContext) {
	localctx = NewObjectIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, ACIParserRULE_objectIdentifier)
	localctx = NewObjectIdentifierValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(445)
		p.Match(ACIParserDelimitedAddress)
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

// IIPV6AddressContext is an interface to support dynamic dispatch.
type IIPV6AddressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIPV6AddressContext differentiates from other interfaces.
	IsIPV6AddressContext()
}

type IPV6AddressContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIPV6AddressContext() *IPV6AddressContext {
	var p = new(IPV6AddressContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_iPV6Address
	return p
}

func InitEmptyIPV6AddressContext(p *IPV6AddressContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_iPV6Address
}

func (*IPV6AddressContext) IsIPV6AddressContext() {}

func NewIPV6AddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IPV6AddressContext {
	var p = new(IPV6AddressContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_iPV6Address

	return p
}

func (s *IPV6AddressContext) GetParser() antlr.Parser { return s.parser }

func (s *IPV6AddressContext) CopyAll(ctx *IPV6AddressContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IPV6AddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IPV6AddressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IPV6AddressValueContext struct {
	IPV6AddressContext
}

func NewIPV6AddressValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IPV6AddressValueContext {
	var p = new(IPV6AddressValueContext)

	InitEmptyIPV6AddressContext(&p.IPV6AddressContext)
	p.parser = parser
	p.CopyAll(ctx.(*IPV6AddressContext))

	return p
}

func (s *IPV6AddressValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IPV6AddressValueContext) DelimitedAddress() antlr.TerminalNode {
	return s.GetToken(ACIParserDelimitedAddress, 0)
}

func (s *IPV6AddressValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterIPV6AddressValue(s)
	}
}

func (s *IPV6AddressValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitIPV6AddressValue(s)
	}
}

func (p *ACIParser) IPV6Address() (localctx IIPV6AddressContext) {
	localctx = NewIPV6AddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, ACIParserRULE_iPV6Address)
	localctx = NewIPV6AddressValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(447)
		p.Match(ACIParserDelimitedAddress)
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

// IIPV4AddressContext is an interface to support dynamic dispatch.
type IIPV4AddressContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIPV4AddressContext differentiates from other interfaces.
	IsIPV4AddressContext()
}

type IPV4AddressContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIPV4AddressContext() *IPV4AddressContext {
	var p = new(IPV4AddressContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_iPV4Address
	return p
}

func InitEmptyIPV4AddressContext(p *IPV4AddressContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_iPV4Address
}

func (*IPV4AddressContext) IsIPV4AddressContext() {}

func NewIPV4AddressContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IPV4AddressContext {
	var p = new(IPV4AddressContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_iPV4Address

	return p
}

func (s *IPV4AddressContext) GetParser() antlr.Parser { return s.parser }

func (s *IPV4AddressContext) CopyAll(ctx *IPV4AddressContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IPV4AddressContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IPV4AddressContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IPV4AddressValueContext struct {
	IPV4AddressContext
}

func NewIPV4AddressValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IPV4AddressValueContext {
	var p = new(IPV4AddressValueContext)

	InitEmptyIPV4AddressContext(&p.IPV4AddressContext)
	p.parser = parser
	p.CopyAll(ctx.(*IPV4AddressContext))

	return p
}

func (s *IPV4AddressValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IPV4AddressValueContext) DelimitedAddress() antlr.TerminalNode {
	return s.GetToken(ACIParserDelimitedAddress, 0)
}

func (s *IPV4AddressValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterIPV4AddressValue(s)
	}
}

func (s *IPV4AddressValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitIPV4AddressValue(s)
	}
}

func (p *ACIParser) IPV4Address() (localctx IIPV4AddressContext) {
	localctx = NewIPV4AddressContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, ACIParserRULE_iPV4Address)
	localctx = NewIPV4AddressValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(449)
		p.Match(ACIParserDelimitedAddress)
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

// ISecurityStrengthFactorContext is an interface to support dynamic dispatch.
type ISecurityStrengthFactorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSecurityStrengthFactorContext differentiates from other interfaces.
	IsSecurityStrengthFactorContext()
}

type SecurityStrengthFactorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySecurityStrengthFactorContext() *SecurityStrengthFactorContext {
	var p = new(SecurityStrengthFactorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_securityStrengthFactor
	return p
}

func InitEmptySecurityStrengthFactorContext(p *SecurityStrengthFactorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_securityStrengthFactor
}

func (*SecurityStrengthFactorContext) IsSecurityStrengthFactorContext() {}

func NewSecurityStrengthFactorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SecurityStrengthFactorContext {
	var p = new(SecurityStrengthFactorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_securityStrengthFactor

	return p
}

func (s *SecurityStrengthFactorContext) GetParser() antlr.Parser { return s.parser }

func (s *SecurityStrengthFactorContext) CopyAll(ctx *SecurityStrengthFactorContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SecurityStrengthFactorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SecurityStrengthFactorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SecurityStrengthFactorValueContext struct {
	SecurityStrengthFactorContext
}

func NewSecurityStrengthFactorValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SecurityStrengthFactorValueContext {
	var p = new(SecurityStrengthFactorValueContext)

	InitEmptySecurityStrengthFactorContext(&p.SecurityStrengthFactorContext)
	p.parser = parser
	p.CopyAll(ctx.(*SecurityStrengthFactorContext))

	return p
}

func (s *SecurityStrengthFactorValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SecurityStrengthFactorValueContext) INT() antlr.TerminalNode {
	return s.GetToken(ACIParserINT, 0)
}

func (s *SecurityStrengthFactorValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterSecurityStrengthFactorValue(s)
	}
}

func (s *SecurityStrengthFactorValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitSecurityStrengthFactorValue(s)
	}
}

func (p *ACIParser) SecurityStrengthFactor() (localctx ISecurityStrengthFactorContext) {
	localctx = NewSecurityStrengthFactorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, ACIParserRULE_securityStrengthFactor)
	localctx = NewSecurityStrengthFactorValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(451)
		p.Match(ACIParserINT)
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

// ITimeOfDayContext is an interface to support dynamic dispatch.
type ITimeOfDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTimeOfDayContext differentiates from other interfaces.
	IsTimeOfDayContext()
}

type TimeOfDayContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeOfDayContext() *TimeOfDayContext {
	var p = new(TimeOfDayContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_timeOfDay
	return p
}

func InitEmptyTimeOfDayContext(p *TimeOfDayContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_timeOfDay
}

func (*TimeOfDayContext) IsTimeOfDayContext() {}

func NewTimeOfDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeOfDayContext {
	var p = new(TimeOfDayContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_timeOfDay

	return p
}

func (s *TimeOfDayContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeOfDayContext) CopyAll(ctx *TimeOfDayContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TimeOfDayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeOfDayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TimeOfDayValueContext struct {
	TimeOfDayContext
}

func NewTimeOfDayValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TimeOfDayValueContext {
	var p = new(TimeOfDayValueContext)

	InitEmptyTimeOfDayContext(&p.TimeOfDayContext)
	p.parser = parser
	p.CopyAll(ctx.(*TimeOfDayContext))

	return p
}

func (s *TimeOfDayValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeOfDayValueContext) INT() antlr.TerminalNode {
	return s.GetToken(ACIParserINT, 0)
}

func (s *TimeOfDayValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterTimeOfDayValue(s)
	}
}

func (s *TimeOfDayValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitTimeOfDayValue(s)
	}
}

func (p *ACIParser) TimeOfDay() (localctx ITimeOfDayContext) {
	localctx = NewTimeOfDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, ACIParserRULE_timeOfDay)
	localctx = NewTimeOfDayValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(453)
		p.Match(ACIParserINT)
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

// IObjectIdentifierArcContext is an interface to support dynamic dispatch.
type IObjectIdentifierArcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode

	// IsObjectIdentifierArcContext differentiates from other interfaces.
	IsObjectIdentifierArcContext()
}

type ObjectIdentifierArcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectIdentifierArcContext() *ObjectIdentifierArcContext {
	var p = new(ObjectIdentifierArcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_objectIdentifierArc
	return p
}

func InitEmptyObjectIdentifierArcContext(p *ObjectIdentifierArcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_objectIdentifierArc
}

func (*ObjectIdentifierArcContext) IsObjectIdentifierArcContext() {}

func NewObjectIdentifierArcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectIdentifierArcContext {
	var p = new(ObjectIdentifierArcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_objectIdentifierArc

	return p
}

func (s *ObjectIdentifierArcContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectIdentifierArcContext) INT() antlr.TerminalNode {
	return s.GetToken(ACIParserINT, 0)
}

func (s *ObjectIdentifierArcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectIdentifierArcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectIdentifierArcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterObjectIdentifierArc(s)
	}
}

func (s *ObjectIdentifierArcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitObjectIdentifierArc(s)
	}
}

func (p *ACIParser) ObjectIdentifierArc() (localctx IObjectIdentifierArcContext) {
	localctx = NewObjectIdentifierArcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, ACIParserRULE_objectIdentifierArc)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(455)
		p.Match(ACIParserINT)
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

// IInheritanceContext is an interface to support dynamic dispatch.
type IInheritanceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsInheritanceContext differentiates from other interfaces.
	IsInheritanceContext()
}

type InheritanceContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInheritanceContext() *InheritanceContext {
	var p = new(InheritanceContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_inheritance
	return p
}

func InitEmptyInheritanceContext(p *InheritanceContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_inheritance
}

func (*InheritanceContext) IsInheritanceContext() {}

func NewInheritanceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InheritanceContext {
	var p = new(InheritanceContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_inheritance

	return p
}

func (s *InheritanceContext) GetParser() antlr.Parser { return s.parser }

func (s *InheritanceContext) CopyAll(ctx *InheritanceContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *InheritanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InheritanceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InheritanceExpressionContext struct {
	InheritanceContext
}

func NewInheritanceExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InheritanceExpressionContext {
	var p = new(InheritanceExpressionContext)

	InitEmptyInheritanceContext(&p.InheritanceContext)
	p.parser = parser
	p.CopyAll(ctx.(*InheritanceContext))

	return p
}

func (s *InheritanceExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InheritanceExpressionContext) INHERITANCEPREFIX() antlr.TerminalNode {
	return s.GetToken(ACIParserINHERITANCEPREFIX, 0)
}

func (s *InheritanceExpressionContext) InheritanceLevels() IInheritanceLevelsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInheritanceLevelsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInheritanceLevelsContext)
}

func (s *InheritanceExpressionContext) RBRAK() antlr.TerminalNode {
	return s.GetToken(ACIParserRBRAK, 0)
}

func (s *InheritanceExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(ACIParserDOT, 0)
}

func (s *InheritanceExpressionContext) AttributeBindTypeOrValue() IAttributeBindTypeOrValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeBindTypeOrValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeBindTypeOrValueContext)
}

func (s *InheritanceExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterInheritanceExpression(s)
	}
}

func (s *InheritanceExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitInheritanceExpression(s)
	}
}

func (p *ACIParser) Inheritance() (localctx IInheritanceContext) {
	localctx = NewInheritanceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, ACIParserRULE_inheritance)
	localctx = NewInheritanceExpressionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(457)
		p.Match(ACIParserINHERITANCEPREFIX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(458)
		p.InheritanceLevels()
	}
	{
		p.SetState(459)
		p.Match(ACIParserRBRAK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(460)
		p.Match(ACIParserDOT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(461)
		p.AttributeBindTypeOrValue()
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

// IInheritanceLevelsContext is an interface to support dynamic dispatch.
type IInheritanceLevelsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsInheritanceLevelsContext differentiates from other interfaces.
	IsInheritanceLevelsContext()
}

type InheritanceLevelsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInheritanceLevelsContext() *InheritanceLevelsContext {
	var p = new(InheritanceLevelsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_inheritanceLevels
	return p
}

func InitEmptyInheritanceLevelsContext(p *InheritanceLevelsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_inheritanceLevels
}

func (*InheritanceLevelsContext) IsInheritanceLevelsContext() {}

func NewInheritanceLevelsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InheritanceLevelsContext {
	var p = new(InheritanceLevelsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_inheritanceLevels

	return p
}

func (s *InheritanceLevelsContext) GetParser() antlr.Parser { return s.parser }

func (s *InheritanceLevelsContext) CopyAll(ctx *InheritanceLevelsContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *InheritanceLevelsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InheritanceLevelsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InheritanceLevelValueContext struct {
	InheritanceLevelsContext
}

func NewInheritanceLevelValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InheritanceLevelValueContext {
	var p = new(InheritanceLevelValueContext)

	InitEmptyInheritanceLevelsContext(&p.InheritanceLevelsContext)
	p.parser = parser
	p.CopyAll(ctx.(*InheritanceLevelsContext))

	return p
}

func (s *InheritanceLevelValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InheritanceLevelValueContext) DelimitedNumbers() antlr.TerminalNode {
	return s.GetToken(ACIParserDelimitedNumbers, 0)
}

func (s *InheritanceLevelValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterInheritanceLevelValue(s)
	}
}

func (s *InheritanceLevelValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitInheritanceLevelValue(s)
	}
}

func (p *ACIParser) InheritanceLevels() (localctx IInheritanceLevelsContext) {
	localctx = NewInheritanceLevelsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, ACIParserRULE_inheritanceLevels)
	localctx = NewInheritanceLevelValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(463)
		p.Match(ACIParserDelimitedNumbers)
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

// IAttributeBindTypeOrValueContext is an interface to support dynamic dispatch.
type IAttributeBindTypeOrValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAttributeBindTypeOrValueContext differentiates from other interfaces.
	IsAttributeBindTypeOrValueContext()
}

type AttributeBindTypeOrValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeBindTypeOrValueContext() *AttributeBindTypeOrValueContext {
	var p = new(AttributeBindTypeOrValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeBindTypeOrValue
	return p
}

func InitEmptyAttributeBindTypeOrValueContext(p *AttributeBindTypeOrValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeBindTypeOrValue
}

func (*AttributeBindTypeOrValueContext) IsAttributeBindTypeOrValueContext() {}

func NewAttributeBindTypeOrValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeBindTypeOrValueContext {
	var p = new(AttributeBindTypeOrValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_attributeBindTypeOrValue

	return p
}

func (s *AttributeBindTypeOrValueContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeBindTypeOrValueContext) CopyAll(ctx *AttributeBindTypeOrValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AttributeBindTypeOrValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeBindTypeOrValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AttributeBindTypeOrValueValueContext struct {
	AttributeBindTypeOrValueContext
}

func NewAttributeBindTypeOrValueValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttributeBindTypeOrValueValueContext {
	var p = new(AttributeBindTypeOrValueValueContext)

	InitEmptyAttributeBindTypeOrValueContext(&p.AttributeBindTypeOrValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*AttributeBindTypeOrValueContext))

	return p
}

func (s *AttributeBindTypeOrValueValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeBindTypeOrValueValueContext) AttributeType() IAttributeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeTypeContext)
}

func (s *AttributeBindTypeOrValueValueContext) HASH() antlr.TerminalNode {
	return s.GetToken(ACIParserHASH, 0)
}

func (s *AttributeBindTypeOrValueValueContext) BINDTYPES() antlr.TerminalNode {
	return s.GetToken(ACIParserBINDTYPES, 0)
}

func (s *AttributeBindTypeOrValueValueContext) AttributeValue() IAttributeValueContext {
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

func (s *AttributeBindTypeOrValueValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAttributeBindTypeOrValueValue(s)
	}
}

func (s *AttributeBindTypeOrValueValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAttributeBindTypeOrValueValue(s)
	}
}

func (p *ACIParser) AttributeBindTypeOrValue() (localctx IAttributeBindTypeOrValueContext) {
	localctx = NewAttributeBindTypeOrValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, ACIParserRULE_attributeBindTypeOrValue)
	localctx = NewAttributeBindTypeOrValueValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(465)
		p.AttributeType()
	}
	{
		p.SetState(466)
		p.Match(ACIParserHASH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserBINDTYPES:
		{
			p.SetState(467)
			p.Match(ACIParserBINDTYPES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ACIParserINT, ACIParserID, ACIParserWildcardString, ACIParserANY:
		{
			p.SetState(468)
			p.AttributeValue()
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

// IAttributeFiltersContext is an interface to support dynamic dispatch.
type IAttributeFiltersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAttributeFiltersContext differentiates from other interfaces.
	IsAttributeFiltersContext()
}

type AttributeFiltersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeFiltersContext() *AttributeFiltersContext {
	var p = new(AttributeFiltersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeFilters
	return p
}

func InitEmptyAttributeFiltersContext(p *AttributeFiltersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeFilters
}

func (*AttributeFiltersContext) IsAttributeFiltersContext() {}

func NewAttributeFiltersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeFiltersContext {
	var p = new(AttributeFiltersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_attributeFilters

	return p
}

func (s *AttributeFiltersContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeFiltersContext) CopyAll(ctx *AttributeFiltersContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AttributeFiltersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeFiltersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AttributeFiltersExpressionContext struct {
	AttributeFiltersContext
}

func NewAttributeFiltersExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttributeFiltersExpressionContext {
	var p = new(AttributeFiltersExpressionContext)

	InitEmptyAttributeFiltersContext(&p.AttributeFiltersContext)
	p.parser = parser
	p.CopyAll(ctx.(*AttributeFiltersContext))

	return p
}

func (s *AttributeFiltersExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeFiltersExpressionContext) AllAttributeFilterSet() []IAttributeFilterSetContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeFilterSetContext); ok {
			len++
		}
	}

	tst := make([]IAttributeFilterSetContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeFilterSetContext); ok {
			tst[i] = t.(IAttributeFilterSetContext)
			i++
		}
	}

	return tst
}

func (s *AttributeFiltersExpressionContext) AttributeFilterSet(i int) IAttributeFilterSetContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeFilterSetContext); ok {
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

	return t.(IAttributeFilterSetContext)
}

func (s *AttributeFiltersExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ACIParserCOMMA)
}

func (s *AttributeFiltersExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserCOMMA, i)
}

func (s *AttributeFiltersExpressionContext) AllSEMI() []antlr.TerminalNode {
	return s.GetTokens(ACIParserSEMI)
}

func (s *AttributeFiltersExpressionContext) SEMI(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserSEMI, i)
}

func (s *AttributeFiltersExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAttributeFiltersExpression(s)
	}
}

func (s *AttributeFiltersExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAttributeFiltersExpression(s)
	}
}

func (p *ACIParser) AttributeFilters() (localctx IAttributeFiltersContext) {
	localctx = NewAttributeFiltersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, ACIParserRULE_attributeFilters)
	var _la int

	var _alt int

	localctx = NewAttributeFiltersExpressionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(471)
		p.AttributeFilterSet()
	}
	p.SetState(474)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = 1 + 1
	for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1 + 1:
			{
				p.SetState(472)
				_la = p.GetTokenStream().LA(1)

				if !(_la == ACIParserSEMI || _la == ACIParserCOMMA) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(473)
				p.AttributeFilterSet()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(476)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext())
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

// IAttributeFilterSetContext is an interface to support dynamic dispatch.
type IAttributeFilterSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAttributeFilterSetContext differentiates from other interfaces.
	IsAttributeFilterSetContext()
}

type AttributeFilterSetContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeFilterSetContext() *AttributeFilterSetContext {
	var p = new(AttributeFilterSetContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeFilterSet
	return p
}

func InitEmptyAttributeFilterSetContext(p *AttributeFilterSetContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeFilterSet
}

func (*AttributeFilterSetContext) IsAttributeFilterSetContext() {}

func NewAttributeFilterSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeFilterSetContext {
	var p = new(AttributeFilterSetContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_attributeFilterSet

	return p
}

func (s *AttributeFilterSetContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeFilterSetContext) CopyAll(ctx *AttributeFilterSetContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AttributeFilterSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeFilterSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AttributeFilterSetExpressionContext struct {
	AttributeFilterSetContext
}

func NewAttributeFilterSetExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttributeFilterSetExpressionContext {
	var p = new(AttributeFilterSetExpressionContext)

	InitEmptyAttributeFilterSetContext(&p.AttributeFilterSetContext)
	p.parser = parser
	p.CopyAll(ctx.(*AttributeFilterSetContext))

	return p
}

func (s *AttributeFilterSetExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeFilterSetExpressionContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *AttributeFilterSetExpressionContext) AddOperation() antlr.TerminalNode {
	return s.GetToken(ACIParserAddOperation, 0)
}

func (s *AttributeFilterSetExpressionContext) DeleteOperation() antlr.TerminalNode {
	return s.GetToken(ACIParserDeleteOperation, 0)
}

func (s *AttributeFilterSetExpressionContext) AllAttributeFilter() []IAttributeFilterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeFilterContext); ok {
			len++
		}
	}

	tst := make([]IAttributeFilterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeFilterContext); ok {
			tst[i] = t.(IAttributeFilterContext)
			i++
		}
	}

	return tst
}

func (s *AttributeFilterSetExpressionContext) AttributeFilter(i int) IAttributeFilterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeFilterContext); ok {
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

	return t.(IAttributeFilterContext)
}

func (s *AttributeFilterSetExpressionContext) AllDoubleAmpersand() []IDoubleAmpersandContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoubleAmpersandContext); ok {
			len++
		}
	}

	tst := make([]IDoubleAmpersandContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoubleAmpersandContext); ok {
			tst[i] = t.(IDoubleAmpersandContext)
			i++
		}
	}

	return tst
}

func (s *AttributeFilterSetExpressionContext) DoubleAmpersand(i int) IDoubleAmpersandContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoubleAmpersandContext); ok {
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

	return t.(IDoubleAmpersandContext)
}

func (s *AttributeFilterSetExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAttributeFilterSetExpression(s)
	}
}

func (s *AttributeFilterSetExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAttributeFilterSetExpression(s)
	}
}

func (p *ACIParser) AttributeFilterSet() (localctx IAttributeFilterSetContext) {
	localctx = NewAttributeFilterSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, ACIParserRULE_attributeFilterSet)
	var _la int

	localctx = NewAttributeFilterSetExpressionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(478)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ACIParserAddOperation || _la == ACIParserDeleteOperation) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(479)
		p.Match(ACIParserEqualTo)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(489)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ACIParserID {
		{
			p.SetState(480)
			p.AttributeFilter()
		}
		p.SetState(486)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ACIParserDAMP {
			{
				p.SetState(481)
				p.DoubleAmpersand()
			}
			{
				p.SetState(482)
				p.AttributeFilter()
			}

			p.SetState(488)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// IDoubleAmpersandContext is an interface to support dynamic dispatch.
type IDoubleAmpersandContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDoubleAmpersandContext differentiates from other interfaces.
	IsDoubleAmpersandContext()
}

type DoubleAmpersandContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoubleAmpersandContext() *DoubleAmpersandContext {
	var p = new(DoubleAmpersandContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_doubleAmpersand
	return p
}

func InitEmptyDoubleAmpersandContext(p *DoubleAmpersandContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_doubleAmpersand
}

func (*DoubleAmpersandContext) IsDoubleAmpersandContext() {}

func NewDoubleAmpersandContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoubleAmpersandContext {
	var p = new(DoubleAmpersandContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_doubleAmpersand

	return p
}

func (s *DoubleAmpersandContext) GetParser() antlr.Parser { return s.parser }

func (s *DoubleAmpersandContext) CopyAll(ctx *DoubleAmpersandContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DoubleAmpersandContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleAmpersandContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DoubleAmpersandDelimiterContext struct {
	DoubleAmpersandContext
}

func NewDoubleAmpersandDelimiterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoubleAmpersandDelimiterContext {
	var p = new(DoubleAmpersandDelimiterContext)

	InitEmptyDoubleAmpersandContext(&p.DoubleAmpersandContext)
	p.parser = parser
	p.CopyAll(ctx.(*DoubleAmpersandContext))

	return p
}

func (s *DoubleAmpersandDelimiterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoubleAmpersandDelimiterContext) DAMP() antlr.TerminalNode {
	return s.GetToken(ACIParserDAMP, 0)
}

func (s *DoubleAmpersandDelimiterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterDoubleAmpersandDelimiter(s)
	}
}

func (s *DoubleAmpersandDelimiterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitDoubleAmpersandDelimiter(s)
	}
}

func (p *ACIParser) DoubleAmpersand() (localctx IDoubleAmpersandContext) {
	localctx = NewDoubleAmpersandContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, ACIParserRULE_doubleAmpersand)
	localctx = NewDoubleAmpersandDelimiterContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(491)
		p.Match(ACIParserDAMP)
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

// IAttributeFilterContext is an interface to support dynamic dispatch.
type IAttributeFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAttributeFilterContext differentiates from other interfaces.
	IsAttributeFilterContext()
}

type AttributeFilterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeFilterContext() *AttributeFilterContext {
	var p = new(AttributeFilterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeFilter
	return p
}

func InitEmptyAttributeFilterContext(p *AttributeFilterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeFilter
}

func (*AttributeFilterContext) IsAttributeFilterContext() {}

func NewAttributeFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeFilterContext {
	var p = new(AttributeFilterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_attributeFilter

	return p
}

func (s *AttributeFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeFilterContext) CopyAll(ctx *AttributeFilterContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AttributeFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AttributeFilterExpressionContext struct {
	AttributeFilterContext
}

func NewAttributeFilterExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttributeFilterExpressionContext {
	var p = new(AttributeFilterExpressionContext)

	InitEmptyAttributeFilterContext(&p.AttributeFilterContext)
	p.parser = parser
	p.CopyAll(ctx.(*AttributeFilterContext))

	return p
}

func (s *AttributeFilterExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeFilterExpressionContext) AttributeType() IAttributeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeTypeContext)
}

func (s *AttributeFilterExpressionContext) COLON() antlr.TerminalNode {
	return s.GetToken(ACIParserCOLON, 0)
}

func (s *AttributeFilterExpressionContext) LDAPFilter() ILDAPFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILDAPFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILDAPFilterContext)
}

func (s *AttributeFilterExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAttributeFilterExpression(s)
	}
}

func (s *AttributeFilterExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAttributeFilterExpression(s)
	}
}

func (p *ACIParser) AttributeFilter() (localctx IAttributeFilterContext) {
	localctx = NewAttributeFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, ACIParserRULE_attributeFilter)
	localctx = NewAttributeFilterExpressionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(493)
		p.AttributeType()
	}
	{
		p.SetState(494)
		p.Match(ACIParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(495)
		p.LDAPFilter()
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

// IDistinguishedNamesContext is an interface to support dynamic dispatch.
type IDistinguishedNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDistinguishedNamesContext differentiates from other interfaces.
	IsDistinguishedNamesContext()
}

type DistinguishedNamesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDistinguishedNamesContext() *DistinguishedNamesContext {
	var p = new(DistinguishedNamesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_distinguishedNames
	return p
}

func InitEmptyDistinguishedNamesContext(p *DistinguishedNamesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_distinguishedNames
}

func (*DistinguishedNamesContext) IsDistinguishedNamesContext() {}

func NewDistinguishedNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistinguishedNamesContext {
	var p = new(DistinguishedNamesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_distinguishedNames

	return p
}

func (s *DistinguishedNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *DistinguishedNamesContext) CopyAll(ctx *DistinguishedNamesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DistinguishedNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistinguishedNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DistinguishedNamesListContext struct {
	DistinguishedNamesContext
}

func NewDistinguishedNamesListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DistinguishedNamesListContext {
	var p = new(DistinguishedNamesListContext)

	InitEmptyDistinguishedNamesContext(&p.DistinguishedNamesContext)
	p.parser = parser
	p.CopyAll(ctx.(*DistinguishedNamesContext))

	return p
}

func (s *DistinguishedNamesListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistinguishedNamesListContext) AllDistinguishedName() []IDistinguishedNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDistinguishedNameContext); ok {
			len++
		}
	}

	tst := make([]IDistinguishedNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDistinguishedNameContext); ok {
			tst[i] = t.(IDistinguishedNameContext)
			i++
		}
	}

	return tst
}

func (s *DistinguishedNamesListContext) DistinguishedName(i int) IDistinguishedNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDistinguishedNameContext); ok {
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

	return t.(IDistinguishedNameContext)
}

func (s *DistinguishedNamesListContext) AllDoublePipe() []IDoublePipeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDoublePipeContext); ok {
			len++
		}
	}

	tst := make([]IDoublePipeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDoublePipeContext); ok {
			tst[i] = t.(IDoublePipeContext)
			i++
		}
	}

	return tst
}

func (s *DistinguishedNamesListContext) DoublePipe(i int) IDoublePipeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDoublePipeContext); ok {
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

	return t.(IDoublePipeContext)
}

func (s *DistinguishedNamesListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterDistinguishedNamesList(s)
	}
}

func (s *DistinguishedNamesListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitDistinguishedNamesList(s)
	}
}

func (p *ACIParser) DistinguishedNames() (localctx IDistinguishedNamesContext) {
	localctx = NewDistinguishedNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, ACIParserRULE_distinguishedNames)
	var _la int

	localctx = NewDistinguishedNamesListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(506)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ACIParserLocalLDAPScheme {
		{
			p.SetState(497)
			p.DistinguishedName()
		}
		p.SetState(503)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ACIParserDPIPE {
			{
				p.SetState(498)
				p.DoublePipe()
			}
			{
				p.SetState(499)
				p.DistinguishedName()
			}

			p.SetState(505)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(508)
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

// IDoublePipeContext is an interface to support dynamic dispatch.
type IDoublePipeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDoublePipeContext differentiates from other interfaces.
	IsDoublePipeContext()
}

type DoublePipeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDoublePipeContext() *DoublePipeContext {
	var p = new(DoublePipeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_doublePipe
	return p
}

func InitEmptyDoublePipeContext(p *DoublePipeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_doublePipe
}

func (*DoublePipeContext) IsDoublePipeContext() {}

func NewDoublePipeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DoublePipeContext {
	var p = new(DoublePipeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_doublePipe

	return p
}

func (s *DoublePipeContext) GetParser() antlr.Parser { return s.parser }

func (s *DoublePipeContext) CopyAll(ctx *DoublePipeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DoublePipeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoublePipeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DoublePipeDelimiterContext struct {
	DoublePipeContext
}

func NewDoublePipeDelimiterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DoublePipeDelimiterContext {
	var p = new(DoublePipeDelimiterContext)

	InitEmptyDoublePipeContext(&p.DoublePipeContext)
	p.parser = parser
	p.CopyAll(ctx.(*DoublePipeContext))

	return p
}

func (s *DoublePipeDelimiterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DoublePipeDelimiterContext) DPIPE() antlr.TerminalNode {
	return s.GetToken(ACIParserDPIPE, 0)
}

func (s *DoublePipeDelimiterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterDoublePipeDelimiter(s)
	}
}

func (s *DoublePipeDelimiterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitDoublePipeDelimiter(s)
	}
}

func (p *ACIParser) DoublePipe() (localctx IDoublePipeContext) {
	localctx = NewDoublePipeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, ACIParserRULE_doublePipe)
	localctx = NewDoublePipeDelimiterContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(510)
		p.Match(ACIParserDPIPE)
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

// ILDAPURIAndBindTypeContext is an interface to support dynamic dispatch.
type ILDAPURIAndBindTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLDAPURIAndBindTypeContext differentiates from other interfaces.
	IsLDAPURIAndBindTypeContext()
}

type LDAPURIAndBindTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLDAPURIAndBindTypeContext() *LDAPURIAndBindTypeContext {
	var p = new(LDAPURIAndBindTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lDAPURIAndBindType
	return p
}

func InitEmptyLDAPURIAndBindTypeContext(p *LDAPURIAndBindTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lDAPURIAndBindType
}

func (*LDAPURIAndBindTypeContext) IsLDAPURIAndBindTypeContext() {}

func NewLDAPURIAndBindTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LDAPURIAndBindTypeContext {
	var p = new(LDAPURIAndBindTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_lDAPURIAndBindType

	return p
}

func (s *LDAPURIAndBindTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *LDAPURIAndBindTypeContext) CopyAll(ctx *LDAPURIAndBindTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LDAPURIAndBindTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LDAPURIAndBindTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UriAndBindTypeContext struct {
	LDAPURIAndBindTypeContext
}

func NewUriAndBindTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UriAndBindTypeContext {
	var p = new(UriAndBindTypeContext)

	InitEmptyLDAPURIAndBindTypeContext(&p.LDAPURIAndBindTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*LDAPURIAndBindTypeContext))

	return p
}

func (s *UriAndBindTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UriAndBindTypeContext) DistinguishedName() IDistinguishedNameContext {
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

func (s *UriAndBindTypeContext) QMARK() antlr.TerminalNode {
	return s.GetToken(ACIParserQMARK, 0)
}

func (s *UriAndBindTypeContext) AttributeBindTypeOrValue() IAttributeBindTypeOrValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeBindTypeOrValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeBindTypeOrValueContext)
}

func (s *UriAndBindTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterUriAndBindType(s)
	}
}

func (s *UriAndBindTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitUriAndBindType(s)
	}
}

func (p *ACIParser) LDAPURIAndBindType() (localctx ILDAPURIAndBindTypeContext) {
	localctx = NewLDAPURIAndBindTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, ACIParserRULE_lDAPURIAndBindType)
	localctx = NewUriAndBindTypeContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(512)
		p.DistinguishedName()
	}
	{
		p.SetState(513)
		p.Match(ACIParserQMARK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(514)
		p.AttributeBindTypeOrValue()
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

// ILDAPURIContext is an interface to support dynamic dispatch.
type ILDAPURIContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLDAPURIContext differentiates from other interfaces.
	IsLDAPURIContext()
}

type LDAPURIContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLDAPURIContext() *LDAPURIContext {
	var p = new(LDAPURIContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lDAPURI
	return p
}

func InitEmptyLDAPURIContext(p *LDAPURIContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lDAPURI
}

func (*LDAPURIContext) IsLDAPURIContext() {}

func NewLDAPURIContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LDAPURIContext {
	var p = new(LDAPURIContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_lDAPURI

	return p
}

func (s *LDAPURIContext) GetParser() antlr.Parser { return s.parser }

func (s *LDAPURIContext) CopyAll(ctx *LDAPURIContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LDAPURIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LDAPURIContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FullyQualifiedLDAPURIContext struct {
	LDAPURIContext
}

func NewFullyQualifiedLDAPURIContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FullyQualifiedLDAPURIContext {
	var p = new(FullyQualifiedLDAPURIContext)

	InitEmptyLDAPURIContext(&p.LDAPURIContext)
	p.parser = parser
	p.CopyAll(ctx.(*LDAPURIContext))

	return p
}

func (s *FullyQualifiedLDAPURIContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullyQualifiedLDAPURIContext) DistinguishedName() IDistinguishedNameContext {
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

func (s *FullyQualifiedLDAPURIContext) URIAttributeList() IURIAttributeListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IURIAttributeListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IURIAttributeListContext)
}

func (s *FullyQualifiedLDAPURIContext) URISearchScopes() IURISearchScopesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IURISearchScopesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IURISearchScopesContext)
}

func (s *FullyQualifiedLDAPURIContext) URISearchFilter() IURISearchFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IURISearchFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IURISearchFilterContext)
}

func (s *FullyQualifiedLDAPURIContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterFullyQualifiedLDAPURI(s)
	}
}

func (s *FullyQualifiedLDAPURIContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitFullyQualifiedLDAPURI(s)
	}
}

func (p *ACIParser) LDAPURI() (localctx ILDAPURIContext) {
	localctx = NewLDAPURIContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, ACIParserRULE_lDAPURI)
	localctx = NewFullyQualifiedLDAPURIContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(516)
		p.DistinguishedName()
	}
	{
		p.SetState(517)
		p.URIAttributeList()
	}
	{
		p.SetState(518)
		p.URISearchScopes()
	}
	{
		p.SetState(519)
		p.URISearchFilter()
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

// IURISearchFilterContext is an interface to support dynamic dispatch.
type IURISearchFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsURISearchFilterContext differentiates from other interfaces.
	IsURISearchFilterContext()
}

type URISearchFilterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyURISearchFilterContext() *URISearchFilterContext {
	var p = new(URISearchFilterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_uRISearchFilter
	return p
}

func InitEmptyURISearchFilterContext(p *URISearchFilterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_uRISearchFilter
}

func (*URISearchFilterContext) IsURISearchFilterContext() {}

func NewURISearchFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *URISearchFilterContext {
	var p = new(URISearchFilterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_uRISearchFilter

	return p
}

func (s *URISearchFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *URISearchFilterContext) CopyAll(ctx *URISearchFilterContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *URISearchFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *URISearchFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UriSearchFilterContext struct {
	URISearchFilterContext
}

func NewUriSearchFilterContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UriSearchFilterContext {
	var p = new(UriSearchFilterContext)

	InitEmptyURISearchFilterContext(&p.URISearchFilterContext)
	p.parser = parser
	p.CopyAll(ctx.(*URISearchFilterContext))

	return p
}

func (s *UriSearchFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UriSearchFilterContext) QMARK() antlr.TerminalNode {
	return s.GetToken(ACIParserQMARK, 0)
}

func (s *UriSearchFilterContext) LDAPFilter() ILDAPFilterContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILDAPFilterContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILDAPFilterContext)
}

func (s *UriSearchFilterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterUriSearchFilter(s)
	}
}

func (s *UriSearchFilterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitUriSearchFilter(s)
	}
}

func (p *ACIParser) URISearchFilter() (localctx IURISearchFilterContext) {
	localctx = NewURISearchFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, ACIParserRULE_uRISearchFilter)
	localctx = NewUriSearchFilterContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(521)
		p.Match(ACIParserQMARK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(522)
		p.LDAPFilter()
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

// IURISearchScopesContext is an interface to support dynamic dispatch.
type IURISearchScopesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsURISearchScopesContext differentiates from other interfaces.
	IsURISearchScopesContext()
}

type URISearchScopesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyURISearchScopesContext() *URISearchScopesContext {
	var p = new(URISearchScopesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_uRISearchScopes
	return p
}

func InitEmptyURISearchScopesContext(p *URISearchScopesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_uRISearchScopes
}

func (*URISearchScopesContext) IsURISearchScopesContext() {}

func NewURISearchScopesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *URISearchScopesContext {
	var p = new(URISearchScopesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_uRISearchScopes

	return p
}

func (s *URISearchScopesContext) GetParser() antlr.Parser { return s.parser }

func (s *URISearchScopesContext) CopyAll(ctx *URISearchScopesContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *URISearchScopesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *URISearchScopesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UriSearchScopesContext struct {
	URISearchScopesContext
}

func NewUriSearchScopesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UriSearchScopesContext {
	var p = new(UriSearchScopesContext)

	InitEmptyURISearchScopesContext(&p.URISearchScopesContext)
	p.parser = parser
	p.CopyAll(ctx.(*URISearchScopesContext))

	return p
}

func (s *UriSearchScopesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UriSearchScopesContext) QMARK() antlr.TerminalNode {
	return s.GetToken(ACIParserQMARK, 0)
}

func (s *UriSearchScopesContext) LDAPSearchScopes() antlr.TerminalNode {
	return s.GetToken(ACIParserLDAPSearchScopes, 0)
}

func (s *UriSearchScopesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterUriSearchScopes(s)
	}
}

func (s *UriSearchScopesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitUriSearchScopes(s)
	}
}

func (p *ACIParser) URISearchScopes() (localctx IURISearchScopesContext) {
	localctx = NewURISearchScopesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, ACIParserRULE_uRISearchScopes)
	var _la int

	localctx = NewUriSearchScopesContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(524)
		p.Match(ACIParserQMARK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(526)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ACIParserLDAPSearchScopes {
		{
			p.SetState(525)
			p.Match(ACIParserLDAPSearchScopes)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IURIAttributeListContext is an interface to support dynamic dispatch.
type IURIAttributeListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsURIAttributeListContext differentiates from other interfaces.
	IsURIAttributeListContext()
}

type URIAttributeListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyURIAttributeListContext() *URIAttributeListContext {
	var p = new(URIAttributeListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_uRIAttributeList
	return p
}

func InitEmptyURIAttributeListContext(p *URIAttributeListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_uRIAttributeList
}

func (*URIAttributeListContext) IsURIAttributeListContext() {}

func NewURIAttributeListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *URIAttributeListContext {
	var p = new(URIAttributeListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_uRIAttributeList

	return p
}

func (s *URIAttributeListContext) GetParser() antlr.Parser { return s.parser }

func (s *URIAttributeListContext) CopyAll(ctx *URIAttributeListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *URIAttributeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *URIAttributeListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UriAttributeListContext struct {
	URIAttributeListContext
}

func NewUriAttributeListContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UriAttributeListContext {
	var p = new(UriAttributeListContext)

	InitEmptyURIAttributeListContext(&p.URIAttributeListContext)
	p.parser = parser
	p.CopyAll(ctx.(*URIAttributeListContext))

	return p
}

func (s *UriAttributeListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UriAttributeListContext) QMARK() antlr.TerminalNode {
	return s.GetToken(ACIParserQMARK, 0)
}

func (s *UriAttributeListContext) AllAttributeType() []IAttributeTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IAttributeTypeContext); ok {
			len++
		}
	}

	tst := make([]IAttributeTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IAttributeTypeContext); ok {
			tst[i] = t.(IAttributeTypeContext)
			i++
		}
	}

	return tst
}

func (s *UriAttributeListContext) AttributeType(i int) IAttributeTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeContext); ok {
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

	return t.(IAttributeTypeContext)
}

func (s *UriAttributeListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ACIParserCOMMA)
}

func (s *UriAttributeListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserCOMMA, i)
}

func (s *UriAttributeListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterUriAttributeList(s)
	}
}

func (s *UriAttributeListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitUriAttributeList(s)
	}
}

func (p *ACIParser) URIAttributeList() (localctx IURIAttributeListContext) {
	localctx = NewURIAttributeListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, ACIParserRULE_uRIAttributeList)
	var _la int

	localctx = NewUriAttributeListContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(528)
		p.Match(ACIParserQMARK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(537)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ACIParserID {
		{
			p.SetState(529)
			p.AttributeType()
		}
		p.SetState(534)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == ACIParserCOMMA {
			{
				p.SetState(530)
				p.Match(ACIParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(531)
				p.AttributeType()
			}

			p.SetState(536)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = ACIParserRULE_distinguishedName
	return p
}

func InitEmptyDistinguishedNameContext(p *DistinguishedNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_distinguishedName
}

func (*DistinguishedNameContext) IsDistinguishedNameContext() {}

func NewDistinguishedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DistinguishedNameContext {
	var p = new(DistinguishedNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_distinguishedName

	return p
}

func (s *DistinguishedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DistinguishedNameContext) CopyAll(ctx *DistinguishedNameContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DistinguishedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistinguishedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DistinguishedNameValueContext struct {
	DistinguishedNameContext
}

func NewDistinguishedNameValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DistinguishedNameValueContext {
	var p = new(DistinguishedNameValueContext)

	InitEmptyDistinguishedNameContext(&p.DistinguishedNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*DistinguishedNameContext))

	return p
}

func (s *DistinguishedNameValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistinguishedNameValueContext) LocalLDAPScheme() antlr.TerminalNode {
	return s.GetToken(ACIParserLocalLDAPScheme, 0)
}

func (s *DistinguishedNameValueContext) AllRelativeDistinguishedName() []IRelativeDistinguishedNameContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRelativeDistinguishedNameContext); ok {
			len++
		}
	}

	tst := make([]IRelativeDistinguishedNameContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRelativeDistinguishedNameContext); ok {
			tst[i] = t.(IRelativeDistinguishedNameContext)
			i++
		}
	}

	return tst
}

func (s *DistinguishedNameValueContext) RelativeDistinguishedName(i int) IRelativeDistinguishedNameContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelativeDistinguishedNameContext); ok {
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

	return t.(IRelativeDistinguishedNameContext)
}

func (s *DistinguishedNameValueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(ACIParserCOMMA)
}

func (s *DistinguishedNameValueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserCOMMA, i)
}

func (s *DistinguishedNameValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterDistinguishedNameValue(s)
	}
}

func (s *DistinguishedNameValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitDistinguishedNameValue(s)
	}
}

func (p *ACIParser) DistinguishedName() (localctx IDistinguishedNameContext) {
	localctx = NewDistinguishedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, ACIParserRULE_distinguishedName)
	var _la int

	localctx = NewDistinguishedNameValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(539)
		p.Match(ACIParserLocalLDAPScheme)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	{
		p.SetState(540)
		p.RelativeDistinguishedName()
	}
	p.SetState(545)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == ACIParserCOMMA {
		{
			p.SetState(541)
			p.Match(ACIParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(542)
			p.RelativeDistinguishedName()
		}

		p.SetState(547)
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

// IRelativeDistinguishedNameContext is an interface to support dynamic dispatch.
type IRelativeDistinguishedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsRelativeDistinguishedNameContext differentiates from other interfaces.
	IsRelativeDistinguishedNameContext()
}

type RelativeDistinguishedNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelativeDistinguishedNameContext() *RelativeDistinguishedNameContext {
	var p = new(RelativeDistinguishedNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_relativeDistinguishedName
	return p
}

func InitEmptyRelativeDistinguishedNameContext(p *RelativeDistinguishedNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_relativeDistinguishedName
}

func (*RelativeDistinguishedNameContext) IsRelativeDistinguishedNameContext() {}

func NewRelativeDistinguishedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelativeDistinguishedNameContext {
	var p = new(RelativeDistinguishedNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_relativeDistinguishedName

	return p
}

func (s *RelativeDistinguishedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RelativeDistinguishedNameContext) CopyAll(ctx *RelativeDistinguishedNameContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *RelativeDistinguishedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelativeDistinguishedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type RelativeDistinguishedNameMacroContext struct {
	RelativeDistinguishedNameContext
}

func NewRelativeDistinguishedNameMacroContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelativeDistinguishedNameMacroContext {
	var p = new(RelativeDistinguishedNameMacroContext)

	InitEmptyRelativeDistinguishedNameContext(&p.RelativeDistinguishedNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelativeDistinguishedNameContext))

	return p
}

func (s *RelativeDistinguishedNameMacroContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelativeDistinguishedNameMacroContext) MacroValue() antlr.TerminalNode {
	return s.GetToken(ACIParserMacroValue, 0)
}

func (s *RelativeDistinguishedNameMacroContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterRelativeDistinguishedNameMacro(s)
	}
}

func (s *RelativeDistinguishedNameMacroContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitRelativeDistinguishedNameMacro(s)
	}
}

type RelativeDistinguishedNameValueContext struct {
	RelativeDistinguishedNameContext
}

func NewRelativeDistinguishedNameValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelativeDistinguishedNameValueContext {
	var p = new(RelativeDistinguishedNameValueContext)

	InitEmptyRelativeDistinguishedNameContext(&p.RelativeDistinguishedNameContext)
	p.parser = parser
	p.CopyAll(ctx.(*RelativeDistinguishedNameContext))

	return p
}

func (s *RelativeDistinguishedNameValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelativeDistinguishedNameValueContext) AttributeType() IAttributeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeTypeContext)
}

func (s *RelativeDistinguishedNameValueContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *RelativeDistinguishedNameValueContext) AttributeValue() IAttributeValueContext {
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

func (s *RelativeDistinguishedNameValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterRelativeDistinguishedNameValue(s)
	}
}

func (s *RelativeDistinguishedNameValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitRelativeDistinguishedNameValue(s)
	}
}

func (p *ACIParser) RelativeDistinguishedName() (localctx IRelativeDistinguishedNameContext) {
	localctx = NewRelativeDistinguishedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, ACIParserRULE_relativeDistinguishedName)
	p.SetState(553)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserID:
		localctx = NewRelativeDistinguishedNameValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(548)
			p.AttributeType()
		}
		{
			p.SetState(549)
			p.Match(ACIParserEqualTo)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(550)
			p.AttributeValue()
		}

	case ACIParserMacroValue:
		localctx = NewRelativeDistinguishedNameMacroContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(552)
			p.Match(ACIParserMacroValue)
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

// ILDAPFilterContext is an interface to support dynamic dispatch.
type ILDAPFilterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLDAPFilterContext differentiates from other interfaces.
	IsLDAPFilterContext()
}

type LDAPFilterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLDAPFilterContext() *LDAPFilterContext {
	var p = new(LDAPFilterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lDAPFilter
	return p
}

func InitEmptyLDAPFilterContext(p *LDAPFilterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lDAPFilter
}

func (*LDAPFilterContext) IsLDAPFilterContext() {}

func NewLDAPFilterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LDAPFilterContext {
	var p = new(LDAPFilterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_lDAPFilter

	return p
}

func (s *LDAPFilterContext) GetParser() antlr.Parser { return s.parser }

func (s *LDAPFilterContext) CopyAll(ctx *LDAPFilterContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LDAPFilterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LDAPFilterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FilterExpressionsContext struct {
	LDAPFilterContext
}

func NewFilterExpressionsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FilterExpressionsContext {
	var p = new(FilterExpressionsContext)

	InitEmptyLDAPFilterContext(&p.LDAPFilterContext)
	p.parser = parser
	p.CopyAll(ctx.(*LDAPFilterContext))

	return p
}

func (s *FilterExpressionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FilterExpressionsContext) AllLDAPFilterExpr() []ILDAPFilterExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILDAPFilterExprContext); ok {
			len++
		}
	}

	tst := make([]ILDAPFilterExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILDAPFilterExprContext); ok {
			tst[i] = t.(ILDAPFilterExprContext)
			i++
		}
	}

	return tst
}

func (s *FilterExpressionsContext) LDAPFilterExpr(i int) ILDAPFilterExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILDAPFilterExprContext); ok {
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

	return t.(ILDAPFilterExprContext)
}

func (s *FilterExpressionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterFilterExpressions(s)
	}
}

func (s *FilterExpressionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitFilterExpressions(s)
	}
}

type ParentheticalFilterExpressionContext struct {
	LDAPFilterContext
}

func NewParentheticalFilterExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalFilterExpressionContext {
	var p = new(ParentheticalFilterExpressionContext)

	InitEmptyLDAPFilterContext(&p.LDAPFilterContext)
	p.parser = parser
	p.CopyAll(ctx.(*LDAPFilterContext))

	return p
}

func (s *ParentheticalFilterExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalFilterExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, 0)
}

func (s *ParentheticalFilterExpressionContext) LDAPFilterExpr() ILDAPFilterExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILDAPFilterExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILDAPFilterExprContext)
}

func (s *ParentheticalFilterExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, 0)
}

func (s *ParentheticalFilterExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalFilterExpression(s)
	}
}

func (s *ParentheticalFilterExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalFilterExpression(s)
	}
}

func (p *ACIParser) LDAPFilter() (localctx ILDAPFilterContext) {
	localctx = NewLDAPFilterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, ACIParserRULE_lDAPFilter)
	var _la int

	p.SetState(565)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParentheticalFilterExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(555)
			p.Match(ACIParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(556)
			p.LDAPFilterExpr()
		}
		{
			p.SetState(557)
			p.Match(ACIParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewFilterExpressionsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(562)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-3)) & ^0x3f) == 0 && ((int64(1)<<(_la-3))&5764607523034234881) != 0 {
			{
				p.SetState(559)
				p.LDAPFilterExpr()
			}

			p.SetState(564)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
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

// ILDAPFilterExprContext is an interface to support dynamic dispatch.
type ILDAPFilterExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsLDAPFilterExprContext differentiates from other interfaces.
	IsLDAPFilterExprContext()
}

type LDAPFilterExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLDAPFilterExprContext() *LDAPFilterExprContext {
	var p = new(LDAPFilterExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lDAPFilterExpr
	return p
}

func InitEmptyLDAPFilterExprContext(p *LDAPFilterExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_lDAPFilterExpr
}

func (*LDAPFilterExprContext) IsLDAPFilterExprContext() {}

func NewLDAPFilterExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LDAPFilterExprContext {
	var p = new(LDAPFilterExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_lDAPFilterExpr

	return p
}

func (s *LDAPFilterExprContext) GetParser() antlr.Parser { return s.parser }

func (s *LDAPFilterExprContext) CopyAll(ctx *LDAPFilterExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *LDAPFilterExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LDAPFilterExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParentheticalFilterExpressionWithOptionalBooleanOperatorContext struct {
	LDAPFilterExprContext
}

func NewParentheticalFilterExpressionWithOptionalBooleanOperatorContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext {
	var p = new(ParentheticalFilterExpressionWithOptionalBooleanOperatorContext)

	InitEmptyLDAPFilterExprContext(&p.LDAPFilterExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*LDAPFilterExprContext))

	return p
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) AllLPAREN() []antlr.TerminalNode {
	return s.GetTokens(ACIParserLPAREN)
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) LPAREN(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserLPAREN, i)
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) AllLDAPFilterExpr() []ILDAPFilterExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ILDAPFilterExprContext); ok {
			len++
		}
	}

	tst := make([]ILDAPFilterExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ILDAPFilterExprContext); ok {
			tst[i] = t.(ILDAPFilterExprContext)
			i++
		}
	}

	return tst
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) LDAPFilterExpr(i int) ILDAPFilterExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILDAPFilterExprContext); ok {
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

	return t.(ILDAPFilterExprContext)
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) AllRPAREN() []antlr.TerminalNode {
	return s.GetTokens(ACIParserRPAREN)
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) RPAREN(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserRPAREN, i)
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) AllAMP() []antlr.TerminalNode {
	return s.GetTokens(ACIParserAMP)
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) AMP(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserAMP, i)
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) AllPIPE() []antlr.TerminalNode {
	return s.GetTokens(ACIParserPIPE)
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) PIPE(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserPIPE, i)
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) AllBANG() []antlr.TerminalNode {
	return s.GetTokens(ACIParserBANG)
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) BANG(i int) antlr.TerminalNode {
	return s.GetToken(ACIParserBANG, i)
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterParentheticalFilterExpressionWithOptionalBooleanOperator(s)
	}
}

func (s *ParentheticalFilterExpressionWithOptionalBooleanOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitParentheticalFilterExpressionWithOptionalBooleanOperator(s)
	}
}

type NegatedFilterExpressionContext struct {
	LDAPFilterExprContext
}

func NewNegatedFilterExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NegatedFilterExpressionContext {
	var p = new(NegatedFilterExpressionContext)

	InitEmptyLDAPFilterExprContext(&p.LDAPFilterExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*LDAPFilterExprContext))

	return p
}

func (s *NegatedFilterExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NegatedFilterExpressionContext) BANG() antlr.TerminalNode {
	return s.GetToken(ACIParserBANG, 0)
}

func (s *NegatedFilterExpressionContext) LDAPFilterExpr() ILDAPFilterExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILDAPFilterExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILDAPFilterExprContext)
}

func (s *NegatedFilterExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterNegatedFilterExpression(s)
	}
}

func (s *NegatedFilterExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitNegatedFilterExpression(s)
	}
}

type AttributeValueAssertionExpressionContext struct {
	LDAPFilterExprContext
}

func NewAttributeValueAssertionExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttributeValueAssertionExpressionContext {
	var p = new(AttributeValueAssertionExpressionContext)

	InitEmptyLDAPFilterExprContext(&p.LDAPFilterExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*LDAPFilterExprContext))

	return p
}

func (s *AttributeValueAssertionExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeValueAssertionExpressionContext) AttributeValueAssertion() IAttributeValueAssertionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeValueAssertionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeValueAssertionContext)
}

func (s *AttributeValueAssertionExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAttributeValueAssertionExpression(s)
	}
}

func (s *AttributeValueAssertionExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAttributeValueAssertionExpression(s)
	}
}

func (p *ACIParser) LDAPFilterExpr() (localctx ILDAPFilterExprContext) {
	localctx = NewLDAPFilterExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, ACIParserRULE_lDAPFilterExpr)
	var _la int

	var _alt int

	p.SetState(581)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ACIParserLPAREN:
		localctx = NewParentheticalFilterExpressionWithOptionalBooleanOperatorContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(574)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = 1 + 1
		for ok := true; ok; ok = _alt != 1 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1 + 1:
				{
					p.SetState(567)
					p.Match(ACIParserLPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(569)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(568)
						_la = p.GetTokenStream().LA(1)

						if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&-2305843009213693952) != 0) {
							p.GetErrorHandler().RecoverInline(p)
						} else {
							p.GetErrorHandler().ReportMatch(p)
							p.Consume()
						}
					}

				} else if p.HasError() { // JIM
					goto errorExit
				}
				{
					p.SetState(571)
					p.LDAPFilterExpr()
				}
				{
					p.SetState(572)
					p.Match(ACIParserRPAREN)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(576)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 43, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case ACIParserBANG:
		localctx = NewNegatedFilterExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(578)
			p.Match(ACIParserBANG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(579)
			p.LDAPFilterExpr()
		}

	case ACIParserID:
		localctx = NewAttributeValueAssertionExpressionContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(580)
			p.AttributeValueAssertion()
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

// IAttributeValueAssertionContext is an interface to support dynamic dispatch.
type IAttributeValueAssertionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
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
	p.RuleIndex = ACIParserRULE_attributeValueAssertion
	return p
}

func InitEmptyAttributeValueAssertionContext(p *AttributeValueAssertionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeValueAssertion
}

func (*AttributeValueAssertionContext) IsAttributeValueAssertionContext() {}

func NewAttributeValueAssertionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeValueAssertionContext {
	var p = new(AttributeValueAssertionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_attributeValueAssertion

	return p
}

func (s *AttributeValueAssertionContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeValueAssertionContext) CopyAll(ctx *AttributeValueAssertionContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AttributeValueAssertionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeValueAssertionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AttributeValueAssertionStatementContext struct {
	AttributeValueAssertionContext
}

func NewAttributeValueAssertionStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttributeValueAssertionStatementContext {
	var p = new(AttributeValueAssertionStatementContext)

	InitEmptyAttributeValueAssertionContext(&p.AttributeValueAssertionContext)
	p.parser = parser
	p.CopyAll(ctx.(*AttributeValueAssertionContext))

	return p
}

func (s *AttributeValueAssertionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeValueAssertionStatementContext) AttributeType() IAttributeTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeTypeContext)
}

func (s *AttributeValueAssertionStatementContext) AttributeOperators() IAttributeOperatorsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAttributeOperatorsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAttributeOperatorsContext)
}

func (s *AttributeValueAssertionStatementContext) AttributeValue() IAttributeValueContext {
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

func (s *AttributeValueAssertionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAttributeValueAssertionStatement(s)
	}
}

func (s *AttributeValueAssertionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAttributeValueAssertionStatement(s)
	}
}

func (p *ACIParser) AttributeValueAssertion() (localctx IAttributeValueAssertionContext) {
	localctx = NewAttributeValueAssertionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, ACIParserRULE_attributeValueAssertion)
	localctx = NewAttributeValueAssertionStatementContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(583)
		p.AttributeType()
	}
	{
		p.SetState(584)
		p.AttributeOperators()
	}
	{
		p.SetState(585)
		p.AttributeValue()
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

// IAttributeTypeContext is an interface to support dynamic dispatch.
type IAttributeTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAttributeTypeContext differentiates from other interfaces.
	IsAttributeTypeContext()
}

type AttributeTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeTypeContext() *AttributeTypeContext {
	var p = new(AttributeTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeType
	return p
}

func InitEmptyAttributeTypeContext(p *AttributeTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeType
}

func (*AttributeTypeContext) IsAttributeTypeContext() {}

func NewAttributeTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeTypeContext {
	var p = new(AttributeTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_attributeType

	return p
}

func (s *AttributeTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeTypeContext) CopyAll(ctx *AttributeTypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AttributeTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AttributeTypeIdentifierContext struct {
	AttributeTypeContext
}

func NewAttributeTypeIdentifierContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttributeTypeIdentifierContext {
	var p = new(AttributeTypeIdentifierContext)

	InitEmptyAttributeTypeContext(&p.AttributeTypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*AttributeTypeContext))

	return p
}

func (s *AttributeTypeIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeTypeIdentifierContext) ID() antlr.TerminalNode {
	return s.GetToken(ACIParserID, 0)
}

func (s *AttributeTypeIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAttributeTypeIdentifier(s)
	}
}

func (s *AttributeTypeIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAttributeTypeIdentifier(s)
	}
}

func (p *ACIParser) AttributeType() (localctx IAttributeTypeContext) {
	localctx = NewAttributeTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, ACIParserRULE_attributeType)
	localctx = NewAttributeTypeIdentifierContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(587)
		p.Match(ACIParserID)
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
	p.RuleIndex = ACIParserRULE_attributeValue
	return p
}

func InitEmptyAttributeValueContext(p *AttributeValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeValue
}

func (*AttributeValueContext) IsAttributeValueContext() {}

func NewAttributeValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeValueContext {
	var p = new(AttributeValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_attributeValue

	return p
}

func (s *AttributeValueContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeValueContext) CopyAll(ctx *AttributeValueContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AttributeValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AttributeAssertionValueContext struct {
	AttributeValueContext
}

func NewAttributeAssertionValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttributeAssertionValueContext {
	var p = new(AttributeAssertionValueContext)

	InitEmptyAttributeValueContext(&p.AttributeValueContext)
	p.parser = parser
	p.CopyAll(ctx.(*AttributeValueContext))

	return p
}

func (s *AttributeAssertionValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeAssertionValueContext) ANY() antlr.TerminalNode {
	return s.GetToken(ACIParserANY, 0)
}

func (s *AttributeAssertionValueContext) ID() antlr.TerminalNode {
	return s.GetToken(ACIParserID, 0)
}

func (s *AttributeAssertionValueContext) INT() antlr.TerminalNode {
	return s.GetToken(ACIParserINT, 0)
}

func (s *AttributeAssertionValueContext) WildcardString() antlr.TerminalNode {
	return s.GetToken(ACIParserWildcardString, 0)
}

func (s *AttributeAssertionValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAttributeAssertionValue(s)
	}
}

func (s *AttributeAssertionValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAttributeAssertionValue(s)
	}
}

func (p *ACIParser) AttributeValue() (localctx IAttributeValueContext) {
	localctx = NewAttributeValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, ACIParserRULE_attributeValue)
	var _la int

	localctx = NewAttributeAssertionValueContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(589)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ACIParserINT || _la == ACIParserID || _la == ACIParserWildcardString || _la == ACIParserANY) {
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

// IAttributeOperatorsContext is an interface to support dynamic dispatch.
type IAttributeOperatorsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EqualTo() antlr.TerminalNode
	ExtensibleRuleAttrMatch() antlr.TerminalNode
	ExtensibleRuleMatch() antlr.TerminalNode
	GreaterThanOrEqual() antlr.TerminalNode
	LessThanOrEqual() antlr.TerminalNode
	ExtensibleRuleDNOIDMatch() antlr.TerminalNode
	ExtensibleRuleDNMatch() antlr.TerminalNode
	ApproximateMatch() antlr.TerminalNode

	// IsAttributeOperatorsContext differentiates from other interfaces.
	IsAttributeOperatorsContext()
}

type AttributeOperatorsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAttributeOperatorsContext() *AttributeOperatorsContext {
	var p = new(AttributeOperatorsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeOperators
	return p
}

func InitEmptyAttributeOperatorsContext(p *AttributeOperatorsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ACIParserRULE_attributeOperators
}

func (*AttributeOperatorsContext) IsAttributeOperatorsContext() {}

func NewAttributeOperatorsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeOperatorsContext {
	var p = new(AttributeOperatorsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ACIParserRULE_attributeOperators

	return p
}

func (s *AttributeOperatorsContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeOperatorsContext) EqualTo() antlr.TerminalNode {
	return s.GetToken(ACIParserEqualTo, 0)
}

func (s *AttributeOperatorsContext) ExtensibleRuleAttrMatch() antlr.TerminalNode {
	return s.GetToken(ACIParserExtensibleRuleAttrMatch, 0)
}

func (s *AttributeOperatorsContext) ExtensibleRuleMatch() antlr.TerminalNode {
	return s.GetToken(ACIParserExtensibleRuleMatch, 0)
}

func (s *AttributeOperatorsContext) GreaterThanOrEqual() antlr.TerminalNode {
	return s.GetToken(ACIParserGreaterThanOrEqual, 0)
}

func (s *AttributeOperatorsContext) LessThanOrEqual() antlr.TerminalNode {
	return s.GetToken(ACIParserLessThanOrEqual, 0)
}

func (s *AttributeOperatorsContext) ExtensibleRuleDNOIDMatch() antlr.TerminalNode {
	return s.GetToken(ACIParserExtensibleRuleDNOIDMatch, 0)
}

func (s *AttributeOperatorsContext) ExtensibleRuleDNMatch() antlr.TerminalNode {
	return s.GetToken(ACIParserExtensibleRuleDNMatch, 0)
}

func (s *AttributeOperatorsContext) ApproximateMatch() antlr.TerminalNode {
	return s.GetToken(ACIParserApproximateMatch, 0)
}

func (s *AttributeOperatorsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeOperatorsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeOperatorsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.EnterAttributeOperators(s)
	}
}

func (s *AttributeOperatorsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ACIParserListener); ok {
		listenerT.ExitAttributeOperators(s)
	}
}

func (p *ACIParser) AttributeOperators() (localctx IAttributeOperatorsContext) {
	localctx = NewAttributeOperatorsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, ACIParserRULE_attributeOperators)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(591)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36028797018971104) != 0) {
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
