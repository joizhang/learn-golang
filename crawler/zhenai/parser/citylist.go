package parser

import (
	"imooc.com/joizhang/learn-golang/crawler/types"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) types.ParseResult {
	re, _ := regexp.Compile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := types.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "City "+string(m[2]))
		result.Requests = append(
			result.Requests,
			types.Request{Url: string(m[1]), ParseFunc: ParseCity})
	}
	return result
}
