// Code generated from ACIParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package aciparser // ACIParser
import "github.com/antlr4-go/antlr/v4"

// BaseACIParserListener is a complete listener for a parse tree produced by ACIParser.
type BaseACIParserListener struct{}

var _ ACIParserListener = &BaseACIParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseACIParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseACIParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseACIParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseACIParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseACIParserListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseACIParserListener) ExitParse(ctx *ParseContext) {}

// EnterInstruction is called when production instruction is entered.
func (s *BaseACIParserListener) EnterInstruction(ctx *InstructionContext) {}

// ExitInstruction is called when production instruction is exited.
func (s *BaseACIParserListener) ExitInstruction(ctx *InstructionContext) {}

// EnterPermissionBindRule is called when production permissionBindRule is entered.
func (s *BaseACIParserListener) EnterPermissionBindRule(ctx *PermissionBindRuleContext) {}

// ExitPermissionBindRule is called when production permissionBindRule is exited.
func (s *BaseACIParserListener) ExitPermissionBindRule(ctx *PermissionBindRuleContext) {}

// EnterPermission is called when production permission is entered.
func (s *BaseACIParserListener) EnterPermission(ctx *PermissionContext) {}

// ExitPermission is called when production permission is exited.
func (s *BaseACIParserListener) ExitPermission(ctx *PermissionContext) {}

// EnterPrivilege is called when production privilege is entered.
func (s *BaseACIParserListener) EnterPrivilege(ctx *PrivilegeContext) {}

// ExitPrivilege is called when production privilege is exited.
func (s *BaseACIParserListener) ExitPrivilege(ctx *PrivilegeContext) {}

// EnterTargetRule is called when production targetRule is entered.
func (s *BaseACIParserListener) EnterTargetRule(ctx *TargetRuleContext) {}

// ExitTargetRule is called when production targetRule is exited.
func (s *BaseACIParserListener) ExitTargetRule(ctx *TargetRuleContext) {}

// EnterTargetOperator is called when production targetOperator is entered.
func (s *BaseACIParserListener) EnterTargetOperator(ctx *TargetOperatorContext) {}

// ExitTargetOperator is called when production targetOperator is exited.
func (s *BaseACIParserListener) ExitTargetOperator(ctx *TargetOperatorContext) {}

// EnterTargetRules is called when production targetRules is entered.
func (s *BaseACIParserListener) EnterTargetRules(ctx *TargetRulesContext) {}

// ExitTargetRules is called when production targetRules is exited.
func (s *BaseACIParserListener) ExitTargetRules(ctx *TargetRulesContext) {}

// EnterBindRule is called when production bindRule is entered.
func (s *BaseACIParserListener) EnterBindRule(ctx *BindRuleContext) {}

// ExitBindRule is called when production bindRule is exited.
func (s *BaseACIParserListener) ExitBindRule(ctx *BindRuleContext) {}

// EnterWordAnd is called when production wordAnd is entered.
func (s *BaseACIParserListener) EnterWordAnd(ctx *WordAndContext) {}

// ExitWordAnd is called when production wordAnd is exited.
func (s *BaseACIParserListener) ExitWordAnd(ctx *WordAndContext) {}

// EnterWordOr is called when production wordOr is entered.
func (s *BaseACIParserListener) EnterWordOr(ctx *WordOrContext) {}

// ExitWordOr is called when production wordOr is exited.
func (s *BaseACIParserListener) ExitWordOr(ctx *WordOrContext) {}

// EnterWordNot is called when production wordNot is entered.
func (s *BaseACIParserListener) EnterWordNot(ctx *WordNotContext) {}

// ExitWordNot is called when production wordNot is exited.
func (s *BaseACIParserListener) ExitWordNot(ctx *WordNotContext) {}

// EnterBindRules is called when production bindRules is entered.
func (s *BaseACIParserListener) EnterBindRules(ctx *BindRulesContext) {}

// ExitBindRules is called when production bindRules is exited.
func (s *BaseACIParserListener) ExitBindRules(ctx *BindRulesContext) {}

