package main

import (
	"regexp"
	"fmt"
)

const text = `
My email is zxlwhom@outlook.com
My email is zxlwhom@outlook.com
My email is         kk@qq.com
My email is         ddd@qqqq.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}
}
