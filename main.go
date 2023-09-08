package main

import (
	"contacts/frontend/frontend"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// to associate the GET http methos and the "/contacts" path to a handler function
	router.GET("/contacts", frontend.ListContacts)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(201, gin.H{"message": "Server working as expected!"})
	})

	// run on port 80
	router.Run(":80")
}
