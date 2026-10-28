package engine

import (
	"os"
	"testing"

	"imooc.com/joizhang/learn-golang/crawler/types"
	"imooc.com/joizhang/learn-golang/crawler/zhenai/parser"
)

const (
	cityListFixture = "../zhenai/parser/citylist_test_data.html"
	seedURL         = "http://www.zhenai.com/zhenghun"
)

// loadCityList 读取本地快照，测试不依赖真实网络。
func loadCityList(t *testing.T) []byte {
	t.Helper()
	b, err := os.ReadFile(cityListFixture)
	if err != nil {
		t.Fatalf("读取快照失败: %v", err)
	}
	return b
}

func TestSimpleEngine_Run(t *testing.T) {
	body := loadCityList(t)

	var fetched []string
	e := SimpleEngine{
		Fetch: func(url string) ([]byte, error) {
			fetched = append(fetched, url)
			return body, nil
		},
		MaxItems: 1,
	}

	e.Run(types.Request{Url: seedURL, ParseFunc: parser.ParseCityList})

	if len(fetched) != 1 || fetched[0] != seedURL {
		t.Errorf("MaxItems=1 时应当只抓种子 %q，实际抓取: %v", seedURL, fetched)
	}
}
