package engine

import (
	"log"

	"imooc.com/joizhang/learn-golang/crawler/scheduler"
	"imooc.com/joizhang/learn-golang/crawler/types"
)

type ConcurrentEngine struct {
	Scheduler   scheduler.Scheduler
	WorkerCount int
	// Fetch 可选，注入自定义抓取函数；为 nil 时使用 fetcher.Fetch
	Fetch FetchFunc
	// MaxItems 可选，抓够这么多 item 后停止；0 表示不限制（会一直爬下去）
	MaxItems int

	// visited 只在 Run 里读写，所以不用加锁。
	// 放进结构体而不是包级变量，避免多个引擎/多次测试互相污染。
	visited map[string]bool
}

// 并发版
func (e *ConcurrentEngine) Run(seeds ...types.Request) {
	out := make(chan types.ParseResult)

	e.visited = make(map[string]bool)
	e.Scheduler.Run()

	fetch := resolveFetch(e.Fetch)

	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler, fetch)
	}

	for _, r := range seeds {
		if e.isDuplicate(r.Url) {
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
			if e.MaxItems > 0 && itemCount >= e.MaxItems {
				return
			}
		}

		for _, request := range result.Requests {
			if e.isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

// URL deduplicate
func (e *ConcurrentEngine) isDuplicate(url string) bool {
	if e.visited[url] {
		return true
	}
	e.visited[url] = true
	return false
}

func (e *ConcurrentEngine) createWorker(in chan types.Request, out chan types.ParseResult, ready scheduler.ReadyNotifier, fetch FetchFunc) {
	go func() {
		for {
			// Tell scheduler I am ready
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(fetch, request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
