package apps

import (
	"fmt"
	"gomodtest/listfileserver/common"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const ListFileReaderURL = "/list/"

type URLNotSupportedError string

func (error URLNotSupportedError) Error() string {
	return string(error)
}

//func (error URLNotSupportedError) Message() string {
//
//}

func ListFileReader(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, ListFileReaderURL) != 0 {
		return URLNotSupportedError(fmt.Sprintf("the URL Path: %s not supported\n", request.URL.Path))
	}

	subPath := request.URL.Path[len(ListFileReaderURL):]
	file, err := os.Open(subPath)

	if err != nil {
		return err
	}
	defer common.HandleCloser(file)

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	_, err = writer.Write(all)
	if err != nil {
		return err
	}

	return nil
}
