package scheduler

import (
	"imooc.com/joizhang/learn-golang/crawler/types"
)

// Request队列和Worker队列
type QueuedScheduler struct {
	requestChan chan types.Request
	workChan    chan chan types.Request
}

func (s *QueuedScheduler) Submit(r types.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerChan() chan types.Request {
	return make(chan types.Request)
}

func (s *QueuedScheduler) WorkerReady(w chan types.Request) {
	s.workChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workChan = make(chan chan types.Request)
	s.requestChan = make(chan types.Request)
	go func() {
		var requestQ []types.Request
		var workQ []chan types.Request
		for {
			var activeRequest types.Request
			var activeWorker chan types.Request
			if len(requestQ) > 0 && len(workQ) > 0 {
				activeWorker = workQ[0]
				activeRequest = requestQ[0]
			}

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r)
			case w := <-s.workChan:
				workQ = append(workQ, w)
			case activeWorker <- activeRequest:
				workQ = workQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
