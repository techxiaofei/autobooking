package query

import (
	"compress/flate"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func ContentEncoding(res *http.Response) (bodyReader io.Reader, err error) {
	switch res.Header.Get("Content-Encoding") {
	case "gzip":
		bodyReader, err = gzip.NewReader(res.Body)
	case "deflate":
		bodyReader = flate.NewReader(res.Body)
	default:
		bodyReader = res.Body
	}
	return
}

func WithSlot(text string) bool {
	if !strings.Contains(text, "slotId") {
		fmt.Println("no slot release")
		return false
	}
	if strings.Contains(text, "session has expired") {
		fmt.Printf("session expired:%v\n", time.Now().Format("2016-01-02 11:11:11"))
		return false
	}

	return true
}
