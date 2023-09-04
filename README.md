[![Go Report Card](https://goreportcard.com/badge/github.com/JesseCoretta/go-aci)](https://goreportcard.com/report/github.com/JesseCoretta/go-aci) [![GoDoc](https://godoc.org/github.com/JesseCoretta/go-aci?status.svg)](https://godoc.org/github.com/JesseCoretta/go-aci) ![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square) [![codecov](https://codecov.io/gh/JesseCoretta/go-aci/graph/badge.svg?token=RLW4DHLKQP)](https://codecov.io/gh/JesseCoretta/go-aci) [![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/JesseCoretta/go-aci/issues) 

## Overview

Package aci implements the complete ACIv3 syntax in a vendor-agnostic manner with rich features.

## License

The aci (go-aci) package, from [`go-aci`](http://github.com/JesseCoretta/go-aci), is available under the terms of the MIT license. For further
details, see the LICENSE file within the aforementioned repository. 

Package aci implements the complete ACIv3 syntax in a vendor-agnostic manner with rich features.

## Status

Current state is EXPERIMENTAL!

This package is in its early stages, and is undergoing active development. It should NOT be used in any production capacity at this time.
capacity at this time. At the moment, this package is primarily intended for R&D, PoCs and other scenarios in which risk
is minimal or nonexistent.

If and when the version reaches 1.0.0, it shall no longer be considered experimental.

## About ACIs

An ACI (access control instruction) is a directive that is used to disclose or withhold information based on a predefined
set of conditions, as well as control the abilities of users with varying levels of granularity. Within the spirit of this
package, the ACI syntax is a means for securing information within an X.500/LDAP directory, and is supported by multiple
directory products on the market today.

## Implementation Compatibility Notice

The ACIv3 syntax, though largely the same across the multiple supporting directory products that have adopted it, does have
a few variations in terms of available keywords and features. Though this is not a comprehensive list, a few of these cases
are listed below:

- TargetRule scoping, through the targetscope keyword
- BindRule roles, through the roledn keyword
- Group attribute-based value matching, through the groupattr keyword
- LDAP Extended Operation OID definitions, through the extop keyword
- LDAP Control OID definitions, through the targetcontrol keyword
- Rights definitions, such as Import and Export
- Permitted 'levels' for inheritance value matching

This package aims to support *ALL* of the facets of the ACIv3 syntax without exception. Users will need to verify, however,
that any ACI definitions generated -- in part or in whole as a result of using this package -- are compatible with their
particular X.500/LDAP product; check the docs!

## License

The aci (go-aci) package, from [`go-aci`](http://github.com/JesseCoretta/go-aci), is available under the terms of the MIT license. For
further details, see the LICENSE file within the aforementioned repository.

# Features

- Intuitive: well-documented with numerous examples
- Efficient: a package-wide cyclomatic complexity factor limit of nine (9) is imposed; the imported go-stackage and go-objectid
packages both exercise similar criteria
- Compatible: package design encompasses the facets of the ACIv3 specification as a whole, as opposed to catering to any
specific directory product implementation
- Flexible: ACI composition can be approached in a variety of ways, without enforcing any particular style; for example,
parenthetical encapsulation can be enabled or disabled for select (and eligible) type instances when desired, or set globally

## Marshaling and Unmarshaling

This package (internally) implements an ANTLR4 parsing subsystem to facilitate the marshaling of ACI textual definitions.

Within the terms of this package, marshaling is defined through a process that reads the user-provided textual ACI definition,
parses the components and generates a proper instance of the package-provided Instruction type.

Conversely, unmarshaling is defined through a process that generates a textual ACI definition based upon the contexts of a
preexisting Instruction type instance.

## Potential Use Cases

This package could conceivably be used in any of the following scenarios:

- For Directory security audits that pertain to, or include, access control review
- For Directory personnel in charge of authoring and/or managing rich documentation
- For Directory personnel who desire a means to author and/or manage sets of ACIs in a more programmatic / automated manner, perhaps
with the aid of a templating system
- For use as an access control framework within an actual (Go-based) Directory System Agent implementation that honors the ACI syntax
- For generalized experimentation within the realm of Directory System Agent access control design and even penetration testing

## Limitations

The go-aci package (straight out of the box, so to speak) is not an access control decision-making framework unto itself -- that
particular functionality would reside in the X.500/LDAP server to be protected *through the use of ACIs*.

However this package could be leveraged to CRAFT such a framework, given all of the syntax-defined types are made available to the
end user. If users wish to approach this concept, they are advised to leverage the underlying go-stackage Stack type's methods for
implementing evaluatory capabilities, such as attribute value assertion checks and the like.  This would conceivably allow the use
of matchingRule and ldapSyntax operations that precede attribute value disclosure/withholding.

## Quotation Schemes

Another trait of this package's flexibility is the ability to handle either of the following quotation schemes when parsing or
building statements that reference a multi-valued expression:

```
"value" || "value" || "value" ...
"value || value || value" ...
```

In particular, these sorts of quotation schemes appear in the following scenarios:

- `targetattr` TargetRules for lists of LDAP attribute types
- `target`, `target_to` and `target_from` TargetRule distinguished names
- `userdn` and `groupdn` BindRule distinguished names
- `extop` and `targetcontrol` TargetRule ASN.1 object identifiers

Users are advised to honor the quotation scheme recommended by their vendor or product documentation. This package aims to
support either of the above schemes with no variance in the end result, but has no official position as to which of these
schemes should be honored by the user.

## Contribution Encouraged

The ACIv3 syntax is fairly complex, rendering its innate flexibility akin to a double-edged sword. As such there may be errors, or
concepts overlooked by the author within this package.  Users are STRONGLY ENCOURAGETO SPEAK UP if they perceive a feature or some
behavioral trait of the package to be suboptimal or incomplete in some manner.

See [issuers](https://github.com/JesseCoretta/go-aci/issues) for all bug reports -- past and present -- as well as a means to file new ones.

## Words of Warning

The concept of access control -- whether related to the security of databases or not -- is an extremely critical component of effective
cybersecurity design as a whole. Permissions, such as ACIs, should never be implemented in an untested or cavalier fashion. Breaches
associated with poor access control models can destroy companies, end careers and maybe even endanger human lives.

Though this package can reduce much of the tedium associated with directory security through the use of permissions, it can just as
easily generate completely bogus rules that will have the opposite intended effect. Even worse, it may generate rules that may expose
sensitive DIT content!

Those who choose to leverage this package are strongly advised to triple-check their work. Make no assumptions. Take no unnecessary risks.
TEST. TEST. TEST and then TEST some more!

Another area of risk is the disposition (or lack thereof) regarding so-called "ACI Best Practices", which do vary across the various
supporting directory products on the market. Users uncertain as to the best approach for a desired action are strongly advised to ask
their vendor, or consult an appropriate online community forum.

By now, it is likely obvious this package aims to provide everything one could possibly need to compose an ACI. However, this package does
not discriminate ACIs that may be overtly "broad" in their influence or entry-matching potential.

One such example of this is careless use of the negated equality operator (!=), which (when used improperly) can disclose myriad attribute
values unintentionally. This particular case is well-documented in vendor manuals for supporting directory products (likely for legal CYA
reasons). Users are advised to LEARN the syntax well enough to know when to take such risks.

