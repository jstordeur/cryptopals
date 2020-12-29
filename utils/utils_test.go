package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexToBase64(t *testing.T) {
	src := []byte("abc")
	res, err := HexToBase64(src)
	assert.NotNil(t, err)
	assert.Equal(t, 0, len(res))
	src = []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	res, err = HexToBase64(src)
	assert.Nil(t, err)
	assert.Equal(t, "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t", res)
}

func TestXor(t *testing.T) {
	buf1 := []byte{4}
	buf2 := []byte{4, 4}
	res, err := Xor(buf1, buf2)
	assert.NotNil(t, err)
	assert.Nil(t, res)
	buf1 = append(buf1, byte(19))
	res, err = Xor(buf1, buf2)
	assert.Nil(t, err)
	assert.Equal(t, []byte{0, 23}, res)
}

func TestRepeatXor(t *testing.T) {
	src := []byte{1, 2, 1, 2}
	key := []byte{1, 2}
	assert.Equal(t, []byte{0, 0, 0, 0}, RepeatXor(src, key))
	src = []byte{1, 2}
	key = []byte{1, 2, 3}
	assert.Equal(t, []byte{0, 0}, RepeatXor(src, key))
	src = []byte{1, 2, 3}
	key = []byte{0, 1}
	assert.Equal(t, []byte{1, 3, 3}, RepeatXor(src, key))
}
