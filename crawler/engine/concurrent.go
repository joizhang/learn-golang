package engine

import (
	"log"
	"imooc.com/joizhang/learn-golang/crawler/types"
	"imooc.com/joizhang/learn-golang/crawler/scheduler"
)

type ConcurrentEngine struct {
	Scheduler   scheduler.Scheduler
	WorkerCount int
}

// 并发版
func (e *ConcurrentEngine) Run(seeds ...types.Request) {
	out := make(chan types.ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	for {
		result := <-out
		for _, item := range result.Items {
			log.Printf("Got item #%d: %v", itemCount, item)
			itemCount++
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan types.Request,
	out chan types.ParseResult,
	ready scheduler.ReadyNotifier) {
	go func() {
		for {
			// tell scheduler i am ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
