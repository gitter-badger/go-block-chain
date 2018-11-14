package blockchain

// Block structure for the Block dataType
type Block struct {
	Hash         []byte
	Data         []byte
	PreviousHash []byte
	Nonce        int
}

// BlockChain structure for the BlockChain dataType
type BlockChain struct {
	Blocks []*Block
}

// DeriveHash from the PreviousHash of the same BlockChain
// func (block *Block) DeriveHash() {
// 	info := bytes.Join([][]byte{block.Data, block.PreviousHash}, []byte{})
// 	hash := sha256.Sum256(info)
// 	block.Hash = hash[:]
// }

// CreateBlock from the PreviousHash of the same BlockChain
func CreateBlock(data string, previousHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), previousHash, 0}
	// block.DeriveHash()
	proofOfWork := NewProof(block)
	nonce, hash := proofOfWork.Run()
	block.Nonce = nonce
	block.Hash = hash
	return block
}

// AddBlock to the existing BlockChain
func (blockChain *BlockChain) AddBlock(data string) {
	previousBlock := blockChain.Blocks[len(blockChain.Blocks)-1]
	newBlock := CreateBlock(data, previousBlock.Hash)
	blockChain.Blocks = append(blockChain.Blocks, newBlock)
}

// Genesis of the BlockChain
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain to initialize the BlockChain
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
