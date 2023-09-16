[![Go Report Card](https://goreportcard.com/badge/github.com/JesseCoretta/go-aci)](https://goreportcard.com/report/github.com/JesseCoretta/go-aci) [![Go Reference](https://pkg.go.dev/badge/github.com/JesseCoretta/go-aci.svg)](https://pkg.go.dev/github.com/JesseCoretta/go-aci) ![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square) [![codecov](https://codecov.io/gh/JesseCoretta/go-aci/graph/badge.svg?token=RLW4DHLKQP)](https://codecov.io/gh/JesseCoretta/go-aci) [![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/JesseCoretta/go-aci/issues) ![Static Badge](https://img.shields.io/badge/experimental-red?logo=simple-icons&labelColor=purple&color=maroon&link=https%3A%2F%2Fgithub.com%2FJesseCoretta%2Fgo-aci)

## Overview

Package aci implements the complete ACIv3 syntax in a vendor-agnostic manner with rich features.

## License

The aci package -- from [`go-aci`](http://github.com/JesseCoretta/go-aci) -- is available under the terms of the MIT license.

For further details, see the [LICENSE](/LICENSE) file within the root of the linked repository.

## Dependencies

This package depends upon the following official (built-in) Go packages:

  - [`crypto/sha1`](https://pkg.go.dev/crypto/sha1)
  - [`encoding/binary`](https://pkg.go.dev/encoding/binary)
  - [`fmt`](https://pkg.go.dev/fmt)
  - [`errors`](https://pkg.go.dev/errors)
  - [`reflect`](https://pkg.go.dev/reflect)
  - [`strconv`](https://pkg.go.dev/strconv)
  - [`strings`](https://pkg.go.dev/strings)
  - [`testing`](https://pkg.go.dev/testing)
  - [`time`](https://pkg.go.dev/time)
  - [`unicode`](https://pkg.go.dev/unicode)

This package depends upon the following third party packages:

  - [`go-antlraci`](https://github.com/JesseCoretta/go-antlraci)\*
  - [`go-objectid`](https://github.com/JesseCoretta/go-objectid)\*
  - [`go-stackage`](https://github.com/JesseCoretta/go-stackage)\*

_\* Conceived and maintained by same author_

## Status and Production Roadmap

Per the badge in the header, this package is to be regarded as highly experimental.

Design for this package began during June of 2023. As such, it is quite infantile and has not been subjected to any known use in the wild.

Because of this, it should NOT be used in any production or mission-critical context. The only appropriate applications for this package, at
the time of this writing, are those confined to R&D, PoC, testing and other low-or-no-risk scenarios.

However, should the [code coverage](https://app.codecov.io/gh/JesseCoretta/go-aci) factor reach __95% or higher__, the package shall be bumped
up to a production release number (e.g.: `v1.0.0`+) and no longer regarded nor advertised as "experimental". The experimental badge within
the header of this document shall also be removed.

## About ACIs

An ACI[^1] is a directive that is used to disclose or withhold information based on a predefined set of conditions, as well
as control the abilities of users with varying levels of granularity. Within the spirit of this package, the ACI syntax is
a means for securing information within an X.500/LDAP[^2] DIT[^3], and is supported by multiple directory products on the market
today.

[^1]: Access Control Instruction
[^2]: ITU-T X-Series 500 / Lightweight Directory Access Protocol ([`X.500`](https://www.itu.int/rec/T-REC-X.500) and [RFC4510](https://datatracker.ietf.org/doc/html/rfc4510), et al)
[^3]: Directory Information Tree

## Implementation Compatibility Notice

The ACIv3 syntax, though largely the same across the multiple supporting directory products that have adopted it, does have
a few variations in terms of available keywords and features. Though this is not a comprehensive list, a few of these cases
are listed below:

- TargetRule scoping, through the targetscope keyword
- BindRule DN[^4] roles, through the roledn keyword
- Group attribute-based value matching, through the groupattr keyword
- LDAP Extended Operation OID[^5] definitions, through the extop keyword
- LDAP Control OID definitions, through the targetcontrol keyword
- Rights definitions, such as Import and Export
- Permitted 'levels' for inheritance value matching

This package aims to support *ALL* of the facets of the ACIv3 syntax without exception. Users will need to verify, however,
that any ACI definitions generated -- in part or in whole as a result of using this package -- are compatible with their
particular X.500/LDAP DSA[^6] implementation; check the docs and always listen to your vendor!

[^4]: (LDAP) Distinguished Name
[^5]: (ASN.1) Object Identifier
[^6]: Directory System Agent, i.e.: your LDAP server(s)

## License

The aci (go-aci) package, from [`go-aci`](http://github.com/JesseCoretta/go-aci), is available under the terms of the MIT
license. For further details, see the LICENSE file within the aforementioned repository.

# Features

- Intuitive
  - Well-documented with numerous examples
  - Written with consideration for the various strata of LDAP engineers
  - Style leans towards "informative" more so than "succinct", and strives to be _more helpful_ than _terse_
- Efficient
  - Easy initialization routines available for those on-the-go
  - Unnecessary string literals are not stored -- for example, the `ldap:///` prefix required of all LDAP DNs per ACIv3 -- yet they will still be _present_ during string representation
    - Thus, less typing for full compliance!
  - Parsers offered via [`go-antlraci`](http://github.com/JesseCoretta/go-antlraci) are wrapped and extended in this package in a "conveniently compartmentalized" manner
    - All-or-nothing `Instruction` parsing is **not required**
    - Users may opt to parse only a _single_ TargetRule or an entire (nested!) BindRules hierarchical expression!
    - All major components allow compartmentalized parsing:
      - `TargetRule`
      - `TargetRules`
      - `BindRule`
      - `BindRules`
      - `PermissionBindRule`
      - `PermissionBindRules`
      - `Instruction`
      - `Instructions`
  - Memory usage was given particular consideration during the initial design phase
    - Certain values, such as an SSF[^7] level, were designed to fit (_completely!_) inside a single byte instance, regardless of the effective value
    - All numerical types utilize only the smallest possible primitive numeral type (e.g.: `uint8`, `uint16`) to suit its needs; never too much, and never too little!
    - Embedded pointer references are leveraged __centrally__ throughout this package; storage of unneeded string literals is avoided wherever possible
  - A package-wide cyclomatic complexity factor limit of nine (9) is imposed
    - We realize the standard recommended maximum is fifteen (15); we feel we can do better!
    - The following imported packages also exercise this philosophy:
      - [`go-stackage`](http://github.com/JesseCoretta/go-stackage)
      - [`go-objectid`](http://github.com/JesseCoretta/go-objectid)
- Compatible
  - Overall package design is meant to honor all of the facets of the ACIv3 specification **_in its entirety_**
  - No single vendor implementation is catered-to exclusively
  - So-called "Happy Mediums" were chosen at all possible points, favoring the exclusion of _no one_ over any vendor-partisan implementation
  - Quite simply, the chances are good that this package supports **_more_** of the ACIv3 syntax than your actual LDAP product does
- Flexible
  - ACI composition can be approached in a variety of ways, without enforcing any particular "style"
  - Parenthetical encapsulation can be enabled or disabled for select (and eligible) type instances when desired, or set globally
  - Two (2) distinct quotation styles are supported for multi-valued DNs, OIDs and ATs present within eligible BindRule and/or TargetRule instances; see [here](#quotation-schemes) for details
  - Fluent-style "method chaining" is supported, but not required
  - Most values can be assigned and set 'piecemeal', or in an 'all-in-one-shot' context

[^7]: Security Strength Factor, which represents a numerical abstract (0-256) based upon the perceived level of confidentiality
employed by the request in question

## Marshaling and Unmarshaling

This package (internally) implements an ANTLR4[^8] parsing subsystem to facilitate the marshaling of ACI textual definitions, but
also allows all value types to be built or assembled _without_ the use of textual ACIv3 syntax expression.

Within the terms of this package, marshaling (parsing) is defined through a process that reads the user-provided textual expression,
parses the components and generates a proper instance of the package-provided type.

Conversely, unmarshaling (or "string representation") is defined through a process that generates a textual ACI definition based upon
the contents of a preexisting type instance, usually a receiver. This type instance may have been created through parsing, or it may
have been assembled manually by the user.

When combined, these two complementary -- yet opposite -- contexts offer a reliable, fully bidirectional solution.

[^8]: ANother Tool for Language Recognition, Version 4 (www.antlr.org)

## Potential Use Cases

This package could conceivably be used in any of the following scenarios:

- For Directory security audits that pertain to, or include, access control review
- For Directory personnel in charge of authoring and/or managing rich documentation
- For Directory personnel who desire a means to author and/or manage sets of ACIs in a more programmatic / automated manner, perhaps
with the aid of a templating system
- For use as an access control framework within an actual (Go-based) Directory System Agent implementation that honors the ACI syntax
- For generalized experimentation within the realm of DSA access control design and even PEN[^9] testing

[^9]: PENetration Testing

## Limitations

The go-aci package (straight out of the box, so to speak) is not an access control decision-making framework unto itself -- that
particular functionality would reside in the X.500/LDAP server to be protected *through the use of ACIs*.

However this package could be leveraged to CRAFT such a framework, given all of the syntax-defined types are made available to the
end user. If users wish to approach this concept, they are advised to leverage the underlying [`go-stackage`](https://github.com/JesseCoretta/go-stackage)
Stack type's methods for implementing evaluatory capabilities, such as attribute value assertion checks and the like. This would
conceivably allow the use of matchingRule and ldapSyntax operations that precede attribute value disclosure/withholding.

## Comparison Operators

Thanks to the import of the [`go-stackage`](https://github.com/JesseCoretta/go-stackage) package, this package gains access to all
of the necessary comparison operators for use in the crafting of ACIv3 compliant BindRule and TargetRule expressions.

Below are two (2) tables containing all of the comparison operators available, as well all of the applicable (and non-applicable)
scenarios for their use.

| Bind Keyword | Eq  | Ne  | Lt  | Le  | Gt  | Ge  |
| ------------- | :---: | :---: | :---: | :---: | :---: | :---: |
| **`ip`**         |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`dns`**        |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`userdn`**     |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`groupdn`**    |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`roledn`**     |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`userattr`**   |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`groupattr`**  |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`authmethod`** |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`dayofweek`**  |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| [**`ssf`**](## "Security Strength Factor")        |  ✅  |  ✅  |  ✅  |  ✅  |  ✅  |  ✅  |
| **`timeofday`**  |  ✅  |  ✅  |  ✅  |  ✅  |  ✅  |  ✅  |

| Target Keyword | Eq  | Ne  | Lt  | Le  | Gt  | Ge  |
| --------------- | :---: | :---: | :---: | :---: | :---: | :---: |
| **`target`**          |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`target_to`**       |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`target_from`**     |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`targetattr`**      |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`targetfilter`**    |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`targattrfilters`** |  ✅  |  ❌  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`targetscope`**     |  ✅  |  ❌  |  ❌ |  ❌ |  ❌ |  ❌ |
| **`targetcontrol`**   |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |
| [**`extop`**](## "LDAP Extended Operation Object Identifier(s)")           |  ✅  |  ✅  |  ❌ |  ❌ |  ❌ |  ❌ |

## Quotation Schemes

Another trait of this package's flexibility is the ability to handle either of the following quotation schemes when parsing or
building statements that reference a multi-valued expression:

```
"value" || "value" || "value" ...
"value || value || value" ...
```

In particular, these sorts of quotation schemes appear in the following scenarios:

- `targetattr` TargetRules for lists of LDAP ATs[^10]
- `target`, `target_to` and `target_from` TargetRule DNs
- `userdn` and `groupdn` BindRule DNs
- `extop` and `targetcontrol` TargetRule OIDs

Users are advised to honor the quotation scheme recommended by their vendor or product documentation. This package aims to
support either of the above schemes with no variance in the end result, but has no official position as to which of these
schemes should be honored by the user except that quotation should always be used _in some form_.

[^10]: (LDAP) AttributeTypes

## Contribution Encouraged

The ACIv3 syntax is fairly complex, rendering its innate flexibility akin to a double-edged sword. As such there may be errors, or
concepts overlooked by the author within this package. Users are STRONGLY ENCOURAGED TO SPEAK UP if they perceive a feature or some
behavioral trait of the package to be suboptimal or incomplete in some manner.

See [issues](https://github.com/JesseCoretta/go-aci/issues) for all bug reports -- past and present -- as well as a means to file
new ones.

## Some Words of Warning

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

One such example of this is careless use of the negated equality operator (`!=`), which (when used improperly) can disclose myriad attribute
values unintentionally. This particular case is well-documented in vendor manuals for supporting directory products (likely for legal CYA[^11]
reasons). Users are advised to LEARN the syntax well enough to know when to take such risks.

[^11]: Cover Your ~~_REDACTED_~~ **Gluteus Maximus**

