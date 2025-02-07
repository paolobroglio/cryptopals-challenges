package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func singleByteXORCipher(data []byte, key byte) []byte {
	decoded := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		decodedByte := data[i] ^ key
		if decodedByte == 0 {
			decodedByte = ' '
		}
		decoded[i] = decodedByte
	}
	return decoded
}

func detectSingleByteXOR(hexString string) (string, float64) {

	data, _ := hex.DecodeString(hexString) // TODO: error handling
	bestScore := 0.0
	bestPlaintext := ""

	for key := 0; key < 256; key++ {
		plaintext := string(singleByteXORCipher(data, byte(key)))
		score := scoreText(plaintext)
		if score > bestScore {
			bestScore = score
			bestPlaintext = plaintext
		}
	}
	return bestPlaintext, bestScore
}

func detectSingleCharacterXORInFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	bestOverallScore := 0.0
	bestOverallPlaintext := ""

	for scanner.Scan() {
		plaintext, score := detectSingleByteXOR(scanner.Text())
		if score > bestOverallScore {
			bestOverallScore = score
			bestOverallPlaintext = plaintext
		}
	}

	fmt.Printf("Best decoded string: %q\n", bestOverallPlaintext)
}

func repeatingKeyXORCipher(plaintext string, key string) string {
	data := []byte(plaintext)
	encrypted := make([]byte, 0)
	for i, b := range data {
		currKeyByte := byte(key[i%len(key)])
		encrypted = append(encrypted, b^currKeyByte)
	}
	hexString := ConvertByteArrayToHexString(encrypted)
	fmt.Printf("Best decoded string: %s\n", hexString)
	return ""
}
