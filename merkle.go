package blockchain

/*
 * Merkle Tree
 */

import (
	"crypto/sha256"
	"fmt"
)

// Merkle : get merkle tree root
func Merkle(nodeList []HashCode) (HashCode, error) {

	// check parameters
	if len(nodeList) == 0 {
		return HashCode{}, fmt.Errorf("input is empty")
	}

	// if len is 1, then get the root, else continue calculate
	for len(nodeList) > 1 {
		nodeLength := len(nodeList)

		// patch element's count to be 2x
		if nodeLength%2 != 0 {
			lastNode := nodeList[nodeLength-1]
			nodeList = append(nodeList, lastNode)
		}

		// log
		// fmt.Println("----------")
		// for _, element := range nodeList {
		// 	fmt.Printf("%x", element)
		// 	fmt.Println()
		// }

		nodeLength = len(nodeList)
		newNodeList := make([]HashCode, 0, nodeLength/2)

		// calculate high level nodes
		for i := 0; i < nodeLength/2; i++ {
			hash1 := nodeList[2*i][:]
			hash2 := nodeList[2*i+1][:]
			data := append(hash1, hash2...)
			hash := sha256.Sum256(data)
			newNodeList = append(newNodeList, hash)
		}

		// set to higher level
		nodeList = newNodeList
	}

	fmt.Printf("merkleroot = %x\n", nodeList[0])

	return nodeList[0], nil
}
