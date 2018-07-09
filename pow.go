package blockchain

/*
 * Prof of Work
 */

import (
	"fmt"
	"math"
)

const maxNonce = ^uint32(0)

// simplified difficulty target, difficult range [0, 256)
func target(difficulty float32) HashCode {
	target := HashCode{} // [32]byte
	bits := uint64(math.Floor(float64(difficulty)))
	bits = bits % 256
	bytePos := bits / 8
	bitPos := bits % 8
	target[bytePos] = 255 >> bitPos
	return target
}

// ProfOfWork mine a block
func ProfOfWork(header *BlockHeader, difficulty float32) error {
	header.Difficulty = difficulty

	target := target(difficulty)
	// fmt.Printf("target = %x\n", target)

	for nonce := uint32(0); nonce < maxNonce; nonce++ {
		header.Nonce = nonce
		hash, err := header.hash()
		if err != nil {
			continue
		} else {
			if hash.compare(&target) < 0 {
				// fmt.Printf("hash = %x\n", hash)
				return nil
			}
		}
	}

	return fmt.Errorf("not found result")
}
