package main

import (
	"fmt"
	"library/database"
	"library/server"
)

func main() {
	fmt.Println("Welcome to the Library...")
	database.LoadDatabaseConfig("DatabaseConfig.json")
	server.SetupRouter()
}