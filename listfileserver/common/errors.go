package common

import (
	"go.uber.org/zap"
	"net/http"
	"os"
)

type Controller func(writer http.ResponseWriter, request *http.Request)

type App func(writer http.ResponseWriter, request *http.Request) error

var zapLog, _ = zap.NewProduction()

type BadRequest interface {
	error
}

func ErrorWrap(app App) Controller {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover()
			if err, ok := r.(error); ok {
				zapLog.Error(err.Error())
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := app(writer, request)
		if err == nil {
			return
		}
		zapLog.Error(err.Error())

		if badRequest, isBadRequest := err.(BadRequest); isBadRequest {
			http.Error(writer, badRequest.Error(), http.StatusBadRequest)
			return
		}

		var code int
		switch {
		case os.IsNotExist(err):
			code = http.StatusNotFound
		case os.IsPermission(err):
			code = http.StatusForbidden
		default:
			code = http.StatusInternalServerError
		}
		http.Error(writer, http.StatusText(code), code)
	}
}
