package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
)

// DONT TOUCH
type Block struct {
	Generation uint64 //index
	Difficulty uint8  //number of null bytes at end of hash
	Data       string
	Hash       []byte
	PrevHash   []byte
	Proof      uint64 //value to hash and validate for last d 0s
}

// Initial function: Create new initial (generation 0) block.
func Initial(difficulty uint8) Block {
	var initialBlock Block
	initialBlock.Generation = 0
	initialBlock.Difficulty = difficulty
	initialBlock.Data = ""
	// initialBlock.Hash = initialBlock.CalcHash()
	initialBlock.PrevHash = []byte{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'}

	// Generation: 0,
	// Difficulty: difficulty,
	// Data:       "",
	// Proof: 242278,
	// PrevHash:   []byte{'\x00', 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
	// PrevHash: []byte{'\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00', '\x00'},
	// PrevHash: []byte{'\x00', '\x00'}}
	// Hash: initialBlock.CalcHash()}

	return initialBlock
}

// Next function: Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	var nextBlock Block
	nextBlock.Generation = prev_block.Generation + 1
	nextBlock.Difficulty = prev_block.Difficulty
	nextBlock.Data = data
	fmt.Printf("Prev block hash = %v\n", prev_block.Hash)
	nextBlock.PrevHash = prev_block.Hash
	nextBlock.Hash = nextBlock.CalcHash()
	return nextBlock
}

// CalcHash function: Calculate the block's hash, uses SHA256
func (blk Block) CalcHash() []byte {
	var hashString string
	// fmt.Printf("Prev Hash: %v\n", blk.PrevHash)
	prevHashString := hex.EncodeToString(blk.PrevHash)
	// fmt.Printf("Prev Hash String: %v\n", prevHashString)
	hashString = prevHashString + ":" + strconv.FormatUint(blk.Generation, 10) + ":" + strconv.Itoa(int(blk.Difficulty)) + ":" + blk.Data + ":" + strconv.FormatUint(blk.Proof, 10)
	fmt.Printf("Converted to hashString = %v\n", hashString)
	// hashStringTest := "29528aaf90e167b2dc248587718caab237a81fd25619a5b18be4986f75f30000:1:2:message:75729"
	hash256 := sha256.New()
	hash256.Write([]byte(hashString))
	// hash256.Write([]byte(hashStringTest))
	hashed := hash256.Sum(nil)
	return hashed
}

// ValidHash function: Check if last d bytes of hash value are null
func (blk Block) ValidHash() bool {
	var res bool
	d := int(blk.Difficulty)
	hashed := blk.CalcHash()
	hashLen := len(hashed)

	validity := make([]byte, d) //byte array of 0s determined by difficulty value
	for i := 0; i < d; i++ {
		validity[i] = 0
	}

	//[(hashLen-d):] selects only last d bytes of array to compare for 0s
	if bytes.Equal(hashed[hashLen-d:], validity) {
		fmt.Printf("bytes.Equal? TRUE | hashed = %v, validity = %v\n", hashed[hashLen-d:], validity)
		res = true
	} else {
		fmt.Printf("bytes.Equal? FALSE | hashed = %v, validity = %v\n", hashed[hashLen-d:], validity)
		res = false
	}
	return res
}

// DONT TOUCH
// SetProof function: Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
