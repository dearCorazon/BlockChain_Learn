package main


const dbFile = "blockchain.db"
const blockBucket  = "blocks"

type Blockchain struct {
	//blocks []*Block
	tip []byte
	db *bolt.DB
}
type BlockchainIterator struct {
	 currentHash []byte
	 db *bolt.DB
}
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
