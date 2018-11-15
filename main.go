package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/the-code-innovator/go-block-chain/blockchain"
)

// CommandLineInterface struct for handling command line interface
type CommandLineInterface struct {
	blockchain *blockchain.BlockChain
}

func (commandLineInterface *CommandLineInterface) printUsage() {
	fmt.Println("USAGE:")
	fmt.Println("add -block BLOCK_DATA - to add a block to the blockchain.")
	fmt.Println("print 				   - prints the blocks in the blockchain.")
}

// ValidateArguments to validate the arguments for the CommandLineInterface
func (commandLineInterface *CommandLineInterface) ValidateArguments() {
	if len(os.Args) < 2 {
		commandLineInterface.printUsage()
		runtime.Goexit()
	}
}

// AddBlock to call inherent AddBlock into the BlockChain from commandLineInterface
func (commandLineInterface *CommandLineInterface) AddBlock(data string) {
	commandLineInterface.blockchain.AddBlock(data)
	fmt.Println("ADDED BLOCK TO BLOCKCHAIN.")
}

// PrintChain to print the Blocks in the BlockChain from commandLineInterface
func (commandLineInterface *CommandLineInterface) PrintChain() {
	iterator := commandLineInterface.blockchain.Iterator()
	for {
		block := iterator.Next()
		fmt.Printf("PREVIOUS HASH: %x\n", block.PreviousHash)
		fmt.Printf("DATA IN BLOCK: %s\n", block.Data)
		fmt.Printf("MAIN HASH: %x\n", block.Hash)
		proofOfWork := blockchain.NewProof(block)
		fmt.Printf("PROOF OF WORK: %s\n", strconv.FormatBool(proofOfWork.Validate()))
		fmt.Println()
		if len(block.PreviousHash) == 0 {
			break
		}
	}
}

func (commandLineInterface *CommandLineInterface) run() {
	commandLineInterface.ValidateArguments()
	addBlockCommand := flag.NewFlagSet("add", flag.ExitOnError)
	printChainCommand := flag.NewFlagSet("print", flag.ExitOnError)
	addBlockData := addBlockCommand.String("block", "", "BLOCK_DATA")
	switch os.Args[1] {
	case "add":
		err := addBlockCommand.Parse(os.Args[2:])
		blockchain.Handle(err)
	case "print":
		err := printChainCommand.Parse(os.Args[2:])
		blockchain.Handle(err)
	case "default":
		commandLineInterface.printUsage()
		runtime.Goexit()
	}
	if addBlockCommand.Parsed() {
		if *addBlockData == "" {
			addBlockCommand.Usage()
			runtime.Goexit()
		}
		commandLineInterface.AddBlock(*addBlockData)
	}
	if printChainCommand.Parsed() {
		commandLineInterface.PrintChain()
	}
}

func main() {
	defer os.Exit(0)
	chain := blockchain.InitBlockChain()
	// chain.AddBlock("1st after genesis.")
	// chain.AddBlock("2nd after genesis.")
	// chain.AddBlock("3rd after genesis.")
	// for _, block := range chain.Blocks {
	// 	fmt.Printf("PREVIOUS HASH: %x\n", block.PreviousHash)
	// 	fmt.Printf("DATA IN BLOCK: %s\n", block.Data)
	// 	fmt.Printf("HASH: %x\n", block.Hash)
	// 	proofOfWork := blockchain.NewProof(block)
	// 	fmt.Printf("ProofOfWork: %s\n", strconv.FormatBool(proofOfWork.Validate()))
	// 	fmt.Println()
	// }
	defer chain.DataBase.Close()
	commandLinenIterface := CommandLineInterface{chain}
	commandLinenIterface.run()
}
