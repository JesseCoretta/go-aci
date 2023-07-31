/*
ACIv3 Lexer Grammar

Implemented by Jesse Coretta ▲

ADVISORY

This is an initial release and is potentially unsuitable for
mission-critical / production environments. At the moment, it
should only be used as a convenient supplement to an already
hardened ACI review/approval/documentation/etc., process. Use
at your own risk.

See further down for LICENSE details.

ABOUT THIS FILE

This ANTLRv4 (4.13.0) lexer grammar implements lexer support for
Version 3.0 of the Access Control Instruction syntax specification
and all of its abstract components. See below for LICENSE details.

This file is sourced via the main ACIParser.g4 grammar file. See
further details there.

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

lexer grammar ACILexer;

BKW_UDN		: 'userdn'		;
BKW_GDN		: 'groupdn'		;
BKW_RDN		: 'roledn'		;
BKW_UAT		: 'userattr'		;
BKW_GAT		: 'groupattr'		;
BKW_SSF		: 'ssf'			;
BKW_IP		: 'ip'			;
BKW_DNS		: 'dns'			;
BKW_AM		: 'authmethod'		;
BKW_TOD		: 'timeofday'		;
BKW_DOW		: 'dayofweek'		;

TKW_TARGET	: 'target'		;
TKW_TO		: 'target_to'		;
TKW_FROM	: 'target_from'		;
TKW_SCOPE	: 'targetscope'		;
TKW_ATTR	: 'targetattr'		;
TKW_FILTER	: 'targetfilter'	;
TKW_AF		: 'targattrfilters'	;
TKW_CTRL	: 'targetcontrol'	;
TKW_EXTOP	: 'extop'		;

PARENT		: 'parent'		;
DQUOTE		: '"'			;
LPAREN		: '('			;
RPAREN		: ')'			;
COMMA		: ','			;
SEMI		: ';'			;
EQ		: '='			;
NE		: '!='			;
LT		: '<'			;
LE		: '<='			;
GT		: '>'			;
GE		: '>='			;

// Symbolic ANDs (&&) are used as delimiter literals within
// ANDed attributeFilterSet instances.
SYMBOLIC_AND: '&&';

// Symbolic ORs (||) are used as delimiter literals within
// ORed lists of attributeTypes, objectIdentifiers and 
// distinguishedNames.
SYMBOLIC_OR: '||';
fragment BANG: '!';

// The "anchor" is a string literal that will always
// appear identical within an ACI (as shown), and acts
// as a suitable starting point for processing and basic
// validation. 
//
// The anchor is preceded by zero (0) or more Target Rules
// and the trailing space (which is intended) is followed
// by a brief, double-quoted, and DIT-unique descriptive
// string label of (ideally) limited ASCII diversity.
ANCHOR
  : 'version 3.0' SEMI SPACE 'acl'
  ;

//////////////////////////////////////
// Target Rule components

//////////////////////////////////////
// Permission and Access Rights components

// Permission is a single statement that defines one or
// more privilege specifiers (e.g.: write, read) along
// with a single disposition (allow, deny). Specifiers
// are always parenthetical, whether singular or not.

READ_PRIV	: 'read'	;
WRITE_PRIV	: 'write'	;
ADD_PRIV	: 'add'		;
CMP_PRIV	: 'compare'	;
SRC_PRIV	: 'search'	;
DEL_PRIV	: 'delete'	;
PRX_PRIV	: 'proxy'	;
EXP_PRIV	: 'export'	;
IMP_PRIV	: 'import'	;
SLF_PRIV	: 'selfwrite'	;
ALL_PRIV	: 'all'		;
NO_PRIV		: 'none'	;

ALLOW:'allow';
DENY: 'deny';

// AND defines statements that mandate all specified
// Bind Rule conditions evaluate as true.
WORD_AND
  : 'AND'
  | 'and'
  ;

// OR defines statements that mandate at least one of
// the specified Bind Rule conditions evaluates as true.
WORD_OR
  : 'OR'
  | 'or'
  ;

// NOT defines Bind Rule statements that negate otherwise
// matchable values. NOT is special, is right-associated
// and MUST include a space between AND and NOT (the 'NOT'
// literal is never used alone in this particular syntax).
WORD_NOT
  : 'AND NOT'
  | 'and not'
  ;

NCTF_TO_WHSP
  : ( '\t' | '\u000C' | ( '\r\n' | '\n' ) )+ -> skip
  ;

QUOTED_STRING
  : DQUOTE  ~["\t\r\n\u000C]+ DQUOTE
  ;

WS
  : [ \t]+ -> skip
  ;
fragment SPACE  : ' '			;

