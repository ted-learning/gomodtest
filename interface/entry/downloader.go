package main

import (
	"fmt"
	_interface "gomodtest/interface"
	"gomodtest/interface/test"
)

type Retriever interface {
	Get(url string) string
}

func getRetriever() Retriever {
	return _interface.Retriever{}
}

func getFakeRetriever() Retriever {
	return test.Retriever{}
}

func main() {
	retriever := getRetriever()
	fmt.Println(retriever.Get("https://coding.imooc.com"))

	retriever2 := getFakeRetriever()
	fmt.Println(retriever2.Get("https://coding.imooc.com"))
}
