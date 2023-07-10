/*
ACIv3 Grammar

Author: Jesse Coretta - 07/06/2023

This ANTLRv4 (4.13.0) file implements parser/lexer support for the
Access Control Instruction Syntax Version 3 and all of its abstract
components.

ACIv3 is a popular expressive access control syntax used by various
directory products on the market today, including Netscape and Oracle
Unified.
*/
parser grammar ACIParser;

options { tokenVocab=ACILexer; }

// Begin Parser directives

parse
     : instruction EOF
     ;

// Instruction is the top level access control instruction definition
instruction
     : LPAREN targetRule version permBindRules RPAREN SEMI	# accessControlInstruction
     ;

version
     : LPAREN ANCHOR Literal SEMI	# accessControlInstructionNameAndVersion
     ;

permBindRules
     : permBindRule*		# permissionBindRules
     ;

permBindRule
     : permission bindRule SEMI	# permissionBindRule
     ;

permission
     : DISPOSITION LPAREN ( RIGHTS ( COMMA RIGHTS )* )? RPAREN 	# permissionExpression
     ;

// targetRules defines a sequence of zero (0) or more
// targetRule instances; use of Target Rules is optional
// in ACIs. Each targetRule instance present within an
// instance of this type must be unique within an ACI.
// In other words, the same Target Keyword -- such as
// userdn -- can only appear in a single Target Rule).
targetRules
     : targetRule*		# targetRuleExpressions
     ;

// targetRule defines any one (1) of nine (9) possible
// Target Rule types.
targetRule
     : targetControl		# targetcontrol
     | targetExtOp		# extop
     | targetFilter		# targetfilter
     | targetAttrFilters	# targattrfilters
     | targetScope		# targetscope
     | targetAttr		# targetattr
     | target			# targetdn
     | targetTo			# targettodn
     | targetFrom		# targetfromDN
     ;

// 'targetcontrol' and 'extop' Target Rule syntaxes contain
// one (1) or more ASN.1 Object Identifiers delimited by a
// double pipe (||).
targetControl
     : LPAREN TargetControlKeyword (EqualTo|NotEqualTo) objectIdentifiers RPAREN	# parentheticalControls
     ;

targetExtOp
     : LPAREN ExtOpKeyword (EqualTo|NotEqualTo) objectIdentifiers RPAREN	# parentheticalExtendedOperations
     ;

// 'targetfilter' Target Rule syntax
targetFilter
     : LPAREN TargetFilterKeyword (EqualTo|NotEqualTo) targetFilterValue	# parentheticalTargetFilterExpression
     ;

targetFilterValue
     : DQUOTE lDAPFilter DQUOTE	# quotedFilterExpression
     ;

// 'target', 'target_to' and 'target_from' Target Rule syntax
target
     : LPAREN TargetKeyword (EqualTo|NotEqualTo) distinguishedNames	# targetDistinguishedNames
     ;

targetTo
     : LPAREN TargetToKeyword (EqualTo|NotEqualTo) distinguishedName	# targetToDistinguishedName
     ;

targetFrom
     : LPAREN TargetFromKeyword (EqualTo|NotEqualTo) distinguishedName	# targetFromDistinguishedName
     ;

// 'targetattrfilters' Target Rule syntax
targetAttrFilters
     : LPAREN TargetAttrFiltersKeyword EqualTo targetAttrFiltersValue RPAREN	# parentheticalTargetAttrFilters
     ;

targetAttrFiltersValue
     : DQUOTE attributeFilters DQUOTE		# quotedAttributeFilters
     | DQUOTE attributeFilterSet DQUOTE		# quotedAttributeFilterSet
     | DQUOTE attributeFilter DQUOTE		# quotedAttributeFilter
     ;

// 'targetscope' Target Rule syntax
targetScope
     : LPAREN TargetScopeKeyword EqualTo TargetRuleSearchScopes RPAREN # targetScopeBindRule
     ;

// 'targetattr' Target Rule syntax
targetAttr
     : LPAREN TargetAttrKeyword (EqualTo|NotEqualTo) attributeType RPAREN	# targetAttrBindRule
     ;

// List is not quoted as a whole, but individual members are.
attributeTypes
     : ( attributeType ( doublePipe attributeType )* )			# attributeTypesList
     ;

// Bind Rule Boolean statements
bindRule
     : bindRuleExpr						   # bindRuleInstance
     | bindRuleExprParen (BooleanWordOperators bindRuleExprParen)* # parentheticalBindRuleInstanceWithRequiredBooleanOperator
     ;

// Bind Rule expressions
bindRuleExprParen
     : LPAREN bindRuleExpr (BooleanWordOperators bindRuleExpr)* RPAREN	# parentheticalBindRuleExpressionWithRequiredBooleanOperatorRecursion
     | <assoc=right> BooleanNotWord bindRuleExpr			# negatedBindRuleExpressionRecursion
     | LPAREN bindRuleExpr RPAREN					# parentheticalBindRuleExpressionRecursion
     | bindRuleExpr							# bindRuleExpressionRecursion
     ;

