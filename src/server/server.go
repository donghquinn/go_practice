package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "strings"

	// "strings"

	mux "github.com/gorilla/mux"
)

// 요청으로 들어올 유저 정보

const port = ":8080"

type UserStruct struct {
	Name        string `json:"name"`
	Age         string `json:"age"`
	Gender      string `json:"gender"`
	Email       string `json:"email"`
	DateOfBirth string `json:"dateOfBirth"`
}

// var Users []UserStruct

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

// 사용자 관련 요청
func userSetReturn(write http.ResponseWriter, read *http.Request) {
	fmt.Println("Endpoint: Users")

	parameters := read.URL.Query()

	var parameterArray []string

	// json.NewEncoder(write).Encode(parameters)

	for params, value := range parameters {
		fmt.Println(params, "=>", value)

		parameterArray = append(parameterArray, value...)
	}

	User := &UserStruct{
		Name:        parameterArray[0],
		Age:         parameterArray[1],
		Gender:      parameterArray[2],
		Email:       parameterArray[3],
		DateOfBirth: parameterArray[4],
	}

	json.NewEncoder(write).Encode(parameterArray)
	json.NewEncoder(write).Encode(User)
	fmt.Fprintf(write, "User Request")
	fmt.Fprintln(write, User)
}

// 라우터 핸들러
func handleRequests(users ...UserStruct) {
	router := mux.NewRouter()

	router.HandleFunc("/users", userSetReturn)
	router.HandleFunc("/", homeReturn)
	router.HandleFunc("/hello", hellowReturn)

	println("Listening: 8080")

	// 에러 핸들링
	err := http.ListenAndServe(port, router)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func main() {

	// fmt.Println(User)

	// log.Default(Users)

	handleRequests()
	// log.Fatal(http.ListenAndServe(":8080", nil))

	// handleRequests(Users...)
}

// 유저 정보 할당
