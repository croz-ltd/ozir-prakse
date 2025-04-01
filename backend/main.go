package main

import (
	"log"
	"net/http"
)

func main() {
	randomWords := []string{"CI/CD", "API", "Kubernetes", "Helm", "AWS", "Terraform", "Ansible", "Jenkins"}

	var i uint32

	i = 0
	n := uint32(len(randomWords))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		log.Println("got request on /, generating a random word...")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(randomWords[i % n]))

		i += 1
	})

	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalln("failed to start API!")
	}
}