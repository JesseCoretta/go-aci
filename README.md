[![Go Report Card](https://goreportcard.com/badge/github.com/JesseCoretta/go-aci)](https://goreportcard.com/report/github.com/JesseCoretta/go-aci) [![OpenSSF Best Practices](https://www.bestpractices.dev/projects/8895/badge)](https://www.bestpractices.dev/projects/8895) [![Go Reference](https://pkg.go.dev/badge/github.com/JesseCoretta/go-aci.svg)](https://pkg.go.dev/github.com/JesseCoretta/go-aci) [![CodeQL](https://github.com/JesseCoretta/go-aci/workflows/CodeQL/badge.svg)](https://github.com/JesseCoretta/go-aci/actions/workflows/github-code-scanning/codeql) [![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](https://github.com/JesseCoretta/go-aci/blob/main/LICENSE) [![codecov](https://codecov.io/gh/JesseCoretta/go-aci/graph/badge.svg?token=RLW4DHLKQP)](https://codecov.io/gh/JesseCoretta/go-aci) [![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/JesseCoretta/go-aci/issues) [![Experimental](https://img.shields.io/badge/experimental-blue?logoColor=blue&label=%F0%9F%A7%AA%20%F0%9F%94%AC&labelColor=blue&color=gray)](https://github.com/JesseCoretta/JesseCoretta/blob/main/EXPERIMENTAL.md) [![GitHub Workflow Status (with event)](https://img.shields.io/github/actions/workflow/status/jessecoretta/go-aci/go.yml?event=push)](https://github.com/JesseCoretta/go-aci/actions/workflows/go.yml) [![Author](https://img.shields.io/badge/author-Jesse_Coretta-darkred?label=%F0%9F%94%BA&labelColor=indigo&color=maroon)](mailto:jesse.coretta@icloud.com) [![GitHub release (with filter)](https://img.shields.io/github/v/release/JesseCoretta/go-aci)](https://github.com/JesseCoretta/go-aci/releases) [![Help Animals](https://img.shields.io/badge/help_animals-gray?label=%F0%9F%90%BE%20%F0%9F%98%BC%20%F0%9F%90%B6&labelColor=yellow)](https://github.com/JesseCoretta/JesseCoretta/blob/main/DONATIONS.md)

## Overview

Package aci implements the complete ACIv3 syntax as an SDK in a vendor-agnostic manner with rich features and a flexible design.

## Maintenance Notice

This package is now archived. The API is stable and it does everything I ever wanted it to do. Enjoy.

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

This package depends upon the following third-party packages:

  - [`go-antlraci`](https://github.com/JesseCoretta/go-antlraci)\*
  - [`go-objectid`](https://github.com/JesseCoretta/go-objectid)\*
  - [`go-stackage`](https://github.com/JesseCoretta/go-stackage)\*
  - [`go-shifty`](https://github.com/JesseCoretta/go-shifty)\*

_\* Conceived and maintained by same author_

## Status and Production Roadmap

Per the badge in the header, this package is to be regarded as highly experimental.

Design for this package began during June of 2023. As such, it is quite infantile and has not been subjected to any known use in the wild.

Because of this, it should NOT be used in any production or mission-critical context without extreme scrutiny.

The only "safe" applications for this package -- at the time of this writing -- are those confined to R&D, PoC, testing and other low-or-no-risk scenarios.

Adoption of this package is NOT a substitute for a comprehensive ACI review process. It is intended to reduce tedium, _not_ assume full responsibility for ACI authorship and management.

If you want to try and tie some A.I into this, thats your business -- but don't complain to me when some emerging intelligence locks you out of your DIT because it considers you a threat. 🤓

## About ACIs

Within the context of the official ACIv3 syntax, an [ACI](## "Access Control Instruction") is an expressive statement or "policy" that is used to define the disclosing or withholding disposition for information within an [X.500](## "ITU-T X-Series 500")/[LDAP](## "Lightweight Directory Access Protocol") [DIT](## "Directory Information Tree") or [DIB](## "Directory Information Base").

In layperson's terms, ACIs are a specific and (largely) non-proprietary form of "LDAP permissions" that govern who can read, write, search, etc. An individual "permission" is typically called an _instruction_.

Generally, instructions are stored within the very entry they're designed to control. For example, an instruction that governs user accounts would (likely) be found assigned to the `ou=People,dc=example,dc=com` entry through the `aci` LDAP attributeType.

Given that an instruction can police the base entry in addition to its subordinate entries, an instruction can "go anywhere", so long as the entry or entries it protects are hierarchically subordinate (or "beneath") said instruction.

In other words, an instruction that governs all user accounts generally should not be _assigned_ to one of those accounts, rather it should be assigned to the immediate parent context, or perhaps to the DIT's root entry in cases where ALL instructions are centrally managed.

Not all LDAP server implementations support the ACIv3 syntax, but several do. Check your vendor or reference material for compatibility information. Though the above information is correct _in general_, you should always yield to the style or technique required or recommended by your vendor or support community.

## ACI Syntax Variant Notice

Please note that this package is solely tailored to honor the ACIv3 syntax. The ACIv3 syntax is **wholly separate and distinct** from the following LDAP ACI implementations:

  - OpenLDAP's _Experimental_ ACI syntax
  - ApacheDS's _Entry, Prescriptive & Subentry_ ACI syntax

Users in need of parsing and/or assembly capabilities for any ACI or ACL syntax **other than** the ACIv3 syntax honored in this package will have to look elsewhere, or author their own (hint: check out [ANTLR4](http://www.antlr.org/) if you need to write your own text-to-instance parser).

## Implementation Compatibility Notice

The ACIv3 syntax, though largely the same across the multiple supporting directory products that have adopted it, does have a few small variations in terms of the available keywords and features.

Though this is not a comprehensive list, a few of these cases are listed below:

- TargetRule scoping, through the `targetscope` keyword
- BindRule [DN](## "LDAP Distinguished Name") roles, through the `roledn` keyword
- Group attribute-based value matching, through the `groupattr` keyword
- LDAP Extended Operation [OID](## "ASN.1 Object Identifier") definitions, through the `extop` keyword
- LDAP Control OID definitions, through the `targetcontrol` keyword
- Rights definitions, such as `import` and `export`
- Number of supported 'levels' for inheritance value matching

This package aims to support *ALL* of the facets of the ACIv3 syntax without exception. Users will need to verify, however, that any ACI definitions generated -- in part or in whole as a result of using this package -- are compatible with their particular X.500/LDAP [DSA](## "Directory System Agent; an LDAP server") implementation; check the docs and always listen to your vendor!

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
    - Certain values, such as an [**SSF**](## "Security Strength Factor") level, were designed to fit (_completely!_) inside a single byte instance, regardless of the effective value
    - All numerical types utilize only the smallest possible primitive numeral type (e.g.: `uint8`, `uint16`) to suit its needs; never too much, and never too little!
    - Embedded pointer references are leveraged __centrally__ throughout this package; storage of unneeded string literals is avoided wherever possible
  - A package-wide cyclomatic complexity factor limit of nine (9) is imposed
    - We realize the standard recommended maximum is fifteen (15); we simply feel we can do better!
    - The following imported packages also exercise this philosophy:
      - [`go-objectid`](https://github.com/JesseCoretta/go-objectid)
      - [`go-stackage`](https://github.com/JesseCoretta/go-stackage)
      - [`go-shifty`](https://github.com/JesseCoretta/go-shifty)
- Compatible
  - Overall package design is meant to honor all of the facets of the ACIv3 specification **_in its entirety_**
  - No single vendor implementation is catered-to exclusively
  - So-called "Happy Mediums" were chosen at all possible points, favoring the exclusion of _no one_ over any vendor-partisan implementation
  - Quite simply: chances are good that this package supports **_more_** of the ACIv3 syntax than your actual LDAP product does
- Flexible
  - ACI composition can be approached in a variety of ways, without enforcing any particular "style"
  - Parenthetical encapsulation, padding and other attributes can be enabled or disabled for select (and eligible) type instances when desired, or even set globally in certain cases
  - Two (2) distinct quotation styles are supported for multi-valued DNs, OIDs and ATs present within eligible BindRule and/or TargetRule instances; see [here](#quotation-schemes) for details
  - Fluent-style "method chaining" is supported, but not required
  - Most values can be assigned and set 'piecemeal', or in an 'all-in-one-shot' context

## Parsing vs. Assembly

This package imports an [ANTLR4](## "ANother Tool for Language Recognition, Version 4") parsing subsystem to facilitate the "conversion" of ACI textual definitions into proper type instances. This functionality is provided by the [`go-antlraci`](https://github.com/JesseCoretta/go-antlraci) package.

The parser reads the user-provided textual expression, processes the components and generates a proper instance of the package-provided type.

For example, this is a basic parse operation in which an abstract SSF context is evaluated as greater than or equal to an SSF value of 128. This expression is to be converted into an instance of BindRule.

```
// raw textual expression
raw := `ssf >= "128"`

// create variable for our
// new BindRule
var br BindRule

// Parse the above raw variable using
// your BindRule instance's Parse
// method ...
if err := br.Parse(raw); err != nil {
        fmt.Println(err) // always check your parser errors
        return
}
```

On the other hand, one need not limit their activities to "parsing of textual expressions"; one can also assemble various expressive object instances and input only the specific values needed.

For example, this is a basic assembly operation that mirrors the outcome of the above:

```
// Create our SSF instance
obj := SSF(128)

// Use object's Ge (Greater Than Or Equal)
// method to create a new BindRule instance

br := obj.Ge()
```

Which technique is "better" depends on the use-case. For individuals writing ACIs themselves, the "assembly" option may seem more appropriate.

Then again, the same sentiment may not apply to situations in which there are no individuals, and the process is wholly automated: the "parsing" option may be far more appealing. Parsing may also be desirable if a simple syntax check is the only objective.

YMMV.

## Potential Use Cases

This package could conceivably be used in any of the following scenarios:

- For Directory security audits that pertain to, or include, access control review
- For Directory personnel in charge of authoring and/or managing rich documentation
- For Directory personnel who desire a means to author and/or manage sets of ACIs in a more programmatic / automated manner, perhaps with the aid of a templating system
- For use as an access control framework within an actual (Go-based) Directory System Agent implementation that honors the ACI syntax
- For generalized experimentation within the realm of DSA access control design

## Limitations

The go-aci package (straight out of the box, so to speak) is not an access control decision-making framework unto itself -- that particular functionality would reside in the X.500/LDAP server to be protected *through the use of ACIs*.

However this package could be leveraged to craft such a framework, given all of the syntax-defined types are made available to the end user. If users wish to approach this concept, they are advised to leverage the underlying [`stackage.Stack`](https://github.com/JesseCoretta/go-stackage) type's methods for implementing evaluatory capabilities, such as attribute value assertion checks and the like. This would conceivably allow the use of `matchingRule` and `ldapSyntax` operations that precede attribute value disclosure/withholding (hint: take a look at [`go-schemax`](https://github.com/JesseCoretta/go-schemax) if this capability interests you).

## Comparison Operators

Thanks to the import of the [`go-stackage`](https://github.com/JesseCoretta/go-stackage) package, this package gains access to all of the necessary comparison operators for use in the crafting of ACIv3 compliant BindRule and TargetRule expressions.

Below are two (2) tables containing all of the comparison operators available, as well all of the applicable (and non-applicable) scenarios for their use. These tables are based upon vendor documentation for multiple adopters of the ACIv3 syntax, therefore it may not reflect the offerings of every product identically.

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

Another trait of this package's flexibility is the ability to handle either of the following quotation schemes when parsing or building statements that reference a multi-valued expression:

```
"value" || "value" || "value" ...
"value || value || value" ...
```

In particular, these sorts of quotation schemes appear in the following scenarios:

- `targetattr` TargetRules for lists of LDAP [ATs](## "LDAP AttributeTypes")
- `target`, `target_to` and `target_from` TargetRule DNs
- `userdn` and `groupdn` BindRule DNs
- `extop` and `targetcontrol` TargetRule OIDs

Users are advised to honor the quotation scheme recommended by their vendor or product documentation. This package aims to support either of the above schemes with no variance in the end result, but has no official position as to which of these schemes should be honored by the user except that quotation should always be used _in some form_.

## Contribution Encouraged

The ACIv3 syntax is fairly complex, rendering its innate flexibility akin to a double-edged sword. As such there may be errors, or concepts overlooked by the author within this package. Users are **strongly encouraged to speak up** if they perceive a feature or some behavioral trait of the package to be suboptimal or incomplete in some manner.

See [issues](https://github.com/JesseCoretta/go-aci/issues) for all bug reports -- past and present -- as well as a means to file new ones.

## Some Words of Warning

The concept of access control -- whether related to the security of databases or not -- is an extremely critical component of effective cybersecurity design as a whole. Permissions, such as ACIs, should never be implemented in an untested or cavalier fashion. Breaches associated with poor access control models can destroy companies, end careers and maybe even endanger human lives.

Though this package can reduce much of the tedium associated with directory security through the use of permissions, it can just as easily generate completely bogus rules that will have the opposite intended effect. Even worse, it may generate rules that may expose sensitive DIT content!

Those who choose to leverage this package are strongly advised to triple-check their work. Make no assumptions. Take no unnecessary risks. TEST. TEST. TEST and then TEST some more!

Another area of risk is the disposition (or lack thereof) regarding so-called "ACI Best Practices", which do vary across the various supporting directory products on the market. Users uncertain as to the best approach for a desired action are strongly advised to ask their vendor, or consult an appropriate online community forum.

By now, it is likely obvious this package aims to provide everything one could possibly need to compose an ACI. However, this package does not discriminate ACIs that may be overtly "broad" in their influence or entry-matching potential.

One such example of this is careless use of the negated equality operator (`!=`), which (when used improperly) can disclose myriad attribute values unintentionally. This particular case is well-documented in vendor manuals for supporting directory products (likely for legal CYA reasons). Users are advised to LEARN the syntax well enough to know when to take such risks.
