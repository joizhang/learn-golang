package parser

import (
	"io/ioutil"
	"testing"
)

// 快照里有 20 个用户链接；页脚还有一个 anchor 文本为 "1" 的分页链接指向
// 城市列表页，ParseCityUserList 也会把它输出成请求，所以总请求数是 21。
const (
	userCount       = 20
	paginationCount = 1
)

func TestParseCityUserList(t *testing.T) {
	content, err := ioutil.ReadFile("cityuserlist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityUserList(content)

	expectedUrls := []string{
		"http://album.zhenai.com/u/1995815593",
		"http://album.zhenai.com/u/1314495053",
		"http://album.zhenai.com/u/1626200317",
	}
	expectedUsers := []string{
		"User 小顺儿", "User 风中的蒲公英", "User 路漫漫",
	}

	if len(result.Requests) != userCount+paginationCount {
		t.Errorf("result should hava %d requests; but had %d",
			userCount+paginationCount, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s",
				i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != userCount {
		t.Errorf("result should hava %d items; but had %d",
			userCount, len(result.Items))
	}
	for i, user := range expectedUsers {
		if result.Items[i].(string) != user {
			t.Errorf("expected user #%d: %s; but was %s",
				i, user, result.Items[i].(string))
		}
	}

}
