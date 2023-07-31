/*
Basic LDAP Filter, DN and URI PARSER Grammar for ANTLR4 (4.13.0)

NOTE: work in progress!

MIT LICENSE

Copyright (c) 2023 Jesse Coretta

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE

ABOUT THIS FILE

This ANTLRv4 (4.13.0) grammar file contains basic parser rules for LDAP Search
URIs and their components: Search Filters, Scopes, Attributes and Distinguished
Names.

SUPPORTED EXPRESSIONS, SYNTAX AND SOME LIMITATIONS

While effective at "decompiling" or "dissecting" an LDAP Search Filter for use
in writing a parser, it should not be considered a complete "parsing" solution.
Certain syntactically-bogus LDAP filters -- such as those missing encapsulation
characters (i.e.: closing parenthesis) -- may be processed sans any noticeable
error output.

That said, this grammar file can aid in said deconstruction of filter statements.

All of the standard attributeAssertionValue "operators" are supported:

 - Equality (=)
 - Approximate (~=)
 - Ordering (>=/<=)
 - ExtensibleRules ( ':dn:2.5.13.5:=', ':caseIgnoreMatch:=', et al)

All of the logical Boolean symbol "operators" are supported:

 - AND (&)
 - OR (|)
 - NOT (!)

Any attributeValueAssertion (AVA) comprised of an attributeType, one (1) or more
operator permutations and a non-zero attributeValue should be processed without
incident ** SO LONG AS THE EXPRESSION IS PARENTHETICAL **.

In other words:

 - GOOD:  (objectClass=*)
 - BAD:   objectClass=*

Though the "BAD" example is perfectly legal in most respected LDAP DUAs, ANTLR's
lexer design (when combining things like DNs and Filters within the same grammar
file) reacts quite badly to ambiguity among varying rules. If you're writing a
parser using this grammar, you can easily compensate for this in your code before
processing begins (e.g.: "if <val> is not parenthetical, make it so").

Substring values (such as 'ab*d') are supported in attributeValue instances, but
the associated attributeType is  obviously not verified for such capabilities in
the ANTLR4 grammar "layer" according to any schema definitions. The same is true
for ordering (GE/LE comparison).

In other words, schema-illogical AVAs would be identified (likely) after the ANTLR
parsing stage. An example of what I mean: '(color>=blue)' is illogical LDAP-wise,
but as far as ANTLR is concerned, perfectly valid per the parser rules.

AttributeType Tags, such as ';binary' or ';lang-CC', are supported as suffix
components to an attributeType instance, e.g.: 'color;lang-fr=bleu'.

LEXER GRAMMAR

See the accompanying (sourced) LDAPLexer.g4 file for lexer grammar rules.
*/

parser grammar LDAPParser;
options { tokenVocab=LDAPLexer; }

// uniformResourceIdentifier expresses a fully-qualified LDAP URI containing a
// combination of the following values:
//
// - DistinguishedName (with 'ldap', 'ldaps' or 'ldapi' scheme prefix using one
//   (1) or more of the following:
//
//   - ldap://<FQDN|IPv4:port|[IPv6]:port>/, OR ...
//   - IPC local-path prefix (ldapi://%2fvar%2frun%2ldapi/), OR ...
//   - implicit localhost prefix (ldap:/// or ldaps:///)
//
// - Search Attributes (one (1) or more comma-delimited LDAP Attribute Types)
// - Search Scopes ('base', 'one', 'sub' or undefined)
// - Search Filter (e.g.: '(&(objectClass=account)(gidNumber=8701))')
//
// Each of the above components are delimited with a QUESTION MARK character (ASCII #34, '?'), whether or
// not the values are specified. In other words, if only a DN were provided, the resultant URI might appear
// as 'ldap:///ou=People,dc=example,dc=com???'.
uniformResourceIdentifier
  : distinguishedName
	uRIDelimiter uRISearchAttributes
	uRIDelimiter searchScope
	uRIDelimiter searchFilter
  ;

uRISearchAttributes: attributeTypeOrMatchingRule ( COMMA attributeTypeOrMatchingRule )*;
uRIDelimiter: QMARK;

// LDAP Search Scope 
searchScope
  : BASEOBJECT_SCOPE   # baseObject_scope
  | SINGLELEVEL_SCOPE  # onelevel_scope
  | WHOLESUBTREE_SCOPE # subtree_scope
  ;

// searchFilter describes the string representation of an RFC4515 LDAP Search Filter.
//
// e.g.: "(&(objectClass=employee)(terminated=FALSE))"
searchFilter
  : openingParenthesis searchFilter closingParenthesis # parenthetical_filter
  | searchFilterExpr              	   	       # filter
  ;

// searchFilterExpr describes a single expressive LDAP filter statement, but may
// not be a complete filter unto itself.
searchFilterExpr
  : openingParenthesis (and searchFilterExpr+)+? closingParenthesis # and_filter_expression
  | openingParenthesis (or searchFilterExpr+)+? closingParenthesis # or_filter_expression
  | <assoc=right> openingParenthesis not searchFilterExpr closingParenthesis # not_filter_expression
  | attributeValueAssertion+  # ava_expression
  ;

// and (&) represents a stack of conditions, all of which must
// evaluate as TRUE
and: AND;

// or (|) represents a stack of conditions, at least one (1) of which
// must evaluate as TRUE
or:  OR;

// not (!) represents a negated stack of conditions
not: NOT;

// attributeValueAssertion is a filter component that specifies a condition involving an
// attributeType, a comparison operator of some kind and an attributeValue.
attributeValueAssertion
  : equalityMatch
  | greaterThanOrEqualMatch
  | lessThanOrEqualMatch
  | approximateMatch
  | extensibleMatch
  ;

// attributeTypeOrMatchingRule represents a single LDAP attributeType or
// matchingRule identifier. In the case of an attributeType, an optional
// attribute tag (e.g.: ';lang-fr') is permitted.
attributeTypeOrMatchingRule: IDENTIFIER;

extensibleMatch
  : openingParenthesis attributeTypeOrMatchingRule? (EXTENSIBLE_RULE (DN|attributeTypeOrMatchingRule|objectIdentifier))* EXTENSIBLE_RULE EQUALITY (distinguishedName|attributeValue) closingParenthesis
  ;
equalityMatch
  : openingParenthesis attributeTypeOrMatchingRule EQUALITY (distinguishedName|attributeValue|ASTERISK) closingParenthesis
  ;
greaterThanOrEqualMatch
  : openingParenthesis attributeTypeOrMatchingRule GREATER_THAN_OR_EQUAL attributeValue	closingParenthesis
  ;
lessThanOrEqualMatch
  : openingParenthesis attributeTypeOrMatchingRule LESS_THAN_OR_EQUAL attributeValue closingParenthesis
  ;
approximateMatch
  : openingParenthesis attributeTypeOrMatchingRule APPROXIMATE attributeValue closingParenthesis
  ;
objectIdentifier
  : OBJECT_IDENTIFIER
  ;
openingParenthesis: LPAREN ;
closingParenthesis: RPAREN ;

// distinguishedName may, or may not, be prefixed with 'ldap://<fqdn>/' or just 'ldap:///'.
distinguishedName: LOCAL_LDAP_SCHEME? ( attributeTypeOrMatchingRule EQUALITY attributeValue ) ( COMMA ( attributeTypeOrMatchingRule EQUALITY attributeValue ) )*;

// attributeValue is a broad catch-all that tries to represent any
// possible LDAP attribute value, without any consideration for
// ldapSyntax, matchingRules, etc.
attributeValue:	(.)+?|.;
