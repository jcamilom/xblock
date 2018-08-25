package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

func main() {
	block0 := NewBlock("This is the genesis block", []byte{})
	fmt.Printf("Prev. hash: %x\n", block0.PrevBlockHash)
	fmt.Printf("Data: %s\n", block0.Data)
	fmt.Printf("Hash: %x\n", block0.Hash)
	fmt.Println()

	block1 := NewBlock("This is the first block", block0.Hash)
	fmt.Printf("Prev. hash: %x\n", block1.PrevBlockHash)
	fmt.Printf("Data: %s\n", block1.Data)
	fmt.Printf("Hash: %x\n", block1.Hash)
	fmt.Println()

}
