/*
Package frequency implements some tools to do charachter frequency analysis.
*/
package frequency

import (
	"fmt"
	"math"
	"unicode"

	"github.com/jstordeur/cryptopals/utils"
)

// Map from character to their frequency (percentage) in the english language.
var LetterFrequencyInEnglish = map[byte]float64{
	32:  18.32, // space
	'e': 10.22,
	't': 7.51,
	'a': 6.55,
	'o': 6.20,
	'i': 5.73,
	'n': 5.70,
	's': 5.33,
	'r': 4.97,
	'h': 4.86,
	'd': 3.35,
	'l': 3.35,
	'u': 2.29,
	'c': 2.26,
	'm': 2.02,
	'f': 1.97,
	'w': 1.69,
	'g': 1.63,
	'p': 1.50,
	'y': 1.47,
	'b': 1.27,
	'v': 0.79,
	'k': 0.57,
	'x': 0.15,
	'j': 0.11,
	'q': 0.09,
	'z': 0.06,
}

// ChiSquared estimates if src is an english text. A low value indicates a
// higher probability of being an english text.
//
// For each letter, computes the square of the difference between the
// occurence of the letter in src and the expected count (using the known
// frequencies of english letters). Normalizes by dividing the difference by
// the expected count and sums over all the letters in the alphabet.
// TODO: deal better with capital letters. They are counted towards 'others' at
// the moment.
func ChiSquared(src []byte) float64 {
	src_counts := make(map[byte]int)
	others_count := 0.0
	for pos, char := range src {
		char = byte(unicode.ToLower(rune(src[pos])))
		src_counts[char] += 1
		// if char is not a letter in LetterFrequencyInEnglish count it towards
		// 'others' which are later assumed to have a frequency of zero. This isn't
		// true obvisouly, it would be better to have a table taking into account
		//  punctuation and spaces but I couldn't find the data easily.
		if _, found := LetterFrequencyInEnglish[char]; !found {
			others_count += 1
		}
	}
	res := 0.0
	for letter, freq := range LetterFrequencyInEnglish {
		expected_count := freq * float64(len(src)) / 100
		res += math.Pow(expected_count-float64(src_counts[letter]), 2) / expected_count
		//fmt.Printf("letter: %s, expected_count: %f, actual: %f, res: %f \n", string(letter), expected_count, float64(src_counts[letter]), res)
	}
	// Don't know exactly how to account for others but they need to penalize the
	// score quite a bit.
	res += others_count * others_count
	//fmt.Printf("Adding %f for others, res: %f\n", others_count, res)
	return res
}

// TestForSingleXor tries to xor src with all possible single byte key and
// returns the most likely key to yield a decrypted text in english and the
// associated chi_squared test value.
func TestForSingleByteXor(enc []byte) (key byte, chi float64) {
	res := byte(0)
	min_chi := math.MaxFloat64
	for i := 0; i <= 255; i++ {
		x := byte(i)
		key := make([]byte, len(enc))
		for pos, _ := range key {
			key[pos] = x
		}
		xored, _ := utils.Xor(enc, key)
		tmp := ChiSquared(xored)
		if tmp < min_chi {
			// fmt.Printf("new min: %f, key: %d\n", tmp, x)
			min_chi = tmp
			res = x
		}
	}
	return res, min_chi
}

// Returns the Hamming distance between two slices of bytes of equal length.
// Returns an error if the lengths are not equal.
func HammingDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("HammingDistance called with two buffers of "+
			"different sizes %d != %d ", len(a), len(b))
	}
	res := 0
	for i, _ := range a {
		xor := a[i] ^ b[i]
		for ; xor > 0; xor = xor >> 1 {
			if xor%2 == 1 {
				res += 1
			}
		}
	}
	return res, nil
}

// FindKeysize estimates the size of the key used to encrypt enc with repeated
// xor encryption. Test all key sizes up to max in bytes.
// Returns the estimated size in number of bytes.
// Returns an error if len(enc) < max*4 as it requires at least four encrypted
// blocks to work.
func FindKeysize(enc []byte, max int) (int, error) {
	if max > len(enc)/2+1 {
		return 0, fmt.Errorf("The encrypted text (len==%d) must be at least "+
			" four times the size of the maximum key size max: %d", len(enc), max)
	}
	min_avg := math.MaxFloat64
	res := 0
	for size := 1; size <= max; size++ {
		sum := 0.0
		// xor all possible pairs of the four blocks
		for i := 0; i < 3; i++ {
			for j := i + 1; j < 4; j++ {
				// fmt.Printf("i: %d, j: %d", i, j)
				d, err := HammingDistance(enc[i*size:(i+1)*size], enc[j*size:(j+1)*size])
				if err != nil {
					return 0, err
				}
				sum += float64(d) / float64(size)
				fmt.Printf("size: %d, hamming: %f\n", size, float64(d)/float64(size))
			}
		}
		if avg := float64(sum) / 6; avg < min_avg {
			min_avg = avg
			res = size
			fmt.Printf("new min: %f, min_size: %d\n", min_avg, size)
		}
	}
	return res, nil
}
