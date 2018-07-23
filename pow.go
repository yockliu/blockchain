package blockchain

/*
 * Prof of Work
 */

import (
	"fmt"
	"runtime"

	. "github.com/yockliu/bitcoinlib"
)

const maxNonce = ^uint32(0)

// simplified difficulty target, difficult range [0, 256)
func target(bits uint32) HashCode {
	target := HashCode{} // [32]byte
	bits = bits % 256
	bytePos := bits / 8
	bitPos := bits % 8
	target[bytePos] = 255 >> bitPos
	return target
}

// ProfOfWork mine a block
func ProfOfWork(block *Block) error {
	target := target(block.Bits)
	// fmt.Printf("target = %x\n", target)

	for nonce := uint32(0); nonce < maxNonce; nonce++ {
		block.Nonce = nonce
		hash := block.Hash()
		// fmt.Printf("hash = %x\n", hash)
		if hash.Compare(&target) < 0 {
			// fmt.Printf("hash = %x\n", hash)
			return nil
		}

		runtime.Gosched()
	}

	return fmt.Errorf("ProfOfWork Not found result")
}
