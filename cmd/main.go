package main

import (
	"database/sql"
	"dr.agenda/database"
	"dr.agenda/factory"
	"dr.agenda/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	const serverPort = ":8085"

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	server := setupServer(db)

	log.Printf("Server is running on port %s", serverPort)
	if err := server.Run(serverPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

func setupServer(db *sql.DB) *gin.Engine {
	server := gin.Default()

	healthInsuranceHandler := factory.InsuranceFactory(db)

	routes.RegisterInsuranceRoutes(server, &healthInsuranceHandler)

	return server
}
