package scheduler

import (
	"imooc.com/joizhang/learn-golang/crawler/types"
)

// 所有Worker公用一个输入
type SimpleScheduler struct {
	workerChan chan types.Request
}

func (s *SimpleScheduler) Submit(r types.Request) {
	// send types down to worker chan
	s.workerChan <- r
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan types.Request) {
	s.workerChan = c
}
