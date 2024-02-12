package main

import (
	"context"
	"log"

	"main/ent"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Update the connection string below with your MySQL database credentials
	databaseURL := "root:@tcp(localhost:3306)/emailchaser_db?parseTime=True"
	client, err := ent.Open("mysql", databaseURL)
	if err != nil {
		log.Fatalf("failed to connect to mysql: %v", err)
	}
	defer client.Close()

	// Automatically migrate your schema
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
