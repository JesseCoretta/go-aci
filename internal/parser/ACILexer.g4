lexer grammar ACILexer;

WHSP
  : ' '+?
  ;

DQUOTE
  : '"'
  ;

LBRAK
  : '['
  ;

LPAREN
  : '('
  ;

RBRAK
  : ']'
  ;

RPAREN
  : ')'
  ;

DOT
  : '.'
  ;

COLON
  : ':'
  ;

TILDE
  : '~'
  ;

EQ
  : '='
  ;

NE
  : BANG EQ
  ;

GT
  : '>'
  ;

LT
  : '<'
  ;

APX
  : TILDE EQ
  ;

GE
  : GT EQ
  ;

LE
  : LT EQ
  ;

EXA
  : COLON EQ
  ;

EXO
  : COLON 'dn' COLON
  ;

EXD
  : COLON 'dn' COLON EQ
  ;

HASH
  : '#'
  ;

// Symbolic ANDs are used as delimiters in ANDed lists
SYMBOLIC_AND
  : AMPERSAND AMPERSAND
  ;

fragment AMPERSAND
  : '&'
  ;

// Symbolic ORs are used as delimiters in ORed lists
SYMBOLIC_OR
  : PIPE PIPE
  ;

fragment PIPE
  : '|'
  ;

fragment BANG
  : '!'
  ;

FILTER_AND
  : AMPERSAND
  ;

FILTER_OR
  : PIPE
  ;

FILTER_NOT
  : BANG
  ;

FILTER_OPERATOR
  : FILTER_AND
  | FILTER_OR
  | FILTER_NOT
  ;

COMMA
  : ','
  ;

SEMI
  : ';'
  ;

STAR
  : '*'
  ;

LOCAL_LDAP_SCHEME
  : 'ldap:///'
  ;

INHERITANCE_PREFIX
  : 'parent'
  ;

ANCHOR
  : 'version 3.0; acl '
  ;

SUNDAY
  : [Ss][Uu][Nn]
  ;

MONDAY
  : [Mm][Oo][Nn]
  ;

TUESDAY
  : [Tt][Uu][Ee][Ss]
  ;

WEDNESDAY
  : [Ww][Ee][Dd]
  ;

THURSDAY
  : [Tt][Hh][Uu][Rr]
  ;

FRIDAY
  : [Ff][Rr][Ii]
  ;

SATURDAY
  : [Ss][Aa][Tt]
  ;

ANONYMOUS
  : [Nn][Oo][Nn][Ee]
  ;
SIMPLE
  : [Ss][Ii][Mm][Pp][Ll][Ee]
  ;
SSL
  : [Ss][Ss][Ll]
  ;
SASL
  : [Ss][Aa][Ss][Ll]
  ;

// Target Rule keywords
TARGET
  : [Tt][Aa][Rr][Gg][Ee][Tt]
  ;

TARGET_TO
  : [Tt][Aa][Rr][Gg][Ee][Tt] '_' [Tt][Oo]
  ;

TARGET_FROM
  : [Tt][Aa][Rr][Gg][Ee][Tt] '_' [Ff][Rr][Oo][Mm]
  ;

TARGET_SCOPE
  : [Tt][Aa][Rr][Gg][Ee][Tt][Ss][Cc][Oo][Pp][Ee]
  ;

TARGET_ATTR
  : [Tt][Aa][Rr][Gg][Ee][Tt][Aa][Tt][Tt][Rr]
  ;

TARGET_FILTER
  : [Tt][Aa][Rr][Gg][Ee][Tt][Ff][Ii][Ll][Tt][Ee][Rr]
  ;

TARGET_ATTR_FILTERS
  : [Tt][Aa][Rr][Gg][Aa][Tt][Tt][Rr][Ff][Ii][Ll][Tt][Ee][Rr][Ss]
  ;

TARGET_CONTROL
  : [Tt][Aa][Rr][Gg][Ee][Tt][Cc][Oo][Nn][Tt][Rr][Oo][Ll]
  ;

