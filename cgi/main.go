package main

import (
	"io/ioutil"
	"net/http"
	"net/http/cgi"
)

func main() {
	err := cgi.Serve(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("https://httpbin.org/get")
		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			defer resp.Body.Close()
			b, _ := ioutil.ReadAll(resp.Body)
			w.Write(b)
		}
	}))

	if err != nil {
		println(err.Error())
	}
}
