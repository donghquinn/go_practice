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

/*
"/" 로 들어온 요청 핸들링
@write는 응답을 저장하는 용도
@read는 요청을 읽는 용도. 읽은 요청은 메모리 포인터로 저장
*/
func homeReturn(write http.ResponseWriter, read *http.Request) {
	fmt.Fprintf(write, "MainPage")
	fmt.Println("EndPoint: Home")

	// 응답 값에 담을 내용을 인코딩
	json.NewEncoder(write).Encode("Welcome")
}

func hellowReturn(write http.ResponseWriter, read *http.Request) {

	fmt.Fprintf(write, "hello")
	fmt.Println("Endpoint: Hello")

	json.NewEncoder(write).Encode("Hello")
}

func userSetReturn(write http.ResponseWriter, read *http.Request) {
	fmt.Println("Endpoint: Users")

	parametr := mux.Vars(read)

	if parametr == nil {
		errMsg := "No Parameter Provided"

		json.NewEncoder(write).Encode(errMsg)
	}

	name := parametr["name"]
	age := parametr["age"]
	gender := parametr["gender"]
	email := parametr["email"]
	dateOfBirth := parametr["dateOfBirth"]

	for _, user := range Users {
		user.age = age
		user.name = name
		user.gender = gender
		user.email = email
		user.dateOfBirth = dateOfBirth

	}
	json.NewEncoder(write).Encode(name)
	json.NewEncoder(write).Encode(age)
	json.NewEncoder(write).Encode(gender)
	json.NewEncoder(write).Encode(email)
	json.NewEncoder(write).Encode(dateOfBirth)

	fmt.Fprintf(write, "User Request")
}

func handleRequests(users ...user) {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/users/{name}/{age}/{gender}/{email}/{dateOfBirth}", userSetReturn)
	router.HandleFunc("/", homeReturn)
	router.HandleFunc("/hello", hellowReturn)

	println("Listening: 8080")
	log.Fatal(http.ListenAndServe(port, router))
}

func main() {

	fmt.Println(Users)

	// log.Default(Users)

	handleRequests()
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// handleRequests(Users...)
}

// 유저 정보 할당
