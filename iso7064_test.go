package iso7064

import (
	"testing"
)

var iso7064tests = []struct {
	in  string
	out error
}{
	{"9845000Q555DEB9E2A69", nil},                  // LEI
	{"9845000Q555DEB9E2A68", errIncorrectChecksum}, // Invalid LEI checksum
	{"WEST12345698765432GB82", nil},                // IBAN
}

func TestIsValid(t *testing.T) {
	for _, tt := range iso7064tests {
		t.Run(tt.in, func(t *testing.T) {
			err := IsValid(tt.in)
			if err != tt.out {
				t.Errorf("%s: got '%v', want '%v'", tt.in, err, tt.out)
			}
		})
	}
}
