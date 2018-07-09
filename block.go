package blockchain

/*
 * Block and Block Header
 */

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
)

// BlockHeader header of block
type BlockHeader struct {
	Version           uint32   // 4 byte
	Previousblockhash HashCode // 32 byte
	Merkleroot        HashCode // 32 byte
	Timestamp         uint32   // 4 byte
	Difficulty        float32  // 4 byte
	Nonce             uint32   // 4 byte
}

func (header *BlockHeader) hash() (HashCode, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, header)

	if err != nil {
		return HashCode{}, err
	}

	return sha256.Sum256(buf.Bytes()), nil
}

// Block block of chain
type Block struct {
	Size     uint32
	Header   BlockHeader
	Bits     uint64
	Tx       []HashCode
	Hash     HashCode
	NextHash HashCode
	Index    uint64
}

const fixedSize = 4 + 80 + 32 + 32 + 8

func (block *Block) calcuSize() {
	block.Size = uint32(fixedSize) + 32*uint32(len(block.Tx))
}
