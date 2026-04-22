package main

import (
	"net/http"
	"os"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("static/index.html")
		if err != nil {
			http.Error(w, "No se pudo leer index.html", http.StatusInternalServerError)
			return
		}

		hostname, err := os.Hostname()
		if err != nil {
			hostname = "desconocido"
		}

		html := string(data)
		html = strings.ReplaceAll(html, "POD_PLACEHOLDER", "Pod: "+hostname)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write([]byte(html))
	})

	http.ListenAndServe(":8080", nil)
}
