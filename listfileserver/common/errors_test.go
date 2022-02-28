package common

import (
	"errors"
	"gomodtest/listfileserver/apps"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func panicError(_ http.ResponseWriter, _ *http.Request) error {
	panic("123")
}

func urlNotSupportedError(_ http.ResponseWriter, _ *http.Request) error {
	return apps.URLNotSupportedError("URL Not Supported Error")
}

func notFoundError(_ http.ResponseWriter, _ *http.Request) error {
	return os.ErrNotExist
}

func permissionError(_ http.ResponseWriter, _ *http.Request) error {
	return os.ErrPermission
}

func internalSeverError(_ http.ResponseWriter, _ *http.Request) error {
	return errors.New("internal sever error")
}

func noError(response http.ResponseWriter, _ *http.Request) error {
	_, err := response.Write([]byte("no error"))
	if err != nil {
		return err
	}
	return nil
}

var givens = []struct {
	app  App
	code int
	msg  string
}{
	{panicError, 500, "Internal Server Error"},
	{urlNotSupportedError, 400, "URL Not Supported Error"},
	{notFoundError, 404, "Not Found"},
	{permissionError, 403, "Forbidden"},
	{internalSeverError, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func assertResponse(t *testing.T, actualResponse io.Reader, actualCode int, given struct {
	app  App
	code int
	msg  string
}) {
	b, _ := ioutil.ReadAll(actualResponse)
	actualBody := strings.Trim(string(b), "\n")
	if actualBody != given.msg || actualCode != given.code {
		t.Errorf("expect (%d %s), but got (%d %s)", given.code, given.msg, actualCode, actualBody)
	}
}

func TestErrorWrap(t *testing.T) {
	for _, given := range givens {
		f := ErrorWrap(given.app)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodGet, "/test", nil)

		f(response, request)
		assertResponse(t, response.Body, response.Code, given)
	}
}

func TestServer(t *testing.T) {
	for _, given := range givens {
		f := ErrorWrap(given.app)
		server := httptest.NewServer(http.HandlerFunc(f))

		response, _ := http.Get(server.URL)
		assertResponse(t, response.Body, response.StatusCode, given)
	}

}

func BenchmarkGetServer(b *testing.B) {
	f := ErrorWrap(noError)
	server := httptest.NewServer(http.HandlerFunc(f))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		response, _ := http.Get(server.URL)
		body, _ := ioutil.ReadAll(response.Body)
		actualBody := strings.Trim(string(body), "\n")

		if response.StatusCode != 200 {
			b.Errorf("expect 200, but got %d %s", response.StatusCode, actualBody)
		}
	}
}
