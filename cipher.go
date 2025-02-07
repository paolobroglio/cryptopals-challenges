package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func singleByteXORCipher(input string, key byte) (string, error) {
	data, err := hex.DecodeString(input)
	if err != nil {
		return "", err
	}
	decoded := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		decodedByte := data[i] ^ key
		if decodedByte == 0 {
			decodedByte = ' '
		}
		decoded[i] = decodedByte
	}
	return string(decoded), nil
}

func detectSingleByteXOR(hexString string) (string, float64) {

	bestScore := 0.0
	bestPlaintext := ""

	for key := 0; key < 256; key++ {
		plaintext, _ := singleByteXORCipher(hexString, byte(key))
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
