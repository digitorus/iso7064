// Package iso7064 is a GO implementation of ISO 7064 used in validation of format like IBAN, ISTC, ISNI, LEI, etc.
//
// The initial version of this package is based the IBAN validator from https://github.com/almerlucke/go-iban
package iso7064

import (
	"errors"
	"math/big"
	"strconv"
)

var errDigitsValidation = errors.New("digits validation failed")
var errIncorrectChecksum = errors.New("incorrect checksum")

// IsValid validates if the given code complies with ISO7064 and tries to
// detect the correct mod automatically.
func IsValid(code string) error {
	// Replace each letter in the string with two digits, thereby expanding the string, where A = 10, B = 11, ..., Z = 35
	mods := ""
	for _, c := range code {
		// Get character code point value
		i := int(c)

		// Check if c is characters A-Z (codepoint 65 - 90)
		if i > 64 && i < 91 {
			// A=10, B=11 etc...
			i -= 55
			// Add int as string to mod string
			mods += strconv.Itoa(i)
		} else {
			mods += string(c)
		}
	}

	// Create bignum from mod string and perform module
	bigVal, success := new(big.Int).SetString(mods, 10)
	if !success {
		return errDigitsValidation
	}

	modVal := new(big.Int).SetInt64(97)
	resVal := new(big.Int).Mod(bigVal, modVal)

	// Check if module is equal to 1
	if resVal.Int64() != 1 {
		return errIncorrectChecksum
	}

	return nil
}
