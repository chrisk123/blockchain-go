package work_queue

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs         chan Worker
	Results      chan interface{}
	StopRequests chan int
	NumWorkers   uint
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)
	// TODO
	return q
}

// A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {
	running := true
	// Run tasks from the queue, unless we have been asked to stop.
	for running {
		// TODO: run tasks from Jobs
		// TODO: stop when a signal is sent on StopRequests
	}
}

func (queue WorkQueue) Enqueue(work Worker) {
	// TODO
}

func (queue WorkQueue) Shutdown() {
	// TODO
}
