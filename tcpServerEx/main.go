package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

// 쿼리한 결과 값
var countResult string

/*
응답 구조체
대문자 - export field로 인식
*/
type ResponseType struct {
	CodeRes int
	DataRes string
	ErrMsg  string
}

func main() {
	// "/result" 라우트에 대한 핸들러
	http.HandleFunc("/result", resultHandler)

	fmt.Println("Server Start")

	// server start
	http.ListenAndServe(":8080", nil)

}

// count result set response
func resultHandler(response http.ResponseWriter, request *http.Request) {
	mysql()

	setResponse(response)
}

// 응답
func setResponse(response http.ResponseWriter) {

	json.NewEncoder(response).Encode("Hi There")

	responseBase := `total Counts: %s`

	responsing := fmt.Sprintf(responseBase, countResult)

	responseData := ResponseType{CodeRes: 200, DataRes: responsing, ErrMsg: ""}

	json.NewEncoder(response).Encode(responseData)
}

// mysql controller
func mysql() {
	getEnv()

	db := openDb()
	queryCount(db)
}

// env loader
func getEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
}

// Database 연결 함수
func openDb() *sql.DB {
	// 환경변수 로드
	dbPasswd := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// fmt.Println(dbUser, dbPasswd, dbName, dbHost, dbHost, dbPort)

	// fmt.Println(dbUser + ":" + dbPasswd + "@" + "tcp" + "(" + dbHost + ":" + dbPort + ")" + "/" + dbName)

	// Database 연결
	database, err := sql.Open("mysql", dbUser+":"+dbPasswd+"@"+"tcp"+"("+dbHost+":"+dbPort+")"+"/"+dbName)

	// Database 연결 설정
	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)

	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	return database
}

func queryCount(database *sql.DB) {
	// get table name
	clientTable := os.Getenv("CLIENT_TABLE")

	selectState := `SELECT COUNT(name) as count FROM %s`

	// make query statement
	queryStatement := fmt.Sprintf(selectState, clientTable)

	// request query
	count, err := database.Query(queryStatement)

	if err != nil {
		panic(err.Error())
	}

	for count.Next() {
		count.Scan(&countResult)

		fmt.Println(countResult)
	}
}
