package main

import "github.com/gin-gonic/gin"

func DeleteEndPoint(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	// Get id endpoint
	id := c.Params.ByName("id")
	var endpoint EndPoint
	// SELECT * FROM endpoints WHERE id = 1;
	db.First(&endpoint, id)

	if endpoint.Id != 0 {
		// DELETE FROM endpoints WHERE id = endpoint.Id
		db.Delete(&endpoint)
		// Display JSON result
		c.JSON(200, gin.H{"success": "EndPoint #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "EndPoint not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/endpoints/1
}

func DeleteMaintenance(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	// Get id endpoint
	id := c.Params.ByName("id")
	var maintenance Maintenance
	// SELECT * FROM endpoints WHERE id = 1;
	db.First(&maintenance, id)

	if maintenance.Id != 0 {
		// DELETE FROM endpoints WHERE id = endpoint.Id
		db.Delete(&maintenance)
		// Display JSON result
		c.JSON(200, gin.H{"success": "maintenance #" + id + " deleted"})
	} else {
		// Display JSON error
		c.JSON(404, gin.H{"error": "maintenance not found"})
	}

	// curl -i -X DELETE http://localhost:8080/api/v1/maintenance/1
}

