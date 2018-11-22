package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"

	"golang.org/x/crypto/ripemd160"
)

const (
	checkSumLength = 4
	version        = byte(0x00)
)

// Wallet structure for the Wallet
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

// Address for finding the address of the Wallet
func (wallet Wallet) Address() []byte {
	publicKeyHash := PublicKeyHash(wallet.PublicKey)
	versionedHash := append([]byte{version}, publicKeyHash...)
	checkSum := GenerateCheckSum(versionedHash)
	completeHash := append(versionedHash, checkSum...)
	address := Base58Encode(completeHash)
	fmt.Printf("PUBLIC KEY: %x\n", wallet.PublicKey)
	fmt.Printf("PUBLIC HASH: %x\n", publicKeyHash)
	fmt.Printf("ADDRESS: %x\n", address)
	return address
}

// NewKeyPair for creating a new KeyPair
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	Handle(err)
	public := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, public
}

func MakeWallet() *Wallet {
	private, public := NewKeyPair()
	wallet := Wallet{private, public}
	return &wallet
}

func PublicKeyHash(publicKey []byte) []byte {
	publicKeyHash := sha256.Sum256(publicKey)
	ripemd160Hasher := ripemd160.New()
	_, err := ripemd160Hasher.Write(publicKeyHash[:])
	Handle(err)
	publicRipeMD160Hash := ripemd160Hasher.Sum(nil)
	return publicRipeMD160Hash
}

func GenerateCheckSum(payload []byte) []byte {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])
	return secondHash[:checkSumLength]
}
