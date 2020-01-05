package main

import (
	"fmt"
	"hash"
	"crypto/sha256"
	"golang.org/x/crypto/sha3"
	"github.com/ebfe/keccak"
	"github.com/jzelinskie/whirlpool"
)

func generateAndPrint(label, text string, h hash.Hash) {
	_, err := h.Write([]byte(text))
	if err != nil {
		fmt.Printf("Failed to write hash: %v\n", err)
	}

	fmt.Printf("%s: %x\n", label, h.Sum(nil))
}

func main(){
	fmt.Println("Enter text")
	var s1 string
	fmt.Scanln(&s1)

	generateAndPrint("SHA-256", s1, sha256.New())
	generateAndPrint("SHA-254", s1, sha256.New224())
	generateAndPrint("SHA3-224", s1, sha3.New224())
	generateAndPrint("SHA3-384", s1, sha3.New384())
	generateAndPrint("Keccak-384", s1, keccak.New384())
	generateAndPrint("Whirlpool (512 bit)", s1, whirlpool.New())

}