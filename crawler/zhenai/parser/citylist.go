package parser

import (
	"imooc.com/joizhang/learn-golang/crawler/types"
	"regexp"
)

var CityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

// 获取城市列表
func ParseCityList(contents []byte) types.ParseResult {
	matches := CityListRe.FindAllSubmatch(contents, -1)
	result := types.ParseResult{}
	for _, m := range matches {
		result.Requests = append(
			result.Requests,
			types.Request{Url: string(m[1]), ParseFunc: ParseCityUserList})
	}
	return result
}
