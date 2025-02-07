package main

import "fmt"

func main() {
	encrypted := repeatingKeyXORCipher("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal", "ICE")
	fmt.Println(encrypted)
}
