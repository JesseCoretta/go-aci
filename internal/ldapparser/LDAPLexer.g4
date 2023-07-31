/*
Basic LDAP Filter, DN and URI LEXER Grammar for ANTLR4 (4.13.0)

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

This ANTLRv4 (4.13.0) grammar file contains basic lexer rules for LDAP Search
URIs and their components: Search Filters, Scopes, Attributes and Distinguished
Names.

PARSER GRAMMAR

See the accompanying LDAPParser.g4 file for parser grammar rules.
*/

lexer grammar LDAPLexer;

LPAREN          : '('              	 ;
RPAREN          : ')'              	 ;
COMMA		: ','			 ;
QMARK		: '?'			 ;
ASTERISK	: '*'			 ;
DN		: 'dn'			 ;
WHITESPACE	: ' '			 ;

fragment LBRAK  : '['              	 ;
fragment RBRAK  : ']'              	 ;
fragment DOLLAR	: '$'		   	 ;
fragment SOLIDUS: '/'		   	 ;
fragment DOT	: '.'		   	 ;
fragment TILDE	: '~'                	 ;
fragment EQ	: '='                  	 ;
fragment EX	: ':'		 	 ;
fragment APX	: TILDE EQ             	 ;
fragment GE	: '>' EQ               	 ;
fragment LE	: '<' EQ               	 ;
fragment ATTR	: 'attr'		 ;
fragment LDAP	: 'ldap'		 ;

BASEOBJECT_SCOPE
  : [Bb][Aa][Ss][Ee]
  ;

SINGLELEVEL_SCOPE
  : [Oo][Nn][Ee]
  ;

WHOLESUBTREE_SCOPE
  : [Ss][Uu][Bb]
  ;

LOCAL_LDAP_SCHEME
  : LDAP ('s'|'i')? EX SOLIDUS SOLIDUS ( (LBRAK|'%2f')? IDENTIFIER ( (DOT|EX|'%2f') IDENTIFIER )* RBRAK? (EX IDENTIFIER)? )? SOLIDUS
  ;

fragment EXCLAMATION
  : '!'
  ;
NOT: EXCLAMATION;

fragment AMPERSAND
  : '&'
  ;
AND: AMPERSAND;

fragment VERTICAL
  : '|'
  ;
OR: VERTICAL;

OBJECT_IDENTIFIER
  : ( [0-9]+ ( DOT [0-9]+ )+ )
  ;

IDENTIFIER
  : [a-zA-Z0-9\-;]+
  ;

EQUALITY
  : EQ
  ;

APPROXIMATE
  : APX
  ;

LESS_THAN_OR_EQUAL
  : LE
  ;

GREATER_THAN_OR_EQUAL
  : GE
  ;

EXTENSIBLE_RULE
  : EX
  ;

EXCLUSIONS
  : [()\t\r\n] -> skip
  ;
