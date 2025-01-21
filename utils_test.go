package main

import (
	"reflect"
	"testing"
)

func TestConvertHexStringToByteArray(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expected := []byte{73, 39, 109, 32, 107, 105, 108, 108, 105, 110, 103, 32, 121, 111, 117, 114, 32, 98, 114, 97, 105, 110, 32, 108, 105, 107, 101, 32, 97, 32, 112, 111, 105, 115, 111, 110, 111, 117, 115, 32, 109, 117, 115, 104, 114, 111, 111, 109}

	result, err := ConvertHexStringToByteArray(input)
	if !reflect.DeepEqual(result, expected) || err != nil {
		t.Fatalf(`ConvertHexStringToByteArray(%s) = %v, %v, want = %v`, input, result, err, expected)
	}
}

func TestFixedXOR(t *testing.T) {
	aBuffer, _ := ConvertHexStringToByteArray("1c0111001f010100061a024b53535009181c")
	bBuffer, _ := ConvertHexStringToByteArray("686974207468652062756c6c277320657965")
	expected, _ := ConvertHexStringToByteArray("746865206b696420646f6e277420706c6179")

	result, err := FixedXOR(aBuffer, bBuffer)
	if err != nil {
		t.Fatalf(`FixedXOR(...,...), error = %v`, err)
	}
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf(`error`)
	}
}
