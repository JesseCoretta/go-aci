/*
Package aci implements bidirectional marshaling and component abstractions pertaining to the third (3rd) version of the ACI
syntax specification.

# Status

This package is in its early stages, and is undergoing active development. It should NOT be used in any mission-critical
capacity at this time. At the moment, this package is primarily intended for R&D, PoCs and other scenarios in which risk
is minimal or nonexistent.

# About ACIs

An ACI (access control instruction) is a directive that is used to disclose or withhold information based on a predefined
set of conditions, as well as control the abilities of users with varying levels of granularity. Within the spirit of this
package, the ACI syntax is a means for securing information within an X.500/LDAP directory, and is supported by multiple
directory products on the market today.

# Implementation Compatibility Notice

The ACIv3 syntax, though largely the same across the multiple supporting directory products that have adopted it, does have
a few variations in terms of available keywords and features. Though this is not a comprehensive list, a few of these cases
are listed below:

• Target Rule scoping, through the targetscope keyword

• Group attribute-based value matching, through the groupattr keyword

• LDAP Extended Operation OID definitions, through the extop keyword

• LDAP Control OID definitions, through the targetcontrol keyword

• Rights definitions, such as Import and Export

This package aims to support all of the facets of the ACIv3 syntax without exception. Users will need to verify, however,
that any ACI definitions generated -- in part or in whole as a result of using this package -- are compatible with their
particular X.500/LDAP product.

# License

The aci (go-aci) package, from http://github.com/JesseCoretta/go-aci, is available under the terms of the MIT license. For
further details, see the LICENSE file within the aforementioned repository.

# Features

• Intuitive: well-documented with numerous examples

• Efficient: a package-wide cyclomatic complexity factor limit of nine (9) is imposed

• Convenient: a Fluent design is implemented where possible, allowing the chaining of certain command sequences

• Compatible: package design encompasses the facets of the ACIv3 specification as a whole, as opposed to catering to any
specific directory product implementation

• Flexible: ACI composition can be approached in a variety of ways, without enforcing any particular style; for example,
parenthetical encapsulation can be enabled or disabled for select (and eligible) type instances when desired

• High visibility: each Stack instance may be equipped with a channel object, thereby allowing real-time debug/error message
output via a user-controlled instance; see the Stack.SetMessageChan method for details

# Marshaling and Unmarshaling

Within the terms of this package, marshaling is defined through a process that reads the user-provided textual ACI definition,
parses the components and generates a proper instance of the package-provided *ACI struct type. The marshaling process can be
used to allow object-oriented manipulation and interrogation of an ACI's individual component, or to simply gauge the syntactical
validity of a text-based definition.

Conversely, unmarshaling is defined through a process that generates a textual ACI definition based upon the contexts of a
preexisting ACI type instance.

# Potential Use Cases

This package could conceivably be used in any of the following scenarios:

• For Directory security audits that pertain to, or include, access control review

• For Directory personnel in charge of authoring and/or managing documentation

• For Directory personnel who desire a means to author and/or manage ACI stacks in a more programmatic / automated manner

• For use as an access control framework within an actual (Go-based) Directory System Agent implementation that honors the ACI syntax

• For general experimentation within the realm of Directory System Agent access control design

# Limitations

The go-aci package (straight out of the box, so to speak) is not an access control decision-making framework unto itself -- that
particular functionality would reside in the X.500/LDAP server to be protected *through the use of ACIs*.

However this package could be leveraged to CRAFT such a framework, given all of the syntax-defined types are made available to the
end user. If users wish to approach this concept, they are advised to leverage the underlying go-stackage Stack type's methods for
implementing evaluatory capabilities, such as attribute value assertion checks and the like.  This would conceivably allow the use
of matchingRule and ldapSyntax operations that precede attribute value disclosure/withholding.

Another limitation is the lack of comprehensive LDAP Search Filter parsing and decompilation into Rule instances. For the moment,
any user-provided LDAP Search Filter (i.e.: when crafting a `targetfilter` Target Rule) is taken at face-value and is NOT verified.
This will be improved in the near future, at which point an LDAP Search Filter string value shall be interrogated, deconstructed
verified and recomposed into (nested) Rule instances correctly (or will return a meaningful error).

# Contribution Encouraged

The ACIv3 syntax is fairly complex, rendering its innate flexibility akin to a double-edged sword. As such there may be errors, or
concepts overlooked by the author, within this package. Users are strongly encouraged to speak up if they perceive a feature or some
behavioral trait of the package to be suboptimal in some manner.

See https://github.com/JesseCoretta/go-aci/issues for current bug reports, as well as a means to file new ones.

# Warning

The concepts of access control -- whether related to the security of databases or not -- is an extremely critical component of effective
cybersecurity design as a whole. Permissions, such as ACIs, should never be implemented in an untested or cavalier fashion. Though this
package can reduce much of the tedium associated with directory security through the use of permissions, it can just as easily generate
completely bogus rules that will have the opposite intended effect. Those who choose to leverage this package are strongly advised to
triple-check their work. Make no assumptions. Take no unnecessary risks. TEST. TEST. TEST and then TEST some more!

Another area of risk is the disposition (or lack thereof) regarding so-called "ACI Best Practices", which do vary across the various
supporting directory products on the market.

Again, this package aims to provide everything one could possibly need to compose an ACI. However, this package does not discriminate
ACIs that may be overtly "broad" in their influence or entry-matching potential.  One example of this is careless use of the negated
equality operator (!=), which (when used improperly) can disclose myriad attribute values unintentionally. This particular case is well
documented in vendor manuals for supporting directory products (likely for legal CYA reasons). Users are advised to LEARN the syntax
well enough to know when to take such risks.
*/
package aci
