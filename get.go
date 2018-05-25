package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetEndPoints(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	var endpoints EndPoints
	// SELECT * FROM endpoints
	db.Find(&endpoints)

	// Display JSON result
	c.JSON(200, endpoints)

	// curl -i http://localhost:8080/api/v1/endpoints
}

func GetMaintenances(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	var maintenance Maintenances
	// SELECT * FROM endpoints
	db.Find(&maintenance)

	fmt.Printf("maintenance##########--------%+v\n", maintenance[0].StartTime)
	// Display JSON result
	c.JSON(200, maintenance)

	// curl -i http://localhost:8080/api/v1/maintenance
}

func GetEndPoint(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("label")
	var endpoint EndPoint
	// SELECT * FROM endpoints WHERE id = 1;
	db.First(&endpoint, id)

	if endpoint.Id != 0 {
		// Display JSON result
		c.JSON(200, endpoint)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Endpoint not found"})
	}
	// curl -i http://localhost:8080/api/v1/endpoints/1
}

func GetValidEndPoints(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("label")
	var endpoint EndPoint
	// SELECT * FROM endpoints WHERE id = 1;
	db.First(&endpoint, id)

	if endpoint.Id != 0 {
		// Display JSON result
		c.JSON(200, endpoint)
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "Endpoint not found"})
	}
	// curl -i http://localhost:8080/api/v1/endpoints/1
}
