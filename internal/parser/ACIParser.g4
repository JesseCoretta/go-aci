/*
ACIv3 Grammar

Author: Jesse Coretta - 07/06/2023

This ANTLRv4 (4.13.0) file implements parser/lexer support for the
Access Control Instruction Syntax Version 3 and all of its abstract
components. It is released under the terms of the MIT License.

ACIv3 is a popular expressive access control syntax used by various
directory products on the market today, including Netscape and Oracle
Unified.
*/

parser grammar ACIParser;

options { tokenVocab=ACILexer; }

parse
  : instruction EOF
  ;

// instruction describes a version 3.0 access control
// instruction (aci) at the top level.
instruction
  : targetRules LPAREN ANCHOR DQUOTE attributeTypeOrValue DQUOTE WHSP*? SEMI WHSP*? permissionBindRules WHSP? RPAREN # aci
  ;

///////////////////////////////////////////////////////////////////////////////
// Begin PERMISSION and RIGHTS

permissionBindRules
  : permissionBindRule*		# permission_bind_rules
  ;

permissionBindRule
  : permission bindRule WHSP? SEMI WHSP?	# permission_and_bind_rule_pair
  ;

permission
  : permissionDisposition WHSP*? LPAREN ( WHSP*? accessPrivileges ( COMMA accessPrivileges WHSP*?)* ) WHSP*? RPAREN	# permission_expression
  ;

permissionDisposition
  : ALLOW_ACCESS	# allow_access
  | DENY_ACCESS		# deny_access
  ;

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
  : targetControl
  | targetExtendedOperation
  | targetFilter
  | targetAttrFilters
  | targetScope
  | targetAttributes
  | target
  | targetTo
  | targetFrom
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

// 'targetcontrol' Target Rule syntax
targetExtendedOperation
  : LPAREN TARGET_EXTENDED_OPERATION (equalTo|notEqualTo) objectIdentifiers RPAREN 			# targetextop_rule
  ;

targetSearchScopes
  : BASE_OBJECT_SCOPE		# base_object_targetscope
  | ONE_LEVEL_TARGET_SCOPE   	# one_level_targetscope
  | SUB_TREE_TARGET_SCOPE	# sub_tree_targetscope
  | SUBORDINATE_TARGET_SCOPE 	# subordinate_targetscope
  ;

objectIdentifiers
  : DQUOTE ( objectIdentifier ( oRDelimiter objectIdentifier )* ) DQUOTE             			# quoted_object_identifier_list
  | ( DQUOTE objectIdentifier DQUOTE ( oRDelimiter (DQUOTE objectIdentifier DQUOTE) )* )		# list_of_quoted_object_identifiers
  ;

targetedAttributes
  : DQUOTE ( attributeTypeOrValue ( oRDelimiter attributeTypeOrValue )* ) DQUOTE     			# quoted_targeted_attributes_list
  | ( DQUOTE attributeTypeOrValue DQUOTE ( oRDelimiter (DQUOTE attributeTypeOrValue DQUOTE) )* )	# list_of_quoted_attributes
  ;

objectIdentifier
  : ( objectIdentifierArc ( DOT objectIdentifierArc )+ )                # object_identifier
  ;

objectIdentifierArc
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
// Ampersands (symbolic AND ('&&')) are always used, never '||'.
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
  : ( attributeFilterSet ( (COMMA|SEMI) attributeFilterSet )+? )     # attribute_filters
  ;

// A sequence of one (1) or more attributeFilter values prefixed with a
// single LDAP Operation (add or delete) and (if need be) joined by a
// double ampersand (&&)
attributeFilterSet
  : attributeFilterOperation equalTo ( attributeFilter ( aNDDelimiter attributeFilter )* )? # attribute_filter_set
  ;

attributeFilterOperation
  : ADD_PRIVILEGE		# add_filter_operation
  | DELETE_PRIVILEGE		# delete_filter_operation
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
  : LPAREN bindRuleExpr RPAREN   # bind_rule_expression
  | bindUserDN                   # userdn_expression
  | bindUserAttr                 # userattr_expression
  | bindGroupDN                  # groupdn_expression
  | bindGroupAttr                # groupattr_expression
  | bindRoleDN                   # roledn_expression
  | bindDNS                      # dns_expression
  | bindIP                       # ip_expression
  | bindTimeOfDay                # timeofday_expression
  | bindDayOfWeek                # dayofweek_expression
  | bindSecurityStrengthFactor   # ssf_expression
  | bindAuthMethod               # authmethod_expression
  ;

