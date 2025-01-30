package main

import (
	"math"
	"sync"
)

type SingleByteXORCipherResult struct {
	score      float64
	deciphered string
}

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

func CharacterFrequencySingleByteXORCipher(input string) string {
	lettersFrequencies := GetLettersFrequencies()

	var wg sync.WaitGroup
	results := make(chan SingleByteXORCipherResult)

	for _, letterFrequency := range lettersFrequencies {
		wg.Add(1)
		key := []byte(string(letterFrequency))
		go func() {
			defer wg.Done()
			data, _ := SingleByteXORCipher(input, key[0])
			score := GetChi2(string(data))
			res := SingleByteXORCipherResult{
				score:      score,
				deciphered: string(data),
			}
			results <- res
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	mininumScore := math.Inf(1)
	mostProbableText := ""
	for res := range results {
		if res.score < mininumScore {
			mininumScore = res.score
			mostProbableText = res.deciphered
		}
	}

	return mostProbableText
}
