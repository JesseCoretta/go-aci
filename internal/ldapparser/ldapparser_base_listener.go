// Code generated from LDAPParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package ldapparser // LDAPParser
import "github.com/antlr4-go/antlr/v4"

// BaseLDAPParserListener is a complete listener for a parse tree produced by LDAPParser.
type BaseLDAPParserListener struct{}

var _ LDAPParserListener = &BaseLDAPParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLDAPParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLDAPParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLDAPParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLDAPParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterUniformResourceIdentifier is called when production uniformResourceIdentifier is entered.
func (s *BaseLDAPParserListener) EnterUniformResourceIdentifier(ctx *UniformResourceIdentifierContext) {
}

// ExitUniformResourceIdentifier is called when production uniformResourceIdentifier is exited.
func (s *BaseLDAPParserListener) ExitUniformResourceIdentifier(ctx *UniformResourceIdentifierContext) {
}

// EnterURISearchAttributes is called when production uRISearchAttributes is entered.
func (s *BaseLDAPParserListener) EnterURISearchAttributes(ctx *URISearchAttributesContext) {}

// ExitURISearchAttributes is called when production uRISearchAttributes is exited.
func (s *BaseLDAPParserListener) ExitURISearchAttributes(ctx *URISearchAttributesContext) {}

// EnterURIDelimiter is called when production uRIDelimiter is entered.
func (s *BaseLDAPParserListener) EnterURIDelimiter(ctx *URIDelimiterContext) {}

// ExitURIDelimiter is called when production uRIDelimiter is exited.
func (s *BaseLDAPParserListener) ExitURIDelimiter(ctx *URIDelimiterContext) {}

// EnterBaseObject_scope is called when production baseObject_scope is entered.
func (s *BaseLDAPParserListener) EnterBaseObject_scope(ctx *BaseObject_scopeContext) {}

// ExitBaseObject_scope is called when production baseObject_scope is exited.
func (s *BaseLDAPParserListener) ExitBaseObject_scope(ctx *BaseObject_scopeContext) {}

// EnterOnelevel_scope is called when production onelevel_scope is entered.
func (s *BaseLDAPParserListener) EnterOnelevel_scope(ctx *Onelevel_scopeContext) {}

// ExitOnelevel_scope is called when production onelevel_scope is exited.
func (s *BaseLDAPParserListener) ExitOnelevel_scope(ctx *Onelevel_scopeContext) {}

// EnterSubtree_scope is called when production subtree_scope is entered.
func (s *BaseLDAPParserListener) EnterSubtree_scope(ctx *Subtree_scopeContext) {}

// ExitSubtree_scope is called when production subtree_scope is exited.
func (s *BaseLDAPParserListener) ExitSubtree_scope(ctx *Subtree_scopeContext) {}

// EnterParenthetical_filter is called when production parenthetical_filter is entered.
func (s *BaseLDAPParserListener) EnterParenthetical_filter(ctx *Parenthetical_filterContext) {}

// ExitParenthetical_filter is called when production parenthetical_filter is exited.
func (s *BaseLDAPParserListener) ExitParenthetical_filter(ctx *Parenthetical_filterContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseLDAPParserListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseLDAPParserListener) ExitFilter(ctx *FilterContext) {}

// EnterAnd_filter_expression is called when production and_filter_expression is entered.
func (s *BaseLDAPParserListener) EnterAnd_filter_expression(ctx *And_filter_expressionContext) {}

// ExitAnd_filter_expression is called when production and_filter_expression is exited.
func (s *BaseLDAPParserListener) ExitAnd_filter_expression(ctx *And_filter_expressionContext) {}

// EnterOr_filter_expression is called when production or_filter_expression is entered.
func (s *BaseLDAPParserListener) EnterOr_filter_expression(ctx *Or_filter_expressionContext) {}

// ExitOr_filter_expression is called when production or_filter_expression is exited.
func (s *BaseLDAPParserListener) ExitOr_filter_expression(ctx *Or_filter_expressionContext) {}

// EnterNot_filter_expression is called when production not_filter_expression is entered.
func (s *BaseLDAPParserListener) EnterNot_filter_expression(ctx *Not_filter_expressionContext) {}

// ExitNot_filter_expression is called when production not_filter_expression is exited.
func (s *BaseLDAPParserListener) ExitNot_filter_expression(ctx *Not_filter_expressionContext) {}

// EnterAva_expression is called when production ava_expression is entered.
func (s *BaseLDAPParserListener) EnterAva_expression(ctx *Ava_expressionContext) {}

// ExitAva_expression is called when production ava_expression is exited.
func (s *BaseLDAPParserListener) ExitAva_expression(ctx *Ava_expressionContext) {}