// 'dayofweek' Bind Rule syntax
bindDayOfWeek
  : LPAREN bindDayOfWeek RPAREN                                			# parenthetical_dayofweek_bind_rule
  | BIND_DAY_OF_WEEK (equalTo|notEqualTo) DQUOTE ( doW ( COMMA doW )* ) DQUOTE	# dayofweek_bind_rule
  ;

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

iPAddresses
  : ( iPAddress ( COMMA iPAddress )* )+?      				# ips
  ;

iPAddress
  : iPv4Address								# ipv4_address
  | iPv6Address								# ipv6_address
  ;

iPv4Address
  : ( INT ( DOT (INT|STAR)* ) )						# ipv4
  ;

iPv6Address
  : ( attributeTypeOrValue ( COLON attributeTypeOrValue )+ COLON? )	# ipv6
  ;

fQDN
  : ( (attributeTypeOrValue|STAR) ( DOT (attributeTypeOrValue|STAR) )+ )	# fqdn
  ;

///////////////////////////////////////////////////////////////////////////////
// Begin LDAP related rules

lDAPURI
     : distinguishedName uRIAttributeList uRISearchScopes uRISearchFilter       # fullyQualifiedLDAPURI
     ;

uRISearchFilter
     : QMARK lDAPFilter                                                 # uriSearchFilter
     ;

uRISearchScopes
     : QMARK (BASE_OBJECT_SCOPE|ONE_LEVEL_SCOPE|SUB_TREE_SCOPE)?        # uriSearchScopes
     ;

uRIAttributeList
     : QMARK ( attributeTypeOrValue ( COMMA attributeTypeOrValue )* )?  # uriAttributeList
     ;

// distinguishedNames contains one or more LDAP Distinguished Names. In the
// case of >1 DNs, the symbolic OR (||) delimiter is used. This applies to
// userdn, groupdn and target rules.
distinguishedNames
  : DQUOTE ( distinguishedName ( oRDelimiter distinguishedName )* ) DQUOTE             			# quoted_distinguished_name_list
  | ( DQUOTE distinguishedName DQUOTE ( oRDelimiter (DQUOTE distinguishedName DQUOTE) )* )		# list_of_quoted_distinguished_names
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

rDNMacros
  : RDN_MACROS							   	# rdn_macro
  ;

lDAPFilter
  : LPAREN lDAPFilterExpr RPAREN     				   	# parenthetical_filter_expression
  | lDAPFilterExpr*                  				   	# filter_expressions
  ;

lDAPFilterExpr
  : (LPAREN (FILTER_AND|FILTER_OR|FILTER_NOT)? lDAPFilterExpr RPAREN)+?    # parenthetical_filter_expression_opt_bool
  | <assoc=right> FILTER_NOT lDAPFilterExpr                 		   # not_filter_expression
  | aVAOrRDN			                               		   # ava_expression
  ;

// aVAOrRDN is two things (for the sake of a simple grammar file)
// - when present within a distinguishedName, it is a relativeDistinguishedName (e.g.: 'uid=jesse')
// - when present within a filter, it is an attributeValueAssertion (e.g.: 'cn~=Jesse' or 'objectClass=*')
aVAOrRDN
  : attributeTypeOrValue attributeComparisonOperator attributeTypeOrValue	# ava_or_rdn
  ;

// Vertical Inheritance (for User/Group Attribute Matching)
inheritance
  : ( PARENT inheritanceLevels DOT attributeBindTypeOrValue )       # inheritance_expression
  ;

inheritanceLevels
  : LBRAK ( INT ( COMMA INT )* )+? RBRAK                              # inheritance_levels
  ;

attributeBindTypeOrValue
  : attributeTypeOrValue HASH (bindTypes|attributeTypeOrValue)      # attr_bind_type_or_value
  ;

bindTypes
  : BINDTYPE_USER_DN	# USERDN
  | BINDTYPE_GROUP_DN	# GROUPDN
  | BINDTYPE_ROLE_DN	# ROLEDN
  | BINDTYPE_SELF_DN	# SELFDN
  | BINDTYPE_LDAP_URL	# LDAPURL
  ;

attributeTypeOrValue
  : KEY_OR_VALUE	# key_or_value
  ;

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

equalTo
  : EQ
  ;

notEqualTo
  : NE
  ;

greaterThan
  : GT
  ;

lessThan
  : LT
  ;

greaterThanOrEqual
  : GE
  ;

lessThanOrEqual
  : LE
  ;

approximate
  : APX
  ;

extensibleRule
  : COLON
  ;

extensibleRuleDNOID
  : EXO
  ;

extensibleRuleDN
  : EXD
  ;

extensibleRuleAttr
  : EXA
  ;

oRDelimiter
  : SYMBOLIC_OR
  ;

aNDDelimiter
  : SYMBOLIC_AND
  ;
