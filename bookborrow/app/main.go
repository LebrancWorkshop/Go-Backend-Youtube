package main

import (
	"os"
)

func main() {
	// Get Environtment Variables. 
	dialect := os.Getenv("DIALECT");
	host := os.Getenv("HOST");
	port := os.Getenv("PORT");
	user := os.Getenv("USER");
	password := os.Getenv("PASSWORD");
	name := os.Getenv("NAME");
}