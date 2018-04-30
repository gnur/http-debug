package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%q", "ok")
		fmt.Println("----------------------------")
		fmt.Println(string(dump))
		fmt.Println("----------------------------")
	})

	log.Println("Now listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
