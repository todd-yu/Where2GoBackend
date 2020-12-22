package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../.env")
	username := os.Getenv("db_username")
	password := os.Getenv("db_password")
	address := os.Getenv("db_hostname")
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)",
		username,
		password,
		address)
	// fmt.Println(username, password, address, dataSourceName)
	fmt.Println("attempting to load database...")
	database, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	fmt.Println("database successfully loaded!")
}

// func main() {
// 	godotenv.Load("../.env")
// 	username := os.Getenv("db_username")
// 	password := os.Getenv("db_password")
// 	address := os.Getenv("db_hostname")
// 	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)",
// 		username,
// 		password,
// 		address)
// 	fmt.Println(username, password, address, dataSourceName)
// }
