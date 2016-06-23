package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Homepage!")
	fmt.Println("Endpoint hit: homePage")
}

func returnArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "returns a specific article")
	fmt.Println("Endpoint hit: returnArticle")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "all articles")
	fmt.Println("Endpoint hit: returnAllArticles")
}

func addArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "add article")
	fmt.Println("Endpoint hit: addArticles")
}

func delArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete article")
	fmt.Println("Endpoint hit: delArticles")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/all", returnAllArticles)
	http.HandleFunc("/single", returnArticle)
	http.HandleFunc("/delete", delArticles)
	http.HandleFunc("/add", addArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
