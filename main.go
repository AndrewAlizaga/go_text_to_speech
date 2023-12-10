package main

import (
	"fmt"

	"github.com/AndrewAlizaga/go-text-to-speech/internal/v1/routes/text"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("invoke init")
	router := gin.Default()

	text.TextRouter(router.Group("/api/text"))
	// Your code here

	router.Run(":8080")
}
