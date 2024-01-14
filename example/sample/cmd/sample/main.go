
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, world!")
	handler1 := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.Path)
		io.WriteString(w, "Hello-1\n")
	}
	handler2 := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL.Path)
		io.WriteString(w, "Hello-2\n")
	}

	http.HandleFunc("/foo", handler1)
	http.HandleFunc("/bar", handler2)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
