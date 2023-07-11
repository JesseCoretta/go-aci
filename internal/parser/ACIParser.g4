/*
ACIv3 Parser Grammar

Author: Jesse Coretta ▲
Date:   07/06/2023

This ANTLRv4 (4.13.0) parser grammar implements parser support for
Version 3.0 of the Access Control Instruction syntax specification
and all of its abstract components. See below for LICENSE details.

ACIv3 is a popular expressive access control syntax used by various
directory products on the market today, including (but not limited
to) Netscape and Oracle Unified. ACIs are also considered "online
rules" in that modifications do not generally require DSA downtime.
Most often they reside within the LDAP entries they are prescribed
to protect, and are stored via the multi-valued 'aci' attributeType.

If you believe this solution lacks certain "syntactical sugars" of
which I am unaware (and you can cite literature to that end), you
are encouraged to open a ticket on the go-aci repository.

See also the accompanying (sourced) ACILexer.g4 file.

MIT License

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
SOFTWARE.
*/

parser grammar ACIParser;

// Because this grammar solution became so large, I have split the
// lexer grammer into its own file, which we source here.
options { tokenVocab=ACILexer; }

// instruction is the main parsing target, which is comprised of
// many constituent "parts", all of which are defined later in this
// grammar file.
parse
  : instruction EOF
  ;

// instruction describes a legal version 3.0 access control instruction
// (aci) at the top level.
//
// I don't want to be required to specify optional WHSP here as
// I am doing, but implementing this has proven difficult.
instruction
  : targetRules LPAREN ANCHOR DQUOTE attributeTypeOrValue DQUOTE WHSP*? SEMI WHSP*? permissionBindRules WHSP? RPAREN # aci
  ;

///////////////////////////////////////////////////////////////////////////////
// Begin PERMISSION and RIGHTS

// permissionBindRules describes one (1) or more permissionBindRule
// values. Values of this kind appear within the top-level of an ACI.
permissionBindRules
  : permissionBindRule*		# permission_bind_rules
  ;

// permissionBindRule describes a permission and Bind Rule pair.
// Values of this kind appear within permissionBindRules values.
permissionBindRule
  : permission bindRule WHSP? SEMI WHSP?	# permission_and_bind_rule_pair
  ;

// permission describes a complete permissive statement for an ACI, which
// may either grant or deny certain privileges.
permission
  : permissionDisposition WHSP*? LPAREN ( WHSP*? accessPrivileges ( COMMA accessPrivileges WHSP*?)* ) WHSP*? RPAREN	# permission_expression
  ;

// permissionDisposition describes the disposition of a given ACI permission
// statement, which may be either 'allow' or 'deny'.
permissionDisposition
  : ALLOW_ACCESS	# allow_access
  | DENY_ACCESS		# deny_access
  ;

// accessPrivileges contains multiple discrete privilege
// identifiers, each of which may be used to define access
// rights granted or withheld within a given ACI.
accessPrivileges
  : SEARCH_PRIVILEGE	# search_privilege
  | READ_PRIVILEGE	# read_privilege
  | COMPARE_PRIVILEGE	# compare_privilege
  | ADD_PRIVILEGE	# add_privilege
  | DELETE_PRIVILEGE	# delete_privilege
  | SELFWRITE_PRIVILEGE	# selfwrite_privilege
  | PROXY_PRIVILEGE	# proxy_privilege
  | IMPORT_PRIVILEGE	# import_privilege
  | EXPORT_PRIVILEGE	# export_privilege
  | ALL_PRIVILEGES	# all_privileges
  ;

///////////////////////////////////////////////////////////////////////////////
// Begin TARGET RULES

// targetRules defines a sequence of zero (0) or more
// targetRule instances; use of Target Rules is optional
// in ACIs.
targetRules
  : targetRule*              # target_rules
  ;

// targetRule defines any one (1) of nine (9) possible
// Target Rule types.
targetRule
  : targetControl		# rule_is_targetcontrol
  | targetExtendedOperation	# rule_is_extop
  | targetFilter		# rule_is_targetfilter
  | targetAttrFilters		# rule_is_targattrfilters
  | targetScope			# rule_is_targetscope
  | targetAttributes		# rule_is_targetattr
  | target			# rule_is_target
  | targetTo			# rule_is_target_to
  | targetFrom			# rule_is_target_from
  ;

