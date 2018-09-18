package engine

import (
	"imooc.com/joizhang/learn-golang/crawler/fetcher"
	"imooc.com/joizhang/learn-golang/crawler/types"
	"log"
)

type SimpleEngine struct{}

// 串行
func (e SimpleEngine) Run(seeds ...types.Request) {
	var requests []types.Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r types.Request) (types.ParseResult, error) {
	// log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s: %v", r.Url, err)
		return types.ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}
