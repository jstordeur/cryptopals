package frequency

import (
	"encoding/base64"
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterFrequencyInEnglish(t *testing.T) {
	total := 0.0
	for _, freq := range LetterFrequencyInEnglish {
		total += freq
	}
	assert.Greater(t, 0.1, math.Abs(100.0-total))
}

// TODO: there's probably a better test than that... Although the challenge 3
// is pretty much a test of this function.
func TestChiSquared(t *testing.T) {
	gibberish := []byte(";kjvreurh;awrhh;evuh;ev")
	eng := []byte("A perfectly valid english sentence.")
	french := []byte("Une phrase correct en francais.")
	assert.Greater(t, ChiSquared(gibberish), ChiSquared(eng))
	assert.Greater(t, ChiSquared(french), ChiSquared(eng))
	non_letters := []byte{6, 42, 42, 46, 44, 43, 34, 101, 8, 6, 98, 54, 101, 41, 44, 46, 32, 101, 36, 101, 53, 42, 48, 43, 33, 101, 42, 35, 101, 39, 36, 38, 42, 43}
	correct := []byte("Cooking MC's like a pound of bacon")
	assert.Greater(t, ChiSquared(non_letters), ChiSquared(correct))
	wrong := []byte("Dhhlni`'JD t'knlb'f'whric'ha'efdhi")
	assert.Greater(t, ChiSquared(wrong), ChiSquared(correct))
}

func TestHammingDistance(t *testing.T) {
	a := []byte{1, 2}
	b := []byte{1}
	res, err := HammingDistance(a, b)
	assert.NotNil(t, err)
	b = []byte{1, 1}
	res, err = HammingDistance(a, b)
	assert.Nil(t, err)
	assert.Equal(t, 2, res)
	a = []byte("this is a test")
	b = []byte("wokka wokka!!!")
	res, err = HammingDistance(a, b)
	assert.Nil(t, err)
	assert.Equal(t, 37, res)
}

func TestFindKeySize(t *testing.T) {
	enc, err := base64.StdEncoding.DecodeString(
		"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")
	assert.Nil(t, err)
	fmt.Printf("len(enc): %d\n", len(enc))
	key_size, err := FindKeysize(enc, 7)
	assert.Nil(t, err)
	assert.Equal(t, 3, key_size)
}
