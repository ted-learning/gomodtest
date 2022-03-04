package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
)

func main() {
	request, err := http.NewRequest(http.MethodGet, "https://www.imooc.com/", nil)
	if err != nil {
		panic(err)
	}

	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect: ", req)
			return nil
		},
	}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(response.Body)

	//all, err := ioutil.ReadAll(response.Body)
	dumpResponse, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", dumpResponse)
}
