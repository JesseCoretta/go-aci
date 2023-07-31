/*
Package ldapparser implements an ANTLR4 Lexer/Parser framework to support the parsing
(marshaling) of LDAP Uniform Resource Identifiers and their constituent components:

- Distinguished Names

- Search Attributes

- Search Scopes

- Search Filters

# Advisory

The ldapparser package (and the associated ANTLR4 grammar) is currently in its initial
stages of development. Use with caution and report any bugs.

# Internal Package

This package is internal and should not be accessed by users directly. The top-level
go-aci package calls the ldapparser package as needed during ACI marshaling. It is 
not intended for use in any other scenario by any person or entity.

# License

The ldapparser package is released as a component in the go-aci suite under the
terms of the MIT License. For full details, see the LICENSE file within the package
repository.

# Filter Features and Limitations

While effective at "decompiling" or "dissecting" an LDAP Search Filter for use
in writing a parser, it should not be considered a complete "parsing" solution.
Certain syntactically-bogus LDAP filters -- such as those missing encapsulation
characters (i.e.: closing parenthesis) -- may be processed sans any noticeable
error output.

That said, this grammar file can aid in said deconstruction of filter statements
as a parsing supplement.

All of the standard attributeAssertionValue "operators" are supported:

  - Equality (=)
  - Approximate (~=)
  - Ordering (>=/<=)
  - ExtensibleRules ( ':dn:2.5.13.5:=', ':caseIgnoreMatch:=', et al)

All of the logical Boolean symbol "operators" are supported:

  - AND (&)
  - OR (|)
  - NOT (!)

Any attributeValueAssertion (AVA) comprised of an attributeType, one (1) or more
operator permutations and a non-zero attributeValue should be processed without
incident ** SO LONG AS THE EXPRESSION IS PARENTHETICAL **.

In other words:

  - GOOD:  (objectClass=*)
  - BAD:   objectClass=*

Though the "BAD" example is perfectly legal in most respected LDAP DUAs, ANTLR's
lexer design (when combining things like DNs and Filters within the same grammar
file) reacts quite badly to ambiguity among varying rules. If you are writing a
parser using this grammar, you can easily compensate for this in your code before
processing begins (e.g.: "if <val> is not parenthetical, make it so").

# Value Matching and Schema-related Limitations

All values should be successfully matched through the attributeValue parsing rule.
An instance where a value incorrectly includes a closing parenthesis (as opposed
to using it strictly as a "value end" marker) is considered a BUG and should be
reported.

Values that are one (1) character or longer are considered valid, however certain
characters, such as WHSP (ASCII #32, ' '), will cause a break in a particular
attributeValue (e.g.: 'Jesse Coretta' returns three (3) values: one (1) for my first
name, one (1) for WHSP, and one (1) for my last name). ) I have not yet devised a
reliable fix for this, but since it is not a "fatal condition" -- rather, any parser
COULD conceivably be fixed to handle this -- I have left it as-is for the time being.

Substring Assertion values (such as 'ab*d') are supported in attributeValue instances,
and a unary asterisk is supported in a so-called Presence Assertion values (e.g.: cn=*).
Keep in mind, however, that the associated attributeType is obviously *NOT* verified for
such capabilities in the ANTLR4 grammar "layer" according to any schema definitions. The
same is true for any order-based (i.e.: GE / LE) comparisons.

In other words, schema-illogical AVAs would be identified (likely) after the ANTLR
parsing stage. An example of what I mean: '(color>=blue)' is illogical LDAP-wise,
but as far as ANTLR is concerned, perfectly valid per the parser rules.

AttributeType Tags, such as ';binary' or ';lang-<LANG>', are supported as suffix
components to an attributeType instance, e.g.: 'color;lang-fr=bleu'.
*/
package ldapparser
