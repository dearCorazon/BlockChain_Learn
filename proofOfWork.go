package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)
var (
	maxNounce = math.MaxInt64 
)

const targetBits = 24 

type ProofOfWork struct {
	block *Block
	target *big.Int
}

//new

func  NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target,uint(256- targetBits))

	pow := &ProofOfWork{b,target} 
	
	return  pow 
}

func (pow *ProofOfWork) prepareData(nounce int) []byte {
	data := bytes.Join( 
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.TimeStamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nounce)),
		},
		[]byte{},
	)
	return data
}

func  (pow *ProofOfWork) Run() (int,[]byte) {
	var hashInt big.Int
	var hash [32]byte
	nounce := 0
	
	fmt.Printf("Mining  the Block containing \"%s\"\n" ,pow.block.Data)
	
	for nounce < maxNounce {
		data := pow.prepareData(nounce)
		
		hash =  sha256.Sum256(data)
		fmt.Printf("\r%x",hash)
		hashInt.SetBytes(hash[:])
		
		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nounce++
		}
		
	}
	fmt.Print("\n\n")
	return nounce, hash[:]
}

func (pow *ProofOfWork ) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid := hashInt.Cmp(pow.target) == -1
	return isValid
}