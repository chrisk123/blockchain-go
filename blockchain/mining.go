package blockchain

import "fmt"

type miningWorker struct {
	// TODO. Should implement work_queue.Worker
}

// DONT TOUCH
type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

type Chunk struct {
	Size  uint64
	Start uint64
	End   uint64
}

// Mine the range of proof values, by breaking up into chunks and checking
// "workers" chunks concurrently in a work queue. Should return shortly after a result
// is found.
// MineRange function: Calculates range of proof values until a valid hash is found
func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
	// TODO: spawn concurrent goroutines and send results back on a channel
	var miningResult MiningResult
	miningBlk := blk

	var mineChannel = make(chan uint64, workers)
	//Split mining range [start,end] into chunks and calculate concurrently
	chunked := CalcChunks(start,end,chunks)


	//While Jobs in Queue exist, pass chunk down channel for workers to hash 
		go func() {
			//Pass chunked[i] range of values to the job queue
			chunked[:]
		}
	
	
	// for i := 0; i < int(chunks-1); i++ {
	// 	next_chunk :=
	// }
	// jobChunks :=

	//Slowest situation, 1 worker mining entire range
	if workers == 1 && chunks == 1 {
		for i := start; i < end; i++ {
			//Assign block the next proof value and check validity of hash
			miningBlk.Proof = i
			miningBlk.CalcHash()
			//Check if proof results in a valid hash and return result if so
			if miningBlk.ValidHash() {
				miningResult.Proof = i
				mineChannel <- miningResult.Proof
				miningResult.Found = true
				return miningResult
			}
		}

	}
	// blk.SetProof(res)
	return miningResult
}

// DONT TOUCH
// Call .MineRange with some reasonable values that will probably find a result.
// Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
// Mine function
func (blk *Block) Mine(workers uint64) bool {
	reasonableRangeEnd := uint64(4 * 1 << (8 * blk.Difficulty)) // 4 * 2^(bits that must be zero)
	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
	if mr.Found {
		blk.SetProof(mr.Proof)
	}
	return mr.Found
}

//CalcChunks function: Helper method to divide range into approximately equal-sized chunks
//Returns an array of Chunk structs (.Size, .Start, .End)
func CalcChunks(start, end, chunks uint64) []Chunk {
	proofRange := end - start
	avgChunkSize := proofRange / chunks
	fmt.Printf("AvgChunkSize %v\n", avgChunkSize)
	remainder := proofRange % chunks
	fmt.Printf("Remainder %v\n", remainder)
	chunk := make([]Chunk, chunks)

	for i := uint64(0); i < chunks; i++ {
		if i == 0 {
			chunk[i].Size = avgChunkSize
			chunk[i].Start = i * avgChunkSize
			chunk[i].End = i*avgChunkSize + avgChunkSize
			fmt.Printf("Chunk[%v] = %v\n", i, chunk[i])
			//Distribute remainders 1 at a time
			if remainder > 0 {
				chunk[i].End++
				chunk[i].Size++
				remainder--
			}
		} else {
			chunk[i].Size = avgChunkSize
			chunk[i].Start = chunk[i-1].End + 1
			chunk[i].End = i*avgChunkSize + avgChunkSize
			fmt.Printf("Chunk[%v] = %v\n", i, chunk[i])
			//Distribute remainders 1 at a time
			if remainder > 0 {
				chunk[i].End++
				chunk[i].Size++
				remainder--
			}
		}
	}
	return chunk
}
