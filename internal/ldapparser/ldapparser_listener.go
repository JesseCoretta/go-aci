// Code generated from LDAPParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package ldapparser // LDAPParser
import "github.com/antlr4-go/antlr/v4"

// LDAPParserListener is a complete listener for a parse tree produced by LDAPParser.
type LDAPParserListener interface {
	antlr.ParseTreeListener

	// EnterUniformResourceIdentifier is called when entering the uniformResourceIdentifier production.
	EnterUniformResourceIdentifier(c *UniformResourceIdentifierContext)

	// EnterURISearchAttributes is called when entering the uRISearchAttributes production.
	EnterURISearchAttributes(c *URISearchAttributesContext)

	// EnterURIDelimiter is called when entering the uRIDelimiter production.
	EnterURIDelimiter(c *URIDelimiterContext)

	// EnterBaseObject_scope is called when entering the baseObject_scope production.
	EnterBaseObject_scope(c *BaseObject_scopeContext)

	// EnterOnelevel_scope is called when entering the onelevel_scope production.
	EnterOnelevel_scope(c *Onelevel_scopeContext)

	// EnterSubtree_scope is called when entering the subtree_scope production.
	EnterSubtree_scope(c *Subtree_scopeContext)

	// EnterParenthetical_filter is called when entering the parenthetical_filter production.
	EnterParenthetical_filter(c *Parenthetical_filterContext)

	// EnterFilter is called when entering the filter production.
	EnterFilter(c *FilterContext)

	// EnterAnd_filter_expression is called when entering the and_filter_expression production.
	EnterAnd_filter_expression(c *And_filter_expressionContext)

	// EnterOr_filter_expression is called when entering the or_filter_expression production.
	EnterOr_filter_expression(c *Or_filter_expressionContext)

	// EnterNot_filter_expression is called when entering the not_filter_expression production.
	EnterNot_filter_expression(c *Not_filter_expressionContext)

	// EnterAva_expression is called when entering the ava_expression production.
	EnterAva_expression(c *Ava_expressionContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterOr is called when entering the or production.
	EnterOr(c *OrContext)

	// EnterNot is called when entering the not production.
	EnterNot(c *NotContext)

	// EnterAttributeValueAssertion is called when entering the attributeValueAssertion production.
	EnterAttributeValueAssertion(c *AttributeValueAssertionContext)

	// EnterAttributeTypeOrMatchingRule is called when entering the attributeTypeOrMatchingRule production.
	EnterAttributeTypeOrMatchingRule(c *AttributeTypeOrMatchingRuleContext)

	// EnterExtensibleMatch is called when entering the extensibleMatch production.
	EnterExtensibleMatch(c *ExtensibleMatchContext)

	// EnterEqualityMatch is called when entering the equalityMatch production.
	EnterEqualityMatch(c *EqualityMatchContext)

	// EnterGreaterThanOrEqualMatch is called when entering the greaterThanOrEqualMatch production.
	EnterGreaterThanOrEqualMatch(c *GreaterThanOrEqualMatchContext)

	// EnterLessThanOrEqualMatch is called when entering the lessThanOrEqualMatch production.
	EnterLessThanOrEqualMatch(c *LessThanOrEqualMatchContext)

	// EnterApproximateMatch is called when entering the approximateMatch production.
	EnterApproximateMatch(c *ApproximateMatchContext)

	// EnterObjectIdentifier is called when entering the objectIdentifier production.
	EnterObjectIdentifier(c *ObjectIdentifierContext)

	// EnterOpeningParenthesis is called when entering the openingParenthesis production.
	EnterOpeningParenthesis(c *OpeningParenthesisContext)

	// EnterClosingParenthesis is called when entering the closingParenthesis production.
	EnterClosingParenthesis(c *ClosingParenthesisContext)

	// EnterDistinguishedName is called when entering the distinguishedName production.
	EnterDistinguishedName(c *DistinguishedNameContext)

	// EnterAttributeValue is called when entering the attributeValue production.
	EnterAttributeValue(c *AttributeValueContext)

	// ExitUniformResourceIdentifier is called when exiting the uniformResourceIdentifier production.
	ExitUniformResourceIdentifier(c *UniformResourceIdentifierContext)

	// ExitURISearchAttributes is called when exiting the uRISearchAttributes production.
	ExitURISearchAttributes(c *URISearchAttributesContext)

	// ExitURIDelimiter is called when exiting the uRIDelimiter production.
	ExitURIDelimiter(c *URIDelimiterContext)

	// ExitBaseObject_scope is called when exiting the baseObject_scope production.
	ExitBaseObject_scope(c *BaseObject_scopeContext)

	// ExitOnelevel_scope is called when exiting the onelevel_scope production.
	ExitOnelevel_scope(c *Onelevel_scopeContext)

	// ExitSubtree_scope is called when exiting the subtree_scope production.
	ExitSubtree_scope(c *Subtree_scopeContext)

	// ExitParenthetical_filter is called when exiting the parenthetical_filter production.
	ExitParenthetical_filter(c *Parenthetical_filterContext)

	// ExitFilter is called when exiting the filter production.
	ExitFilter(c *FilterContext)

	// ExitAnd_filter_expression is called when exiting the and_filter_expression production.
	ExitAnd_filter_expression(c *And_filter_expressionContext)

	// ExitOr_filter_expression is called when exiting the or_filter_expression production.
	ExitOr_filter_expression(c *Or_filter_expressionContext)

	// ExitNot_filter_expression is called when exiting the not_filter_expression production.
	ExitNot_filter_expression(c *Not_filter_expressionContext)

	// ExitAva_expression is called when exiting the ava_expression production.
	ExitAva_expression(c *Ava_expressionContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitOr is called when exiting the or production.
	ExitOr(c *OrContext)

	// ExitNot is called when exiting the not production.
	ExitNot(c *NotContext)

	// ExitAttributeValueAssertion is called when exiting the attributeValueAssertion production.
	ExitAttributeValueAssertion(c *AttributeValueAssertionContext)

	// ExitAttributeTypeOrMatchingRule is called when exiting the attributeTypeOrMatchingRule production.
	ExitAttributeTypeOrMatchingRule(c *AttributeTypeOrMatchingRuleContext)

	// ExitExtensibleMatch is called when exiting the extensibleMatch production.
	ExitExtensibleMatch(c *ExtensibleMatchContext)

	// ExitEqualityMatch is called when exiting the equalityMatch production.
	ExitEqualityMatch(c *EqualityMatchContext)

	// ExitGreaterThanOrEqualMatch is called when exiting the greaterThanOrEqualMatch production.
	ExitGreaterThanOrEqualMatch(c *GreaterThanOrEqualMatchContext)

	// ExitLessThanOrEqualMatch is called when exiting the lessThanOrEqualMatch production.
	ExitLessThanOrEqualMatch(c *LessThanOrEqualMatchContext)

	// ExitApproximateMatch is called when exiting the approximateMatch production.
	ExitApproximateMatch(c *ApproximateMatchContext)

	// ExitObjectIdentifier is called when exiting the objectIdentifier production.
	ExitObjectIdentifier(c *ObjectIdentifierContext)

	// ExitOpeningParenthesis is called when exiting the openingParenthesis production.
	ExitOpeningParenthesis(c *OpeningParenthesisContext)

	// ExitClosingParenthesis is called when exiting the closingParenthesis production.
	ExitClosingParenthesis(c *ClosingParenthesisContext)

	// ExitDistinguishedName is called when exiting the distinguishedName production.
	ExitDistinguishedName(c *DistinguishedNameContext)

	// ExitAttributeValue is called when exiting the attributeValue production.
	ExitAttributeValue(c *AttributeValueContext)
}
