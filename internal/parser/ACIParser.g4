/*
ACIv3 Parser Grammar

Implemented by Jesse Coretta ▲

ADVISORY

This is an initial release and is potentially unsuitable for
mission-critical / production environments. At the moment, it
should only be used as a convenient supplement to an already
hardened ACI review/approval/documentation/etc., process. Use
at your own risk.

See further down for LICENSE details.

ABOUT THIS FILE

This ANTLRv4 (4.13.0) parser grammar implements parser support for
Version 3.0 of the Access Control Instruction syntax specification
and all of its abstract components.

ABOUT THE ACI SYNTAX

ACIv3 is a popular expressive access control syntax used by various
directory products on the market today, including (but not limited
to) Netscape and Oracle Unified. ACIs are also considered "online
rules" in that modifications do not generally require DSA downtime.
Most often they reside within the LDAP entries they are prescribed
to protect, and are stored via the multi-valued 'aci' attributeType.

CONTRIBUTIONS and LIMITATIONS

If you believe this solution lacks a certain "syntactical sugar" of
which I am unaware (and you can cite literature to that end), then
you are encouraged to open a new ticket within the github repository
in which the parser resides.

Please note that, at this time, this solution does NOT cover these
ACI syntax "variants":

 - Apache DS "Entry, Prescriptive & Subentry ACIs"
 - OpenLDAP "Experimental ACIs"

The main reason is because they're so incredibly different from the
syntax honored here that I am uncertain as to the ideal means for
integration with this solution.

I may try to tackle this in the near future, but it is extremely low 
priority and I suspect there would be little to no demand for it, as:

 - (a) Apache DS is widely considered lackluster, much of the reference
       material is empty, labeled "TODO" or contain ToC-only pages that
       lead nowhere. Plus its Java. Java makes me want to take a shower

 - (b) OpenLDAP recommends users leverage their proprietary configuration
       based "ACL" syntax with or without dynamic configuration involved
       
Regarding (b): The OpenLDAP ACL syntax is actually quite nice, and far more
appealing -- if for no reason other than its popularity -- for incorporation
into this package when compared to the above (non-implemented) variants.

LEXER CONTENTS

See also the accompanying (sourced) ACILexer.g4 file for lexers.

LICENSE

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
// many "constituents" -- all of which are defined later in this
// grammar file (and its accompanying lexer file).
parse
  : instruction EOF
  ;

// instruction describes a legal version 3.0 access control instruction
// (aci) at the top-level. This is the parser component any given user
// is most likely to use directly, unless they are testing one or more
// constituent components only.
//
// Users are cautioned when using various non-alphanumeric chars for the
// DQUOTE'd attributeTypeOrValue parameter because I'm not 100% certain
// as to what the specification explicitly allows (AND, any given LDAP
// vendor may choose to enforce or relax such constraints for their own
// reasons). When given a choice, it is recommended one confine any string
// label to alphanumeric, whitespace, hyphen and underscore chars (while
// enforcing uniqueness of said value within the entire DIT!).
//
// YMMV and KISS.
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
//
// e.g.: allow(read,search,compare)
permission
  : permissionDisposition WHSP*? LPAREN ( WHSP*? accessPrivileges ( COMMA accessPrivileges WHSP*?)* ) WHSP*? RPAREN	# permission_expression
  ;

// permissionDisposition describes the disposition of a given ACI permission
// statement, which may be either 'allow' or 'deny'.
permissionDisposition
  : allow	# allow_access
  | deny	# deny_access
  ;

allow: ALLOW_ACCESS;
deny: DENY_ACCESS;

// accessPrivileges contains multiple discrete privilege
// identifiers, each of which may be used to define access
// rights granted or withheld within a given ACI.
accessPrivileges
  : searchPrivilege	# search_privilege
  | readPrivilege	# read_privilege
  | comparePrivilege	# compare_privilege
  | addPrivilege	# add_privilege
  | deletePrivilege	# delete_privilege
  | selfWritePrivilege	# selfwrite_privilege
  | proxyPrivilege	# proxy_privilege
  | importPrivilege	# import_privilege
  | exportPrivilege	# export_privilege
  | allPrivileges	# all_privileges
  | noPrivileges	# no_privileges
  ;

searchPrivilege:	SEARCH_PRIVILEGE;
readPrivilege:		READ_PRIVILEGE;
comparePrivilege:	COMPARE_PRIVILEGE;
addPrivilege:		ADD_PRIVILEGE;
deletePrivilege:	DELETE_PRIVILEGE;
selfWritePrivilege:	SELFWRITE_PRIVILEGE;
proxyPrivilege:		PROXY_PRIVILEGE;
exportPrivilege:	EXPORT_PRIVILEGE;
importPrivilege:	IMPORT_PRIVILEGE;

// Grant or withhold no privileges within the DSA.
// NOTE: alias of ANONYMOUS due to the string literal
// value they share.
noPrivileges: ANONYMOUS ;

// Grant or withhold all privileges within the DSA.
// NOTE: alias of ALL_USERS due to the string literal
// value they share.
allPrivileges: ALL_USERS ;

///////////////////////////////////////////////////////////////////////////////
// Begin TARGET RULES

// targetRules defines a sequence of zero (0) or more
// targetRule instances; use of Target Rules is optional
// in ACIs, however Target Rule statements are *ALWAYS*
// parenthetical, unlike Bind Rules which may be either.
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
// NOTE: ineligible for negation (!=).
targetAttrFilters
  : LPAREN TARGET_ATTR_FILTERS equalTo DQUOTE targetAttrFiltersValue DQUOTE RPAREN    			# targattrfilters_rule
  ;

// 'targetscope' Target Rule syntax
// NOTE: ineligible for negation (!=).
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
  : baseTargetScope		# base_object_targetscope
  | oneLevelTargetScope   	# one_level_targetscope
  | subTreeTargetScope		# sub_tree_targetscope
  | subordinateTargetScope 	# subordinate_targetscope
  ;

// baseTargetScope is an alias of BASE_OBJECT_SCOPE
// due to the string literal value they share.
baseTargetScope:	BASE_OBJECT_SCOPE;
oneLevelTargetScope:	ONE_LEVEL_TARGET_SCOPE;
subTreeTargetScope:	SUB_TREE_TARGET_SCOPE;
subordinateTargetScope:	SUBORDINATE_TARGET_SCOPE;

// objectIdentifiers is used by 'targetcontrol' and 'extop' Target Rules and will
// manifest a value in one (1) of the following forms:
//
// - "<n>.<n>.<n>.<...> || <n>.<n>.<n>.<...>"
// - "<n>.<n>.<n>.<...>" || "<n>.<n>.<n>.<...>"
// - "<n>.<n>.<n>.<...>"
//
// Values of this kind are used to represent an ORed (||) list of ASN.1 Object
// Identifiers -- specifically LDAP Control and Extended Operation OIDs.
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
//
// Values of this kind are used to represent an ORed (||) list of attributeTypes,
// except when "*" is used, which should be all-or-nothing in its application.
targetedAttributes
  : DQUOTE ( attributeTypeOrValue ( oRDelimiter attributeTypeOrValue )* ) DQUOTE     			# quoted_targeted_attributes_list
  | ( DQUOTE attributeTypeOrValue DQUOTE ( oRDelimiter (DQUOTE attributeTypeOrValue DQUOTE) )* )	# list_of_quoted_attributes
  | DQUOTE STAR DQUOTE											# all_attributes
  ;

// objectIdentifier contains a dot notation ASN.1 Object Identifier. Values of this
// kind shall be used within 'targetcontrol' and 'extop' Target Rules.
//
// e.g.: 2.16.840.1.113730.3.4.18
objectIdentifier
  : ( numberForm ( DOT numberForm )+ )                			# object_identifier
  ;

// numberForm represents a single arc of any unsigned magnitude within a given ASN.1
// Object Identifier in dot notation.
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
// with a SEMI, others use COMMA: both are supported here :)
//
// Double ampersands (symbolic AND ('&&')) are always used, never '||'.
//
// Note this is behaving a little quirky, not sure I've got it nailed down yet ;/
targetAttrFiltersValue
  : attributeFilters           # attribute_filters_sets
  | attributeFilterSet         # attribute_filters_set
  | attributeFilter            # attribute_filter_single
  ;

// A sequence of one (1) or more attributeFilterAnd values
// that are joined by a COMMA or SEMI (depending on vendor
// implementation).
//
// Note: linebreak added to example below for readability.
//
// e.g.: add=userCertificate:(&(objectClass=employee)(terminated=FALSE)) && userCertificate:(objectClass=shareholder); \
// 	delete=userCertificate:(&(objectClass=executive)(isJerk=TRUE)) && userCertificate:(&(objectClass=marketing)(dweeb=TRUE))
attributeFilters
  : attributeFilterSet (COMMA|SEMI) attributeFilterSet     # attribute_filters
  ;

// A sequence of one (1) or more attributeFilter values prefixed with a
// single LDAP Operation (add or delete) and (if need be) joined by a
// double ampersand (&&)
//
// e.g.: add=userCertificate:(&(objectClass=employee)(terminated=FALSE)) && userCertificate(objectClass=shareholder)
attributeFilterSet
  : attributeFilterOperation attributeFilter ( aNDDelimiter attributeFilter )*		# attribute_filter_set
  ;

// attributeFilterOperation describes the "operational intent" behind the
// specified attributeType:filter pair.
attributeFilterOperation
  : addFilterOperation equalTo	# add_filter_operation
  | delFilterOperation equalTo	# delete_filter_operation
  ;

// addFilterOperation is an alias of ADD_PRIVILEGE
// due to the common string literal value they share.
addFilterOperation: ADD_PRIVILEGE;

// delFilterOperation is an alias of DELETE_PRIVILEGE
// due to the common string literal value they share.
delFilterOperation: DELETE_PRIVILEGE;

// attributeFilter is an attributeType and LDAP Search
// Filter joined by a COLON.
//
// e.g.: userCertificate:(objectClass=employee)
attributeFilter
  : attributeTypeOrValue COLON lDAPFilter	# attribute_filter
  ;

///////////////////////////////////////////////////////////////////////////////
// Begin BIND RULES

// Bind Rule Boolean statements
//
// e.g.:
//	- (timeofday >= "1730" AND timeofday < "2400")
//	- authmethod = "SASL"
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
//
// e.g.: ssf >= "128"
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
//
// e.g.: dayofweek="Mon,Tues,Fri"
bindDayOfWeek
  : LPAREN bindDayOfWeek RPAREN                                			# parenthetical_dayofweek_bind_rule
  | BIND_DAY_OF_WEEK (equalTo|notEqualTo) DQUOTE ( doW ( COMMA doW )* ) DQUOTE	# dayofweek_bind_rule
  ;

// doW describes the individual days of the week. Instances containing these
// values are used within 'dayofweek' Bind Rules.
doW
  : sun	 # Sunday
  | mon	 # Monday
  | tues # Tuesday
  | wed	 # Wednesday
  | thur # Thurday
  | fri	 # Friday
  | sat	 # Saturday
  ;

sun:	SUNDAY;
mon:	MONDAY;
tues:	TUESDAY;
wed:	WEDNESDAY;
thur:	THURSDAY;
fri:	FRIDAY;
sat:	SATURDAY;

// 'authmethod' Bind Rule syntax
//
// e.g.: authmethod != "none"
bindAuthMethod
  : LPAREN bindAuthMethod RPAREN                                       		# parenthetical_authentication_method
  | BIND_AUTH_METHOD (equalTo|notEqualTo) DQUOTE authenticationMethods DQUOTE 	# authentication_method
  ;

// authenticationMethods describes each of the possible authentication
// mechanisms (or lack thereof) a requestor may leverage during LDAP
// communication between a DUA and DSA.
authenticationMethods
  : anonAuth	# none
  | simpleAuth	# simple
  | sSLAuth	# ssl
  | sASLAuth	# sasl
  ;

anonAuth:	ANONYMOUS;
simpleAuth:	SIMPLE;
sSLAuth:	SSL;
sASLAuth:	SASL;

// 'userdn' Bind Rule syntax
//
// e.g.: userdn="ldap:///uid=someone,ou=People,dc=example,dc=com"
bindUserDN
  : LPAREN bindUserDN RPAREN								# parenthetical_bind_userdn 
  | BIND_USER_DN (equalTo|notEqualTo) WHSP? (distinguishedNames|DQUOTE lDAPURI DQUOTE)	# bind_userdn
  ;

// 'roledn' Bind Rule syntax
//
// e.g.: roledn="ldap:///uid=someone,ou=People,dc=example,dc=com"
bindRoleDN
  : LPAREN bindRoleDN RPAREN					# parenthetical_bind_roledn
  | BIND_ROLE_DN (equalTo|notEqualTo) distinguishedNames 	# bind_roledn
  ;

// 'groupdn' Bind Rule syntax
//
// e.g.: groupdn="ldap:///cn=X.500 Administrators,ou=Groups,dc=example,dc=com"
bindGroupDN
  : LPAREN bindGroupDN RPAREN				      				# parenthetical_bind_groupdn
  | BIND_GROUP_DN (equalTo|notEqualTo) (distinguishedNames|DQUOTE lDAPURI DQUOTE) 	# bind_groupdn
  ;

// 'userattr' Bind Rule syntax
//
// e.g.: userattr="owner#USERDN"
bindUserAttr
  : LPAREN bindUserAttr RPAREN                 							# parenthetical_bind_userattr
  | BIND_USER_ATTR (equalTo|notEqualTo) DQUOTE (attributeBindTypeOrValue|inheritance) DQUOTE	# bind_userattr
  ;

// 'groupattr' Bind Rule syntax
//
// e.g.: groupattr="manager#LDAPURL"
bindGroupAttr
  : LPAREN bindGroupAttr RPAREN                                        				# parenthetical_bind_groupattr
  | BIND_GROUP_ATTR (equalTo|notEqualTo) DQUOTE (attributeBindTypeOrValue|inheritance) DQUOTE	# bind_groupattr
  ;

// 'ssf' Bind Rule syntax
//
// e.g.: ssf != "0"
bindSecurityStrengthFactor
  : LPAREN bindSecurityStrengthFactor RPAREN                             					# parenthetical_ssf
  | BIND_SSF (equalTo|notEqualTo|greaterThan|greaterThanOrEqual|lessThan|lessThanOrEqual) DQUOTE INT DQUOTE	# bind_ssf
  ;

// 'timeofday' Bind Rule syntax
//
// e.g.: (timeofday >= "1730" AND timeofday < "2400")
bindTimeOfDay
  : LPAREN bindTimeOfDay RPAREN												# parenthetical_bind_timeofday
  | BIND_TIME_OF_DAY (equalTo|notEqualTo|greaterThan|greaterThanOrEqual|lessThan|lessThanOrEqual) DQUOTE INT DQUOTE	# bind_timeofday
  ;

// 'ip' Bind Rule syntax
//
// e.g.: ip = "192.168.0,12.3.45.*,10.0.0.0/8"
bindIP
  : LPAREN bindIP RPAREN                       			# parenthetical_bind_ip
  | BIND_IP (equalTo|notEqualTo) DQUOTE iPAddresses DQUOTE	# bind_ip
  ;

// 'dns' Bind Rule syntax
//
// e.g.: dns = "www.example.com"
bindDNS
  : LPAREN bindDNS RPAREN                      				# parenthetical_bind_dns
  | BIND_DNS (equalTo|notEqualTo) DQUOTE fQDN DQUOTE			# dns_bind_rule
  ;

// iPAddresses contains a sequence of one (1) or more IPv4 or IPv6
// addresses, delimited by COMMA as needed.
//
// e.g.: '192.168.0,12.3.45.*,10.0.0.0/8'
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
//
// e.g.: '192.168.*'
iPv4Address
  : ( INT ( DOT (INT|STAR)* ) )						# ipv4
  ;

// iPv6Address describes a single IPv6 address, which may include a
// STAR for octet wildcard statements.
//
// e.g.: '2001:470:dead:beef::'
iPv6Address
  : ( attributeTypeOrValue ( COLON attributeTypeOrValue )+ COLON? )	# ipv6
  ;

// fQDN describes a single fully-qualified domain name, which may 
// include a STAR for label wildcard statements.
//
// e.g.: 'www.example.com' or '*.example.com'
fQDN
  : ( attributeTypeOrValue ( DOT attributeTypeOrValue )+ )		# fqdn
  ;

///////////////////////////////////////////////////////////////////////////////
// Begin LDAP related rules

// lDAPURI describes a fully-qualified LDAP URI, which will include either of
// the following conditions
//
// LDAP Search Parameters:
// 	- Comma-delimited attributeType list, and/or ..
// 	- A standard LDAP Search Scope (base, one, sub), and/or ..
// 	- An LDAP Search Filter
//
// ... OR ...
//
// A value that matches the attributeBindTypeOrValue parser type
//
// The prefix of values of this kind shall ALWAYS be a distinguishedName which
// bears local LDAP scheme (ldap:///), regardless of which structure above was
// qualified.
//
// e.g.:
//	- 'ldap:///uid=courtney,ou=People,dc=example,dc=com?cn,sn,givenName?sub?(objectClass=employee)'
//	- 'ldap:///uid=courtney,ou=People,dc=example,dc=com??sub?(objectClass=employee)'
//
// ... OR ...
//
//	- 'ldap:///uid=courtney,ou=People,dc=example,dc=com?manager#USERDN'
lDAPURI
  : distinguishedName uRIAttributeList uRISearchScopes uRISearchFilter	# fully_qualified_ldapuri
  | distinguishedName QMARK attributeBindTypeOrValue			# fully_qualified_ldapuri_attr_bindtype_or_value
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

// uRIAttributeList describes a list of zero (0) or more COMMA-delimited
// attributeType names to be requested for the search operation. Values
// of this kind appear immediately right of the distinguishedName.
uRIAttributeList
  : QMARK ( attributeTypeOrValue ( COMMA attributeTypeOrValue )* )?  # uriAttributeList
  ;

// distinguishedNames contains one (1) or more LDAP Distinguished Names. In the
// case of >1 DNs, the symbolic OR (||) delimiter is used. This applies to
// 'userdn', 'groupdn' and 'target' rules.
//
// e.g.:
//	- "ldap:///uid=courtney,ou=People,dc=example,dc=com" || "ldap:///uid=jane,ou=People,dc=example,dc=com"
//	- "ldap:///uid=courtney,ou=People,dc=example,dc=com || ldap:///uid=jane,ou=People,dc=example,dc=com"
//      - "ldap:///uid=jane,ou=People,dc=example,dc=com"
distinguishedNames
  : DQUOTE ( distinguishedName ( oRDelimiter distinguishedName )* ) DQUOTE             		# quoted_distinguished_name_list
  | ( DQUOTE distinguishedName DQUOTE ( oRDelimiter (DQUOTE distinguishedName DQUOTE) )* )	# list_of_quoted_distinguished_names
  ;

// distinguishedName is a sequence of aVAOrRDN values. Macro
// variable declarations for [$dn], ($dn) and ($attr.<atname>)
// are supported, as well as DN aliases for abstract contexts
// such as parent, self, etc.
//
// e.g.:
// 	- "ldap:///uid=courtney,ou=People,dc=example,dc=com"
//	- "ldap:///anyone"
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
//
// e.g.: "ldap:///uid=courtney,($attr.ou),ou=People,dc=example,dc=com"
rDNMacros
  : RDN_MACROS							   	# rdn_macro
  ;

// lDAPFilter describes the string representation of an RFC4515 LDAP Search Filter.
// Values of this kind are used in several areas, including 'targetfilter' Target Rules,
// LDAP URIs, et al.
//
// e.g.: "(&(objectClass=employee)(terminated=FALSE))"
lDAPFilter
  : LPAREN lDAPFilterExpr RPAREN       	# parenthetical_filter
  | lDAPFilterExpr*                    	# filter
  ;

// lDAPFilterExpr describes a single expressive LDAP filter statement, but may
// not be a complete filter unto itself.
lDAPFilterExpr
  : (LPAREN (FILTER_AND|FILTER_OR|FILTER_NOT)? lDAPFilterExpr RPAREN)+?    # parenthetical_filter_expression_opt_bool
  | <assoc=right> FILTER_NOT lDAPFilterExpr                 		   # not_filter_expression
  | aVAOrRDN			                               		   # ava_expression
  ;

// aVAOrRDN is one (1) of two (2) things for the sake of a simple grammar file:
//
// - When present within a distinguishedName, it is a relative distinguished name, e.g.:
//	- 'ou=Protocol'
//	- 'l=Palm Springs'
//	- 'cn=Courtney Tolana'
//
// - When present within a filter, it is an attribute value assertion, e.g.:
//	- 'color;lang-fr=bleu'
//	- 'objectClass=*'
//
// This is an absolutely critical parser component.
aVAOrRDN
  : LPAREN attributeTypeOrValue attributeComparisonOperator attributeTypeOrValue RPAREN	# parenthetical_ava_or_rdn
  | attributeTypeOrValue attributeComparisonOperator attributeTypeOrValue	 	# ava_or_rdn
  ;

// Vertical Inheritance (for User/Group Attribute Matching)
//
// e.g.: 'parent[0,1,3].owner#USERDN'
inheritance
  : ( PARENT inheritanceLevels DOT attributeBindTypeOrValue )       # inheritance_expression
  ;

// inheritanceLevels describe one (1) of five (5) possible "depth levels":
//
// - Level Zero (0) ("base" (baseObject) scope; zero depth)
// - Level One (1) (one level below baseObject; "onelevel" scope)
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
// e.g.: 'manager#GROUPDN' or 'nickname#squatcobbler'
//
// Values of this kind are used in certain 'userattr' and 'groupattr' Bind
// Rules, and will also be found within certain lDAPURI instances from time
// to time.
attributeBindTypeOrValue
  : attributeTypeOrValue HASH (bindTypes|attributeTypeOrValue)      # attr_bind_type_or_value
  ;

// bindTypes describes one (1) of five (5) possible BIND TYPES to be
// specified for certain 'userattr' and 'groupattr' Bind Rules.
bindTypes
  : userDN	# btUSERDN
  | groupDN	# btGROUPDN
  | roleDN	# btROLEDN
  | selfDN	# btSELFDN
  | lDAPURL	# btLDAPURL
  ;

userDN: BINDTYPE_USER_DN;
roleDN: BINDTYPE_ROLE_DN;
selfDN: BINDTYPE_SELF_DN;
groupDN: BINDTYPE_GROUP_DN;
lDAPURL: BINDTYPE_LDAP_URL;

// attributeTypeOrValue describes a general attributeType OR
// assertion value. Values of this kind MAY manifest as STAR
// for wildcard statements or presence checks.
//
// See the lexer KEY_OR_VALUE definition for additional notes
// on the topic of this value.
attributeTypeOrValue
  : KEY_OR_VALUE	# key_or_value
  | STAR		# presence_key_or_value
  ;

// attributeComparisonOperator describes one (1) of eight (8)
// possible comparison operators to be used in LDAP AVAs.
//
// Note that gt/lt are not valid operators for LDAP AVAs, and
// thus they are not listed here. Only ge/le are available for
// (numerical) ordering matches.
attributeComparisonOperator
  : equalTo		# equal_to
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
