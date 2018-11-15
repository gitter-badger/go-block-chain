package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

// Transaction structure for the Transaction
type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

// TxOutput for Output for the BlockChainw
type TxOutput struct {
	Value     int
	PublicKey string
}

// TxInput for Input for the BlockChain
type TxInput struct {
	ID        []byte
	Output    int
	Signature string
}

// SetID for setting the ID for the Transaction
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte
	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	Handle(err)
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// CoinBaseTx for the coin base transaction
func CoinBaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("COINS TO %s\n", to)
	}
	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{100, to}
	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	tx.SetID()
	return &tx
}

// IsCoinBase to check for CoinBase Transaction
func (tx *Transaction) IsCoinBase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Output == -1
}

// CanUnlock for checking who can unlock the coinbase transaction
func (in *TxInput) CanUnlock(data string) bool {
	return in.Signature == data
}

// CanBeUnlocked to check whether the transaction can be unlocked
func (out *TxOutput) CanBeUnlocked(data string) bool {
	return out.PublicKey == data
}

// NewTransaction for creating a new Transaction in the BlockChain
func NewTransaction(from, to string, amount int, blockchain *BlockChain) *Transaction {
	var inputs []TxInput
	var outputs []TxOutput
	accumulator, validOutputs := blockchain.FindSpendableOutputs(from, amount)
	if accumulator < amount {
		log.Panic("ERROR: NOT ENOUGH FUNDS !")
	}
	for txID, outputs := range validOutputs {
		txIDString, err := hex.DecodeString(txID)
		Handle(err)
		for _, output := range outputs {
			input := TxInput{txIDString, output, from}
			inputs = append(inputs, input)
		}
	}
	outputs = append(outputs, TxOutput{amount, to})
	if accumulator > amount {
		outputs = append(outputs, TxOutput{accumulator - amount, from})
	}
	tx := Transaction{nil, inputs, outputs}
	tx.SetID()
	return &tx
}
