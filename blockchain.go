package blockchain

/*
 * blockchain
 */

import (
	"time"

	. "github.com/yockliu/bitcoinlib"
)

// Version version constant
const Version = uint32(1)

// BlockChain block chain element that help to generate block
type BlockChain struct {
	Current *Block
}

// GenerateBlock generator new block by tx and difficulty
func (blockChain *BlockChain) GenerateBlock(contents []Cell, bits uint32) *Block {
	// check the current block

	prevBlock := blockChain.Current
	prevHash := &HashCode{}
	if prevBlock != nil {
		prevHash = prevBlock.Hash()
	}

	block := Block{}
	block.Version = 1
	block.PrevHash = prevHash
	block.Timestamp = uint32(time.Now().Unix())
	block.Bits = bits
	block.MerkleRoot = merkle(contents)
	block.Contents = contents

	// prof of work
	ProfOfWork(&block)

	if prevBlock != nil {
		prevBlock.NextHash = block.Hash()
	}

	return &block
}

func merkle(contents []Cell) *HashCode {
	nodeList := []HashCode{}
	for _, content := range contents {
		nodeList = append(nodeList, *content.Hash())
	}
	merkleRoot := Merkle(nodeList)
	return merkleRoot
}
