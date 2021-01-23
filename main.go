package main

import (
	"fmt"
	"library/database"
	"library/server"
	"library/web"
)

func main() {
	fmt.Println("Welcome to the Library...")
	database.LoadDatabaseConfig("DatabaseConfig.json")
	router := server.SetupRouter()
	fmt.Println("Load Web")
	web.SetupWebRouter(router)
	router.Run(":8445")
}
