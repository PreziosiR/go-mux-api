package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)
func main() {
	router := mux.NewRouter()
	const port string = ":8000"
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "../mux-api-98344-firebase-adminsdk-nj3uj-2baadb873d.json")
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")
	router.HandleFunc("/posts", addPosts).Methods("POST")
	log.Println("Server listening on port : ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
