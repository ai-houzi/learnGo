package ayy

import (
	"time"
	"net/http"
	"net/http/httputil"
)

type Retriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (t Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	bytes, e := httputil.DumpResponse(resp, true)
	resp.Body.Close()
	if err != nil {
		panic(e)
	}

	return string(bytes)
}
