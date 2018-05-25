package main

import (
	"time"
	//"github.com/gocql/gocql"
	//"github.com/gocql/gocql"
	//"github.com/jinzhu/gorm"
)

type Service struct {
	ServiceName string `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
}

type EndPoint struct {
	Service
	Id           int    `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Label        string `gorm:"not null" form:"label" json:"label"`
	Host         string `gorm:"not null" form:"host" json:"host"`
	Port         string `gorm:"not null" form:"port" json:"port"`
	Uri          string `gorm:"not null" form:"uri" json:"uri"`
	HttpType     string `gorm: default:http json:"http_type"`
	BasicAuth    bool   `gorm: json:"basic_auth"`
	AuthUser     string `gorm: json:"auth_user"`
	AuthPassword string `gorm: json:"auth_password"`
}

type EndPoints []EndPoint

type Maintenance struct {
	//gorm.Model
	Id        int       `gorm:"AUTO_INCREMENT" form:"id" json:"id"`
	Label     string    `gorm:"not null" form:"label" json:"label"`
	StartTime time.Time `gorm:type:time" "not null" form:"starttime" json:"starttime"`
	EndTime   time.Time `gorm:type:time" "not null" form:"endtime" json:"endtime"`
}

type Maintenances []Maintenance
