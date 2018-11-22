package wallet

import (
	"bytes"
	"crypto/elliptic"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
)

const walletFile = "./tmp/wallets.data"

type Wallets struct {
	Wallets map[string]*Wallet
}

func CreateWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	err := wallets.LoadFile()
	return (&wallets), err
}

func (wallets *Wallets) AddWallet() string {
	wallet := MakeWallet()
	address := fmt.Sprintf("%s", wallet.Address())
	wallets.Wallets[address] = wallet
	return address
}

func (wallets *Wallets) GetAllAddresses() []string {
	var addresses []string
	for address := range wallets.Wallets {
		addresses = append(addresses, address)
	}
	return addresses
}

func (wallets Wallets) GetWallet(address string) Wallet {
	return *wallets.Wallets[address]
}

func (wallets *Wallets) LoadFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}
	var walletsLocal Wallets
	fileContent, err := ioutil.ReadFile(walletFile)
	ReturnError(err)
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&walletsLocal)
	ReturnError(err)
	wallets.Wallets = walletsLocal.Wallets
	return nil
}

func (wallets *Wallets) SaveFile() {
	var content bytes.Buffer
	gob.Register(elliptic.P256())
	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(wallets)
	Handle(err)
	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	Handle(err)
}

func ReturnError(err error) error {
	if err != nil {
		return err
	} else {
		return nil
	}
}
