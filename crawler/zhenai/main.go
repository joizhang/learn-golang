package main

import (
	"imooc.com/joizhang/learn-golang/crawler/engine"
	"imooc.com/joizhang/learn-golang/crawler/scheduler"
	"imooc.com/joizhang/learn-golang/crawler/types"
	"imooc.com/joizhang/learn-golang/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	e.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
