package blockchain

import (
	"fmt"
	"os"
	"runtime"

	"github.com/dgraph-io/badger"
)

const (
	dbPath      = "./tmp/blocks"
	dbFile      = "./tmp/blocks/MANIFEST"
	genesisData = "FIRST TRANSACTION FROM GENESIS."
)

// BlockChain structure for the BlockChain dataType
type BlockChain struct {
	LastHash []byte
	DataBase *badger.DB
}

// ChainIterator to go through all the Blocks in badger.DB
type ChainIterator struct {
	CurrentHash []byte
	DataBase    *badger.DB
}

func badgerDBExists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}
	return true
}

// ContinueBlockChain to continue runnings through blockchain validation
func ContinueBlockChain(address string) *BlockChain {
	if badgerDBExists() == false {
		fmt.Println("NO EXISTING BLOCKCHAIN FOUND.\nCREATE ONE.")
		runtime.Goexit()
	}
	var lastHash []byte
	options := badger.DefaultOptions
	options.Dir = dbPath
	options.ValueDir = dbPath
	database, err := badger.Open(options)
	Handle(err)
	err = database.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.Value()
		return err
	})
	Handle(err)
	blockchain := BlockChain{lastHash, database}
	return &blockchain
}

// InitBlockChain to initialize the BlockChain
func InitBlockChain(address string) *BlockChain {
	if badgerDBExists() {
		fmt.Println("BLOCKCHAIN ALREADY EXISTS.")
		runtime.Goexit()
	}
	var lastHash []byte

	options := badger.DefaultOptions
	options.Dir = dbPath
	options.ValueDir = dbPath
	database, err := badger.Open(options)
	Handle(err)
	err = database.Update(func(txn *badger.Txn) error {
		coinBaseTransaction := CoinBaseTx(address, genesisData)
		genesis := Genesis(coinBaseTransaction)
		fmt.Println("GENESIS CREATED.")
		err = txn.Set(genesis.Hash, genesis.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), genesis.Hash)
		lastHash = genesis.Hash
		return err
	})
	Handle(err)
	blockchain := BlockChain{lastHash, database}
	return &blockchain
}

// AddBlock to the existing BlockChain
func (blockChain *BlockChain) AddBlock(data string) {
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
