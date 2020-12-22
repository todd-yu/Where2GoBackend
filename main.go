package main

import (
	"fmt"
	"net/http"
)

func loadDatabase() {
	// godotenv.Load("../.env")
	// username := os.Getenv("db_username")
	// password := os.Getenv("db_password")
	// address := os.Getenv("db_hostname")
	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)",
	// 	username,
	// 	password,
	// 	address)
	// fmt.Println(username, password, address, dataSourceName)
	// fmt.Println("attempting to load database...")
	// database, err := sql.Open("mysql", dataSourceName)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("database successfully loaded!")
}

func handler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("hello world"))
	fmt.Println("yaw yeet")
}

func main() {
	loadDatabase()
	http.HandleFunc("/hello", handler)
	fmt.Println("Listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
