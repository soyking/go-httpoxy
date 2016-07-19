package main

import (
	"net/http"
	"net/http/httputil"
)

var port = ":7071"

func main() {
	http.ListenAndServe(port, http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			b, err := httputil.DumpRequest(r, true)
			if err != nil {
				println(err.Error())
			} else {
				println(string(b))
			}
			w.Write([]byte("fake response"))
		},
	))
}
