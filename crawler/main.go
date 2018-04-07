package main

import (
	"imooc.com/joizhang/learn-golang/crawler/engine"
	"imooc.com/joizhang/learn-golang/crawler/zhenai/parser"
	"imooc.com/joizhang/learn-golang/crawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
