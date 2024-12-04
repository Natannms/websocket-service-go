package web

import (
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	})

	log.Println("Starting server on http://localhost:8080...")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
