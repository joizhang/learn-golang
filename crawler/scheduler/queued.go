package scheduler

import (
	"imooc.com/joizhang/learn-golang/crawler/types"
)

// Request队列和Worker队列

type QueuedScheduler struct {
	requestCHan chan types.Request
	workCHan    chan chan types.Request
}

func (s *QueuedScheduler) Submit(r types.Request) {
	s.requestCHan <- r
}

func (s *QueuedScheduler) WorkerChan() chan types.Request {
	return make(chan types.Request)
}

func (s *QueuedScheduler) WorkerReady(w chan types.Request) {
	s.workCHan <- w
}

func (s *QueuedScheduler) Run() {
	s.workCHan = make(chan chan types.Request)
	s.requestCHan = make(chan types.Request)
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
			case r := <-s.requestCHan:
				requestQ = append(requestQ, r)
			case w := <-s.workCHan:
				workQ = append(workQ, w)
			case activeWorker <- activeRequest:
				workQ = workQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
