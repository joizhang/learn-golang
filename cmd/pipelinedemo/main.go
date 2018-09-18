package main

import (
	"bufio"
	"fmt"
	"imooc.com/joizhang/learn-golang/pipeline"
	"os"
)

func main() {
	//mergeDemo()
	const filename = "small.in"
	const n = 64
	writeAndReadDemo(filename, n)
}

func writeAndReadDemo(filename string, n int) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p := pipeline.RandomSource(n)
	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()
	file, err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	p = pipeline.ReaderSource(bufio.NewReader(file), -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}

func mergeDemo() {
	// 不会等待
	p := pipeline.Merge(
		pipeline.InMemorySort(pipeline.ArraySource(3, 2, 6, 7, 4)),
		pipeline.InMemorySort(pipeline.ArraySource(7, 4, 0, 1, 5)))
	// 在此处等待数据到来
	// 使用 range 则 chan 必须 close
	for v := range p {
		fmt.Println(v)
	}
}
