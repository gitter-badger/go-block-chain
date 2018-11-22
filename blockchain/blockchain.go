package blockchain

import (
	"encoding/hex"
	"fmt"
	"os"
	"runtime"

	"github.com/dgraph-io/badger"
)

// constants for the blockchain
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

// badgerDBExists for checking the availability of database
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
	err1 := database.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		return err
	})
	Handle(err1)
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
func (blockChain *BlockChain) AddBlock(transactions []*Transaction) {
	var lastHash []byte
	err1 := blockChain.DataBase.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		lastHash, err = item.Value()
		return err
	})
	Handle(err1)
	newBlock := CreateBlock(transactions, lastHash)
	err2 := blockChain.DataBase.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)
		blockChain.LastHash = newBlock.Hash
		return err
	})
	Handle(err2)
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

// FindUnspentTransactions in the BlockChain
func (blockchain *BlockChain) FindUnspentTransactions(address string) []Transaction {
	var unSpentTransactions []Transaction
	spentTxns := make(map[string][]int)
	iterator := blockchain.Iterator()
	for {
		block := iterator.Next()
		for _, tx := range block.Transactions {
			txID := hex.EncodeToString(tx.ID)
		OutputIterate:
			for outID, out := range tx.Outputs {
				if spentTxns[txID] != nil {
					for _, spentOut := range spentTxns[txID] {
						if spentOut == outID {
							continue OutputIterate
						}
					}
				}
				if out.CanBeUnlocked(address) {
					unSpentTransactions = append(unSpentTransactions, *tx)
				}
			}
			if tx.IsCoinBase() == false {
				for _, in := range tx.Inputs {
					if in.CanUnlock(address) {
						inTxID := hex.EncodeToString(in.ID)
						spentTxns[inTxID] = append(spentTxns[inTxID], in.Output)
					}
				}
			}
		}
		if len(block.PreviousHash) == 0 {
			break
		}
	}
	return unSpentTransactions
}

// FindUnspentTransactionsOutputs for getting the unspent transactions in BlockChain
func (blockchain *BlockChain) FindUnspentTransactionsOutputs(address string) []TxOutput {
	var unSpentTransactionOutputs []TxOutput
	unSpentTransactions := blockchain.FindUnspentTransactions(address)
	for _, tx := range unSpentTransactions {
		for _, output := range tx.Outputs {
			if output.CanBeUnlocked(address) {
				unSpentTransactionOutputs = append(unSpentTransactionOutputs, output)
			}
		}
	}
	return unSpentTransactionOutputs
}

// FindSpendableOutputs to find the spendable outputs in the BlockChain
func (blockchain *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unSpentOutputs := make(map[string][]int)
	unSpentTransactions := blockchain.FindUnspentTransactions(address)
	accumulated := 0
Work:
	for _, tx := range unSpentTransactions {
		txID := hex.EncodeToString(tx.ID)
		for outID, output := range tx.Outputs {
			if output.CanBeUnlocked(address) && accumulated < amount {
				accumulated += output.Value
				unSpentOutputs[txID] = append(unSpentOutputs[txID], outID)
				if accumulated >= amount {
					break Work
				}
			}
		}
	}
	return accumulated, unSpentOutputs
}
