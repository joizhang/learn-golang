package engine

import (
	"log"

	"imooc.com/joizhang/learn-golang/crawler/types"
)

type SimpleEngine struct {
	// Fetch 可选，注入自定义抓取函数；为 nil 时使用 fetcher.Fetch
	Fetch FetchFunc
	// MaxItems 可选，抓够这么多 item 后停止；0 表示不限制
	MaxItems int
}

// 串行
func (e SimpleEngine) Run(seeds ...types.Request) {
	var requests []types.Request
	requests = append(requests, seeds...)

	fetch := resolveFetch(e.Fetch)

	itemCount := 0
	for len(requests) > 0 {
		if e.MaxItems > 0 && itemCount >= e.MaxItems {
			return
		}

		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(fetch, r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
			itemCount++
			if e.MaxItems > 0 && itemCount >= e.MaxItems {
				return
			}
		}
	}
}
