package main

import (
	"log"
	"net/http"
	"net/http/cgi"
	"os"
)

var (
	port   = ":7070"
	GOEXEC = "/usr/local/bin/go"
	PWD    string
)

func init() {
	PWD, _ = os.Getwd()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler := new(cgi.Handler)
		handler.Path = GOEXEC
		handler.Dir = PWD + "/cgi/"
		args := []string{"run", PWD + "/cgi/main.go"}
		handler.Args = append(handler.Args, args...)
		handler.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
