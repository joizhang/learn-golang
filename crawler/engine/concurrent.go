package engine

import (
	"imooc.com/joizhang/learn-golang/crawler/scheduler"
	"imooc.com/joizhang/learn-golang/crawler/types"
	"log"
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
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			continue
		}
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
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var visitedUrl = make(map[string]bool)

// URL deduplicate
func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}

func (ConcurrentEngine) createWorker(in chan types.Request, out chan types.ParseResult, ready scheduler.ReadyNotifier) {
	go func() {
		for {
			// Tell scheduler I am ready
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
