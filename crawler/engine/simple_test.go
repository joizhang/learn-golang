package engine

import (
	"imooc.com/joizhang/learn-golang/crawler/types"
	"imooc.com/joizhang/learn-golang/crawler/zhenai/parser"
	"testing"
)

func TestSimpleEngine_Run(t *testing.T) {
	SimpleEngine{}.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
