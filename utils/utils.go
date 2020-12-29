/*
Package utils implements the basic crypto operations used in the challenges:
  - convert from hex to base64 (set1 challenge1)
  - xor of two equal length buffers
	- repeating xor encryption
*/
package utils

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// HexToBase64 converts an hex encoded src to a base64 string.
//
// Returns an error if src is not valid.
func HexToBase64(src []byte) (string, error) {
	raw := make([]byte, hex.DecodedLen(len(src)))
	if _, err := hex.Decode(raw, src); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(raw), nil
}

// Xor two buffers of the same size, return an error if the sizes are different.
func Xor(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("Xor called with two buffers of different sizes %d != %d ", len(a), len(b))
	}
	res := make([]byte, len(a))
	for i, v := range a {
		res[i] = v ^ b[i]
	}
	return res, nil
}

// RepeatXor xor src with key, key is used repeateadly as many time as needed
// to cover the length of src.
func RepeatXor(src, key []byte) []byte {
	expanded_key := make([]byte, len(src))
	for i, _ := range expanded_key {
		expanded_key[i] = key[i%len(key)]
	}
	res, _ := Xor(src, expanded_key)
	return res
}
