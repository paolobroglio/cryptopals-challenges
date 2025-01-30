package main

import "fmt"

func main() {
	res := CharacterFrequencySingleByteXORCipher("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Printf("Decrypted: %s\n", res)
}
