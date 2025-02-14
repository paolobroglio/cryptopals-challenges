package main

import "testing"

func TestEncode(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	result, err := encode(input)
	if result != expected || err != nil {
		t.Fatalf(`Encode(%s) = %q, %v, want = %s`, input, result, err, expected)
	}
}
