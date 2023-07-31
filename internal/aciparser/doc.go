/*
Package aciparser is the top-level ANTLR4 (4.13.0) framework implementation designed
for consumption of ACIv3 instruction expressions.

# Advisory

The aciparser package (and the associated ANTLR4 grammar) is currently in its initial
stages of development. Use with caution and report any bugs.

# Internal Package

This package is internal and should not be accessed by users directly. The top-level
go-aci package calls the aciparser package as needed during ACI marshaling. It is
not intended for use in any other scenario by any person or entity.

# License

The ldapparser package is released as a component in the go-aci suite under the
terms of the MIT License. For full details, see the LICENSE file within the package
repository.

# Expected Operation

This package, as imported and triggered by the top-level go-aci suite, shall parse
an ACIv3 definition, such as ...

 (targetfilter="(&(objectClass=employee)(objectClass=engineering))")(targetscope="onelevel")(version 3.0; acl "Allow onelevel searches for employees"; allow(read,search,compare) userdn = "ldap:///anyone";)

... into discrete object instances (made possible through type definitions provided
by go-aci) that may be further analyzed, interrogated or modified as required.
*/
package aciparser
