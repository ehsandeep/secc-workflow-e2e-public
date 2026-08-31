package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/files/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/files/"):]
		data, err := os.ReadFile("data/" + name)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		_, _ = w.Write(data)
	})
	_ = http.ListenAndServe(":8080", nil)
}
