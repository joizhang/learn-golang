package main

import (
	"imooc.com/joizhang/learn-golang/crawler/engine"
	"imooc.com/joizhang/learn-golang/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
