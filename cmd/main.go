package main

import (
	"dr.agenda/database"
	"dr.agenda/factory"
	"github.com/gin-gonic/gin"
)

func main() {
	const serverPort string = ":8080"
	server := gin.Default()

	dbConnection, dbErr := database.ConnectDB()

	if dbErr != nil {
		panic(dbErr)
	}

	HealthInsuranceHandler := factory.HealthInsuranceFactory(dbConnection)

	server.GET("/api/health-insurance", HealthInsuranceHandler.GetInsurances)
	server.GET("/api/health-insurance/:id", HealthInsuranceHandler.GetInsuranceById)
	server.POST("/api/health-insurance", HealthInsuranceHandler.CreateHealthInsurance)

	server.Run(serverPort)
}
