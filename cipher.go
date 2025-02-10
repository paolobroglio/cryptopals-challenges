package main

import (
	"bufio"
	"fmt"
	"os"
)

type SingleByteXORResult struct {
	plaintext string
	score     float64
}

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

func detectSingleByteXOR(hexString string) (SingleByteXORResult, error) {

	data, err := ConvertHexStringToByteArray(hexString)
	if err != nil {
		return SingleByteXORResult{}, err
	}
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

	result := SingleByteXORResult{
		plaintext: bestPlaintext,
		score:     bestScore,
	}
	return result, nil
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
		result, err := detectSingleByteXOR(scanner.Text())
		if err != nil {
			fmt.Println("Error during decryption:", err)
			return
		}
		if result.score > bestOverallScore {
			bestOverallScore = result.score
			bestOverallPlaintext = result.plaintext
		}
	}

	fmt.Printf("Best decoded string: %q\n", bestOverallPlaintext)
}

func repeatingKeyXORCipher(plaintext string, key string) string {
	data := []byte(plaintext)
	encrypted := make([]byte, len(data))
	for i, b := range data {
		keyIndex := i % len(key)
		currKeyByte := byte(key[keyIndex])
		xor := b ^ currKeyByte
		encrypted[i] = xor
	}
	hexString := ConvertByteArrayToHexString(encrypted)
	return hexString
}
