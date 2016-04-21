package main

import (
        "github.com/gin-gonic/gin"
        "./dbapi"
        )

func main() {

	//dbmap = initDb()
	router := gin.Default()
	router.GET("/user",DB.Createtable)
	router.GET("/user/:id",DB.Selectid)
	router.POST("/user",DB.Insert)
	router.Run(":2221")
}