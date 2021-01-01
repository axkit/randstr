// Package randstr provides functions for generation random strings using presets of allowed characters.
package randstr

import (
	"math/rand"
	"time"
)

var (
	digits     = "0123456789"
	lowerchars = "abcdefghijklmnopqrstuvwxyz"
	upperchars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	symbols    = "~!@#$^&*()-+[]"
)

var (
	// AlphaDigit holds 0..9, A..Z, a..z.
	AlphaDigit = []byte(digits + lowerchars + upperchars)

	// LowerAlphaDigit holds a..z.
	LowerAlphaDigit = []byte(digits + lowerchars)

	// UpperAlphaDigit holds A..Z
	UpperAlphaDigit = []byte(digits + upperchars)

	// DecimalDigit holds 0..9.
	DecimalDigit = []byte(digits)

	// HexDigit holds 0..9, A..F.
	HexDigit = []byte(digits + "ABCDEF")

	// Password hold 0..9, A..Z, a..z and  ~!@#$^&*()-+[]
	Password = []byte(digits + lowerchars + upperchars + symbols)
)

var r *rand.Rand

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomString returns random string. Parameter src is used as a list of allowed characters, size result length.
func RandomString(src []byte, size int) string {

	res := make([]byte, size, size)
	limit := int32(len(src))

	for i := 0; i < size; i++ {
		pos := r.Int31n(limit)
		res[i] = src[pos]
	}
	return string(res)
}
