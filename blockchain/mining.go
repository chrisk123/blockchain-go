package blockchain

type miningWorker struct {
	// TODO. Should implement work_queue.Worker
}

// DONT TOUCH
type MiningResult struct {
	Proof uint64 // proof-of-work value, if found.
	Found bool   // true if valid proof-of-work was found.
}

// // Mine the range of proof values, by breaking up into chunks and checking
// // "workers" chunks concurrently in a work queue. Should return shortly after a result
// // is found.
// // MineRange function
// func (blk Block) MineRange(start uint64, end uint64, workers uint64, chunks uint64) MiningResult {
// 	// TODO: spawn concurrent goroutines and send results back on a channel
// }

// // DONT TOUCH
// // Call .MineRange with some reasonable values that will probably find a result.
// // Good enough for testing at least. Updates the block's .Proof and .Hash if successful.
// // Mine function
// func (blk *Block) Mine(workers uint64) bool {
// 	reasonableRangeEnd := uint64(4 * 1 << (8 * blk.Difficulty)) // 4 * 2^(bits that must be zero)
// 	mr := blk.MineRange(0, reasonableRangeEnd, workers, 4321)
// 	if mr.Found {
// 		blk.SetProof(mr.Proof)
// 	}
// 	return mr.Found
// }
