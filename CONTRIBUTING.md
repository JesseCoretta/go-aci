# Welcome to the go-aci contributing guide <!-- omit in toc -->

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
   - This does NOT apply to example content (e.g.: `_examples/guiapp/main.go`)
 - Imports
   - 3rd party package imports introduced as a result of a contribution will require some kind of technical justification
   - Only 3rd party imports released under the MIT license shall be considered
   - This does NOT apply to example content (e.g.: `_examples/guiapp/main.go`)
 - Unit Tests - Contributed content shall be accompanied by sufficiently scaled unit tests
   - A massive code coverage % drop as a result of a pull request would be undesirable
 - Comments - All exported (public) functions, methods, constants and global variables are to be reasonably well-documented
   - I reserve the right to correct grammar, if and when needed
   - If English is not your preferred language, don't be afraid -- I'll take care of the polishing, just give me the highlights
   - This is not necessarily REQUIRED for example content, but it would be appreciated
 - Target Audience - Keep in mind this is an unusual tool intended for that rare specimen of LDAP architect/admin: the kind who know this sort of thing by heart, but wish there was an easier way to do certain things
   - This is not intended to augment the operation of any directory server
   - This is not intended to make decisions as to best practices in the context of rule authorship
   - This package is only as good as the LDAP engineer wielding it

## A Note on Example Contributions

The `_examples` folder is intended to house user-authored implementations of this SDK. There are very few guidelines and restrictions for content of this kind, as it is meant to resemble an app the end-user would write for their own use.

- Be as creative as you want
  - So long as this SDK is the relevant centerpiece (or one of multiple), code to your heart's content
- Import whatever packages you want
  - The above "Imports" guideline need not apply here
  - So long as the imported package(s) aren't malicious or fraught with vulnerabilities, etc., any imports are fine if you need them
- Comments would be nice
  - While nowhere near as strict as the above generalized "Technical Guidelines", commentary within your contributed app code would be greatly appreciated by the end user
  - Don't hold yourself to the same writing style as that of the SDK -- you're trying to "sell" the idea of your app to other individuals of varied experience levels; it need not read like a science journal
- All contributed example content shall immediately become subject to the terms of the MIT license
  - The MIT license text block should appear prominently at the top of your `main.go` file as inline text; it need not be pushed as a separate file itself
