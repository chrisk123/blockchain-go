package blockchain

import "bytes"

//Blockchain struct: Chain is an array of Blocks
type Blockchain struct {
	Chain []Block
}

//Add function: Appends block to end of existing blockchain
func (chain *Blockchain) Add(blk Block) {
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}
	// TODO
	chain.Chain = append(chain.Chain, blk)
}

//IsValid function: Verification checks of blocks in the blockchain
func (chain Blockchain) IsValid() bool {
	validityBytes := make([]byte, 32)
	prevHash := chain.Chain[0].PrevHash
	generation := chain.Chain[0].Generation
	difficulty := chain.Chain[0].Difficulty
	hash := chain.Chain[0].Hash

	//Check hash of first block
	if !bytes.Equal(hash, chain.Chain[0].CalcHash()) {
		return false
	}

	//Check validity of first block
	if !chain.Chain[0].ValidHash() {
		return false
	}

	//Check, if not the first block (generation not 0), then prevHash should be all null = validityBytes
	if !bytes.Equal(prevHash, validityBytes) || generation != 0 {
		return false
	}

	//For subsequent blocks in the chain...
	for i := 1; i < len(chain.Chain); i++ {
		//Check consistent difficulty
		if chain.Chain[i].Difficulty != difficulty {
			return false
		}
		//Check generation incrementing
		if chain.Chain[i].Generation != generation+uint64(i) {
			return false
		}
		//Check previous hash matches
		if !bytes.Equal(chain.Chain[i].PrevHash, hash) {
			return false
		}
		hash = chain.Chain[i].Hash
		//Check hash calculated correctly
		if !bytes.Equal(hash, chain.Chain[i].CalcHash()) {
			return false
		}
		//Check hash validity
		if !chain.Chain[i].ValidHash() {
			return false
		}
	}
	return true
}
