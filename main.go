package main

import (
	"fmt"
	"library/server"
)

func main() {
	fmt.Println("Welcome to the Library...")
	server.SetupRouter()
}
