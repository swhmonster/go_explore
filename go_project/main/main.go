package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "/Users/sunwenhao/Downloads/temp/1.txt"
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(data)
	fmt.Fprintf(w, filePath)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/file", fileHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
