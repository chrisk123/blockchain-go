package blockchain

import "bytes"

// DONT TOUCH
type Blockchain struct {
	Chain []Block
}

func (chain *Blockchain) Add(blk Block) {
	if !blk.ValidHash() {
		panic("adding block with invalid hash")
	}
	// TODO
	chain.Chain = append(chain.Chain, blk)
}

func (chain Blockchain) IsValid() bool {
	// TODO
	validityBytes := make([]byte, 32)
	prevHash := chain.Chain[0].PrevHash
	generation := chain.Chain[0].Generation
	difficulty := chain.Chain[0].Difficulty
	hash := chain.Chain[0].Hash

	if !bytes.Equal(hash, chain.Chain[0].CalcHash()) {
		return false
	}

	if !chain.Chain[0].ValidHash() {
		return false
	}

	if !bytes.Equal(prevHash, validityBytes) || generation != 0 {
		return false
	}

	for i := 1; i < len(chain.Chain); i++ {
		if chain.Chain[i].Difficulty != difficulty {
			return false
		}

		if chain.Chain[i].Generation != generation+uint64(i) {
			return false
		}

		if !bytes.Equal(chain.Chain[i].PrevHash, hash) {
			return false
		}
		hash = chain.Chain[i].Hash

		if !bytes.Equal(hash, chain.Chain[i].CalcHash()) {
			return false
		}

		if !chain.Chain[i].ValidHash() {
			return false
		}
	}
	return true
}
