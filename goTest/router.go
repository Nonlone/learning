package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RouterInit(){

	fmt.Println("router init start ...")
	r := gin.Default()
	defer r.Run(":12312")
	r.GET("/testA", getTestA)
	r.PUT("/testA", putTestA)
	r.POST("/testA",postTestA)
	r.DELETE("/testA",delTestA)

}

func getTestA(c * gin.Context){
	testA := &TestA{}
	engine.Get(testA)
	c.JSON(http.StatusOK,testA)
}


func putTestA(c *gin.Context) {
	c.String(http.StatusOK,"putTestA")

}

func postTestA(c *gin.Context){
	c.String(http.StatusOK,"postTestA")
}

func delTestA(c *gin.Context){
	c.String(http.StatusOK,"delTestA")
}



