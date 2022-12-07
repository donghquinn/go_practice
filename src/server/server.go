package server

import (
	"fmt"
	"log"
	"net/http"
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

func init() {

	http.ListenAndServe(port, nil)
}

func homeReturn(write http.ResponseWriter, read *http.Request) {
	fmt.Fprintf(write, "MainPage")
	fmt.Println("EndPoint: Home")
}

func hellowReturn(write http.ResponseWriter, read *http.Request) {
	fmt.Fprintf(write, "hello")
	fmt.Println("Hellow Page")
}

func handleRequests(users ...user) {
	http.HandleFunc("/", homeReturn)
	http.HandleFunc("/hello", hellowReturn)
	println("Listening: 3000")
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	var Users []user

	Users = []user{
		{name: "kim", age: "20", gender: "mail", email: "mail@mail.com", dateOfBirth: "dong"},
	}

	fmt.Println(Users)

	// log.Fatal(http.ListenAndServe(":8080", nil))

	handleRequests(Users...)
}

// 유저 정보 할당
