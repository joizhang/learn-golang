package parser

import (
	"imooc.com/joizhang/learn-golang/crawler/zhenai/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseProfile(contents, "惠儿")
	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element; but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)
	expected := model.Profile{
		Name:       "惠儿",
		Age:        50,
		Height:     156,
		Weight:     0,
		Income:     "3000元以下",
		Gender:     "女",
		Xinzuo:     "魔羯座",
		Marriage:   "离异",
		Education:  "高中及以下",
		Occupation: "销售总监",
		Hokou:      "四川阿坝",
		House:      "租房",
		Car:        "未购车",
	}
	if profile != expected {
		t.Errorf("expected %v, but was %v", expected, profile)
	}
}
