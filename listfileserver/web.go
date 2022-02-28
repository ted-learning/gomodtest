package main

import (
	"gomodtest/listfileserver/apps"
	"gomodtest/listfileserver/common"
	"net/http"
)

func httpHandleFunc(url string, app common.App) {
	http.HandleFunc(url, common.ErrorWrap(app))
}

func main() {
	httpHandleFunc("/", apps.ListFileReader)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
