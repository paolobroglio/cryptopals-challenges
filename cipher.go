package main

import (
	"fmt"
	"sync"
)

func SingleByteXORCipher(input string, key byte) ([]byte, error) {
	data, err := ConvertHexStringToByteArray(input)
	if err != nil {
		return nil, err
	}
	decoded := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		decodedByte := data[i] ^ key
		decoded[i] = decodedByte
	}
	return decoded, nil
}

func CharacterFrequencySingleByteXORCipher(input string) {
	lettersFrequencies := GetLettersFrequencies()

	var wg sync.WaitGroup
	results := make(chan []byte)

	for _, letterFrequency := range lettersFrequencies {
		wg.Add(1)
		key := []byte(string(letterFrequency))
		go func() {
			defer wg.Done()
			data, _ := SingleByteXORCipher(input, key[0])
			results <- data
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Printf("Deciphered: %s\n", string(res))
	}

}
