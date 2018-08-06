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
	initialBlock.PrevHash = make([]byte, 32)
	return initialBlock
}

// Next function: Create new block to follow this block, with provided data.
func (prev_block Block) Next(data string) Block {
	var nextBlock Block
	nextBlock.Generation = prev_block.Generation + 1
	nextBlock.Difficulty = prev_block.Difficulty
	nextBlock.Data = data
	fmt.Printf("\nPrev block hash = %v\n", prev_block.Hash) //unable to read previous block's hash?
	nextBlock.PrevHash = prev_block.Hash
	// nextBlock.Hash = prev_block.CalcHash()
	return nextBlock
}

// CalcHash function: Calculate the block's hash, uses SHA256
func (blk Block) CalcHash() []byte {
	var hashString string
	prevHashString := hex.EncodeToString(blk.PrevHash)
	hashString = prevHashString + ":" + strconv.FormatUint(blk.Generation, 10) + ":" + strconv.Itoa(int(blk.Difficulty)) + ":" + blk.Data + ":" + strconv.FormatUint(blk.Proof, 10)
	fmt.Printf("Converted to hashString = %v\n", hashString)
	hash256 := sha256.New()
	hash256.Write([]byte(hashString))
	hashed := hash256.Sum(nil)
	return hashed
}

// ValidHash function: Check if last d bytes of hash value are null
func (blk Block) ValidHash() bool {
	var res bool
	d := int(blk.Difficulty)
	hashed := blk.CalcHash()
	hashLen := len(hashed)
	validityBytes := make([]byte, d) //byte array of 0s determined by difficulty value

	//Initialize validity array containing all null
	for i := 0; i < d; i++ {
		validityBytes[i] = '\x00'
	}

	//Select only the last d bytes of hashed and check for null
	if bytes.Equal(hashed[hashLen-d:], validityBytes) {
		fmt.Printf("Validity Check: hashed = %v, validity = %v\n", hashed[hashLen-d:], validityBytes)
		res = true
	} else {
		fmt.Printf("Validity Check:  hashed = %v, validity = %v\n", hashed[hashLen-d:], validityBytes)
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
