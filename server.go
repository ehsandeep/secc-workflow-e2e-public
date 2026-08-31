package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	http.HandleFunc("/files/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Path[len("/files/"):]
		clean := filepath.Clean("/" + name)
		if strings.Contains(clean, "..") {
			http.Error(w, "invalid path", http.StatusBadRequest)
			return
		}
		data, err := os.ReadFile(filepath.Join("data", clean))
		if err != nil {
			http.NotFound(w, r)
			return
		}
		_, _ = w.Write(data)
	})
	_ = http.ListenAndServe(":8080", nil)
}