// 'target' Target Rule syntax
target
  : LPAREN TARGET (equalTo|notEqualTo) distinguishedNames RPAREN     					# target_dn_rule
  ;

// 'target_to' Target Rule syntax
targetTo
  : LPAREN TARGET_TO (equalTo|notEqualTo) DQUOTE distinguishedName DQUOTE RPAREN     			# target_to_rule
  ;

// 'target_from' Target Rule syntax
targetFrom
  : LPAREN TARGET_FROM (equalTo|notEqualTo) DQUOTE distinguishedName DQUOTE RPAREN   			# target_from_rule
  ;

// 'targetfilter' Target Rule syntax
targetFilter
  : LPAREN TARGET_FILTER (equalTo|notEqualTo) DQUOTE lDAPFilter DQUOTE RPAREN        			# targetfilter_rule
  ;

// 'targattrfilters' Target Rule syntax
targetAttrFilters
  : LPAREN TARGET_ATTR_FILTERS equalTo DQUOTE targetAttrFiltersValue DQUOTE RPAREN    			# targattrfilters_rule
  ;

// 'targetscope' Target Rule syntax
targetScope
  : LPAREN TARGET_SCOPE equalTo DQUOTE targetSearchScopes DQUOTE RPAREN 				# targetscope_rule
  ;

// 'targetattr' Target Rule syntax
targetAttributes
  : LPAREN TARGET_ATTR (equalTo|notEqualTo) targetedAttributes RPAREN 					# targetattr_rule
  ;

// 'targetcontrol' Target Rule syntax
targetControl
  : LPAREN TARGET_CONTROL (equalTo|notEqualTo) objectIdentifiers RPAREN	 				# targetcontrol_rule
  ;

// 'extop' Target Rule syntax
targetExtendedOperation
  : LPAREN TARGET_EXTENDED_OPERATION (equalTo|notEqualTo) objectIdentifiers RPAREN 			# targetextop_rule
  ;

// targetSearchScopes contain the standard LDAP scopes contexts
// using the more "distinguished" names. For example, "subtree"
// is used instead of "sub", and "onelevel" instead of "one".
//
// Instances of this kind are used singularly within 'targetscope'
// Target Rules.
targetSearchScopes
  : BASE_OBJECT_SCOPE		# base_object_targetscope
  | ONE_LEVEL_TARGET_SCOPE   	# one_level_targetscope
  | SUB_TREE_TARGET_SCOPE	# sub_tree_targetscope
  | SUBORDINATE_TARGET_SCOPE 	# subordinate_targetscope
  ;

// objectIdentifiers is used by 'targetcontrol' and 'extop' Target Rules and will
// manifest a value in one (1) of the following forms:
//
// - "<n>.<n>.<n>.<...> || <n>.<n>.<n>.<...>"
// - "<n>.<n>.<n>.<...>" || "<n>.<n>.<n>.<...>"
//
// Values of this kind are used to represent an ORed (||) list of OIDs.
objectIdentifiers
  : DQUOTE ( objectIdentifier ( oRDelimiter objectIdentifier )* ) DQUOTE             			# quoted_object_identifier_list
  | ( DQUOTE objectIdentifier DQUOTE ( oRDelimiter (DQUOTE objectIdentifier DQUOTE) )* )		# list_of_quoted_object_identifiers
  ;

// targetedAttributes is used by 'targetattr' Target Rules, and will manifest a value
// in one (1) of the following forms:
// 
// - "attr || attr || attr"
// - "attr" || "attr" || "attr"
// - "*"
// Values of this kind are used to represent an ORed (||) list of attributeTypes.
targetedAttributes
  : DQUOTE ( attributeTypeOrValue ( oRDelimiter attributeTypeOrValue )* ) DQUOTE     			# quoted_targeted_attributes_list
  | ( DQUOTE attributeTypeOrValue DQUOTE ( oRDelimiter (DQUOTE attributeTypeOrValue DQUOTE) )* )	# list_of_quoted_attributes
  | DQUOTE STAR DQUOTE											# all_attributes
  ;

