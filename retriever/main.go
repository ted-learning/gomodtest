package main

import (
	"fmt"
	"gomodtest/retriever/mock"
	"gomodtest/retriever/real"
	"time"
)

type retriever interface {
	Get(url string) string
}

type poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPoster interface {
	retriever
	poster
}

const url = "https://www.imooc.com"

func download(retriever retriever) {
	html := retriever.Get(url)
	fmt.Println(html)
}

func post(poster poster) {
	html := poster.Post(url, map[string]string{
		"Contents": "post lalala",
	})
	fmt.Println(html)
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"Contents": "This is other fake  content",
	})
	return s.Get(url)
}

func main() {

	r1 := &mock.Retriever{Content: "This is fake imooc."}
	inspect(r1)
	//download(retriever)
	r2 := &real.Retriever{TimeOut: 2 * time.Minute}
	inspect(r2)
	//download(retriever2)
	var r retriever
	r = r2
	if r3, ok := r.(*mock.Retriever); ok {
		fmt.Printf(r3.Content)
	} else {
		fmt.Printf("this is not a mock")
	}

	fmt.Println("Try a session")
	fmt.Println(session(r1))
}

func inspect(r retriever) {
	fmt.Println("Inspect:", r)
	fmt.Printf("> %T %v\n", r, r)
	fmt.Print("> Type switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("This is mock, content:", v.Content)
	case *real.Retriever:
		fmt.Println("This is real, timeout:", v.TimeOut)
	}
	fmt.Println()
}
