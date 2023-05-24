package main

import (
	"fmt"
	db2 "github.com/AnwarSaginbai/restapi-with-docker/internal/db"
)

func Run() error {
	fmt.Println("Starting up our application")

	db, err := db2.NewDatabase()
	if err != nil {
		fmt.Println("failed to connect to the database")
		return err
	}

	fmt.Println("Successfully connected and pinged database")
	return nil
}

func main() {
	fmt.Println("Starting up our application")
	Run()
}
