package blockchain

/*
 * blockchain
 */

import (
	"fmt"
	"time"

	. "github.com/yockliu/bitcoinlib"
)

// Version version constant
const Version = uint32(1)

// BlockChain block chain element that help to generate block
type BlockChain struct {
	blocks  []*Block
	indexes map[string]*Block
}

// NewBlockChain new BlockChain
func NewBlockChain() *BlockChain {
	blockChain := BlockChain{}
	blockChain.blocks = []*Block{}
	blockChain.indexes = map[string]*Block{}
	return &blockChain
}

// Current get the current block of the chain
func (blockChain *BlockChain) Current() *Block {
	len := len(blockChain.blocks)
	if len == 0 {
		return nil
	}
	return blockChain.blocks[len-1]
}

// BlockOfHash get Block by Hash
func (blockChain *BlockChain) BlockOfHash(hash *HashCode) *Block {
	hashStr := fmt.Sprintf("%x", hash)
	return blockChain.indexes[hashStr]
}

// BlockOfHeight get Block by Height
func (blockChain *BlockChain) BlockOfHeight(index int) *Block {
	if index < 0 || index >= len(blockChain.blocks) {
		return nil
	}
	return blockChain.blocks[index]
}

// GenerateBlock generator new block by tx and difficulty
func (blockChain *BlockChain) GenerateBlock(contents []Cell, bits uint32) *Block {
	// check the current block
	prevBlock := blockChain.Current()
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

	currentHash := block.Hash()

	if prevBlock != nil {
		prevBlock.NextHash = currentHash
	}

	blockChain.blocks = append(blockChain.blocks, &block)
	hashStr := fmt.Sprintf("%x", currentHash)
	blockChain.indexes[hashStr] = &block

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
