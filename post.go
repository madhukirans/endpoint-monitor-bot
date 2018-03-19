package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
	//"encoding/json"
	"encoding/json"
	"time"
)

func PostMaintenance(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	var maintenance Maintenance
	///var duration Duration

	label := c.Params.ByName("label")
	var endpoint EndPoint
	db.Where("label = ?", label).First(&endpoint)

	fmt.Printf("hi--------%+v\n", endpoint)

	if endpoint.Id != 0 {
		// Display JSON result
		//c.JSON(200, endpoint)
	} else {
		c.JSON(404, gin.H{"error": "Endpoint not found"})
		return
		// Display JSON error
		//
	}


	x, _ := ioutil.ReadAll(c.Request.Body)
	//fmt.Printf("%s \n", x)
	var dat map[string]interface{}
	json.Unmarshal(x, &dat)
	fmt.Println(dat)

	StartTime, _ := time.Parse( time.RFC3339, dat["starttime"].(string))
	EndTime, _ := time.Parse( time.RFC3339, dat["endtime"].(string))

	maintenance.Label = endpoint.Label
	maintenance.StartTime = StartTime
	maintenance.EndTime = EndTime

	if maintenance.EndTime.Sub(maintenance.StartTime).Seconds() > 0 {
		// INSERT INTO "endpoint" (name) VALUES (endpoint.Label);
		db.Create(&maintenance)
		// Display error
		c.JSON(201, gin.H{"success": maintenance})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/endpoints
}

func PostEndPoint(c *gin.Context) {
	db := InitDb()
	defer db.Close()

	print ("Madhu")
	var endpoint EndPoint
	print(c.Params)
	c.Bind(&endpoint)

	print("Madhu:" )
	print(endpoint.Label)

	if endpoint.Label != "" && endpoint.Host != "" {
		if (endpoint.HttpType == ""){
			endpoint.HttpType = "http"
		}
		// INSERT INTO "endpoint" (name) VALUES (endpoint.Label);
		db.Create(&endpoint)
		// Display error
		c.JSON(201, gin.H{"success": endpoint})
	} else {
		// Display error
		c.JSON(422, gin.H{"error": "Fields are empty"})
	}

	// curl -i -X POST -H "Content-Type: application/json" -d "{ \"firstname\": \"Thea\", \"lastname\": \"Queen\" }" http://localhost:8080/api/v1/endpoints
}