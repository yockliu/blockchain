package blockchain

/*
 * blockchain
 */

import (
	"fmt"
	"strings"
	"time"
)

// Version version constant
const Version = uint32(1)

// HashCode sha256 Code 32 byte
type HashCode [32]byte

func (hash *HashCode) compare(anotherHash *HashCode) int {
	hexStr := fmt.Sprintf("%x", hash)
	anotherHexStr := fmt.Sprintf("%x", anotherHash)
	return strings.Compare(hexStr, anotherHexStr)
}

// BlockChain block chain element that help to generate block
type BlockChain struct {
	Current *Block
}

// GenerateBlock generator new block by tx and difficulty
func (blockChain *BlockChain) GenerateBlock(tx []HashCode, difficulty float32) (*Block, error) {
	// check the current block
	if blockChain.Current == nil {
		return nil, fmt.Errorf("need createion")
	}
	current := blockChain.Current

	// calculate merkle root
	merkleroot, merkleErr := Merkle(tx)
	if merkleErr != nil {
		return nil, merkleErr
	}

	// setup header
	header := BlockHeader{}
	header.Version = Version
	header.Previousblockhash = current.Hash
	header.Merkleroot = merkleroot
	header.Timestamp = uint32(time.Now().Unix())

	powErr := ProfOfWork(&header, difficulty)
	if powErr != nil {
		return nil, powErr
	}

	hash, hashError := header.hash()
	if hashError != nil {
		return nil, hashError
	}

	// setup new block
	block := Block{}
	block.Header = header
	block.Bits = uint64(len(tx))
	block.Tx = tx
	block.Hash = hash
	block.calcuSize()
	block.Index = current.Index + 1

	// set current's next
	current.NextHash = block.Hash

	// set current to point to the new block
	blockChain.Current = &block

	return &block, nil
}

// Creation create the first block of the chain
// todo: add bitcoin incoming tx
func (blockChain *BlockChain) Creation() error {
	if blockChain.Current != nil {
		return fmt.Errorf("the chain is not empty, can't do the creation")
	}

	difficulty := float32(20)
	header := BlockHeader{}
	header.Version = Version
	header.Previousblockhash = [32]byte{}
	header.Merkleroot = [32]byte{}
	header.Timestamp = uint32(time.Now().Unix())
	errPow := ProfOfWork(&header, difficulty)
	if errPow != nil {
		return errPow
	}

	block := Block{}
	block.Header = header
	hash, errHeaderHash := header.hash()
	if errHeaderHash != nil {
		return errHeaderHash
	}
	block.Hash = hash
	block.Bits = uint64(0)
	block.Index = 1
	block.calcuSize()

	blockChain.Current = &block

	return nil
}
