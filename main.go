package main

import (
	"fmt"
	"strconv"

	"github.com/the-code-innovator/go-block-chain/blockchain"
)

func main() {
	chain := blockchain.InitBlockChain()
	chain.AddBlock("1st after genesis.")
	chain.AddBlock("2nd after genesis.")
	chain.AddBlock("3rd after genesis.")
	for _, block := range chain.Blocks {
		fmt.Printf("PREVIOUS HASH: %x\n", block.PreviousHash)
		fmt.Printf("DATA IN BLOCK: %s\n", block.Data)
		fmt.Printf("HASH: %x\n", block.Hash)
		proofOfWork := blockchain.NewProof(block)
		fmt.Printf("ProofOfWork: %s\n", strconv.FormatBool(proofOfWork.Validate()))
		fmt.Println()
	}
}
