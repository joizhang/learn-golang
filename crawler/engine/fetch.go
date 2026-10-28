package engine

import (
	"fmt"
	"log"

	"imooc.com/joizhang/learn-golang/crawler/fetcher"
	"imooc.com/joizhang/learn-golang/crawler/types"
)

// FetchFunc 抓取一个 URL 并返回页面内容。
// 抽成函数类型是为了让测试可以注入本地快照，不依赖真实网络。
type FetchFunc func(url string) ([]byte, error)

// resolveFetch 返回实际使用的抓取函数，没有注入时就用默认的 fetcher.Fetch。
func resolveFetch(f FetchFunc) FetchFunc {
	if f != nil {
		return f
	}
	return fetcher.Fetch
}

func worker(fetch FetchFunc, r types.Request) (types.ParseResult, error) {
	// log.Printf("Fetching %s", r.Url)
	body, err := fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return types.ParseResult{}, fmt.Errorf("fetch %s: %w", r.Url, err)
	}
	return r.ParseFunc(body), nil
}
