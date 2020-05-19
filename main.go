package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on: http://localhost:9000")
	mainServer := http.NewServeMux()
	mainServer.HandleFunc("/", serveMain)
	server1 := http.NewServeMux()
	server1.HandleFunc("/", serve1)

	server2 := http.NewServeMux()
	server2.HandleFunc("/", serve2)

	go func() {
		log.Println("Server started on: http://localhost:9001")
		http.ListenAndServe(":9001", server1)
	}()

	go func() {
		log.Println("Server started on: http://localhost:9002")
		http.ListenAndServe(":9002", server2)
	}()

	http.ListenAndServe(":9000", mainServer)
}
func serve1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "running on server 1")
}
func serve2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "running on server 2")
}
func serveMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "running on main server ")
}