// objectIdentifier contains a dot notation ASN.1 Object Identifier. Values of this
// kind shall be used within 'targetcontrol' and 'extop' Target Rules.
objectIdentifier
  : ( numberForm ( DOT numberForm )+ )                			# object_identifier
  ;

// numberForm represents a single arc within a given ASN.1 Object Identifier in dot notation.
numberForm
  : INT                                                                 # number_form
  ;

// targattrfilters value can be one of the following
//
//   [add|delete]=attributeType:filter
//   [add|delete]=attributeType:filter && attributeType:filter ...
//   add=attributeType:filter && attributeType:filter ... (;|,) delete=attributeType:filter && attributeType:filter ...
//
// Note that some directory implementations delimit the latter set
// with semicolon, others use comma: both are supported here :)
//
// Double ampersands (symbolic AND ('&&')) are always used, never '||'.
//
targetAttrFiltersValue
  : attributeFilters           # attribute_filters_sets
  | attributeFilterSet         # attribute_filters_set
  | attributeFilter            # attribute_filter_single
  ;

// A sequence of one (1) or more attributeFilterAnd values
// that are joined by a comma or semicolon (depending on
// the vendor implementation).
attributeFilters
  : attributeFilterSet (COMMA|SEMI) attributeFilterSet     # attribute_filters
  ;

// A sequence of one (1) or more attributeFilter values prefixed with a
// single LDAP Operation (add or delete) and (if need be) joined by a
// double ampersand (&&)
attributeFilterSet
  : attributeFilterOperation ( attributeFilter ( aNDDelimiter attributeFilter )* )? # attribute_filter_set
  ;

// attributeFilterOperation describes the "operational intent" behind the
// specified attributeType:filter pair. Valid values are "add" and "delete".
attributeFilterOperation
  : ADD_PRIVILEGE equalTo	# add_filter_operation
  | DELETE_PRIVILEGE equalTo	# delete_filter_operation
  ;

// attributeFilter is an attributeType and LDAP Search
// Filter delimited by a single colon (:).
attributeFilter
  : attributeTypeOrValue COLON lDAPFilter	# attribute_filter
  ;

///////////////////////////////////////////////////////////////////////////////
// Begin BIND RULES

// Bind Rule Boolean statements
bindRule
  : bindRuleExpr                                                		# bind_rule
  | bindRuleExprParen ((BOOLEAN_AND|BOOLEAN_OR|BOOLEAN_NOT) bindRuleExprParen)* # parenthetical_bind_rule
  ;

// Parenthetical Bind Rule expressions
bindRuleExprParen
  : LPAREN bindRuleExpr ((BOOLEAN_AND|BOOLEAN_OR|BOOLEAN_NOT) bindRuleExpr)* RPAREN	# parenthetical_bind_rule_req_bool_op
  | <assoc=right> BOOLEAN_NOT bindRuleExpr						# negated_bind_rule_expression
  | LPAREN bindRuleExpr RPAREN                                       			# parenthetical_bind_rule_expression
  | bindRuleExpr                                                     			# bind_rule_expression_recursion
  ;

// bindRuleExpr contains a single Bind Rule in the form of
// <bind_keyword> <comparison_operator> <assertion_value>
bindRuleExpr
  : LPAREN bindRuleExpr RPAREN   # rule_is_parenthetical
  | bindUserDN                   # rule_is_userdn
  | bindUserAttr                 # rule_is_userattr
  | bindGroupDN                  # rule_is_groupdn
  | bindGroupAttr                # rule_is_groupattr
  | bindRoleDN                   # rule_is_roledn
  | bindDNS                      # rule_is_dns
  | bindIP                       # rule_is_ip
  | bindTimeOfDay                # rule_is_timeofday
  | bindDayOfWeek                # rule_is_dayofweek
  | bindSecurityStrengthFactor   # rule_is_ssf
  | bindAuthMethod               # rule_is_authmethod
  ;

// 'dayofweek' Bind Rule syntax
bindDayOfWeek
  : LPAREN bindDayOfWeek RPAREN                                			# parenthetical_dayofweek_bind_rule
  | BIND_DAY_OF_WEEK (equalTo|notEqualTo) DQUOTE ( doW ( COMMA doW )* ) DQUOTE	# dayofweek_bind_rule
  ;

