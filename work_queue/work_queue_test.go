package work_queue

import (
	"sync/atomic"
	"testing"
	"time"
)

type testWorker struct {
	nRuns *uint64 // shared counter, to count the number of jobs that actually run
}

const delay = 25 * time.Millisecond
const correctResult = 123456

func (w testWorker) Run() interface{} {
	time.Sleep(delay)            // simulate some work without really doing anything
	atomic.AddUint64(w.nRuns, 1) // count that we ran, for testing
	return correctResult
}

func newWorker(counter *uint64) testWorker {
	w := new(testWorker)
	w.nRuns = counter
	return *w
}

// Test that the work queue can do jobs and get correct results back.
func TestQueueBasics(t *testing.T) {
	nTasks := uint(20)
	nThreads := uint(2)
	nRun := uint64(0)

	q := Create(nThreads, nTasks)
	for i := uint(0); i < nTasks; i++ {
		q.Enqueue(newWorker(&nRun))
	}

	// If <nTasks results are returned, this will deadlock.
	for i := uint(0); i < nTasks; i++ {
		r := <-q.Results
		if r != correctResult {
			t.Error("Got incorrect result from test function:", r)
		}
	}

	time.Sleep(3 * delay) // Give leftover workers time to complete, but there shouldn't be any.

	if len(q.Results) != 0 {
		t.Error("More results that expected from tasks.")
	}
	if uint(nRun) != nTasks {
		t.Errorf("Unexpected number of tasks completed: expected %d; found %d.", nTasks, nRun)
	}
}

// Test that the work queue stops processing jobs when asked to do so.
func TestQueueStop(t *testing.T) {
	nTasks := uint(20)
	nThreads := uint(4)
	nRun := uint64(0)

	q := Create(nThreads, nTasks)
	for i := uint(0); i < nTasks; i++ {
		q.Enqueue(newWorker(&nRun))
	}

	nResults := uint(0)
	for r := range q.Results {
		if r != correctResult {
			t.Error("Got incorrect result from test function:", r)
		}
		nResults += 1
		if nResults > nThreads { // Pretend nThreads tasks give result that don't mean "complete". (*)
			q.Shutdown() // After that, tell the queue we're done and to stop processing tasks;
			break        // and we're done.
		}
	}

	time.Sleep(2 * time.Duration(nTasks/nThreads) * delay) // give workers long enough to do whatever they're going to do

	t.Log("nRun =", nRun)
	// We expect:
	// nThreads tasks for "incomplete" work (* above);
	// nThreads running when we send the shutdown signal;
	// up to nThreads executing while shutting down.
	if uint(nRun) > nThreads*3 {
		t.Error("too many tasks executed:", nRun)
	}
	if uint(nRun) < nThreads*2 {
		t.Error("not enough tasks executed:", nRun)
	}
}
