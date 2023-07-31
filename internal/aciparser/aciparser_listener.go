// Code generated from ACIParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package aciparser // ACIParser
import "github.com/antlr4-go/antlr/v4"

// ACIParserListener is a complete listener for a parse tree produced by ACIParser.
type ACIParserListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterInstruction is called when entering the instruction production.
	EnterInstruction(c *InstructionContext)

	// EnterPermissionBindRule is called when entering the permissionBindRule production.
	EnterPermissionBindRule(c *PermissionBindRuleContext)

	// EnterPermission is called when entering the permission production.
	EnterPermission(c *PermissionContext)

	// EnterPrivilege is called when entering the privilege production.
	EnterPrivilege(c *PrivilegeContext)

	// EnterTargetRule is called when entering the targetRule production.
	EnterTargetRule(c *TargetRuleContext)

	// EnterTargetOperator is called when entering the targetOperator production.
	EnterTargetOperator(c *TargetOperatorContext)

	// EnterTargetRules is called when entering the targetRules production.
	EnterTargetRules(c *TargetRulesContext)

	// EnterBindRule is called when entering the bindRule production.
	EnterBindRule(c *BindRuleContext)

	// EnterWordAnd is called when entering the wordAnd production.
	EnterWordAnd(c *WordAndContext)

	// EnterWordOr is called when entering the wordOr production.
	EnterWordOr(c *WordOrContext)

	// EnterWordNot is called when entering the wordNot production.
	EnterWordNot(c *WordNotContext)

	// EnterBindRules is called when entering the bindRules production.
	EnterBindRules(c *BindRulesContext)

	// EnterQuotedValue is called when entering the quotedValue production.
	EnterQuotedValue(c *QuotedValueContext)

	// EnterBindOperator is called when entering the bindOperator production.
	EnterBindOperator(c *BindOperatorContext)

	// EnterLessThan is called when entering the lessThan production.
	EnterLessThan(c *LessThanContext)

	// EnterLessThanOrEqual is called when entering the lessThanOrEqual production.
	EnterLessThanOrEqual(c *LessThanOrEqualContext)

	// EnterGreaterThan is called when entering the greaterThan production.
	EnterGreaterThan(c *GreaterThanContext)

	// EnterGreaterThanOrEqual is called when entering the greaterThanOrEqual production.
	EnterGreaterThanOrEqual(c *GreaterThanOrEqualContext)

	// EnterEqualTo is called when entering the equalTo production.
	EnterEqualTo(c *EqualToContext)

	// EnterNotEqualTo is called when entering the notEqualTo production.
	EnterNotEqualTo(c *NotEqualToContext)

	// EnterBindKeyword is called when entering the bindKeyword production.
	EnterBindKeyword(c *BindKeywordContext)

	// EnterTargetKeyword is called when entering the targetKeyword production.
	EnterTargetKeyword(c *TargetKeywordContext)

	// EnterOpeningParenthesis is called when entering the openingParenthesis production.
	EnterOpeningParenthesis(c *OpeningParenthesisContext)

	// EnterClosingParenthesis is called when entering the closingParenthesis production.
	EnterClosingParenthesis(c *ClosingParenthesisContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitInstruction is called when exiting the instruction production.
	ExitInstruction(c *InstructionContext)

	// ExitPermissionBindRule is called when exiting the permissionBindRule production.
	ExitPermissionBindRule(c *PermissionBindRuleContext)

	// ExitPermission is called when exiting the permission production.
	ExitPermission(c *PermissionContext)

	// ExitPrivilege is called when exiting the privilege production.
	ExitPrivilege(c *PrivilegeContext)

	// ExitTargetRule is called when exiting the targetRule production.
	ExitTargetRule(c *TargetRuleContext)

	// ExitTargetOperator is called when exiting the targetOperator production.
	ExitTargetOperator(c *TargetOperatorContext)

	// ExitTargetRules is called when exiting the targetRules production.
	ExitTargetRules(c *TargetRulesContext)

	// ExitBindRule is called when exiting the bindRule production.
	ExitBindRule(c *BindRuleContext)

	// ExitWordAnd is called when exiting the wordAnd production.
	ExitWordAnd(c *WordAndContext)

	// ExitWordOr is called when exiting the wordOr production.
	ExitWordOr(c *WordOrContext)

	// ExitWordNot is called when exiting the wordNot production.
	ExitWordNot(c *WordNotContext)

	// ExitBindRules is called when exiting the bindRules production.
	ExitBindRules(c *BindRulesContext)

	// ExitQuotedValue is called when exiting the quotedValue production.
	ExitQuotedValue(c *QuotedValueContext)

	// ExitBindOperator is called when exiting the bindOperator production.
	ExitBindOperator(c *BindOperatorContext)

	// ExitLessThan is called when exiting the lessThan production.
	ExitLessThan(c *LessThanContext)

	// ExitLessThanOrEqual is called when exiting the lessThanOrEqual production.
	ExitLessThanOrEqual(c *LessThanOrEqualContext)

	// ExitGreaterThan is called when exiting the greaterThan production.
	ExitGreaterThan(c *GreaterThanContext)

	// ExitGreaterThanOrEqual is called when exiting the greaterThanOrEqual production.
	ExitGreaterThanOrEqual(c *GreaterThanOrEqualContext)

	// ExitEqualTo is called when exiting the equalTo production.
	ExitEqualTo(c *EqualToContext)

	// ExitNotEqualTo is called when exiting the notEqualTo production.
	ExitNotEqualTo(c *NotEqualToContext)

	// ExitBindKeyword is called when exiting the bindKeyword production.
	ExitBindKeyword(c *BindKeywordContext)

	// ExitTargetKeyword is called when exiting the targetKeyword production.
	ExitTargetKeyword(c *TargetKeywordContext)

	// ExitOpeningParenthesis is called when exiting the openingParenthesis production.
	ExitOpeningParenthesis(c *OpeningParenthesisContext)

	// ExitClosingParenthesis is called when exiting the closingParenthesis production.
	ExitClosingParenthesis(c *ClosingParenthesisContext)
}
