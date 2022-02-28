package apps

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const ListFileReaderURL = "/list/"

type URLNotSupportedError string

func (error URLNotSupportedError) IsBadRequest() bool {
	return true
}

func (error URLNotSupportedError) Error() string {
	return string(error)
}

func ListFileReader(writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, ListFileReaderURL) != 0 {
		return URLNotSupportedError(fmt.Sprintf("the URL Path: %s not supported", request.URL.Path))
	}

	subPath := request.URL.Path[len(ListFileReaderURL):]
	file, err := os.Open(subPath)

	if err != nil {
		return err
	}
	defer handleCloser(file)

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

func handleCloser(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		panic(err)
	}
}
