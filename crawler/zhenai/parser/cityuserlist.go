package parser

import (
	"imooc.com/joizhang/learn-golang/crawler/types"
	"regexp"
)

var cityUserListRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

// 获取城市里面的用户列表
func ParseCityUserList(contents []byte) types.ParseResult {
	matches := cityUserListRe.FindAllSubmatch(contents, -1)
	result := types.ParseResult{}
	for _, m := range matches {
		name := string(m[2])
		result.Requests = append(
			result.Requests,
			types.Request{
				Url: string(m[1]),
				ParseFunc: func(bytes []byte) types.ParseResult {
					return ParseProfile(bytes, name)
				},
			})
	}

	// 获取用户列表页面的城市
	matches = CityListRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		result.Requests = append(
			result.Requests,
			types.Request{
				Url:       string(m[1]),
				ParseFunc: ParseCityList,
			})
	}

	return result
}
