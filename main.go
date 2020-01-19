package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func helloHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "Hello", nil)
	if err != nil {
		log.Printf("Unexepected err: %v", err)
	}
}

func main() {
	fmt.Println("Hello Word")
	tmdbKey := os.Getenv("TMDB_KEY")
	fmt.Println(tmdbKey)
	http.HandleFunc("/", helloHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println(err)
	}
}
