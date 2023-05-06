package main

import (
	"net/http"
	"time"
)

var startedAt = time.Now()

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(startedAt)
		if duration.Seconds() < 10 {
			w.WriteHeader(500)
			w.Write([]byte("Erro interno"))
		} else {
			w.WriteHeader(200)
			w.Write([]byte("Estou no DevOpsDays Aracaju"))
		}
	})
	http.ListenAndServe(":8080", nil)
}
