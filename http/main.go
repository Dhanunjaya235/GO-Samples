package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func hello(w http.ResponseWriter, req *http.Request) {

	reponse, err := json.Marshal(Message{"Alice", "Hello", 1294706395881547000})
	if err != nil {
		fmt.Fprintf(w, "hello\n")
		return
	}
	w.Write(reponse)
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
