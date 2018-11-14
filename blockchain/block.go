package blockchain

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Block structure for the Block dataType
type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
	Nonce        int
}

// DeriveHash from the PreviousHash of the same BlockChain
// func (block *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{block.Data, block.PreviousHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	block.Hash = hash[:]
// }

// CreateBlock from the PreviousHash of the same BlockChain
func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash, 0}
	// block.DeriveHash()
	proofOfWork := NewProof(block)
	nonce, hash := proofOfWork.Run()
	block.Nonce = nonce
	block.Hash = hash
	return block
}

// Genesis of the BlockChain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// Serialize for serializing the output for BadgerDB
func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	Handle(err)
	return result.Bytes()
}

// Deserialize from input BadgerDB for Block
func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	Handle(err)
	return &block
}

// Handle to handle the error
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
