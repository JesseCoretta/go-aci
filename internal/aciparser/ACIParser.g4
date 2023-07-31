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

CONTRIBUTIONS

If you believe this solution lacks a certain "syntactical sugar" of
which I am unaware (and you can cite literature to that end), then
you are encouraged to open a new ticket within the github repository
in which the parser resides.

LIMITATIONS

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
// lexer grammar into its own file, which we source here.
options { tokenVocab=ACILexer; }

// instruction is the main parsing target, which is comprised of
// many "constituents" -- all of which are defined later in this
// grammar file (and its accompanying lexer file).
parse
  : instruction EOF
  ;

instruction
  : targetRules? LPAREN ANCHOR quotedValue SEMI permissionBindRule+ RPAREN
  ;

permissionBindRule
  : permission bindRules SEMI
  ;

permission
  : (ALLOW|DENY) openingParenthesis privilege ( COMMA privilege )* closingParenthesis
  ;

privilege
  : READ_PRIV
  | WRITE_PRIV
  | SLF_PRIV
  | CMP_PRIV
  | SRC_PRIV
  | PRX_PRIV
  | ADD_PRIV
  | DEL_PRIV
  | IMP_PRIV
  | EXP_PRIV
  | ALL_PRIV
  | NO_PRIV
  ;

targetRule
  : openingParenthesis targetKeyword targetOperator quotedValue closingParenthesis
  ;

targetOperator
  : equalTo
  | notEqualTo
  ;

targetRules: targetRule+ ;

/*
//bindRule
//  : bindKeyword compOperator quotedValue
//  | openingParenthesis bindKeyword compOperator quotedValue closingParenthesis
//  ;
//
//bindRules
//  : bindRule
//  | ( bindRule ( wordOperator bindRule )* )
//  | openingParenthesis bindRule+ closingParenthesis
//  ;
*/

bindRule
  : bindKeyword bindOperator quotedValue
  | openingParenthesis bindKeyword bindOperator quotedValue closingParenthesis
  ;

wordAnd: WORD_AND;
wordOr: WORD_OR;
wordNot: WORD_NOT;

/*
bindRules defines one or more bindRule expressions that shall reside within
the bindRules field of a permissionBindRule "pair".
*/
bindRules
  : openingParenthesis bindRules closingParenthesis
  | <assoc=right> bindRules wordNot bindRules
  | bindRule ( ( wordAnd | wordOr ) bindRule )+
  | bindRule
  ;

quotedValue: QUOTED_STRING;

bindOperator
  : equalTo
  | notEqualTo
  | lessThan
  | lessThanOrEqual
  | greaterThan
  | greaterThanOrEqual
  ;

lessThan		: LT	;
lessThanOrEqual		: LE	;
greaterThan		: GT	;
greaterThanOrEqual	: GE	;
equalTo			: EQ	;
notEqualTo		: NE	;

/*
bindKeyword values are used by any given bindRule as the
first field in a bind rule expression ( kw op val ).
*/
bindKeyword
  : BKW_UDN
  | BKW_GDN
  | BKW_RDN
  | BKW_UAT
  | BKW_GAT
  | BKW_SSF
  | BKW_IP
  | BKW_DNS
  | BKW_DOW
  | BKW_TOD
  | BKW_AM
  ;

/*
targetKeyword values are used by any given targetRule as the
first field in a target rule expression ( kw op val ).
*/
targetKeyword
  : TKW_TARGET
  | TKW_TO
  | TKW_FROM
  | TKW_SCOPE
  | TKW_ATTR
  | TKW_FILTER
  | TKW_AF
  | TKW_CTRL
  | TKW_EXTOP
  ;

openingParenthesis: LPAREN;
closingParenthesis: RPAREN;