// bindRuleExpr contains a single Bind Rule in the form of
// <bind_keyword> <comparison_operator> <assertion_value>
bindRuleExpr
     : LPAREN bindRuleExpr RPAREN	# bindRuleExpression
     | bindRuleUserDN			# userdn
     | bindRuleUserAttr			# userattr
     | bindRuleGroupDN			# groupdn
     | bindRuleGroupAttr		# groupattr
     | bindRuleRoleDN			# roledn
     | bindRuleDNS			# dns
     | bindRuleIP			# ip
     | bindRuleTimeOfDay		# timeofday
     | bindRuleDayOfWeek		# dayofweek
     | bindRuleSecurityStrengthFactor	# ssf
     | bindRuleAuthMethod		# authmethod
     ;

// User/Group/Role Distinguished Name Matching (Bind Rule)
bindRuleUserDN
     : LPAREN bindRuleUserDN RPAREN			# parentheticalUserDistinguishedName
     | UserDNKeyword (EqualTo|NotEqualTo) distinguishedName	# userDistinguishedNameExpression
     ;

bindRuleRoleDN
     : LPAREN bindRuleRoleDN RPAREN			# parentheticalRoleDistinguishedName
     | RoleDNKeyword (EqualTo|NotEqualTo) distinguishedName	# roleDistinguishedNameExpression
     ;

bindRuleGroupDN
     : LPAREN bindRuleGroupDN RPAREN			# parentheticalGroupDistinguishedName
     | GroupDNKeyword (EqualTo|NotEqualTo) distinguishedName	# groupDistinguishedNameExpression
     ;

// User/Group Attribute Matching (Bind Rule)
bindRuleUserAttr
     : LPAREN bindRuleUserAttr RPAREN						# parentheticalUserAttributes
     | UserAttrKeyword (EqualTo|NotEqualTo) (attributeBindTypeOrValue|inheritance)	# userAttributesExpression
     ;

bindRuleGroupAttr
     : LPAREN bindRuleGroupAttr RPAREN						# parentheticalGroupAttributes
     | GroupAttrKeyword (EqualTo|NotEqualTo) (attributeBindTypeOrValue|inheritance)	# groupAttributesExpression
     ;

// Authentication Method Matching (Bind Rule)
bindRuleAuthMethod
     : LPAREN bindRuleAuthMethod RPAREN						# parentheticalAuthenticationMethod
     | AuthMethodKeyword (EqualTo|NotEqualTo) (ANONYMOUS|SIMPLE|SSL|SASL)	# authenticationMethodExpression
     ;

// DNS Hostname Matching (w/ wildcard support) (Bind Rule)
bindRuleDNS
     : LPAREN bindRuleDNS RPAREN			# parentheticalDNS
     | DNSKeyword (EqualTo|NotEqualTo) fQDN		# dNSBindRule
     ;

// TimeOfDay Matching (Bind Rule)
bindRuleTimeOfDay
     : LPAREN bindRuleTimeOfDay RPAREN					# parentheticalTimeOfDay
     | TimeOfDayKeyword (EqualTo|NotEqualTo|GreaterThan|GreaterThanOrEqual|LessThan|LessThanOrEqual) DQUOTE timeOfDay DQUOTE	# timeOfDayBindRule
     ;

// DayOfWeek Matching (Bind Rule)
bindRuleDayOfWeek
     : LPAREN bindRuleDayOfWeek RPAREN					# parentheticalDayOfWeek
     | DayOfWeekKeyword (EqualTo|NotEqualTo) DQUOTE dayOfWeek DQUOTE	# dayOfWeekExpression
     ;

// IPv4/IPv6 Address Matching (w/ wildcard support) (Bind Rule)
bindRuleIP
     : LPAREN bindRuleIP RPAREN				# parentheticalIPAddress
     | IPKeyword (EqualTo|NotEqualTo) (iPV4Address|iPV6Address)	# ipAddressBindRule
     ;

// Security Strength Factor Matching (Bind Rule)
bindRuleSecurityStrengthFactor
     : LPAREN bindRuleSecurityStrengthFactor RPAREN				# parentheticalSecurityStrengthFactor
     | SSFKeyword (EqualTo|NotEqualTo|GreaterThan|GreaterThanOrEqual|LessThan|LessThanOrEqual) securityStrengthFactor	# securityStrengthFactorExpression
     ;

dayOfWeek
     : ( DAYS ( COMMA DAYS )* )?				# dayOfWeekValue
     ;

fQDN
     : DelimitedAddress					# fullyQualifiedDomainNameValue
     ;