// doW describes the individual days of the week. Instances containing these
// values are used within 'dayofweek' Bind Rules.
doW
  : SUNDAY	# Sun
  | MONDAY	# Mon
  | TUESDAY	# Tues
  | WEDNESDAY	# Wed
  | THURSDAY	# Thur
  | FRIDAY	# Fri
  | SATURDAY	# Sat
  ;

// 'authmethod' Bind Rule syntax
bindAuthMethod
  : LPAREN bindAuthMethod RPAREN                                       		# parentheticalAuthenticationMethod
  | BIND_AUTH_METHOD (equalTo|notEqualTo) DQUOTE authenticationMethods DQUOTE 	# authentication_method
  ;

// authenticationMethods describes each of the possible authentication
// mechanisms (or lack thereof) a requestor may leverage during LDAP
// communication between a DUA and DSA.
authenticationMethods
  : ANONYMOUS	# none
  | SIMPLE	# simple
  | SSL		# ssl
  | SASL	# sasl
  ;

// 'userdn' Bind Rule syntax
bindUserDN
  : LPAREN bindUserDN RPAREN								# parenthetical_bind_userdn 
  | BIND_USER_DN (equalTo|notEqualTo) WHSP? (distinguishedNames|DQUOTE lDAPURI DQUOTE)	# bind_userdn
  ;

// 'roledn' Bind Rule syntax
bindRoleDN
  : LPAREN bindRoleDN RPAREN					# parenthetical_bind_roledn
  | BIND_ROLE_DN (equalTo|notEqualTo) distinguishedNames 	# bind_roledn
  ;

// 'groupdn' Bind Rule syntax
bindGroupDN
  : LPAREN bindGroupDN RPAREN				      				# parenthetical_bind_groupdn
  | BIND_GROUP_DN (equalTo|notEqualTo) (distinguishedNames|DQUOTE lDAPURI DQUOTE) 	# bind_groupdn
  ;

// 'userattr' Bind Rule syntax
bindUserAttr
  : LPAREN bindUserAttr RPAREN                 					# parenthetical_bind_userattr
  | BIND_USER_ATTR (equalTo|notEqualTo) DQUOTE (attributeBindTypeOrValue|inheritance) DQUOTE  # bind_userattr
  ;

// 'groupattr' Bind Rule syntax
bindGroupAttr
  : LPAREN bindGroupAttr RPAREN                                        		# parenthetical_bind_groupattr
  | BIND_GROUP_ATTR (equalTo|notEqualTo) DQUOTE (attributeBindTypeOrValue|inheritance) DQUOTE	# bind_groupattr
  ;

// 'ssf' Bind Rule syntax
bindSecurityStrengthFactor
  : LPAREN bindSecurityStrengthFactor RPAREN                             					# parenthetical_ssf
  | BIND_SSF (equalTo|notEqualTo|greaterThan|greaterThanOrEqual|lessThan|lessThanOrEqual) DQUOTE INT DQUOTE	# bind_ssf
  ;

// 'timeofday' Bind Rule syntax
bindTimeOfDay
  : LPAREN bindTimeOfDay RPAREN												# parenthetical_bind_timeofday
  | BIND_TIME_OF_DAY (equalTo|notEqualTo|greaterThan|greaterThanOrEqual|lessThan|lessThanOrEqual) DQUOTE INT DQUOTE	# bind_timeofday
  ;

// 'ip' Bind Rule syntax
bindIP
  : LPAREN bindIP RPAREN                       			# parenthetical_bind_ip
  | BIND_IP (equalTo|notEqualTo) DQUOTE iPAddresses DQUOTE	# bind_ip
  ;

// 'dns' Bind Rule syntax
bindDNS
  : LPAREN bindDNS RPAREN                      				# parenthetical_bind_dns
  | BIND_DNS (equalTo|notEqualTo) DQUOTE fQDN DQUOTE			# dns_bind_rule
  ;

// iPAddresses contains a sequence of one (1) or more IPv4 or IPv6
// addresses, delimited by COMMA as needed.
iPAddresses
  : ( iPAddress ( COMMA iPAddress )* )+?      				# ips
  ;

// iPAddress describes any single IPv4 or IPv6 address, and may include
// STAR for octet wildcard statements.
iPAddress
  : iPv4Address								# ipv4_address
  | iPv6Address								# ipv6_address
  ;

