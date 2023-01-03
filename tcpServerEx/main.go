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

var countResult string

func main() {
	http.HandleFunc("/result", resultHandler)

	fmt.Println("Server Start")
	// server start
	http.ListenAndServe(":8080", nil)

}

func resultHandler(response http.ResponseWriter, request *http.Request) {
	mysql()

	json.NewEncoder(response).Encode("Hi There")

	responseBase := `total Counts: %s`

	responsing := fmt.Sprintf(responseBase, countResult)

	json.NewEncoder(response).Encode(responsing)

}

func mysql() {
	getEnv()

	db := openDb()
	queryCount(db)
}

func getEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

}

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
	clientTable := os.Getenv("CLIENT_TABLE")

	selectState := `SELECT COUNT(name) as count FROM %s`

	queryStatement := fmt.Sprintf(selectState, clientTable)

	count, err := database.Query(queryStatement)

	if err != nil {
		panic(err.Error())
	}

	for count.Next() {
		count.Scan(&countResult)

		fmt.Println(countResult)
	}
}
