package main

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"

	_ "github.com/mattn/go-sqlite3"
	"net/http"
	_ "github.com/gocql/gocql"
	_ "time"
)


func main() {
	r := gin.Default()
	r.Use(Cors())

	v1 := r.Group("api/v1")
	{
		v1.POST("/endpoint", PostEndPoint)
		v1.GET("/endpoints", GetEndPoints)
		v1.GET("/endpoint/:label", GetEndPoint)
		v1.PUT("/endpoint/:label", UpdateEndPoint)
		v1.DELETE("/endpoint/:label", DeleteEndPoint)
		v1.POST("/maintenance/:label", PostMaintenance)
		v1.GET("/maintenances", GetMaintenances)
		v1.DELETE("/maintenance/:label", DeleteMaintenance)
	}

	r.Run(":8080")
}

func InitDb() *gorm.DB {
	// Openning file
	db, err := gorm.Open("sqlite3", "./data.db") //, "gorm:gorm@/gorm?charset=utf8&parseTime=True")
	// Display SQL queries
	db.LogMode(true)

	// Error
	if err != nil {
		panic(err)
	}

	//db.DropTable(&EndPoint{})
	// Creating the table
	if !db.HasTable(&EndPoint{}) {
		db.CreateTable(&EndPoint{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&EndPoint{})
	}

	//db.DropTable(&Maintenance{})

	if !db.HasTable(&Maintenance{}) {
		db.CreateTable(&Maintenance{})
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Maintenance{})
	}

	return db
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func processFormField(r *http.Request, field string) (string, string) {
	fieldData := r.PostFormValue(field)
	if len(fieldData) == 0 {
		return "", "Missing '" + field + "' parameter, cannot continue"
	}
	return fieldData, ""
}

func OptionsEndPoint(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE,POST, PUT")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Next()
}
