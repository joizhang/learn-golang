package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityUserList(t *testing.T) {
	content, err := ioutil.ReadFile("cityuserlist_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityUserList(content)
	const resultSize = 20

	expectedUrls := []string{
		"http://album.zhenai.com/u/1995815593",
		"http://album.zhenai.com/u/1314495053",
		"http://album.zhenai.com/u/1626200317",
	}
	expectedUsers := []string{
		"User 小顺儿", "User 风中的蒲公英", "User 路漫漫",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("result should hava %d requests; but had %d",
			resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s",
				i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should hava %d items; but had %d",
			resultSize, len(result.Items))
	}
	for i, user := range expectedUsers {
		if result.Items[i].(string) != user {
			t.Errorf("expected city #%d: %s; but was %s",
				i, user, result.Items[i].(string))
		}
	}

}
