package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// // Existing code from above
// func handleRequests() {
// 	// creates a new instance of a mux router
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	// replace http.HandleFunc with myRouter.HandleFunc
// 	myRouter.HandleFunc("/", homePage)
// 	myRouter.HandleFunc("/all", returnAllArticles)
// 	// finally, instead of passing in nil, we want
// 	// to pass in our newly created router as the second
// 	// argument
// 	log.Fatal(http.ListenAndServe(":10000", myRouter))
// }

// func main() {
// 	fmt.Println("Rest API v2.0 - Mux Routers")

// 	Articles = []Article{
// 		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
// 		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
// 	}

// 	handleRequests()
// }
