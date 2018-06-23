package engine

import (
	"testing"
	"imooc.com/joizhang/learn-golang/crawler/zhenai/parser"
	"imooc.com/joizhang/learn-golang/crawler/types"
)

func TestSimpleEngine_Run(t *testing.T) {

	SimpleEngine{}.Run(types.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}
