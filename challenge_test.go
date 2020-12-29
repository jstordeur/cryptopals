package challenge

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jstordeur/cryptopals/frequency"
	"github.com/jstordeur/cryptopals/utils"
)

func TestChallenge1(t *testing.T) {
	input := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	res, err := utils.HexToBase64(input)
	assert.Nil(t, err)
	assert.Equal(t, "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t", res)
}

func TestChallenge2(t *testing.T) {
	input1, err := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	assert.Nil(t, err)
	input2, err := hex.DecodeString("686974207468652062756c6c277320657965")
	assert.Nil(t, err)
	expected, err := hex.DecodeString("746865206b696420646f6e277420706c6179")
	assert.Nil(t, err)
	res, err := utils.Xor(input1, input2)
	assert.Nil(t, err)
	assert.Equal(t, expected, res)
}

func TestChallenge3(t *testing.T) {
	enc, err := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	assert.Nil(t, err)
	xor, _ := frequency.TestForSingleByteXor(enc)
	key := make([]byte, len(enc))
	for pos, _ := range key {
		key[pos] = xor
	}
	dec, err := utils.Xor(enc, key)
	assert.Nil(t, err)
	assert.Equal(t, "Cooking MC's like a pound of bacon", string(dec))
}

func TestChallenge4(t *testing.T) {
	file, err := os.Open("./input/challenge4.txt")
	assert.Nil(t, err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	min_chi := math.MaxFloat64
	var dec []byte
	for scanner.Scan() {
		input := scanner.Text()
		enc, err := hex.DecodeString(input)
		xor, chi := frequency.TestForSingleByteXor(enc)
		if chi < min_chi {
			min_chi = chi
			key := make([]byte, len(enc))
			for pos, _ := range key {
				key[pos] = xor
			}
			dec, err = utils.Xor(enc, key)
			assert.Nil(t, err)
		}
	}
	err = scanner.Err()
	assert.Nil(t, err)
	assert.Equal(t, "Now that the party is jumping\n", string(dec))
}

func TestChallenge5(t *testing.T) {
	input := []byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal")
	key := []byte("ICE")
	expected, err := hex.DecodeString("0b3637272a2b2e63622c2e69692a23693a2a3" +
		"c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333" +
		"a653e2b2027630c692b20283165286326302e27282f")
	assert.Nil(t, err)
	assert.Equal(t, expected, utils.RepeatXor(input, key))
}

func TestChallenge6(t *testing.T) {
	file, err := os.Open("./input/challenge6.txt")
	assert.Nil(t, err)
	defer file.Close()

	stat, err := file.Stat()
	assert.Nil(t, err)
	enc := make([]byte, stat.Size()*8+1)
	n, err := file.Read(enc)
	assert.Nil(t, err)
	assert.Equal(t, stat.Size(), int64(n))
	key_size, err := frequency.FindKeysize(enc, 20)
	assert.Nil(t, err)
	fmt.Printf("size: %d\n", key_size)
}
