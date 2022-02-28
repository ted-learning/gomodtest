package common

import "io"

func HandleCloser(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		panic(err)
	}
}
