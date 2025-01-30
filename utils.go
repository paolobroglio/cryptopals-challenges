package main

import (
	"fmt"
	"math"
	"strconv"
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

func FixedXOR(aBuffer []byte, bBuffer []byte) ([]byte, error) {
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

var englishFreq = [26]float64{
	0.08167, 0.01492, 0.02782, 0.04253, 0.12702, 0.02228, 0.02015, // A-G
	0.06094, 0.06966, 0.00153, 0.00772, 0.04025, 0.02406, 0.06749, // H-N
	0.07507, 0.01929, 0.00095, 0.05987, 0.06327, 0.09056, 0.02758, // O-U
	0.00978, 0.02360, 0.00150, 0.01974, 0.00074, // V-Z
}

func GetChi2(str string) float64 {
	count := [26]int{}
	ignored := 0

	for _, c := range str {
		if c >= 'A' && c <= 'Z' {
			count[c-'A']++
		} else if c >= 'a' && c <= 'z' {
			count[c-'a']++
		} else if c >= ' ' && c <= '~' {
			ignored++
		} else if c == '\t' || c == '\n' || c == '\r' {
			ignored++
		} else {
			return math.Inf(1)
		}
	}

	lenStr := len(str) - ignored
	if lenStr == 0 {
		return math.Inf(1)
	}

	chi2 := 0.0
	for i := 0; i < 26; i++ {
		observed := float64(count[i])
		expected := float64(lenStr) * englishFreq[i]
		difference := observed - expected
		if expected > 0 {
			chi2 += (difference * difference) / expected
		}
	}

	return chi2
}
