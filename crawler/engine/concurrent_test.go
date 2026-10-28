package engine

import (
	"sync"
	"testing"

	"imooc.com/joizhang/learn-golang/crawler/scheduler"
	"imooc.com/joizhang/learn-golang/crawler/types"
	"imooc.com/joizhang/learn-golang/crawler/zhenai/parser"
)

func TestConcurrentEngine_simpleScheduler(t *testing.T) {
	runConcurrentEngine(t, &scheduler.SimpleScheduler{})
}

func TestConcurrentEngine_queuedScheduler(t *testing.T) {
	runConcurrentEngine(t, &scheduler.QueuedScheduler{})
}

// runConcurrentEngine 用本地快照替代真实请求，并且抓够 MaxItems 就返回。
// 这样测试是确定性的，不会因为珍爱网改版或被墙而一直卡住。
func runConcurrentEngine(t *testing.T, sch scheduler.Scheduler) {
	t.Helper()
	body := loadCityList(t)

	var mu sync.Mutex
	fetches := 0
	e := ConcurrentEngine{
		Scheduler:   sch,
		WorkerCount: 10,
		MaxItems:    1,
		Fetch: func(string) ([]byte, error) {
			mu.Lock()
			fetches++
			mu.Unlock()
			return body, nil
		},
	}

	e.Run(types.Request{Url: seedURL, ParseFunc: parser.ParseCityList})

	mu.Lock()
	defer mu.Unlock()
	if fetches != 1 {
		t.Errorf("MaxItems=1 时应当只抓 1 次，实际抓了 %d 次", fetches)
	}
}
