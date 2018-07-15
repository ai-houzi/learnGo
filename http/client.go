package main

import (
	"net/http"
	"net/http/httputil"
	"fmt"
)

func main() {

	request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	if err != nil {
		panic(err.Error())
	}

	request.Header.Add("User-Agent",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("redirect:",request)
			return nil
		},
	}

	resp, err := client.Do(request)

	if err != nil {
		panic(err.Error())
	}

	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%s\n", s)
}
