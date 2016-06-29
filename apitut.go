package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	fmt.Println("Endpoint hit: returnAllArticles method:" + r.Method)
	json.NewEncoder(w).Encode(articles)
}

func returnOneArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	fmt.Fprintln(w, "key: "+key)
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
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{key}", returnOneArticle)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}
