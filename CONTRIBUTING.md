# Welcome to the go-schemax contributing guide <!-- omit in toc -->

First, welcome to the go-aci repository.

## Contributor guide

A few things should be reviewed before submitting a contribution to this repository:

 1. Read our [Code of Conduct](./CODE_OF_CONDUCT.md) to keep our community approachable and respectable
 2. Review the Oracle Unified Directory and RedHat Directory administrative guides (specifically the chapters that cover their implementation of the ACIv3 syntax)
 3. Review the main [![GoDoc](https://godoc.org/github.com/JesseCoretta/go-aci?status.svg)](https://godoc.org/github.com/JesseCoretta/go-aci) page, which provides the entire suite of useful documentation rendered in Go's typically slick manner 😎.
 4. Review the [Collaborating with pull requests](https://docs.github.com/en/github/collaborating-with-pull-requests) document, unless you're already familiar with its concepts ...

Once you've accomplished the above items, you're probably ready to start making contributions. For this, I thank you.

## Technical Guidelines

This section contains a few guidelines that I've imposed. This list may change at any time.

 - Cyclomatics - A maximum cyclomatic complexity factor of nine (9) is imposed
   - This means that no function or method provided as contributed content shall exceed this limit
 - Unit Tests - Contributed content shall be accompanied by sufficiently scaled unit tests
   - A massive code coverage % drop as a result of a pull request would be undesirable
 - Comments - All exported (public) functions, methods, constants and global variables are to be reasonably well-documented
   - I reserve the right to correct grammar, if and when needed
 - Target Audience - Keep in mind this is an unusual tool intended for that rare specimen of LDAP architect/admin: the kind who know this sort of thing by heart, but wish there was an easier way to do certain things
   - This is not intended to augment the operation of any directory server
   - This is not intended to make decisions as to best practices in the context of rule authorship
   - This package is only as good as the LDAP engineer wielding it