// iPv4Address describes a single IPv4 address, which may include a
// STAR for octet wildcard statements.
iPv4Address
  : ( INT ( DOT (INT|STAR)* ) )						# ipv4
  ;

// iPv6Address describes a single IPv6 address, which may include a
// STAR for octet wildcard statements.
iPv6Address
  : ( attributeTypeOrValue ( COLON attributeTypeOrValue )+ COLON? )	# ipv6
  ;

// fQDN describes a single fully-qualified domain name, which may 
// include a STAR for label wildcard statements.
fQDN
  : ( attributeTypeOrValue ( DOT attributeTypeOrValue )+ )		# fqdn
  ;

///////////////////////////////////////////////////////////////////////////////
// Begin LDAP related rules

// lDAPURI describes a fully-qualified LDAP URI, which may include one (1) or
// more of the following LDAP Search parameters in the following order:
//
// - Comma-delimited attributeType list
// - A standard LDAP Search Scope (base, one, sub)
// - An LDAP Search Filter
//
// The prefix of values of this kind shall ALWAYS be a distinguishedName which
// bears local LDAP scheme (ldap:///).
lDAPURI
  : distinguishedName uRIAttributeList uRISearchScopes uRISearchFilter       # fullyQualifiedLDAPURI
  ;

// uRISearchFilter describes a single LDAP Search Filter with a QMARK (?)
// delimiter prefix. Values of this kind shall appear in the final (far-right)
// field of an lDAPURI.
uRISearchFilter
  : QMARK lDAPFilter                                                 # uriSearchFilter
  ;

// uRISearchScopes describes one (1) of three (3) possible choices for an
// LDAP Search Scope: base, one or sub.  Values of this kind shall appear
// in the middle field of the lDAPURI search parameters.
uRISearchScopes
  : QMARK (BASE_OBJECT_SCOPE|ONE_LEVEL_SCOPE|SUB_TREE_SCOPE)?        # uriSearchScopes
  ;

// uRIAttributeList describes a list of zero (0) or more comma-delimited
// attributeType names to be requested for the search operation. Values
// of this kind appear immediately right of the distinguishedName.
uRIAttributeList
  : QMARK ( attributeTypeOrValue ( COMMA attributeTypeOrValue )* )?  # uriAttributeList
  ;

// distinguishedNames contains one or more LDAP Distinguished Names. In the
// case of >1 DNs, the symbolic OR (||) delimiter is used. This applies to
// 'userdn', 'groupdn' and 'target' rules.
distinguishedNames
  : DQUOTE ( distinguishedName ( oRDelimiter distinguishedName )* ) DQUOTE             		# quoted_distinguished_name_list
  | ( DQUOTE distinguishedName DQUOTE ( oRDelimiter (DQUOTE distinguishedName DQUOTE) )* )	# list_of_quoted_distinguished_names
  ;

// distinguishedName is a sequence of aVAOrRDN values. Macro
// variable declarations for [$dn], ($dn) and ($attr.<atname>)
// are supported.
distinguishedName
  : ( LOCAL_LDAP_SCHEME aVAOrRDN ( COMMA (aVAOrRDN|rDNMacros) )* )	# dn
  | LOCAL_LDAP_SCHEME ANYONE						# anonymous_dn_alias
  | LOCAL_LDAP_SCHEME ALL_USERS						# any_user_dn_alias
  | LOCAL_LDAP_SCHEME SELF						# self_dn_alias
  | LOCAL_LDAP_SCHEME PARENT						# parent_dn_alias
  ;

// rDNMacros contains macro variables for DSA interpolation. The possible
// variables are as follows:
//
// - [$dn]
// - ($dn)
// - ($attr.<at>)
//
// Instances of this type are used within 'userdn' and 'groupdn' Bind Rules
// (assuming the LDAP DSA product supports Macro ACIs). Please note that
// Macro ACIs are not part of the universally-adopted syntax components.
rDNMacros
  : RDN_MACROS							   	# rdn_macro
  ;

// lDAPFilter describes the string representation of an RFC4515 LDAP Search Filter.
// Values of this kind are used in several areas, including 'targetfilter' Target Rules,
// LDAP URIs, et al.
lDAPFilter
  : LPAREN lDAPFilterExpr RPAREN       	# parenthetical_filter_expression
  | lDAPFilterExpr*                    	# filter_expressions
  ;

