### Golang ISO 7064 validator

[![Build Status](https://travis-ci.org/digitorus/iso7064.svg?branch=master)](https://travis-ci.org/digitorus/iso7064)
[![Go Report Card](https://goreportcard.com/badge/github.com/digitorus/iso7064)](https://goreportcard.com/report/github.com/digitorus/iso7064)
[![Coverage Status](https://codecov.io/gh/digitorus/iso7064/branch/master/graph/badge.svg)](https://codecov.io/gh/digitorus/iso7064)
[![GoDoc](https://godoc.org/github.com/digitorus/iso7064?status.svg)](https://godoc.org/github.com/digitorus/iso7064)


This package providers a validtor for the ISO 7064 format like IBAN, ISTC, ISNI, LEI, etc.
Currently this package is hardcoded to validate against MOD 97-10 using numeric strings with two check digits. Please feel free to contribute other ISO 7064 algorithms such as MOD 11-2, 37-2, 661-26 etc.

The initial version of this package is based the IBAN validator from https://github.com/almerlucke/go-iban

#### Contributions

We are looking for contribution to add the following functionality:
- Calculate a new checksum value
- Exctract the checksum value of a given value
- Support other modes ISO 7064 algorithms such as MOD 11-2, 37-2, 661-26 etc.
- Provider helpers (rearrangers) for specific implementations like IBAN
- Add more implementation tests

#### Example

	package main

	import (
		"fmt"
		"github.com/digitorus/iso7064"
	)
	
	func main() {
		if !iso7064.IsValid("9845000Q555DEB9E2A69") {
			fmt.Printf("%v\n", err)
			return
		}

        	fmt.Println("Valid ISO 7064 value")
	}
