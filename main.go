package main

import (
	"context"
	"log"
	controller "main/controllers"
	"main/ent"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Ensure you're using the blank identifier to import if not directly using the package.
)

func main() {
	// Open a connection to your database.
	client, err := ent.Open("mysql", "root:@tcp(localhost:3306)/emailchaser_db?parseTime=True")
	if err != nil {
		log.Fatalf("failed to connect to mysql: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool to create your schema.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Initialize Gin router.
	r := gin.Default()

	// Setup your routes.
	r.POST("/users", controller.CreateUserHandler(client))
	// Here you're passing the client to InviteUser, which returns a gin.HandlerFunc compatible function.
	r.POST("/invite", controller.InviteUser(client))

	// Start the Gin server.
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
	// The log.Fatalf("Okay") line would never be reached; if r.Run returns an error, it'll be caught by the if statement above.
}
