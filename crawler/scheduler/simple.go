package scheduler

import (
	"imooc.com/joizhang/learn-golang/crawler/types"
)

// 所有Worker共用一个输入
type SimpleScheduler struct {
	workerChan chan types.Request
}

func (s *SimpleScheduler) Submit(r types.Request) {
	// send types down to worker chan
	go func() { s.workerChan <- r }()
}

func (s *SimpleScheduler) WorkerChan() chan types.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan types.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan types.Request)
}
