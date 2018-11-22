package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
)

// CommandLineInterface struct for handling command line interface
type CommandLineInterface struct{}

// PrintUsage for printing usage instructions
func (commandLineInterface *CommandLineInterface) PrintUsage() {
	fmt.Println("USAGE:")
	fmt.Println("    build  : go run main.go   <OPTIONS>")
	fmt.Println("    release: ./go-block-chain <OPTIONS>")
	fmt.Println("getbalance -address ADDRESS           - get balance for address.")
	fmt.Println("createblockchain -address ADDRESS     - creates a blockchain.")
	fmt.Println("printchain                            - prints the blocks in the blockchain.")
	fmt.Println("send -from FROM -to TO -amount AMOUNT - send amount from an address to an address.")
}

// ValidateArguments to validate the arguments for the CommandLineInterface
func (commandLineInterface *CommandLineInterface) ValidateArguments() {
	if len(os.Args) < 2 {
		commandLineInterface.PrintUsage()
		runtime.Goexit()
	}
}

// PrintChain to print the Blocks in the BlockChain from commandLineInterface
func (commandLineInterface *CommandLineInterface) PrintChain() {
	chain := blockchain.ContinueBlockChain("")
	defer chain.DataBase.Close()
	iterator := chain.Iterator()
	for {
		block := iterator.Next()
		fmt.Printf("PREVIOUS HASH: %x\n", block.PreviousHash)
		fmt.Printf("MAIN HASH: %x\n", block.Hash)
		proofOfWork := blockchain.NewProof(block)
		fmt.Printf("PROOF OF WORK: %s\n", strconv.FormatBool(proofOfWork.Validate()))
		if len(block.PreviousHash) == 0 {
			break
		}
	}
}

func (commandLineInterface *CommandLineInterface) CreateBlockChain(address string) {
	chain := blockchain.InitBlockChain(address)
	chain.DataBase.Close()
	fmt.Println("FINISHED CREATING BLOCKCHAIN.")
}

func (commandLineInterface *CommandLineInterface) GetBalance(address string) {
	chain := blockchain.ContinueBlockChain(address)
	defer chain.DataBase.Close()
	balance := 0
	unSpentTransactionOutputs := chain.FindUnspentTransactionsOutputs(address)
	for _, output := range unSpentTransactionOutputs {
		balance += output.Value
	}
	fmt.Printf("Balance of %s: %d\n", address, balance)
}

func (commandLineInterface *CommandLineInterface) Send(from, to string, amount int) {
	chain := blockchain.ContinueBlockChain(from)
	defer chain.DataBase.Close()
	tx := blockchain.NewTransaction(from, to, amount, chain)
	chain.AddBlock([]*blockchain.Transaction{tx})
	fmt.Println("SUCCESS.")
}

func (commandLineInterface *CommandLineInterface) Run() {
	commandLineInterface.ValidateArguments()
	getBalanceCommand := flag.NewFlagSet("getbalance", flag.ExitOnError)
	createBlockChainCommand := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCommand := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCommand := flag.NewFlagSet("printchain", flag.ExitOnError)
	getBalanceAddress := getBalanceCommand.String("address", "", "The Address to find Balance.")
	createBlockChainAddress := createBlockChainCommand.String("address", "", "The Address to send Reward to.")
	sendFrom := sendCommand.String("from", "", "Source Wallet Address")
	sendTo := sendCommand.String("to", "", "Destination Wallet Address")
	sendAmount := sendCommand.Int("amount", 0, "Amount To Send")
	switch os.Args[1] {
	case "getbalance":
		err := getBalanceCommand.Parse(os.Args[2:])
		blockchain.Handle(err)
	case "createblockchain":
		err := createBlockChainCommand.Parse(os.Args[2:])
		blockchain.Handle(err)
	case "printchain":
		err := printChainCommand.Parse(os.Args[2:])
		blockchain.Handle(err)
	case "send":
		err := sendCommand.Parse(os.Args[2:])
		blockchain.Handle(err)
	default:
		commandLineInterface.PrintUsage()
		runtime.Goexit()
	}
	if getBalanceCommand.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCommand.Usage()
			runtime.Goexit()
		}
		commandLineInterface.GetBalance(*getBalanceAddress)
	}
	if createBlockChainCommand.Parsed() {
		if *createBlockChainAddress == "" {
			createBlockChainCommand.Usage()
			runtime.Goexit()
		}
		commandLineInterface.CreateBlockChain(*createBlockChainAddress)
	}
	if printChainCommand.Parsed() {
		commandLineInterface.PrintChain()
	}
	if sendCommand.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCommand.Usage()
			runtime.Goexit()
		}
		commandLineInterface.Send(*sendFrom, *sendTo, *sendAmount)
	}
}

func main() {
	defer os.Exit(0)
	commandLineIterface := CommandLineInterface{}
	commandLineIterface.Run()
}
