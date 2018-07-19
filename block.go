package blockchain

/*
 * Block and Block Header
 */

import (
	"crypto/sha256"

	. "github.com/yockliu/bitcoinlib"
)

// Block the cell of the block chain
type Block struct {
	Cell
	Version    uint32
	MHash      *HashCode
	PrevHash   HashCode
	MerkleRoot HashCode // MerkleRoot of Contents
	Timestamp  uint32
	Bits       uint32 // difficulty
	Nonce      uint32
	Contents   []Cell
	NextHash   HashCode
}

// Serialize Serializable
func (block *Block) Serialize() []byte {
	header := block.serializeHead()
	body := block.serializedBody()
	size := uint32(len(header) + len(body))

	blockBytes := ConcatAppend([][]byte{Uint32ToBytes(size), header, body})

	return blockBytes
}

func (block *Block) serializeHead() []byte {
	items := [][]byte{
		Uint32ToBytes(block.Version),   // version
		block.PrevHash[:],              // preHash
		block.MerkleRoot[:],            // merkleRoot
		Uint32ToBytes(block.Timestamp), // timestamp
		Uint32ToBytes(block.Bits),      // bits difficulty
		Uint32ToBytes(block.Nonce),     // nonce
	}

	data := ConcatAppend(items)

	return data
}

func (block *Block) serializedBody() []byte {
	contentBytes := []byte{}

	contentN := CompactSizeUint{Value: uint64(len(block.Contents))}
	contentBytes = append(contentBytes, contentN.Bytes()...)

	for _, content := range block.Contents {
		contentBytes = append(contentBytes, content.Serialize()...)
	}

	return contentBytes
}

// Deserialize Serializable
func (block *Block) Deserialize([]byte) {
	// TODO
}

// Hash Hashable
func (block *Block) Hash() HashCode {
	if block.MHash != nil {
		return *block.MHash
	}
	bytes := block.Serialize()
	hash := HashCode(sha256.Sum256(bytes))
	block.MHash = &hash
	return *block.MHash
}
