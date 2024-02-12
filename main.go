package main

import "fmt"

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Sending 1 BTC to someone")
	bc.AddBlock("Sending another amount of BTC to other person")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
