package main

import (
	"fmt"
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
