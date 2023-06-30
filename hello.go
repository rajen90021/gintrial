package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
)

// func main() {
// 	//	testuuid := uuid.NewRandom()
// 	router := gin.Default()                       //it will create agin router with default middle ware for loger and recovery
// 	router.GET("/getdata", func(c *gin.Context) { ///contain resquest and response
// 		c.JSON(200, gin.H{
// 			"data": " hi i am gin framework",
// 		})

// 	})
// 	router.Run()

// 	//gfmt.Println(testuuid)

// }///*****************or****************
func main() {
	router := gin.Default()
	// router.GET("/getdata", getdata)
	router.POST("/getdatapost", getdatapost)
	router.Run()
}

// note 200 is response code
//
//	func getdata(c *gin.Context) {
//		c.JSON(200, gin.H{"data": "hi i am rajen trying to learn gin framework get "})
//	}
func getdatapost(c *gin.Context) {
	c.JSON(200, gin.H{"data": "hi i am rajen trying to learn gin framework post "})
}
