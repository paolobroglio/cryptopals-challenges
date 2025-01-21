package main

import "testing"

func TestSingleByteXORCipher(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	expected := "Cooking MC's like a pound of bacon"
	key := []byte("X")
	result, err := SingleByteXORCipher(input, key[0])
	if string(result) != expected || err != nil {
		t.Fatalf(`Encode(%s) = %q, %v, want = %s`, input, result, err, expected)
	}
}