TARGET_EXTENDED_OPERATION
  : [Ee][Xx][Tt][Oo][Pp]
  ;

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

BINDTYPE_USER_DN
  : 'USERDN'
  ;

BINDTYPE_GROUP_DN
  : 'GROUPDN'
  ;

BINDTYPE_ROLE_DN
  : 'ROLEDN'
  ;

BINDTYPE_SELF_DN
  : 'SELFDN'
  ;

BINDTYPE_LDAP_URL
  : 'LDAPURL'
  ;

// BASE is the same for a targetscope, or a regular scope
BASE_OBJECT_SCOPE
  : [Bb][Aa][Ss][Ee]
  ;

ONE_LEVEL_SCOPE
  : [Oo][Nn][Ee]
  ;

ONE_LEVEL_TARGET_SCOPE
  : [Oo][Nn][Ee][Ll][Ee][Vv][Ee][Ll]
  ;

SUB_TREE_SCOPE
  : [Ss][Uu][Bb]
  ;

SUB_TREE_TARGET_SCOPE
  : [Ss][Uu][Bb][Tt][Rr][Ee][Ee]
  ;

SUBORDINATE_TARGET_SCOPE
  : [Ss][Uu][Bb][Oo][Rr][Dd][Ii][Nn][Aa][Tt][Ee]
  ;

ALLOW_ACCESS
  : WHSP? [Aa][Ll][Ll][Oo][Ww] WHSP?
  ;

DENY_ACCESS
  : [Dd][Ee][Nn][Yy]
  ;

SEARCH_PRIVILEGE
  : [Ss][Ee][Aa][Rr][Cc][Hh]
  ;

READ_PRIVILEGE
  : [Rr][Ee][Aa][Dd]
  ;

COMPARE_PRIVILEGE
  : [Cc][Oo][Mm][Pp][Aa][Rr][Ee]
  ;

ADD_PRIVILEGE
  : [Aa][Dd][Dd]
  ;

DELETE_PRIVILEGE
  : [Dd][Ee][Ll][Ee][Tt][Ee]
  ;

SELFWRITE_PRIVILEGE
  : [Ss][Ee][Ll][Ff][Ww][Rr][Ii][Tt][Ee]
  ;

PROXY_PRIVILEGE
  : [Pp][Rr][Oo][Xx][Yy]
  ;

IMPORT_PRIVILEGE
  : [Ii][Mm][Pp][Oo][Rr][Tt]
  ;

EXPORT_PRIVILEGE
  : [Ee][Xx][Pp][Oo][Rr][Tt]
  ;

// ALL_PRIVILEGES defines all possible privileges
// **EXCEPT** PROXY_PRIVILEGE.
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
// ... might be used to expand ou=Contractors, etc.
RDN_MACROS
  : '[$dn]'
  | '($dn)'
  | '($attr.' KEY_OR_VALUE ')'
  ;

// AND defines statements that mandate all conditions
// evaluate as true.
BOOLEAN_AND
  : [Aa][Nn][Dd]
  ;

// OR defines statements that mandate at least one
// condition evaluates as true.
BOOLEAN_OR
  : [Oo][Rr]
  ;

// NOT defines statements that negate otherwise
// matchable values. NOT is special, is right
// associated and MUST include a space between
// AND and NOT.
BOOLEAN_NOT
  : [Aa][Nn][Dd] ' ' [Nn][Oo][Tt]
  ;

// Whitespace characters are dumped from
// here on out.  I know this is supposed
// to be at the bottom of the lexer file,
// but all hell breaks loose when it is.
WHITESPACE
  : [ \t\r\n\u000C]+ -> skip
  ;

INT
  : [0-9]+
  ;

// KEY_OR_VALUE can more or less be anything,
// but will be verified in the Go visitor.
//
// I really wish I could split this into two
// lexers that WON'T collide, but I've given
// up on that hopeless cause. To be honest,
// I'm not even sure THIS is right.
KEY_OR_VALUE
  : ~["\\,.:=![\]()#|&<>~\t\r\n]+
  ;

