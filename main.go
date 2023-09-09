package main

import (
	"contacts/frontend/frontend"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// to associate the GET http method and the "/contacts" path to a handler function
	router.GET("/contacts", frontend.ListContacts)
	// to bind the POST http method and the "/contacts" path to a handler function
	router.POST("/contacts", frontend.CreateContact)
	// to bind the GET http method and the "/contact/:id" path to a handler function
	router.GET("/contacts/:id", frontend.GetContact)
	// to bind the POST http method and the "/contact/:id" path to a handler function
	router.POST("/contacts/:id", frontend.UpdateContact)
	// to bind the DELETE http method and the "/contact/:id" path to a handler function
	router.DELETE("/contacts/:id", frontend.DeleteContact)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(201, gin.H{"message": "Server working as expected!"})
	})

	// run on port 80
	router.Run(":80")
}