// lDAPFilterExpr describes a single expressive LDAP filter statement, but may
// not be a complete filter unto itself.
lDAPFilterExpr
  : (LPAREN (FILTER_AND|FILTER_OR|FILTER_NOT)? lDAPFilterExpr RPAREN)+?    # parenthetical_filter_expression_opt_bool
  | <assoc=right> FILTER_NOT lDAPFilterExpr                 		   # not_filter_expression
  | aVAOrRDN			                               		   # ava_expression
  ;

// aVAOrRDN is two things (for the sake of a simple grammar file)
// - when present within a distinguishedName, it is a relativeDistinguishedName (e.g.: 'uid=jesse')
// - when present within a filter, it is an attributeValueAssertion (e.g.: 'cn~=Jesse' or 'objectClass=*')
aVAOrRDN
  : LPAREN attributeTypeOrValue attributeComparisonOperator attributeTypeOrValue RPAREN	# parenthetical_ava_or_rdn
  | attributeTypeOrValue attributeComparisonOperator attributeTypeOrValue	 	# ava_or_rdn
  ;

// Vertical Inheritance (for User/Group Attribute Matching)
inheritance
  : ( PARENT inheritanceLevels DOT attributeBindTypeOrValue )       # inheritance_expression
  ;

// inheritanceLevels describe one (1) of five (5) possible "depth levels":
//
// - Level Zero (0) (baseObject)
// - Level One (1) (one level below baseObject)
// - Level Two (2) (two levels below baseObject)
// - Level Three (3) (three levels below baseObject)
// - Level Four (4) (four levels below baseObject)
//
// Values of this kind are used in inheritance-based 'userattr' and 'groupattr'
// Bind Rules.
inheritanceLevels
  : LBRAK ( INT ( COMMA INT )* )+? RBRAK	# inheritance_levels
  ;

// attributeBindTypeOrValue describes a value of the following syntax:
//
// <attributeType>#<bindType_or_value>
//
// Values of this kind are used in certain 'userattr' and 'groupattr' Bind
// Rules.
attributeBindTypeOrValue
  : attributeTypeOrValue HASH (bindTypes|attributeTypeOrValue)      # attr_bind_type_or_value
  ;

// bindTypes describes one (1) of five (5) possible BIND TYPES to be
// specified for certain 'userattr' and 'groupattr' Bind Rules.
bindTypes
  : BINDTYPE_USER_DN	# USERDN
  | BINDTYPE_GROUP_DN	# GROUPDN
  | BINDTYPE_ROLE_DN	# ROLEDN
  | BINDTYPE_SELF_DN	# SELFDN
  | BINDTYPE_LDAP_URL	# LDAPURL
  ;

// attributeTypeOrValue describes a general attributeType OR
// assertion value. Values of this kind MAY manifest as STAR
// for wildcard statements or presence checks.
attributeTypeOrValue
  : KEY_OR_VALUE	# key_or_value
  | STAR		# presence_key_or_value
  ;

// attributeComparisonOperator describes one (1) of nine (9)
// possible comparison operators to be used in AVAs and RDNs.
attributeComparisonOperator
  : equalTo		# equal_to
  | notEqualTo		# not_equal_to
  | greaterThanOrEqual	# greater_than_or_equal
  | lessThanOrEqual	# less_than_or_equal
  | approximate		# approx
  | extensibleRule	# extensible_rule
  | extensibleRuleDN	# extensible_rule_with_dn
  | extensibleRuleAttr	# extensible_rule_with_attrs
  | extensibleRuleDNOID	# extensible_rule_with_dn_oid
  ;

equalTo				: EQ;
notEqualTo			: NE;
greaterThan			: GT;
lessThan			: LT;
greaterThanOrEqual		: GE;
lessThanOrEqual			: LE;
approximate			: APX;
extensibleRule			: COLON;
extensibleRuleDNOID		: EXO;
extensibleRuleDN		: EXD;
extensibleRuleAttr		: EXA;
oRDelimiter			: SYMBOLIC_OR;
aNDDelimiter			: SYMBOLIC_AND;
