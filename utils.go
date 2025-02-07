package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertHexStringToByteArray(hexString string) ([]byte, error) {
	byteArray := make([]byte, len(hexString)/2)
	for i := 0; i < len(hexString); i += 2 {
		b, err := strconv.ParseUint(hexString[i:i+2], 16, 8)
		if err != nil {
			return nil, err
		}
		byteArray[i/2] = byte(b)
	}
	return byteArray, nil
}

func ConvertByteArrayToHexString(bytes []byte) string {
	hex := ""
	for _, b := range bytes {
		hex += fmt.Sprintf("%02x", b)
	}
	return hex
}

func fixedXOR(aBuffer []byte, bBuffer []byte) ([]byte, error) {
	if len(aBuffer) != len(bBuffer) {
		return nil, fmt.Errorf("cannot do fixed XOR on different length buffers")
	}
	cBuffer := make([]byte, len(aBuffer))
	for i := 0; i < len(aBuffer); i++ {
		aByte := aBuffer[i]
		bByte := bBuffer[i]

		xor := aByte ^ bByte
		cBuffer[i] = xor
	}
	return cBuffer, nil
}

// Function to compute English letter frequency score
func scoreText(text string) float64 {
	frequency := map[rune]float64{
		'e': 12.70, 't': 9.06, 'a': 8.17, 'o': 7.51, 'i': 6.97,
		'n': 6.75, 's': 6.33, 'h': 6.09, 'r': 5.99, 'd': 4.25,
		'l': 4.03, 'c': 2.78, 'u': 2.76, 'm': 2.41, 'w': 2.36,
		'f': 2.23, 'g': 2.02, 'y': 1.97, 'p': 1.93, 'b': 1.29,
		'v': 0.98, 'k': 0.77, 'j': 0.15, 'x': 0.15, 'q': 0.10, 'z': 0.07,
	}
	score := 0.0
	for _, char := range strings.ToLower(text) {
		if val, exists := frequency[char]; exists {
			score += val
		}
	}
	return score
}
