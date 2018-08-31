package fetcher

import (
	"fmt"
	"testing"
)

func TestFetch(t *testing.T) {
	s, err := Fetch("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", s)
}
