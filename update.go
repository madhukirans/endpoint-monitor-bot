package main

import "github.com/gin-gonic/gin"

func UpdateEndPoint(c *gin.Context) {
	// Connection to the database
	db := InitDb()
	// Close connection database
	defer db.Close()

	// Get id endpoint
	id := c.Params.ByName("label")
	var endpoint EndPoint
	// SELECT * FROM endpoints WHERE id = 1;

	db.First(&endpoint, id)

	if endpoint.Label != "" && endpoint.Host != "" {

		if endpoint.Id != 0 {
			var newEndPoint EndPoint
			c.Bind(&newEndPoint)

			result := EndPoint{
				Id:        endpoint.Id,
				Label: newEndPoint.Label,
				Host:  newEndPoint.Host,
			}

			// UPDATE endpoints SET firstname='newEndPoint.Firstname', lastname='newEndPoint.Lastname' WHERE id = endpoint.Id;
			db.Save(&result)
			// Display modified data in JSON message "success"
			c.JSON(200, gin.H{"success": result})
		} else {
			// Display JSON error
			c.JSON(404, gin.H{"error": "EndPoint not found"})
		}

	} else {
		// Display JSON error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X PUT -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Merlyn\" }" http://localhost:8080/api/v1/endpoints/1
}

