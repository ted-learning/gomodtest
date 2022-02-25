package _interface

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Retriever struct {
}

func (Retriever) Get(url string) string {
	get, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(get.Body)

	all, err := ioutil.ReadAll(get.Body)

	if err != nil {
		panic(err)
	}

	return string(all)
}
