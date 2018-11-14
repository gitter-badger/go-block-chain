package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
}

type BlockChain struct {
	blocks []*Block
}

func (block *Block) DeriveHash() {
	info := bytes.Join([][]byte{block.Data, block.PreviousHash}, []byte{})
	hash := sha256.Sum256(info)
	block.Hash = hash[:]
}

func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash}
	block.DeriveHash()
	return block
}

func (blockChain *BlockChain) AddBlock(data string) {
	previousBlock := blockChain.blocks[len(blockChain.blocks)-1]
	newBlock := CreateBlock(data, previousBlock.Hash)
	blockChain.blocks = append(blockChain.blocks, newBlock)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()
	chain.AddBlock("1st after genesis.")
	chain.AddBlock("2nd after genesis.")
	chain.AddBlock("3rd after genesis.")
	for _, block := range chain.blocks {
		fmt.Printf("PREVIOUS HASH: %x\n", block.PreviousHash)
		fmt.Printf("DATA IN BLOCK: %s\n", block.Data)
		fmt.Printf("HASH: %x\n", block.Hash)
	}
}
