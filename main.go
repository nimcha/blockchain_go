package main

import (
	"fmt"
	"time"
)

func main() {
	bc := NewBlockchain()
	time.Sleep(1000 * time.Millisecond)
	bc.AddBlock("Send 1 BTC to Ivan")
	time.Sleep(1000 * time.Millisecond)
	bc.AddBlock("Send 2 more BTC to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Time: %d\n", block.Timestamp)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
