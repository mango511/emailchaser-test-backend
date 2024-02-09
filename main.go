package main

import (
	"context"
	"log"
	controller "main/controllers"
	"main/ent"

	"github.com/gin-gonic/gin"
)

func main() {
	client, err := ent.Open("mysql", "user:password@tcp(localhost:3306)/dbname?parseTime=True")
	if err != nil {
		log.Fatalf("failed to connect to mysql: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	r := gin.Default()

	r.POST("/users", controller.CreateUserHandler(client))
	r.POST("/invite", controller.InvitationHandler())

	r.Run(":8080")
}
