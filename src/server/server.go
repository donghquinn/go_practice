package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mux "github.com/gorilla/mux"
)

// 요청으로 들어올 유저 정보

const port = ":8080"

type user struct {
	name        string
	age         string
	gender      string
	email       string
	dateOfBirth string
}

var Users []user

// func init() {

//		http.ListenAndServe(port)
//	}
func homeReturn(write http.ResponseWriter, read *http.Request) {
	fmt.Fprintf(write, "MainPage")
	fmt.Println("EndPoint: Home")

	json.NewEncoder(write).Encode(Users)
}

func hellowReturn(write http.ResponseWriter, read *http.Request) {
	fmt.Fprintf(write, "hello")
	fmt.Println("Hellow Page")

	json.NewEncoder(write).Encode(Users)
}

func handleRequests(users ...user) {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/", homeReturn)
	router.HandleFunc("/hello", hellowReturn)
	println("Listening: 8080")
	log.Fatal(http.ListenAndServe(port, router))
}

func main() {

	Users = []user{
		{name: "kim", age: "20", gender: "mail", email: "mail@mail.com", dateOfBirth: "dong"},
	}

	fmt.Println(Users)

	// log.Default(Users)

	handleRequests()
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// handleRequests(Users...)
}

// 유저 정보 할당
