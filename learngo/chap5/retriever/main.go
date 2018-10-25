package main

import (
	"fmt"
	"gostudy/learngo/chap5/retriever/mock"
)

type Retriever interface {
	Get(url string) string
}

func downLoad(r Retriever) string {
	return r.Get("wwww.imooc.com")
}
func main() {
	var r Retriever
	r = mock.Retriever{"this is haha"}
	fmt.Println(downLoad(r))
}
