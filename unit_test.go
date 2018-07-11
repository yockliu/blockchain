package blockchain

import (
	"crypto/sha256"
	"testing"

	. "github.com/yockliu/bitcoinlib"
)

func TestBlockChainCreation(t *testing.T) {
	blockChain := BlockChain{}
	if blockChain.Current == nil {
		blockChain.Creation()
	}

	if blockChain.Current != nil {
		t.Log(blockChain.Current)
	} else {
		t.Error("creation failed")
	}
}

func TestBlockChainGenerate(t *testing.T) {
	nodeArray := [8]HashCode{
		sha256.Sum256([]byte("0000000000000000000000000000000000000000000000000000000000000000")),
		sha256.Sum256([]byte("0000000000000000000000000000000000000000000000000000000000000111")),
		sha256.Sum256([]byte("0000000000000000000000000000000000000000000000000000000000000222")),
		sha256.Sum256([]byte("0000000000000000000000000000000000000000000000000000000000000333")),
		sha256.Sum256([]byte("0000000000000000000000000000000000000000000000000000000000000444")),
		sha256.Sum256([]byte("0000000000000000000000000000000000000000000000000000000000000555")),
		sha256.Sum256([]byte("000000000000000000000000000000000000000000000000000000000000066666")),
		sha256.Sum256([]byte("000000000000000000000000000000000000000000000000000000000000077")),
	}

	blockChain := BlockChain{}
	if blockChain.Current == nil {
		blockChain.Creation()
	}

	block, err := blockChain.GenerateBlock(nodeArray[:], 20)
	if err != nil {
		t.Error(err)
	} else {
		if blockChain.Current.Header.Version != 0 {
			blockChain.Current.NextHash = block.Hash
			blockChain.Current = block
		}
		t.Log(blockChain.Current)
	}
}
