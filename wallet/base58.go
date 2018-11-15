package wallet

import (
	"github.com/mr-tron/base58"
	"log"
)

// Base58Encode to assist in encoding the value
func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)
	return []byte(encode)
}

// Base58Decode to assist in getting the decoded value
func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input[:]))
	Handle(err)
	return decode

}

// Handle to handle errors
func Handle(err error) {
	if err != nil {
		log.Panic(err)
	}
}
