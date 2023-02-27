package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello xxx"))
	})

	http.ListenAndServe(":8080", nil)

	fmt.Println("success!")
}
