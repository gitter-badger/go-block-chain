package blockchain

import (
	"fmt"

	"github.com/dgraph-io/badger"
)

const (
	dbPath = "./tmp/blocks"
)

// BlockChain structure for the BlockChain dataType
type BlockChain struct {
	// Blocks []*Block
	LastHash []byte
	DataBase *badger.DB
}

// ChainIterator to go through all the Blocks in badger.DB
type ChainIterator struct {
	CurrentHash []byte
	DataBase    *badger.DB
}

// InitBlockChain to initialize the BlockChain
func InitBlockChain() *BlockChain {
	// return &BlockChain{[]*Block{Genesis()}}
	var lastHash []byte
	options := badger.DefaultOptions
	options.Dir = dbPath
	options.ValueDir = dbPath
	database, err := badger.Open(options)
	Handle(err)
	err = database.Update(func(txn *badger.Txn) error {
		if _, err := txn.Get([]byte("lh")); err == badger.ErrKeyNotFound {
			fmt.Println("NO EXISTING BLOCKCHAIN FOUND.")
			genesis := Genesis()
			fmt.Println("GENESIS PROVED.")
			err = txn.Set(genesis.Hash, genesis.Serialize())
			Handle(err)
			err = txn.Set([]byte("lh"), genesis.Hash)
			lastHash = genesis.Hash
			return err
		}
		item, err := txn.Get([]byte("lh"))
		lastHash, err = item.Value()
		return err
	})
	Handle(err)
	blockchain := BlockChain{lastHash, database}
	return &blockchain
}

// AddBlock to the existing BlockChain
func (blockChain *BlockChain) AddBlock(data string) {
	// previousBlock := blockChain.Blocks[len(blockChain.Blocks)-1]
	// newBlock := CreateBlock(data, previousBlock.Hash)
	// blockChain.Blocks = append(blockChain.Blocks, newBlock)
	var lastHash []byte
	err := blockChain.DataBase.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.Value()
		return err
	})
	Handle(err)
	newBlock := CreateBlock(data, lastHash)
	err = blockChain.DataBase.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)
		blockChain.LastHash = newBlock.Hash
		return err
	})
	Handle(err)
}

// Iterator to get the Iterator over the badgerDB
func (blockChain *BlockChain) Iterator() *ChainIterator {
	iterator := &ChainIterator{blockChain.LastHash, blockChain.DataBase}
	return iterator
}

// Next to iterate to the next Block in the badgerDB
func (iterator *ChainIterator) Next() *Block {
	var block *Block
	err := iterator.DataBase.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iterator.CurrentHash)
		Handle(err)
		encodedBlock, err := item.Value()
		block = Deserialize(encodedBlock)
		return err
	})
	Handle(err)
	iterator.CurrentHash = block.PreviousHash
	return block
}
