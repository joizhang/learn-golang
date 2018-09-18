package engine

import (
	"imooc.com/joizhang/learn-golang/crawler/scheduler"
	"imooc.com/joizhang/learn-golang/crawler/types"
	"imooc.com/joizhang/learn-golang/crawler/zhenai/parser"
	"testing"
)

func TestConcurrentEngine_simpleScheduler(t *testing.T) {
	e := ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}

func TestConcurrentEngine_queuedScheduler(t *testing.T) {
	e := ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
