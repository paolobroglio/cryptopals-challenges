package main

import (
	"encoding/hex"
	"testing"
)

func TestSingleByteXORCipher(t *testing.T) {
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expected := "Cooking MC's like a pound of bacon"
	key := []byte("X")
	input, _ := hex.DecodeString(hexString)
	result := singleByteXORCipher(input, key[0])
	if string(result) != expected {
		t.Fatalf(`Encode(%s) = %q, want = %s`, input, result, expected)
	}
}

func TestRepeatingKeyXORCipher(t *testing.T) {
	inputOne := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	inputTwo := "ICE"

	result := repeatingKeyXORCipher(inputOne, inputTwo)
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	if result != expected {
		t.Fatalf(`repeatingKeyXORCipher(%s, %s) = %s, wanted = %s`, inputOne, inputTwo, result, expected)
	}
}
