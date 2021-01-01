// Package randstr provides functions for generation random strings using presets of allowed characters.
package randstr

import (
	"bytes"
	"testing"
)

func TestRandomString(t *testing.T) {

	tc := []struct {
		src  []byte
		size int
	}{{DecimalDigit, 1}, {AlphaDigit, 2}}

	for i := range tc {
		res := RandomString(tc[i].src, tc[i].size)
		if rl := len(res); rl != tc[i].size {
			t.Logf("Test case %d failed. Expected result's length %d, got %d", i, tc[i].size, rl)
		}
		if !bytes.ContainsAny(tc[i].src, res) {
			t.Logf("Test case %d failed. Result %s does not contais in %s", i, res, string(tc[i].src))
		}
	}
}
