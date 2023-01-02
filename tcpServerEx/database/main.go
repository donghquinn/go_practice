package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

var countResult string

func main() {
	db := openDb()
	queryCount(db)
	getEnv()
}

func getEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}

}

func openDb() *sql.DB {
	dbPasswd := os.Getenv("DB_PASS")
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	fmt.Println(dbUser, dbPasswd, dbName, dbHost, dbHost, dbPort)

	fmt.Println(dbUser + ":" + dbPasswd + "@" + "tcp" + "(" + dbHost + ":" + dbPort + ")" + "/" + dbName)

	database, err := sql.Open("mysql", dbUser+":"+dbPasswd+"@"+"tcp"+"("+dbHost+":"+dbPort+")"+"/"+dbName)

	database.SetMaxOpenConns(10)
	database.SetMaxIdleConns(10)

	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	return database
}

func queryCount(database *sql.DB) {
	count, err := database.Query("SELECT COUNT(name) as count FROM client")

	if err != nil {
		panic(err.Error())
	}

	for count.Next() {
		count.Scan(&countResult)

		fmt.Println(countResult)
	}
}
