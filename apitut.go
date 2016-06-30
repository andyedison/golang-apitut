package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	r "gopkg.in/dancannon/gorethink.v2"
)

var (
	session *r.Session
)

//Article our model for our articles API
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	// Id      string `gorethink:"id,omitempty"`
}

func init() {
	var err error

	session, err = r.Connect(r.ConnectOpts{
		Address:  "192.168.99.100:32781",
		Database: "blog",
		MaxOpen:  40,
	})

	if err != nil {
		log.Fatalln(err.Error())
	}
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

func returnAllArticles(w http.ResponseWriter, req *http.Request) {
	articles := []Article{}

	res, err := r.Table("article").Run(session)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer res.Close()
	err = res.All(&articles)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Endpoint hit: returnAllArticles method:" + req.Method)
	json.NewEncoder(w).Encode(articles)
}

func returnOneArticle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	key := vars["key"]

	fmt.Fprintln(w, "key: "+key)
}

func addArticles(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "add article")
	fmt.Println("Endpoint hit: addArticles")
}

func delArticles(w http.ResponseWriter, req *http.Request) {
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
	fmt.Println("Rest API v3.0 - Mux Routers")
	handleRequests()
}
