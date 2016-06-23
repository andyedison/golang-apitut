package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Article our model for our articles API
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

//Articles slice of Article
type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Homepage!")
	fmt.Println("Endpoint hit: homePage")
}

func returnArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "returns a specific article")
	fmt.Println("Endpoint hit: returnArticle")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Desc", Content: "Lorem Ipsum"},
		Article{Title: "Hello 2", Desc: "Article Desc", Content: "Lorem Ipsum"},
	}
	fmt.Println("Endpoint hit: returnAllArticles")
	json.NewEncoder(w).Encode(articles)
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