// ASN.1 Object Identifier(s) for LDAP Controls
// and LDAP Extended Operations
objectIdentifiers
     : ( objectIdentifier ( doublePipe objectIdentifier )+ )	# objectIdentifierValues
     ;

objectIdentifier
     : DelimitedAddress						# objectIdentifierValue
     ;

iPV6Address
     : DelimitedAddress						# iPV6AddressValue
     ;

iPV4Address
     : DelimitedAddress						# iPV4AddressValue
     ;

securityStrengthFactor
     : INT							# securityStrengthFactorValue
     ;

timeOfDay
     : INT							# timeOfDayValue
     ;

objectIdentifierArc: INT;

// Vertical Inheritance (for User/Group Attribute Matching)
inheritance
     : ( INHERITANCEPREFIX inheritanceLevels RBRAK DOT attributeBindTypeOrValue )	# inheritanceExpression
     ;

inheritanceLevels
     : DelimitedNumbers						# inheritanceLevelValue
     ;

attributeBindTypeOrValue
     : attributeType HASH ( BINDTYPES | attributeValue )	# attributeBindTypeOrValueValue
     ;

// A sequence of one (1) or more attributeFilterAnd values
// that are joined by a comma or semicolon (depending on
// the vendor implementation).
attributeFilters
     : ( attributeFilterSet ( (COMMA|SEMI) attributeFilterSet )+? )	# attributeFiltersExpression
     ;

// A sequence of one (1) or more attributeFilter values prefixed with a
// single LDAP Operation (add or delete) and (if need be) joined by a 
// double ampersand (&&)
attributeFilterSet
     : (AddOperation|DeleteOperation) EqualTo ( attributeFilter ( doubleAmpersand attributeFilter )* )?	# attributeFilterSetExpression
     ;

doubleAmpersand
     : DAMP			# doubleAmpersandDelimiter
     ;

// A single attr:filter value
attributeFilter
     : (attributeType COLON lDAPFilter)	# attributeFilterExpression
     ;

// distinguishedNames contains one or more distinguishedName
// values, and is delimited with a double pipe (||). This
// parser applies to groupdn and userdn Bind Rules which cite
// multiple DNs.
distinguishedNames
     : ( distinguishedName ( doublePipe distinguishedName )* )+	# distinguishedNamesList
     ;

doublePipe
     : DPIPE	# doublePipeDelimiter
     ;

lDAPURIAndBindType
     : distinguishedName QMARK attributeBindTypeOrValue				# uriAndBindType
     ;

lDAPURI
     : distinguishedName uRIAttributeList uRISearchScopes uRISearchFilter 	# fullyQualifiedLDAPURI
     ;

uRISearchFilter
     : QMARK lDAPFilter							# uriSearchFilter
     ;

uRISearchScopes
     : QMARK LDAPSearchScopes?						# uriSearchScopes
     ;

uRIAttributeList
     : QMARK ( attributeType ( COMMA attributeType )* )?		# uriAttributeList
     ;

// distinguishedName is a sequence of relativeDistinguishedName values. Macro
// variable declarations for [$dn], ($dn) and ($attr.<atname>) are supported.
distinguishedName
     : ( LocalLDAPScheme ( relativeDistinguishedName ( COMMA relativeDistinguishedName )* ) )	# distinguishedNameValue
     ;

// relativeDistinguishedName is functionally identical to attributeValueAssertion, but
// also includes so-called Macro ACI support, e.g.: 'uid=jesse' vs. ($attr.uid).
relativeDistinguishedName
     : ( attributeType EqualTo attributeValue )	# relativeDistinguishedNameValue
     | MacroValue				# relativeDistinguishedNameMacro
     ;

lDAPFilter
     : LPAREN lDAPFilterExpr RPAREN	# parentheticalFilterExpression
     | lDAPFilterExpr*			# filterExpressions
     ;

lDAPFilterExpr
     : (LPAREN (AMP|PIPE|BANG)? lDAPFilterExpr RPAREN)+?  # parentheticalFilterExpressionWithOptionalBooleanOperator
     | <assoc=right> BANG lDAPFilterExpr		  # negatedFilterExpression
     | attributeValueAssertion				  # attributeValueAssertionExpression
     ;

attributeValueAssertion
     : attributeType attributeOperators attributeValue	# attributeValueAssertionStatement
     ;

attributeType
     : ID						# attributeTypeIdentifier
     ;

attributeValue
     : ( ANY | ID | INT | WildcardString )		# attributeAssertionValue
     ;

attributeOperators
     : (EqualTo|ExtensibleRuleAttrMatch|ExtensibleRuleMatch|GreaterThanOrEqual|LessThanOrEqual|ExtensibleRuleDNOIDMatch|ExtensibleRuleDNMatch|ApproximateMatch)
     ;
