lexer grammar ACILexer;

LEVELS: [0-4]+;

INT
 : DecimalDigit+
 ;

ID
 : Identifier+
 ;

COLON: ':';
ExtensibleRuleDNMatch: COLON DN COLON EqualTo;
ExtensibleRuleDNOIDMatch: COLON DN COLON;
ExtensibleRuleAttrMatch: COLON EqualTo;
ExtensibleRuleMatch: COLON;
ApproximateMatch: '~=';
NotEqualTo: '!=';
GreaterThanOrEqual: '>=';
LessThanOrEqual: '<=';

DAMP:  AMP AMP;
DPIPE: PIPE PIPE ;

// Boolean operator words
BooleanWordOperators
    : BooleanAndWord
    | BooleanOrWord
    | BooleanNotWord
    ;

BooleanAndWord
    : 'AND'
    | 'and'
    ;

BooleanOrWord
    : 'OR'
    | 'or'
    ;

BooleanNotWord
    : 'AND NOT'
    | 'and not'
    ;

DAYS: 'Sun'
    | 'Mon'
    | 'Tues'
    | 'Wed'
    | 'Thur'
    | 'Fri'
    | 'Sat'
    ;

// Target Rule keywords
TargetKeyword: 'target';
TargetScopeKeyword: 'targetscope';
TargetToKeyword: 'target_to';
TargetFromKeyword: 'target_from';
TargetAttrKeyword: 'targetattr';
TargetFilterKeyword: 'targetfilter';
TargetAttrFiltersKeyword: 'targattrfilters';
TargetControlKeyword: 'targetcontrol';
ExtOpKeyword: 'extop';

ANONYMOUS: 'none';
SIMPLE: 'simple';
SASL: 'SASL';
SSL: 'SSL';

// Bind Rule keywords
RoleDNKeyword: 'roledn';
UserDNKeyword: 'userdn';
GroupDNKeyword: 'groupdn';
UserAttrKeyword: 'userattr';
GroupAttrKeyword: 'groupattr';
SSFKeyword: 'ssf';
DNSKeyword: 'dns';
IPKeyword: 'ip';
AuthMethodKeyword: 'authmethod';
TimeOfDayKeyword: 'timeofday';
DayOfWeekKeyword: 'dayofweek';

DISPOSITION: 'allow' | 'deny' ;
RIGHTS: 'search'
     |  'read'
     |  'compare'
     |  'add'
     |  'delete'
     |  'selfwrite'
     |  'proxy'
     |  'import'
     |  'export'
     |  'all'
     ;

AttributeFilterOperation
     : AddOperation
     | DeleteOperation
     ;

AddOperation:    'add' EqualTo;
DeleteOperation: 'delete' EqualTo;

TargetRuleSearchScopes
     : '"base"'
     | '"onelevel"'
     | '"subtree"'
     | '"subordinate"'
     ;

LDAPSearchScopes
     : 'base'
     | 'one'
     | 'sub'
     ;

BINDTYPES
     : 'USERDN'
     | 'GROUPDN'
     | 'ROLEDN'
     | 'SELFDN'
     | 'LDAPURL'
     ;

INHERITANCEPREFIX: 'parent[';

LocalLDAPScheme: LDAP COLON SOLIDUS SOLIDUS SOLIDUS;

// Miscellaneous literals/characters used frequently
SEMI: ';';
EqualTo: '=';
GreaterThan: '>';
LessThan: '<';
LDAP: 'ldap';
ANCHOR: 'version 3.0; acl ' ;
DN: 'dn';
AMP:  '&' ;
PIPE: '|' ;
BANG: '!' ;
SOLIDUS: '/';
LPAREN: '(' ;
LBRAC: '{' ;
LBRAK: '[' ;
RPAREN: ')' ;
RBRAK: ']' ;
RBRAC: '}' ;
DQUOTE: '"' ;
COMMA: ',';
TILDE: '~';
HASH: '#';
DOLLAR: '$';
ATSIGN: '@';
DOT: '.';
DASH: '-';
STAR: '*';
QMARK: '?';

LineTerminator
 : [\r\n\u2028\u2029] -> channel(HIDDEN)
 ;

WhiteSpaces
 : [\t\u000B\u000C\u0020\u00A0]+ -> channel(HIDDEN)
 ;

NumericLiteral
 : DecimalDigit
 | HexDigit
 ;

Literal
 : StringLiteral
 | NumericLiteral
 ;

DelimitedAddress
 : ( (ID|'*') ( COLON (ID|'*')* )+ )
 | ( (ID|'*') ( DOT (ID|'*')* )+ )
 ;

DelimitedNumbers
 : ( INT (COMMA INT)* )+
 ;

StringLiteral
 : '"' DoubleStringCharacter* '"'
 | '\'' SingleStringCharacter* '\''
 ;

WildcardString
 : (Letter|DecimalDigit|'*')+
 ;

MacroValue
 : '[$dn]'
 | '($dn)'
 | '($attr.' Identifier+ ')'
 ;

fragment DoubleStringCharacter
 : ~["\\\r\n]
 | '\\' EscapeSequence
 | LineContinuation
 ;

fragment SingleStringCharacter
 : ~['\\\r\n]
 | '\\' EscapeSequence
 | LineContinuation
 ;

fragment EscapeSequence
 : CharacterEscapeSequence
 | '0' // no digit ahead! TODO
 | HexEscapeSequence
 | UnicodeEscapeSequence
 ;

fragment CharacterEscapeSequence
 : SingleEscapeCharacter
 | NonEscapeCharacter
 ;

fragment HexEscapeSequence
 : 'x' HexDigit HexDigit
 ;

fragment UnicodeEscapeSequence
 : 'u' HexDigit HexDigit HexDigit HexDigit
 ;

fragment SingleEscapeCharacter
 : ['"\\bfnrtv]
 ;

fragment NonEscapeCharacter
 : ~['"\\bfnrtv0-9xu\r\n]
 ;

fragment EscapeCharacter
 : SingleEscapeCharacter
 | DecimalDigit
 | [xu]
 ;

fragment LineContinuation
 : '\\' LineTerminatorSequence
 ;

fragment LineTerminatorSequence
 : '\r\n'
 | LineTerminator
 ;

fragment Identifier
 : Letter
 | DecimalDigit
 | DASH
 ;

fragment Letter
 : LowercaseLetter
 | UppercaseLetter
 ;

fragment LowercaseLetter
 : [a-z]
 ;

fragment UppercaseLetter
 : [A-Z]
 ;

fragment DecimalDigit
 : [0-9]
 ;

fragment HexDigit
 : [0-9a-fA-F]
 ;

//fragment Value
// :  ~[\r\n,"]+
// ;

ANY: .;