// EnterAnd is called when production and is entered.
func (s *BaseLDAPParserListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production and is exited.
func (s *BaseLDAPParserListener) ExitAnd(ctx *AndContext) {}

// EnterOr is called when production or is entered.
func (s *BaseLDAPParserListener) EnterOr(ctx *OrContext) {}

// ExitOr is called when production or is exited.
func (s *BaseLDAPParserListener) ExitOr(ctx *OrContext) {}

// EnterNot is called when production not is entered.
func (s *BaseLDAPParserListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production not is exited.
func (s *BaseLDAPParserListener) ExitNot(ctx *NotContext) {}

// EnterAttributeValueAssertion is called when production attributeValueAssertion is entered.
func (s *BaseLDAPParserListener) EnterAttributeValueAssertion(ctx *AttributeValueAssertionContext) {}

// ExitAttributeValueAssertion is called when production attributeValueAssertion is exited.
func (s *BaseLDAPParserListener) ExitAttributeValueAssertion(ctx *AttributeValueAssertionContext) {}

// EnterAttributeTypeOrMatchingRule is called when production attributeTypeOrMatchingRule is entered.
func (s *BaseLDAPParserListener) EnterAttributeTypeOrMatchingRule(ctx *AttributeTypeOrMatchingRuleContext) {
}

// ExitAttributeTypeOrMatchingRule is called when production attributeTypeOrMatchingRule is exited.
func (s *BaseLDAPParserListener) ExitAttributeTypeOrMatchingRule(ctx *AttributeTypeOrMatchingRuleContext) {
}

// EnterExtensibleMatch is called when production extensibleMatch is entered.
func (s *BaseLDAPParserListener) EnterExtensibleMatch(ctx *ExtensibleMatchContext) {}

// ExitExtensibleMatch is called when production extensibleMatch is exited.
func (s *BaseLDAPParserListener) ExitExtensibleMatch(ctx *ExtensibleMatchContext) {}

// EnterEqualityMatch is called when production equalityMatch is entered.
func (s *BaseLDAPParserListener) EnterEqualityMatch(ctx *EqualityMatchContext) {}

// ExitEqualityMatch is called when production equalityMatch is exited.
func (s *BaseLDAPParserListener) ExitEqualityMatch(ctx *EqualityMatchContext) {}

// EnterGreaterThanOrEqualMatch is called when production greaterThanOrEqualMatch is entered.
func (s *BaseLDAPParserListener) EnterGreaterThanOrEqualMatch(ctx *GreaterThanOrEqualMatchContext) {}

// ExitGreaterThanOrEqualMatch is called when production greaterThanOrEqualMatch is exited.
func (s *BaseLDAPParserListener) ExitGreaterThanOrEqualMatch(ctx *GreaterThanOrEqualMatchContext) {}

// EnterLessThanOrEqualMatch is called when production lessThanOrEqualMatch is entered.
func (s *BaseLDAPParserListener) EnterLessThanOrEqualMatch(ctx *LessThanOrEqualMatchContext) {}

// ExitLessThanOrEqualMatch is called when production lessThanOrEqualMatch is exited.
func (s *BaseLDAPParserListener) ExitLessThanOrEqualMatch(ctx *LessThanOrEqualMatchContext) {}

// EnterApproximateMatch is called when production approximateMatch is entered.
func (s *BaseLDAPParserListener) EnterApproximateMatch(ctx *ApproximateMatchContext) {}

// ExitApproximateMatch is called when production approximateMatch is exited.
func (s *BaseLDAPParserListener) ExitApproximateMatch(ctx *ApproximateMatchContext) {}

// EnterObjectIdentifier is called when production objectIdentifier is entered.
func (s *BaseLDAPParserListener) EnterObjectIdentifier(ctx *ObjectIdentifierContext) {}

// ExitObjectIdentifier is called when production objectIdentifier is exited.
func (s *BaseLDAPParserListener) ExitObjectIdentifier(ctx *ObjectIdentifierContext) {}

// EnterOpeningParenthesis is called when production openingParenthesis is entered.
func (s *BaseLDAPParserListener) EnterOpeningParenthesis(ctx *OpeningParenthesisContext) {}

// ExitOpeningParenthesis is called when production openingParenthesis is exited.
func (s *BaseLDAPParserListener) ExitOpeningParenthesis(ctx *OpeningParenthesisContext) {}

// EnterClosingParenthesis is called when production closingParenthesis is entered.
func (s *BaseLDAPParserListener) EnterClosingParenthesis(ctx *ClosingParenthesisContext) {}

// ExitClosingParenthesis is called when production closingParenthesis is exited.
func (s *BaseLDAPParserListener) ExitClosingParenthesis(ctx *ClosingParenthesisContext) {}

// EnterDistinguishedName is called when production distinguishedName is entered.
func (s *BaseLDAPParserListener) EnterDistinguishedName(ctx *DistinguishedNameContext) {}

// ExitDistinguishedName is called when production distinguishedName is exited.
func (s *BaseLDAPParserListener) ExitDistinguishedName(ctx *DistinguishedNameContext) {}

// EnterAttributeValue is called when production attributeValue is entered.
func (s *BaseLDAPParserListener) EnterAttributeValue(ctx *AttributeValueContext) {}

// ExitAttributeValue is called when production attributeValue is exited.
func (s *BaseLDAPParserListener) ExitAttributeValue(ctx *AttributeValueContext) {}
