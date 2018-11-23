# golang blockchain
[<img src="https://raw.githubusercontent.com/the-code-innovator/go-block-chain/master/images/mascot.png" width=40 height=40>](https://golang.org)

## [Blockchain](https://en.wikipedia.org/wiki/Blockchain):
* A blockchain, originally block chain, is a growing list of records, called blocks, which are linked using cryptography.
* Each block contains a cryptographic hash of the previous block, a timestamp, and transaction data (generally represented as a merkle tree root hash).

## Description:
a block chain in golang with command line interface builtin.

## Language:
* [golang](https://golang.org)

## Utilities:
* getbalance:
   ```bash
   $ $EXECUTABLE getbalance -address ADDRESS
   ```
   to get balance of address 'ADDRESS'.
* createblockchain:
   ```bash
   $ $EXECUTABLE createblockchain -address ADDRESS
   ```
   to create a blockchain and send reward to the address 'ADDRESS'.
* printchain:
   ```bash
   $ $EXECUTABLE printchain
   ```
   to print the blocks in the blockchain.
* send:
   ```bash
   $ $EXECUTABLE send -from FROM -to TO -amount AMOUNT
   ```
   to send amount AMOUNT from address FROM to address TO.
* createwallet:
   ```bash
   $ $EXECUTABLE createwallet
   ```
   to create a wallet and store it in the wallets database.
* listaddresses:
   ```bash
   $ $EXECUTABLE listaddresses
   ```
   to list all public addresses in the wallets database.
* $EXECUTABLE evaluvates to:
   - build:
      ```bash
      $ go build
      ```
      build the module.
      ```bash
      $ go run main.go
      ```
   - release:
      ```bash
      $PWD/go-block-chain
      ```