// EnterQuotedValue is called when production quotedValue is entered.
func (s *BaseACIParserListener) EnterQuotedValue(ctx *QuotedValueContext) {}

// ExitQuotedValue is called when production quotedValue is exited.
func (s *BaseACIParserListener) ExitQuotedValue(ctx *QuotedValueContext) {}

// EnterBindOperator is called when production bindOperator is entered.
func (s *BaseACIParserListener) EnterBindOperator(ctx *BindOperatorContext) {}

// ExitBindOperator is called when production bindOperator is exited.
func (s *BaseACIParserListener) ExitBindOperator(ctx *BindOperatorContext) {}

// EnterLessThan is called when production lessThan is entered.
func (s *BaseACIParserListener) EnterLessThan(ctx *LessThanContext) {}

// ExitLessThan is called when production lessThan is exited.
func (s *BaseACIParserListener) ExitLessThan(ctx *LessThanContext) {}

// EnterLessThanOrEqual is called when production lessThanOrEqual is entered.
func (s *BaseACIParserListener) EnterLessThanOrEqual(ctx *LessThanOrEqualContext) {}

// ExitLessThanOrEqual is called when production lessThanOrEqual is exited.
func (s *BaseACIParserListener) ExitLessThanOrEqual(ctx *LessThanOrEqualContext) {}

// EnterGreaterThan is called when production greaterThan is entered.
func (s *BaseACIParserListener) EnterGreaterThan(ctx *GreaterThanContext) {}

// ExitGreaterThan is called when production greaterThan is exited.
func (s *BaseACIParserListener) ExitGreaterThan(ctx *GreaterThanContext) {}

// EnterGreaterThanOrEqual is called when production greaterThanOrEqual is entered.
func (s *BaseACIParserListener) EnterGreaterThanOrEqual(ctx *GreaterThanOrEqualContext) {}

// ExitGreaterThanOrEqual is called when production greaterThanOrEqual is exited.
func (s *BaseACIParserListener) ExitGreaterThanOrEqual(ctx *GreaterThanOrEqualContext) {}

// EnterEqualTo is called when production equalTo is entered.
func (s *BaseACIParserListener) EnterEqualTo(ctx *EqualToContext) {}

// ExitEqualTo is called when production equalTo is exited.
func (s *BaseACIParserListener) ExitEqualTo(ctx *EqualToContext) {}

// EnterNotEqualTo is called when production notEqualTo is entered.
func (s *BaseACIParserListener) EnterNotEqualTo(ctx *NotEqualToContext) {}

// ExitNotEqualTo is called when production notEqualTo is exited.
func (s *BaseACIParserListener) ExitNotEqualTo(ctx *NotEqualToContext) {}

// EnterBindKeyword is called when production bindKeyword is entered.
func (s *BaseACIParserListener) EnterBindKeyword(ctx *BindKeywordContext) {}

// ExitBindKeyword is called when production bindKeyword is exited.
func (s *BaseACIParserListener) ExitBindKeyword(ctx *BindKeywordContext) {}

// EnterTargetKeyword is called when production targetKeyword is entered.
func (s *BaseACIParserListener) EnterTargetKeyword(ctx *TargetKeywordContext) {}

// ExitTargetKeyword is called when production targetKeyword is exited.
func (s *BaseACIParserListener) ExitTargetKeyword(ctx *TargetKeywordContext) {}

// EnterOpeningParenthesis is called when production openingParenthesis is entered.
func (s *BaseACIParserListener) EnterOpeningParenthesis(ctx *OpeningParenthesisContext) {}

// ExitOpeningParenthesis is called when production openingParenthesis is exited.
func (s *BaseACIParserListener) ExitOpeningParenthesis(ctx *OpeningParenthesisContext) {}

// EnterClosingParenthesis is called when production closingParenthesis is entered.
func (s *BaseACIParserListener) EnterClosingParenthesis(ctx *ClosingParenthesisContext) {}

// ExitClosingParenthesis is called when production closingParenthesis is exited.
func (s *BaseACIParserListener) ExitClosingParenthesis(ctx *ClosingParenthesisContext) {}
