/*
ACIv3 Lexer Grammar

Author: Jesse Coretta ▲
Date:   07/06/2023

This ANTLRv4 (4.13.0) lexer grammar implements lexer support for
Version 3.0 of the Access Control Instruction syntax specification
and all of its abstract components. See below for LICENSE details.

This file is sourced via the main ACIParser.g4 grammar file.

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

lexer grammar ACILexer;

WHSP: ' '+?;
QMARK: '?';
DQUOTE: '"';
LBRAK: '[';
LPAREN: '(';
RBRAK: ']';
RPAREN: ')';
DOT: '.';
COLON: ':';
TILDE: '~';
EQ: '=';
NE: BANG EQ;
GT: '>';
LT: '<';
APX: TILDE EQ;
GE: GT EQ;
LE: LT EQ;
EXA: COLON EQ;
EXO: COLON 'dn' COLON;
EXD: COLON 'dn' COLON EQ;
HASH: '#';

// Symbolic ANDs (&&) are used as delimiter literals within
// ANDed attributeFilterSet instances.
SYMBOLIC_AND: AMPERSAND AMPERSAND;

fragment AMPERSAND: '&';

// Symbolic ORs (||) are used as delimiter literals within
// ORed lists of attributeTypes, objectIdentifiers and 
// distinguishedNames.
SYMBOLIC_OR: PIPE PIPE;

fragment PIPE: '|';
fragment BANG: '!';

// FILTER_OR (&) describes two (2) or more ANDed attribute
// value assertion statements. All of the attribute value
// assertion statements must evaluate as true. These are
// only used within lDAPFilter instances.
FILTER_AND: AMPERSAND;

// FILTER_OR (|) describes two (2) or more ORed attribute
// value assertion statements. At least one (1) of the
// statements must evaluate as true. These are only used
// within lDAPFilter instances.
FILTER_OR: PIPE;

// FILTER_NOT (!) defines statements that negate otherwise
// matchable attribute value assertion statements. FILTER_NOT
// is right-associated, but will be read in the order input
// by the user. It is up to the DSA to extract and evaluate
// negated conditions prior to ANDs and/or ORs, NOT the DUA
// or client SDK.
FILTER_NOT: BANG;

// FILTER_OPERATOR defines all of the possible Boolean
// operators that may influence an LDAP Search Filter.
FILTER_OPERATOR
  : FILTER_AND
  | FILTER_OR
  | FILTER_NOT
  ;

COMMA: ',';
SEMI: ';';
STAR: '*';

// The so-called "Local LDAP scheme" (ldap:///) forces
// the context of any distinguishedName to be implicitly
// local (meaning one cannot define "proxy-leveraging"
// ACIs by including a hostname/port after the second
// solidus character (/)). This is a design decision
// common to multiple implementations that leverage the
// ACIv3 syntax and is mainly due to security reasons.
//
// If you try to override this, you're probably out of
// your mind and you need to be stopped.
LOCAL_LDAP_SCHEME: 'ldap:///';

// 'parent' is a convenient DN "alias" supported
// by the DSA to describe the immediate superior
// entry of any given entry. This string literal
// is also used within inheritance 'userattr' and
// 'groupattr' Bind Rule statements. In any scenario,
// it is prefixed with the above Local LDAP Scheme
// string literal.
PARENT
  : [Pp][Aa][Rr][Ee][Nn][Tt]
  ;

// 'parent' is a convenient DN "alias" supported
// by the DSA to describe any user DN -- whether
// authenticated (known) or anonymous (null). It
// is prefixed with the above Local LDAP Scheme
// string literal.
ANYONE
  : [Aa][Nn][Yy][Oo][Nn][Ee]
  ;

// 'all' is a convenient DN "alias" supported
// by the DSA to describe any authenticated (known)
// user DN. This does NOT include anonymous. It
// is prefixed with the above Local LDAP Scheme
// string literal.
ALL_USERS
  : [Aa][Ll][Ll]
  ;

// 'self' is a convenient DN "alias" supported by
// the DA to describe ones own DN or entry. It is
// prefixed with the above Local LDAP Scheme string
// literal.
SELF
  : [Ss][Ee][Ll][Ff]
  ;

// The "anchor" is a string literal that will always
// appear identical within an ACI (as shown), and acts
// as a suitable starting point for processing and basic
// validation.
//
// The anchor is preceded by zero (0) or more Target Rules
// and followed by one (1) or more Permission+Bind statements.
ANCHOR
  : 'version 3.0; acl '
  ;

// Sunday is day one (1).
SUNDAY
  : [Ss][Uu][Nn]
  ;

// Monday is day two (2).
MONDAY
  : [Mm][Oo][Nn]
  ;

// Tuesday is day three (3).
TUESDAY
  : [Tt][Uu][Ee][Ss]
  ;

// Wednesday is day four (4).
WEDNESDAY
  : [Ww][Ee][Dd]
  ;

// Thursday is day five (5).
THURSDAY
  : [Tt][Hh][Uu][Rr]
  ;

// Friday is day six (6).
FRIDAY
  : [Ff][Rr][Ii]
  ;

// Saturday is day seven (7).
SATURDAY
  : [Ss][Aa][Tt]
  ;

// 'none' describes an ANONYMOUS LDAP bind.
ANONYMOUS
  : [Nn][Oo][Nn][Ee]
  ;

// 'simple' describes an authenticated LDAP bind
// using weak authentication (DN + clear-text).
SIMPLE
  : [Ss][Ii][Mm][Pp][Ll][Ee]
  ;

// 'ssl' describes an authenticated LDAP bind 
// using weak authentication (DN + clear-text)
// using TLS confidentiality.
SSL
  : [Ss][Ss][Ll]
  ;

// 'sasl' describes an authenticated LDAP bind
// using strong authentication (TLS mutual auth,
// Kerberos, et al) and (almost certainly) using
// TLS confidentiality.
SASL
  : [Ss][Aa][Ss][Ll]
  ;

//////////////////////////////////////
// Target Rule keywords

// 'target' keyword
TARGET
  : [Tt][Aa][Rr][Gg][Ee][Tt]
  ;

// 'target_to' keyword
TARGET_TO
  : [Tt][Aa][Rr][Gg][Ee][Tt] '_' [Tt][Oo]
  ;

// 'target_from' keyword
TARGET_FROM
  : [Tt][Aa][Rr][Gg][Ee][Tt] '_' [Ff][Rr][Oo][Mm]
  ;

// 'targetscope' keyword
TARGET_SCOPE
  : [Tt][Aa][Rr][Gg][Ee][Tt][Ss][Cc][Oo][Pp][Ee]
  ;

// 'targetattr' keyword
TARGET_ATTR
  : [Tt][Aa][Rr][Gg][Ee][Tt][Aa][Tt][Tt][Rr]
  ;

// 'targetfilter' keyword
TARGET_FILTER
  : [Tt][Aa][Rr][Gg][Ee][Tt][Ff][Ii][Ll][Tt][Ee][Rr]
  ;

// 'targattrfilters' keyword
TARGET_ATTR_FILTERS
  : [Tt][Aa][Rr][Gg][Aa][Tt][Tt][Rr][Ff][Ii][Ll][Tt][Ee][Rr][Ss]
  ;

// 'targetcontrol' keyword
TARGET_CONTROL
  : [Tt][Aa][Rr][Gg][Ee][Tt][Cc][Oo][Nn][Tt][Rr][Oo][Ll]
  ;

// 'extop' keyword
TARGET_EXTENDED_OPERATION
  : [Ee][Xx][Tt][Oo][Pp]
  ;

// TODO - WHSP shouldn't need to be specified manually
// fix me plz 
BIND_USER_DN
  : WHSP? 'userdn' WHSP?
  ;

BIND_GROUP_DN
  : 'groupdn'
  ;

BIND_ROLE_DN
  : 'roledn'
  ;

BIND_USER_ATTR
  : 'userattr'
  ;

BIND_GROUP_ATTR
  : 'groupattr'
  ;

BIND_SSF
  : 'ssf'
  ;

BIND_DNS
  : 'dns'
  ;

BIND_IP
  : 'ip'
  ;

BIND_AUTH_METHOD
  : 'authmethod'
  ;

BIND_TIME_OF_DAY
  : 'timeofday'
  ;

BIND_DAY_OF_WEEK
  : 'dayofweek'
  ;

// USERDN string literal is used within 'userattr' and 'groupattr'
// Bind Rule statements.
BINDTYPE_USER_DN
  : 'USERDN'
  ;

// GROUPDN string literal is used within 'userattr' and 'groupattr'
// Bind Rule statements.
BINDTYPE_GROUP_DN
  : 'GROUPDN'
  ;

// ROLEDN string literal is used within 'userattr' and 'groupattr'
// Bind Rule statements.
BINDTYPE_ROLE_DN
  : 'ROLEDN'
  ;

// SELFDN string literal is used within 'userattr' and 'groupattr'
// Bind Rule statements.
BINDTYPE_SELF_DN
  : 'SELFDN'
  ;

// LDAPURL string literal is used within 'userattr' and 'groupattr'
// Bind Rule statements.
BINDTYPE_LDAP_URL
  : 'LDAPURL'
  ;

// BASE is the same for 'targetscope' Target Rules as for lDAPURI
// search parameters and is used the same in either scenario.
BASE_OBJECT_SCOPE
  : [Bb][Aa][Ss][Ee]
  ;

// This is used exclusively within LDAP Search Parameter statements,
// such as those that appear within an lDAPURI. This is not used
// within 'targetscope' Target Rules.
ONE_LEVEL_SCOPE
  : [Oo][Nn][Ee]
  ;

// This is used exclusively within 'targetscope' Target
// Rules and NOT lDAPURI instances.
ONE_LEVEL_TARGET_SCOPE
  : [Oo][Nn][Ee][Ll][Ee][Vv][Ee][Ll]
  ;

// This is used exclusively within LDAP Search Parameter statements,
// such as those that appear within an lDAPURI. This is not used
// within 'targetscope' Target Rules.
SUB_TREE_SCOPE
  : [Ss][Uu][Bb]
  ;

// This is used exclusively within 'targetscope' Target
// Rules and NOT lDAPURI instances.
SUB_TREE_TARGET_SCOPE
  : [Ss][Uu][Bb][Tt][Rr][Ee][Ee]
  ;

// This is used exclusively within 'targetscope' Target
// Rules and NOT lDAPURI instances.
SUBORDINATE_TARGET_SCOPE
  : [Ss][Uu][Bb][Oo][Rr][Dd][Ii][Nn][Aa][Tt][Ee]
  ;

// The disposition of a permission is to grant some level(s)
// of access to the directory.
ALLOW_ACCESS
  : WHSP? [Aa][Ll][Ll][Oo][Ww] WHSP?
  ;

// The disposition of a permission is to deny some level(s)
// of access to the directory.
DENY_ACCESS
  : [Dd][Ee][Nn][Yy]
  ;

// Grant or withhold LDAP search access to the DSA.
SEARCH_PRIVILEGE
  : [Ss][Ee][Aa][Rr][Cc][Hh]
  ;

// Grant or withhold LDAP read access to the DSA.
READ_PRIVILEGE
  : [Rr][Ee][Aa][Dd]
  ;

// Grant or withhold LDAP compare access to the DSA.
COMPARE_PRIVILEGE
  : [Cc][Oo][Mm][Pp][Aa][Rr][Ee]
  ;

// Grant or withhold LDAP entry-creation access to the DSA.
ADD_PRIVILEGE
  : [Aa][Dd][Dd]
  ;

// Grant or withhold LDAP entry-deletion access to the DSA.
DELETE_PRIVILEGE
  : [Dd][Ee][Ll][Ee][Tt][Ee]
  ;

// Grant or withhold LDAP modifications to ones own entry within the DSA.
SELFWRITE_PRIVILEGE
  : [Ss][Ee][Ll][Ff][Ww][Rr][Ii][Tt][Ee]
  ;

// Grant or withhold LDAP remote proxy capabilities within the DSA.
PROXY_PRIVILEGE
  : [Pp][Rr][Oo][Xx][Yy]
  ;

// Grant or withhold LDAP DIT import capabilities within the DSA.
IMPORT_PRIVILEGE
  : [Ii][Mm][Pp][Oo][Rr][Tt]
  ;

// Grant or withhold LDAP DIT export capabilities within the DSA.
EXPORT_PRIVILEGE
  : [Ee][Xx][Pp][Oo][Rr][Tt]
  ;

// Grant or withhold all privileges within the DSA **EXCEPT** for
// proxy privileges.
ALL_PRIVILEGES
  : [Aa][Ll][Ll]
  ;

// Certain directory implementations allow the use
// of macro statements within distinguished names
// to allow extended flexibility in terms of value
// matching in ACIs. For instance:
//
//   [$dn],ou=People,dc=example,dc=com
//
// ... might be used to expand ou=Contractors, and
// ou=Accounting, etc.
//
// Please note these are string literals and users
// should not expect to see any interpolation (that
// is what the DSA does, NOT the DUA or client SDK).
RDN_MACROS
  : '[$dn]'
  | '($dn)'
  | '($attr' DOT KEY_OR_VALUE ')'
  ;

// AND defines statements that mandate all specified
// Bind Rule conditions evaluate as true.
BOOLEAN_AND
  : [Aa][Nn][Dd]
  ;

// OR defines statements that mandate at least one of
// the specified Bind Rule conditions evaluates as true.
BOOLEAN_OR
  : [Oo][Rr]
  ;

// NOT defines Bind Rule statements that negate otherwise
// matchable values. NOT is special, is right-associated
// and MUST include a space between AND and NOT (the 'NOT'
// literal is never used alone in this particular syntax).
BOOLEAN_NOT
  : [Aa][Nn][Dd] ' ' [Nn][Oo][Tt]
  ;

// Whitespace characters are dumped from here on out. I
// know this is supposed to be at the bottom of the lexer
// file, but all hell breaks loose when it is :(
WHITESPACE
  : [ \t\r\n\u000C]+ -> skip
  ;

// INT represents any unsigned integer of any magnitude.
// At no point in this solution are negative integers used.
INT
  : [0-9]+
  ;

// KEY_OR_VALUE can more or less be anything, but will be
// verified in the Go visitor.
//
// This is used in a variety of areas - most importantly
// within aVAOrRDN instances - to describe a key/value
// statement (which is something you see VERY OFTEN in the
// LDAP world). This was particularly tricky to implement
// due to the extensive comparison operators that must be
// supported -- operators that go well beyond the typical
// eq, ge, lt operators -- such as ':=', ':dn:=', et al.
//
// I REALLY wish I could split this into two (2) lexers that
// WON'T collide, e.g.:
//
// - KEY:   [a-z][a-zA-z0-9\-]* [a-z]*
//
//    ... and ...
//
// - VALUE: ~["\\,.:=!?[\]()#|&<>~\t\r\n]+
//
// ... but I've given up on that for the moment. Every attempt
// to do wreaks havoc within this otherwise functional setup.
// 
// The (negated!) characters below are specified due to their
// special nature elsewhere in this implementation, i.e.: '&'
// in Boolean lists, and (probably?) shouldn't appear in values
// such as the 'acl' (ACI label).
//
// To be honest, I'm quite sure this is NOT an ideal solution
// (likely will barf on certain otherwise harmless characters
// in a value), but it DOES seem to work for the moment ...
KEY_OR_VALUE
  : ~["\\,.:=!?[\]()#|&<>~\t\r\n]+
  ;

