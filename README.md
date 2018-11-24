# [<img src="https://raw.githubusercontent.com/the-code-innovator/go-block-chain/master/images/mascot.png" width=40 height=40>](https://golang.org)**golang** Block Chain

[<img src="https://raw.githubusercontent.com/the-code-innovator/go-block-chain/master/images/github.png" width=40 height=40>](https://github.com/the-code-innovator/go-block-chain)
[<img src="https://raw.githubusercontent.com/the-code-innovator/go-block-chain/master/images/gitter.png" width=40 height=40>](https://gitter.im/go-block-chain/Lobby)

## [Blockchain](https://en.wikipedia.org/wiki/Blockchain):
* A blockchain, originally block chain, is a growing list of records, called blocks, which are linked using cryptography.
* Each block contains a cryptographic hash of the previous block, a timestamp, and transaction data (generally represented as a merkle tree root hash).

## Description:
a block chain in [**golang**](https://golang.org) with command line interface builtin.

## Language:
* [golang](https://golang.org)

## Usage:
```
USAGE:
    -> dev   : go run main.go   <OPTIONS>
    -> build : ./go-block-chain <OPTIONS>
 • getbalance -address ADDRESS           - get balance for address.
 • createblockchain -address ADDRESS     - creates a blockchain.
 • printchain                            - prints the blocks in the blockchain.
 • send -from FROM -to TO -amount AMOUNT - send amount from an address to an address.
 • createwallet                          - creates a new wallet.
 • listaddresses                         - lists the addresses in our wallet file.
```

## Utilities:
* getbalance:
   ```bash
   $ $EXECUTABLE getbalance -address ADDRESS
   ```
   - To get balance of address 'ADDRESS'.
* createblockchain:
   ```bash
   $ $EXECUTABLE createblockchain -address ADDRESS
   ```
   - To create a blockchain and send reward to the address 'ADDRESS'.
* printchain:
   ```bash
   $ $EXECUTABLE printchain
   ```
   - To print the blocks in the blockchain.
* send:
   ```bash
   $ $EXECUTABLE send -from FROM -to TO -amount AMOUNT
   ```
   - To send amount AMOUNT from address 'FROM' to address 'TO'.
* createwallet:
   ```bash
   $ $EXECUTABLE createwallet
   ```
   - To create a wallet and store it in the wallets database.
* listaddresses:
   ```bash
   $ $EXECUTABLE listaddresses
   ```
   - To list all public addresses in the wallets database.

`$EXECUTABLE` evaluvates to:
   - dev:
      ```
      go run main.go
      ```
   - build:
      - First build the module.
      ```bash
      $ go build
      ```
      - After building the module.
      ```bash
      $PWD/go-block-chain
      ```
